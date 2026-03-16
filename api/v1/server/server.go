package server

import (
	"fmt"
	"net/http"
	apiservices "sofia-backend/api/v1/api-services"
	"sofia-backend/api/v1/controllers"
	"sofia-backend/config"
	"sofia-backend/domain/facades"
	"sofia-backend/types"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func NewV1Server(cfg *config.Config, appContainer *facades.FacadeContainer) *gin.Engine {
	// Crea una nueva instancia de Gin

	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	if cfg.Debug {
		router.Use(gin.Logger())

		router.Use(cors.New(cors.Config{
			AllowAllOrigins:  true, // igual que AllowOrigins: []string{"*"}
			AllowMethods:     []string{"*"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"*"},
			AllowCredentials: false, // debe ser false si AllowAllOrigins es true
			AllowWildcard:    true,  // permite el uso de "*"
		}))
	}
	router.Use(getRecovery())
	router.Use(getErrorHandler())

	// Configura las rutas y controladores, aca configuramos la api v1
	api := router.Group("/api/v1")

	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

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
	//purchaseController := controllers.NewPurchaseController(appContainer.PurchaseFacade)
	//purchaseController.RegisterRoutes(api)

	// Supplier controller
	supplierController := controllers.NewSupplierController(
		appContainer.SupplierFacade,
	)
	supplierController.RegisterRoutes(api)

	// Delivery Purchase Note controller
	//deliveryPurchaseNoteController := controllers.NewDeliveryPurchaseNoteController(appContainer.DeliveryPurchaseNoteFacade)
	//deliveryPurchaseNoteController.RegisterRoutes(api)

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
