package recipe

type MarkNotificationsAsReadRecipe struct {
	NotificationIDs []string `json:"notificationIds" binding:"required"`
	WasRead         bool     `json:"wasRead" binding:"omitempty"`
}

type SendNotificationRecipe struct {
	From             string      `json:"from" binding:"required"`
	To               string      `json:"to" binding:"required"`
	NotificationType string      `json:"notificationType" binding:"required"`
	Payload          interface{} `json:"payload" binding:"required"`
}
