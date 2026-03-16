package ports

import "sofia-backend/domain/models"

type PortSee interface {
	SendMessage(notification models.NotificationModel)
	GetClients() []models.ClientModel
	AddClient(client models.ClientModel)
	RemoveClient(client models.ClientModel)
	BroadcastMessage(notification models.NotificationModel)
	GetClient(userID string) chan models.NotificationModel
}
