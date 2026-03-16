package models

import (
	"time"
)

type UserAccountModel struct {
	ID           string     `json:"id"`
	UserName     string     `json:"userName"`
	UserEmail    string     `json:"userEmail"`
	Description  string     `json:"description"`
	UserPassword string     `json:"userPassword"`
	Available    bool       `json:"available"`
	IsNewAccount bool       `json:"isNewAccount"` // Indicates if the account is new
	DeletedAt    *time.Time `json:"deletedAt"`    // Nullable field for soft delete
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}
