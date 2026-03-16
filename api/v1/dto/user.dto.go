package dto

import (
	"sofia-backend/domain/models"
	"time"
)

type UserDTO struct {
	ID           string     `json:"id"`
	UserName     string     `json:"userName"`
	UserEmail    string     `json:"userEmail"`
	Description  string     `json:"description"`
	Available    bool       `json:"available"`
	IsNewAccount bool       `json:"isNewAccount"`
	DeletedAt    *time.Time `json:"deletedAt"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

type UserWithProfilesDTO struct {
	ID           string                       `json:"id"`
	UserName     string                       `json:"userName"`
	UserEmail    string                       `json:"userEmail"`
	Description  string                       `json:"description"`
	Available    bool                         `json:"available"`
	IsNewAccount bool                         `json:"isNewAccount"`
	DeletedAt    *time.Time                   `json:"deletedAt"`
	CreatedAt    time.Time                    `json:"createdAt"`
	UpdatedAt    time.Time                    `json:"updatedAt"`
	Profiles     []models.ProfileAccountModel `json:"profiles"`
}
