package ports

import (
	"sofia-backend/domain/models"

	"github.com/jmoiron/sqlx"
)

type PortWarehouse interface {
	//GetAll() ([]models.ModelWarehouse, error)
	GetWarehouseByID(id string) (*models.ModelWarehouse, error)
	GetWarehousesByStoreId(id string) ([]models.ModelWarehouse, error)
	GetTransitionWarehouseByStoreID(storeID string) (*models.ModelWarehouse, error)
	CreateWarehouse(warehouse *models.ModelWarehouse) (*models.ModelWarehouse, error)
	UpdateWarehouse(warehouse *models.ModelWarehouse) (*models.ModelWarehouse, error)

	//DeleteWarehouse(id string) error

	CreateFirstWarehouse(tx *sqlx.Tx, store *models.StoreModel) error
}
