package services

import (
	"context"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
)

type InventoryService struct {
	PowerChecker
	repo ports.PortInventory
}

func NewInventoryService(inventoryPort ports.PortInventory) *InventoryService {
	return &InventoryService{
		repo: inventoryPort,
	}
}

func (s *InventoryService) GetProductStoreInventoryByStoreID(ctx context.Context, storeID string, warehouseIDs []string) ([]models.ModelWarehouseProductStock, error) {
	return s.repo.GetProductStoreInventoryByStoreID(storeID, warehouseIDs)
}

func (s *InventoryService) GetProductTransitByReferences(ctx context.Context, storeID string, warehouseIDs []string) ([]models.ModelWarehouseProductTransit, error) {
	return s.repo.GetProductTransitByReferences(storeID, warehouseIDs)
}

func (s *InventoryService) GetSingleProductStock(ctx context.Context, storeID string, warehouseID string, storeProductID string) (*models.ModelWarehouseProductStock, error) {
	return s.repo.GetSingleProductStock(storeID, warehouseID, storeProductID)
}

func (s *InventoryService) GetSingleProductTransit(ctx context.Context, storeID string, warehouseID string, storeProductID string) ([]models.ModelWarehouseProductTransit, error) {
	return s.repo.GetSingleProductTransit(storeID, warehouseID, storeProductID)
}

func (s *InventoryService) GetCurrentStock(ctx context.Context, companyId, storeId, warehouseId string) (*models.ModelWarehouseProductStock, error) {
	s.PowerChecker.EveryPower(ctx,
		PowerPrefixCompany+companyId,
		PowerPrefixStore+storeId,
	)

	return s.repo.GetCurrentStock(companyId, storeId, warehouseId)
}
