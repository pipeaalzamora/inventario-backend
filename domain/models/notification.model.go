package models

import (
	"time"
)

type ClientModel struct {
	// Unique Client ID
	ID string
	// Client channel
	Channel chan NotificationModel
}

type NotificationModel struct {
	ID               string      `json:"id"`
	From             *string     `json:"from"`
	To               string      `json:"to"`
	ReadAt           *time.Time  `json:"readAt"`
	SendAt           time.Time   `json:"sendAt"`
	NotificationType string      `json:"notificationType"`
	Payload          interface{} `json:"payload"`
}
