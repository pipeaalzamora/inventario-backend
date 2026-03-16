package facades

import (
	"context"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/shared"
)

type NotificationFacade struct {
	appService *services.ServiceContainer
}

func NewNotificationFacade(appService *services.ServiceContainer) *NotificationFacade {
	return &NotificationFacade{
		appService: appService,
	}
}

func (f *NotificationFacade) GetNotifications(ctx context.Context, page int, size int, filter *map[string]interface{}) (shared.PaginationResponse[models.NotificationModel], error) {
	userId, ok := f.appService.NotificationService.GetUserIDFromContext(ctx)
	if !ok {
		return shared.PaginationResponse[models.NotificationModel]{}, nil
	}
	notifications, total, err := f.appService.NotificationService.GetNotificationsByUserID(userId, page, size, filter)
	if err != nil {
		return shared.PaginationResponse[models.NotificationModel]{}, err
	}

	return shared.NewPagination(notifications, total, page, size), nil
}

func (f *NotificationFacade) MarkNotificationsAsRead(ctx context.Context, body *recipe.MarkNotificationsAsReadRecipe) error {
	userId, ok := f.appService.NotificationService.GetUserIDFromContext(ctx)
	if !ok {
		return nil
	}
	err := f.appService.NotificationService.MarkNotificationsAsRead(userId, body.NotificationIDs, body.WasRead)
	if err != nil {
		return err
	}
	return nil
}

func (f *NotificationFacade) SendNotification(body *recipe.SendNotificationRecipe) error {
	_, err := f.appService.NotificationService.SendMessage(body)
	if err != nil {
		return err
	}

	return nil
}

func (f *NotificationFacade) BroadcastNotification(message interface{}) error {
	return f.appService.NotificationService.BroadcastMessage(message)
}

func (f *NotificationFacade) RemoveClient(ctx context.Context) {
	userId, ok := f.appService.NotificationService.GetUserIDFromContext(ctx)
	if !ok {
		return
	}
	f.appService.NotificationService.RemoveClient(userId)
}

func (f *NotificationFacade) RegisterClient(ctx context.Context) chan models.NotificationModel {
	userId, ok := f.appService.NotificationService.GetUserIDFromContext(ctx)
	if !ok {
		return nil
	}
	return f.appService.NotificationService.RegisterClient(userId)
}
