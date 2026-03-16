package ports

import (
	"sofia-backend/domain/models"
	"time"
)

type PortInventoryCount interface {
	GetAll() ([]models.ModelInventoryCount, error)
	GetByID(id string) (*models.ModelInventoryCount, error)
	GetCompletedByID(itemsMetadata []models.ModelInventoryCountMetadata) ([]models.ModelInventoryCountItem, error)
	GetItemsByInventoryCountID(id string) ([]models.ModelInventoryCountItem, error)
	GetAllByUserId(userId string) ([]models.ModelInventoryCount, error)
	Create(report *models.ModelInventoryCount) (*models.ModelInventoryCount, error)
	Update(model *models.ModelInventoryCount) (*models.ModelInventoryCount, error)
	Delete(id string) error
	Commit(model *models.ModelInventoryCount) (*models.ModelInventoryCount, error)
	ChangeAssigned(id string, userId *string) (*models.ModelInventoryCount, error)
	ChangeDate(id string, newDate time.Time) (*models.ModelInventoryCount, error)
	ChangeState(id string, newState string) (*models.ModelInventoryCount, error)
	Reject(model *models.ModelInventoryCount) (*models.ModelInventoryCount, error)
	GetIncidenceByProduct(countId string, productId string) (*models.ModelInventoryCountItem, error)
	SaveIncidence(countId string, productId string, imageUrl *string, observation *string) error
	DeleteIncidenceImage(countId string, productId string, observation *string) error
}
