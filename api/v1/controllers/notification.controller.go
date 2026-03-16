package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	apiservices "sofia-backend/api/v1/api-services"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/facades"
	"sofia-backend/shared"
	"time"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	notificationFacade *facades.NotificationFacade
}

func NewNotificationController(notificationFacade *facades.NotificationFacade) *NotificationController {
	return &NotificationController{notificationFacade: notificationFacade}
}

func (ctrl *NotificationController) RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/notifications")

	r.GET("", ctrl.getNotifications)
	r.GET("/subscribe", apiservices.SetupSSe(), ctrl.subscribe)
	r.POST("/mark-as-read", ctrl.markNotificationsAsRead)
}

func (ctrl *NotificationController) subscribe(gctx *gin.Context) {
	// Set headers first
	w := gctx.Writer

	// Flush headers immediately
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	clientChannel := ctrl.notificationFacade.RegisterClient(gctx)
	defer ctrl.notificationFacade.RemoveClient(gctx)

	for {
		select {
		case <-gctx.Done():
			fmt.Println("Client disconnected (context done)")
			return

		case <-ticker.C:
			// Send proper SSE ping with retry logic
			w.Write([]byte("data: ping\n\n"))
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}

		case msg, ok := <-clientChannel:
			if !ok {
				fmt.Println("Client channel closed")
				return
			}
			fmt.Printf("Streaming message: %+v\n", msg)
			jsonMsg, err := json.Marshal(msg)
			if err != nil {
				fmt.Printf("Error marshalling message: %v\n", err)
				continue
			}
			// Proper SSE format with error handling
			_, writeErr := w.Write([]byte(fmt.Sprintf("data: %s\n\n", jsonMsg)))
			if writeErr != nil {
				fmt.Printf("Error writing to client: %v\n", writeErr)
				return
			}
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		}
	}
}

func (ctrl *NotificationController) markNotificationsAsRead(gctx *gin.Context) {
	// Validar el cuerpo de la solicitud
	var body recipe.MarkNotificationsAsReadRecipe
	if err := gctx.ShouldBindBodyWithJSON(&body); err != nil {
		gctx.Error(err)
		return
	}

	err := ctrl.notificationFacade.MarkNotificationsAsRead(gctx, &body)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.Status(http.StatusNoContent)
}

func (ctrl *NotificationController) getNotifications(gctx *gin.Context) {
	// Validar los parámetros de consulta
	var queryParams shared.PageQueryParams
	if err := gctx.ShouldBindQuery(&queryParams); err != nil {
		gctx.Error(err)
		return
	}

	notifications, err := ctrl.notificationFacade.GetNotifications(gctx, queryParams.Page, queryParams.Size, queryParams.Filter)
	if err != nil {
		gctx.Error(err)
		return
	}

	gctx.JSON(http.StatusOK, notifications)
}
