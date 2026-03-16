package dto

import "sofia-backend/domain/models"

type ProfileAccountWithPowersDTO struct {
	ID          string                     `json:"id"`
	ProfileName string                     `json:"profileName"`
	Description string                     `json:"description"`
	Powers      []models.PowerAccountModel `json:"powers"`
}
