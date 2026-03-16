package services

import (
	"context"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
)

type WarehouseService struct {
	PowerChecker
	warehouseRepo ports.PortWarehouse
	storeRepo     ports.PortStore
}

func NewWarehouseService(warehouseRepo ports.PortWarehouse, storeRepo ports.PortStore) *WarehouseService {
	return &WarehouseService{
		warehouseRepo: warehouseRepo,
		storeRepo:     storeRepo,
	}
}

/*
func (s *WarehouseService) GetAll(ctx context.Context) ([]models.ModelWarehouse, error) {
	powers, ok := s.GetPowersFromContext(ctx)
	if !ok {
		return nil, nil
	}

	// Extract store IDs from powers
	filteredStoreID := map[string]bool{}
	for _, p := range powers {
		if strings.HasPrefix(p, PowerPrefixStore) {
			storeId := p[len(PowerPrefixStore):]
			filteredStoreID[storeId] = true
		}
	}

	allWarehouses, err := s.warehouseRepo.GetAll()
	if err != nil {
		return nil, err
	}

	result := []models.ModelWarehouse{}
	for _, warehouse := range allWarehouses {
		if filteredStoreID[warehouse.StoreId] {
			result = append(result, warehouse)
		}
	}

	return result, nil
}
*/

func (s *WarehouseService) GetWarehousesByStoreId(ctx context.Context, id string) ([]models.ModelWarehouse, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+id); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.warehouseRepo.GetWarehousesByStoreId(id)
}

func (s *WarehouseService) GetTransitionWarehouseByStoreID(ctx context.Context, storeID string) (*models.ModelWarehouse, error) {
	// No require permissions check since it's used internally by the system
	// for automatic purchase movements
	return s.warehouseRepo.GetTransitionWarehouseByStoreID(storeID)
}

func (s *WarehouseService) GetWarehouseByID(ctx context.Context, id string) (*models.ModelWarehouse, error) {
	warehouse, err := s.warehouseRepo.GetWarehouseByID(id)
	if err != nil {
		return nil, err
	}
	if warehouse == nil {
		return nil, nil
	}

	// Obtain store to get company ownership
	store, err := s.storeRepo.GetStoreByID(warehouse.StoreId)
	if err != nil {
		return nil, err
	}
	if store == nil {
		return nil, types.ThrowMsg("Tienda no encontrada")
	}

	// Validate company and store ownership
	if ok := s.EveryPower(ctx, PowerPrefixCompany+store.CompanyID, PowerPrefixStore+warehouse.StoreId); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta bodega")
	}

	return warehouse, nil
}

func (s *WarehouseService) CreateWarehouse(ctx context.Context, recipe recipe.WarehouseRecipe) (*models.ModelWarehouse, error) {
	if ok := s.EveryPower(ctx, PowerWarehouseCreate); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear bodegas")
	}

	if ok := s.EveryPower(ctx, PowerPrefixCompany+recipe.CompanyId, PowerPrefixStore+recipe.StoreId); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}

	warehouse := &models.ModelWarehouse{
		StoreId:             recipe.StoreId,
		WarehouseName:       recipe.WarehouseName,
		Description:         recipe.Description,
		WarehouseAddress:    recipe.WarehouseAddress,
		IsMomeventWarehouse: false,
	}

	return s.warehouseRepo.CreateWarehouse(warehouse)
}

func (s *WarehouseService) UpdateWarehouse(ctx context.Context, id string, recipe recipe.WarehouseRecipe) (*models.ModelWarehouse, error) {
	if ok := s.EveryPower(ctx, PowerWarehouseUpdate); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar bodegas")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+recipe.CompanyId, PowerPrefixStore+recipe.StoreId); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}

	warehouse := &models.ModelWarehouse{
		ID:                  id,
		StoreId:             recipe.StoreId,
		WarehouseName:       recipe.WarehouseName,
		Description:         recipe.Description,
		WarehouseAddress:    recipe.WarehouseAddress,
		IsMomeventWarehouse: false,
	}

	return s.warehouseRepo.UpdateWarehouse(warehouse)
}

/*
func (s *WarehouseService) DeleteWarehouse(ctx context.Context, id string) error {
	// First get the warehouse to validate ownership
	warehouse, err := s.warehouseRepo.GetWarehouseByID(id)
	if err != nil {
		return err
	}
	if warehouse == nil {
		return types.ThrowMsg("Bodega no encontrada")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+warehouse.CompanyId, PowerPrefixStore+warehouse.StoreId); !ok {
		return types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}
	return s.warehouseRepo.DeleteWarehouse(id)
}
*/
