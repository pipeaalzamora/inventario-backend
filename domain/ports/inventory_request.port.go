package ports

import (
	"sofia-backend/domain/models"
	"sofia-backend/infraestructure/entities"
)

type PortInventoryRequest interface {
	GetInventoryRequestsByStore(storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelInventoryRequest, int, error)
	UpdateInventoryRequest(id string, model *models.ModelInventoryRequest, obs *string, userID string) (*models.ModelInventoryRequest, error)
	ChangeInventoryRequestStatus(id string, newStatus entities.RequestStatus, obs *string, userID string) (*models.ModelInventoryRequest, error)
	CreateInventoryRequest(model *models.ModelInventoryRequest, obs *string) (*models.ModelInventoryRequest, error)
	GetInventoryRequestByID(id string) (*models.ModelInventoryRequest, error)
	ApproveAndUpdateRequest(id string, model *models.ModelInventoryRequest, obs *string, userID string) (*models.ModelInventoryRequest, error)
}
