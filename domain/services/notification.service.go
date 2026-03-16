package services

import (
	"fmt"
	"log"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
)

type NotificationService struct {
	PowerChecker
	notificationRepo ports.PortNotification
	sseService       ports.PortSee
}

func NewNotificationService(repo ports.PortNotification, sseService ports.PortSee) *NotificationService {
	return &NotificationService{
		notificationRepo: repo,
		sseService:       sseService,
	}
}

func (s *NotificationService) GetNotificationsByUserID(userID string, page int, size int, filter *map[string]interface{}) (data []models.NotificationModel, total int, err error) {
	notifications, total, err := s.notificationRepo.GetNotificationsByUserID(userID, page, size, filter)

	if err != nil {
		return nil, 0, err
	}

	return notifications, total, nil
}

func (s *NotificationService) MarkNotificationsAsRead(userId string, ids []string, wasRead bool) error {
	err := s.notificationRepo.UpdateRead(userId, ids, wasRead)
	if err != nil {
		return err
	}
	return nil
}

func (s *NotificationService) RemoveClient(userId string) {
	clientChannel := s.sseService.GetClient(userId)
	if clientChannel != nil {
		s.sseService.RemoveClient(models.ClientModel{ID: userId, Channel: clientChannel})
	}
}

func (s *NotificationService) RegisterClient(userId string) chan models.NotificationModel {
	client := s.sseService.GetClient(userId)
	if client != nil {
		log.Printf("Client %s already registered", userId)
		return client
	}

	channel := models.ClientModel{
		ID:      userId,
		Channel: make(chan models.NotificationModel, 100), // Buffered channel to prevent blocking
	}
	s.sseService.AddClient(channel)
	return channel.Channel
}

func (s *NotificationService) SendMessage(notification *recipe.SendNotificationRecipe) (*models.NotificationModel, error) {
	input := &models.NotificationModel{
		From:             &notification.From,
		To:               notification.To,
		Payload:          notification.Payload,
		NotificationType: notification.NotificationType,
		ReadAt:           nil,
	}

	notificationCreated, err := s.notificationRepo.CreateNotification(input)
	if err != nil {
		return nil, err
	}

	s.sseService.SendMessage(*notificationCreated)

	return notificationCreated, nil
}

func (s *NotificationService) SendMultipleMessages(notifications []recipe.SendNotificationRecipe) ([]models.NotificationModel, error) {
	var createdNotifications []models.NotificationModel

	for _, noti := range notifications {
		createdNoti, err := s.SendMessage(&noti)
		if err != nil {
			log.Printf("Error sending notification to %s: %v", noti.To, err)
			continue
		}
		createdNotifications = append(createdNotifications, *createdNoti)
	}

	if len(createdNotifications) == 0 {
		return nil, types.ThrowMsg("No se pudieron enviar las notificaciones")
	}

	return createdNotifications, nil
}

func (s *NotificationService) BroadcastMessage(message interface{}) error {
	clients := s.sseService.GetClients()
	if len(clients) == 0 {
		return types.ThrowMsg("No hay clientes conectados para enviar la notificación.")
	}

	for _, client := range clients {
		noti := &models.NotificationModel{
			From:             nil,
			To:               client.ID,
			Payload:          message,
			NotificationType: "broadcast",
			ReadAt:           nil,
		}
		notificationModel, err := s.notificationRepo.CreateNotification(noti)
		if err != nil {
			fmt.Printf("Error creating notification: %v\n", err)
			continue
		}

		if notificationModel == nil {
			continue
		}

		// marshal is only required for db not for sending
		s.sseService.SendMessage(*notificationModel)
	}
	return nil
}
