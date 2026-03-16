package ports

import (
	"sofia-backend/domain/models"
)

type PortProfile interface {
	GetProfilesByUserID(userId string) ([]models.ProfileAccountModel, error)
	GetProfileByID(profileId string) (*models.ProfileAccountModel, error)
	GetProfiles() ([]models.ProfileAccountModel, error)
	CreateProfile(input *models.ProfileAccountModel) (*models.ProfileAccountModel, error)
	UpdateProfile(id string, input *models.ProfileAccountModel) (*models.ProfileAccountModel, error)
	CreateProfileWithPowers(input *models.ProfileAccountModel, powerIDs []string) (*models.ProfileAccountModel, error)
	UpdateProfileWithPowers(id string, input *models.ProfileAccountModel, powerIDs []string) (*models.ProfileAccountModel, error)
	DeleteProfile(id string) error
	GetAllProfilesWithUserId() ([]models.ProfileAccountUserAccountModel, error)
	GetProfilesByPowerID(powerName string) ([]models.ProfileAccountModel, error)
}
