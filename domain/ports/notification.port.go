package ports

import (
	"sofia-backend/domain/models"
)

type PortNotification interface {
	GetNotificationsByUserID(userID string, page int, size int, filter *map[string]interface{}) ([]models.NotificationModel, int, error)
	CreateNotification(notification *models.NotificationModel) (*models.NotificationModel, error)
	UpdateRead(userId string, notificationIDs []string, wasRead bool) error
}
