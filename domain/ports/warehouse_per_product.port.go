package ports

import "sofia-backend/domain/models"

// PortWarehouseProduct define las operaciones para gestionar stock de productos en bodegas.
// Ahora usa store_product_id (producto por tienda) en lugar de product_company_id.
type PortWarehouseProduct interface {
	GetAll() ([]models.ModelProductWarehouse, error)
	GetAllByWarehouse(warehouseId string) ([]models.ModelProductWarehouse, error)
	GetAllByStoreId(storeId string) ([]models.ModelProductWarehouse, error) // Reemplaza GetAllByCompanyId
	GetById(id string) (*models.ModelProductWarehouse, error)
	GetByProductAndWarehouseId(storeProductId, warehouseId string) (*models.ModelProductWarehouse, error)

	CreateWPP(productWarehouse *models.ModelProductWarehouse) (*models.ModelProductWarehouse, error)
	UpdateWPP(productWarehouse *models.ModelProductWarehouse) (*models.ModelProductWarehouse, error)

	// storeProductId es el ID de product_per_store
	MakeProductTransfer(storeProductId string, warehouseFrom, warehouseTo *models.ModelProductWarehouse) error
	MakeProductInput(storeProductId string, warehouse *models.ModelProductWarehouse) error
	MakeProductWaste(storeProductId string, warehouse *models.ModelProductWarehouse) error
}
