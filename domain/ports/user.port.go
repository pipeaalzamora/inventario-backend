package ports

import (
	"sofia-backend/domain/models"
)

type PortUser interface {
	UpdateUser(id string, user *models.UserAccountModel) (*models.UserAccountModel, error)
	UpdateUserWithProfiles(userId string, input *models.UserAccountModel, profileIDs []string) (*models.UserAccountModel, error)
	UpdateUserPassword(userId string, password string) error
	CreateUserWithProfiles(input *models.UserAccountModel, profileIDs []string) (*models.UserAccountModel, error)
	CreateUser(user *models.UserAccountModel) (*models.UserAccountModel, error)
	GetUsersByProfileID(page int, size int, filter *map[string]interface{}, profileID string) ([]models.UserAccountModel, int, error)
	GetAllUsers(page int, size int, filter *map[string]interface{}) ([]models.UserAccountModel, int, error)
	GetUsersByProfileIDs(profileIDs []string) ([]models.UserAccountModel, error)
	GetUsers() ([]models.UserAccountModel, error)
	GetUserByEmail(email string) (*models.UserAccountModel, error)
	GetUserByID(id string) (*models.UserAccountModel, error)
	DeactivateUser(id string) error
	ActivateUser(id string) error
}
