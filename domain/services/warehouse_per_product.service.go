package services

import (
	"context"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
)

type WarehousePerProductService struct {
	PowerChecker
	repo ports.PortWarehouseProduct
}

func NewWarehousePerProductService(repo ports.PortWarehouseProduct) *WarehousePerProductService {
	return &WarehousePerProductService{
		repo: repo,
	}
}

func (s *WarehousePerProductService) GetAll(ctx context.Context) ([]models.ModelProductWarehouse, error) {
	return s.repo.GetAll()
}

func (s *WarehousePerProductService) GetAllByWarehouse(ctx context.Context, warehouseId string) ([]models.ModelProductWarehouse, error) {
	return s.repo.GetAllByWarehouse(warehouseId)
}

func (s *WarehousePerProductService) GetAllByStoreId(ctx context.Context, storeId string) ([]models.ModelProductWarehouse, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeId); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.repo.GetAllByStoreId(storeId)
}

func (s *WarehousePerProductService) GetById(ctx context.Context, id string) (*models.ModelProductWarehouse, error) {
	return s.repo.GetById(id)
}

func (s *WarehousePerProductService) GetWarehousePerProductByStoreProductAndWarehouse(ctx context.Context, storeProductId, warehouseId string) (*models.ModelProductWarehouse, error) {
	// No permissions check since this is used internally by system
	return s.repo.GetByProductAndWarehouseId(storeProductId, warehouseId)
}

func (s *WarehousePerProductService) CreateNewWPP(ctx context.Context, model *models.ModelProductWarehouse) (*models.ModelProductWarehouse, error) {
	return s.repo.CreateWPP(model)
}

func (s *WarehousePerProductService) UpdateWPP(ctx context.Context, model *models.ModelProductWarehouse) (*models.ModelProductWarehouse, error) {
	return s.repo.UpdateWPP(model)
}

func (s *WarehousePerProductService) MakeProductTransfer(productId string, warehouseFrom, warehouseTo *models.ModelProductWarehouse) error {
	return s.repo.MakeProductTransfer(productId, warehouseFrom, warehouseTo)
}

func (s *WarehousePerProductService) MakeProductInput(productId string, warehouseTo *models.ModelProductWarehouse) error {
	return s.repo.MakeProductInput(productId, warehouseTo)
}

func (s *WarehousePerProductService) MakeProductWaste(productId string, warehouseFrom *models.ModelProductWarehouse) error {
	return s.repo.MakeProductWaste(productId, warehouseFrom)
}
