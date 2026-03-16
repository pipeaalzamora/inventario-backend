package ports

import (
	"sofia-backend/domain/models"
)

// PortProductMovement define las operaciones para gestionar movimientos de productos.
// Ahora usa StoreProductID (producto por tienda) en lugar de ProductCompanyID.
type PortProductMovement interface {
	GetAllProductMovements() ([]models.ModelProductMovement, error)
	GetByProductMovementId(movementId string) (*models.ModelProductMovement, error)
	GetAllProductMovementsByStoreProductID(storeProductId string) ([]models.ModelProductMovement, error)
	GetAllProductMovementsByStoreID(storeId string) ([]models.ModelProductMovement, error)
	GetAllProductMovementsByWarehouseIDs(warehouseIDs []string) ([]models.ModelProductMovement, error)
	GetAllProductMovementsByDateRange(warehouseID string) ([]models.ModelProductMovement, error)

	CreateNewSingleMovement(model models.ModelProductMovement) (*models.ModelProductMovement, error)
	CreateNewMultiplesMovements(models []models.ModelProductMovement) ([]models.ModelProductMovement, error)

	// CreateTransferMovements crea movimientos de transferencia entre bodegas
	// Actualiza el stock en ambas bodegas y el costo promedio en la bodega destino
	CreateTransferMovements(movements []models.ModelProductMovement, newAvgCosts map[string]float32) ([]models.ModelProductMovement, error)
}
