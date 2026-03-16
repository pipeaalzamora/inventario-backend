package ports

import "sofia-backend/domain/models"

type PortInventory interface {
	GetProductStoreInventoryByStoreID(storeID string, warehouseIDs []string) ([]models.ModelWarehouseProductStock, error)
	GetProductTransitByReferences(storeID string, warehouseIDs []string) ([]models.ModelWarehouseProductTransit, error)
	GetSingleProductStock(storeID string, warehouseID string, storeProductID string) (*models.ModelWarehouseProductStock, error)
	GetSingleProductTransit(storeID string, warehouseID string, storeProductID string) ([]models.ModelWarehouseProductTransit, error)
	GetCurrentStock(companyId, storeId, warehouseId string) (*models.ModelWarehouseProductStock, error)
}
