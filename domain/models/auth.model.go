package models

import "time"

type RecoveryCode struct {
	Code     string    `json:"code"`
	TTL      time.Time `json:"ttl"`
	Verified bool      `json:"verified"`
}
