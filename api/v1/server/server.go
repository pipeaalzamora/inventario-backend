package server

import (
	"fmt"
	"log/slog"
	"net/http"
	apiservices "sofia-backend/api/v1/api-services"
	"sofia-backend/api/v1/controllers"
	"sofia-backend/config"
	"sofia-backend/domain/facades"
	"sofia-backend/types"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var metrics = &httpMetrics{}

type httpMetrics struct {
	totalRequests atomic.Uint64
	totalErrors   atomic.Uint64
	inFlight      atomic.Int64
	latencyNs     atomic.Uint64
	status2xx     atomic.Uint64
	status3xx     atomic.Uint64
	status4xx     atomic.Uint64
	status5xx     atomic.Uint64
}

func NewV1Server(cfg *config.Config, appContainer *facades.FacadeContainer) *gin.Engine {
	// Crea una nueva instancia de Gin

	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(requestLogger())
	router.Use(metricsMiddleware())
	router.Use(cors.New(corsConfig(cfg)))
	router.Use(getRecovery())
	router.Use(getErrorHandler())

	// Configura las rutas y controladores, aca configuramos la api v1
	api := router.Group("/api/v1")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	api.GET("/metrics", metricsHandler)

	// auth controller
	authController := controllers.NewAuthController(appContainer.AuthFacade, cfg)
	authController.RegisterRoutes(api)

	// supplier_oc public
	//supplierOCController := controllers.NewSupplierOCController(appContainer.SupplierOCFacade)
	//supplierOCController.RegisterRoutes(api)

	// Auth routes
	api.Use(apiservices.AuthMiddleware(appContainer.AuthFacade, cfg.JwtSecret))

	// User controller
	userController := controllers.NewUserController(appContainer.UserFacade)
	userController.RegisterRoutes(api)

	// EntityNotification controller
	notificationController := controllers.NewNotificationController(appContainer.NotificationFacade)
	notificationController.RegisterRoutes(api)

	// Profile account controller
	profileController := controllers.NewProfileAccountController(appContainer.ProfileFacade)
	profileController.RegisterRoutes(api)

	// Product controller
	productController := controllers.NewProductController(appContainer.ProductFacade)
	productController.RegisterRoutes(api)

	// Company controller
	companyController := controllers.NewCompanyController(appContainer.CompanyFacade)
	companyController.RegisterRoutes(api)

	// Store controller
	storeController := controllers.NewStoreController(appContainer.StoreFacade)
	storeController.RegisterRoutes(api)

	// Warehouse controller
	warehouseController := controllers.NewWarehouseController(appContainer.WarehouseFacade)
	warehouseController.RegisterRoutes(api)

	// Inventory  controller
	inventoryController := controllers.NewInventoryController(appContainer.InventoryFacade)
	inventoryController.RegisterRoutes(api)

	// Inventory Request controller
	//inventoryRequestController := controllers.NewInventoryRequestController(appContainer.InventoryRequestFacade)
	//inventoryRequestController.RegisterRoutes(api)

	// Inventory Report controller
	inventoryReportController := controllers.NewInventoryCountController(*appContainer.InventoryCountFacade)
	inventoryReportController.RegisterRoutes(api)

	// Product Company controller
	productCompanyController := controllers.NewProductCompanyController(appContainer.ProductFacade)
	productCompanyController.RegisterRoutes(api)

	// Purchase controller
	purchaseController := controllers.NewPurchaseController(appContainer.PurchaseFacade)
	purchaseController.RegisterRoutes(api)

	// Supplier controller
	supplierController := controllers.NewSupplierController(
		appContainer.SupplierFacade,
	)
	supplierController.RegisterRoutes(api)

	// Delivery Purchase Note controller
	deliveryPurchaseNoteController := controllers.NewDeliveryPurchaseNoteController(appContainer.DeliveryPurchaseNoteFacade)
	deliveryPurchaseNoteController.RegisterRoutes(api)

	// Product Movement controller
	productMovementController := controllers.NewProductMovementController(appContainer.ProductMovementFacade)
	productMovementController.RegisterRoutes(api)

	// Measurement controller
	measurementController := controllers.NewMeasurementController(appContainer.MeasurementFacade)
	measurementController.RegisterRoutes(api)

	// Store Product controller
	storeProductController := controllers.NewStoreProductController(appContainer.StoreProductFacade)
	storeProductController.RegisterRoutes(api)

	requestController := controllers.NewRequestController(appContainer.RequestFacade)
	requestController.RegisterRoutes(api)

	return router
}

func corsConfig(cfg *config.Config) cors.Config {
	if cfg.Debug {
		return cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Authorization", "Content-Type", "Accept"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: false,
			AllowWildcard:    true,
		}
	}

	origins := []string{}
	if strings.TrimSpace(cfg.FrontUrl) != "" {
		origins = append(origins, strings.TrimRight(cfg.FrontUrl, "/"))
	}

	return cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}
}

func requestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()

		slog.Info("http_request",
			"method", ctx.Request.Method,
			"path", ctx.FullPath(),
			"status", ctx.Writer.Status(),
			"latency_ms", time.Since(start).Milliseconds(),
			"client_ip", ctx.ClientIP(),
		)
	}
}

func metricsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		metrics.inFlight.Add(1)
		defer metrics.inFlight.Add(-1)

		ctx.Next()

		status := ctx.Writer.Status()
		metrics.totalRequests.Add(1)
		metrics.latencyNs.Add(uint64(time.Since(start).Nanoseconds()))
		switch {
		case status >= 500:
			metrics.status5xx.Add(1)
			metrics.totalErrors.Add(1)
		case status >= 400:
			metrics.status4xx.Add(1)
			metrics.totalErrors.Add(1)
		case status >= 300:
			metrics.status3xx.Add(1)
		default:
			metrics.status2xx.Add(1)
		}
	}
}

func metricsHandler(ctx *gin.Context) {
	total := metrics.totalRequests.Load()
	avgLatency := float64(0)
	if total > 0 {
		avgLatency = float64(metrics.latencyNs.Load()) / float64(total) / float64(time.Second)
	}

	body := strings.Join([]string{
		"# HELP sofia_http_requests_total Total HTTP requests.",
		"# TYPE sofia_http_requests_total counter",
		"sofia_http_requests_total " + strconv.FormatUint(total, 10),
		"# HELP sofia_http_errors_total Total HTTP requests ending in 4xx or 5xx.",
		"# TYPE sofia_http_errors_total counter",
		"sofia_http_errors_total " + strconv.FormatUint(metrics.totalErrors.Load(), 10),
		"# HELP sofia_http_in_flight_requests Current in-flight HTTP requests.",
		"# TYPE sofia_http_in_flight_requests gauge",
		"sofia_http_in_flight_requests " + strconv.FormatInt(metrics.inFlight.Load(), 10),
		"# HELP sofia_http_request_duration_seconds_avg Average HTTP request duration in seconds since process start.",
		"# TYPE sofia_http_request_duration_seconds_avg gauge",
		"sofia_http_request_duration_seconds_avg " + strconv.FormatFloat(avgLatency, 'f', 6, 64),
		"# HELP sofia_http_responses_total HTTP responses by status class.",
		"# TYPE sofia_http_responses_total counter",
		"sofia_http_responses_total{class=\"2xx\"} " + strconv.FormatUint(metrics.status2xx.Load(), 10),
		"sofia_http_responses_total{class=\"3xx\"} " + strconv.FormatUint(metrics.status3xx.Load(), 10),
		"sofia_http_responses_total{class=\"4xx\"} " + strconv.FormatUint(metrics.status4xx.Load(), 10),
		"sofia_http_responses_total{class=\"5xx\"} " + strconv.FormatUint(metrics.status5xx.Load(), 10),
		"",
	}, "\n")

	ctx.Data(http.StatusOK, "text/plain; version=0.0.4; charset=utf-8", []byte(body))
}

func getErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			var errs []types.ErrorResponse
			status := http.StatusBadRequest

			for _, e := range ctx.Errors {
				if ae, ok := e.Err.(types.AppError); ok {
					withParam := false
					// Mapear tipo de dominio a código HTTP
					switch ae.Type() {
					case types.DataError:
						status = 422
					case types.PowerError:
						status = 403
					case types.RecipeError:
						status = 400
						withParam = true
					default:
						status = 400
					}

					if withParam {
						errs = append(errs, types.ErrorResponse{
							Message: ae.Error(),
							Param:   ae.Param(),
						})
						continue
					}

					errs = append(errs, types.ErrorResponse{
						Message: ae.Error(),
					})

					continue
				}

				if ve, ok := e.Err.(validator.ValidationErrors); ok {
					for _, fe := range ve {
						//el mensaje es un error de validacion que podria venir en formato json
						errs = append(errs, types.ErrorResponse{Param: fe.Field(), Message: fe.Error()})
					}
				} else {
					// Para cualquier otro tipo de error
					errs = append(errs, types.ErrorResponse{Message: e.Error()})
				}

			}

			ctx.JSON(status, errs)
			ctx.Abort()
		}
	}
}

func getRecovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ctx.JSON(http.StatusInternalServerError, []types.ErrorResponse{
					{Message: fmt.Sprintf("Error interno: %v", r)},
				})
				ctx.Abort()

				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		ctx.Next()
	}
}
