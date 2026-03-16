package entities

import (
	"encoding/json"
	"time"
)

type EntityNotification struct {
	ID               string          `db:"id" json:"id"`
	From             *string         `db:"from_user" json:"from"`
	To               string          `db:"to_user" json:"to"`
	ReadAt           *time.Time      `db:"read_at" json:"read_at"`
	SendAt           time.Time       `db:"send_at" json:"send_at"`
	NotificationType string          `db:"notification_type" json:"notification_type"`
	Payload          json.RawMessage `db:"payload" json:"payload"`
}

type FilterNotification struct {
	WasRead *bool `json:"wasRead"`
}
