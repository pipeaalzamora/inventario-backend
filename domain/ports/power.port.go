package ports

import (
	"sofia-backend/domain/models"

	"github.com/jmoiron/sqlx"
)

type PortPower interface {
	GetPowersByProfile(profileId string) ([]models.PowerAccountModel, error)
	GetPowerCategories() ([]models.PowerAccountCategoryModel, error)
	GetPowerCategoryByID(categoryId string) (*models.PowerAccountCategoryModel, error)
	GetPowerAccountsByCategory(categoryId string) ([]models.PowerAccountModel, error)
	GetPowers() ([]models.PowerAccountModel, error)
	GetPowersByUserID(userId string) ([]models.PowerAccountModel, error)
	GetAllPowersWithProfileId() ([]models.PowerAccountProfileModel, error)
	// CreatePowerCategory(input *recipe.PowerCategoryRecipe) (*models.PowerAccountCategoryModel, error)
	// CreatePowerAccount(input *recipe.PowerAccountRecipe) (*models.PowerAccountModel, error)

	AddOwnPowerToProfileTx(tx *sqlx.Tx, profileIds []string, power *models.PowerAccountModel) error
}
