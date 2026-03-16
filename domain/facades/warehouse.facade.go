package facades

import (
	"context"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
)

type WarehouseFacade struct {
	appServices *services.ServiceContainer
}

func NewWarehouseFacade(appServices *services.ServiceContainer) *WarehouseFacade {
	return &WarehouseFacade{
		appServices: appServices,
	}
}

func (f *WarehouseFacade) GetWarehousesByStoreId(ctx context.Context, id string) ([]models.ModelWarehouse, error) {
	return f.appServices.WarehouseService.GetWarehousesByStoreId(ctx, id)
}

func (f *WarehouseFacade) CreateWarehouse(ctx context.Context, recipe recipe.WarehouseRecipe) (*models.ModelWarehouse, error) {
	return f.appServices.WarehouseService.CreateWarehouse(ctx, recipe)
}

func (f *WarehouseFacade) UpdateWarehouse(ctx context.Context, id string, recipe recipe.WarehouseRecipe) (*models.ModelWarehouse, error) {
	return f.appServices.WarehouseService.UpdateWarehouse(ctx, id, recipe)
}

/*
func (f *WarehouseFacade) DeleteWarehouse(ctx context.Context, id string) error {
	return f.appServices.WarehouseService.DeleteWarehouse(ctx, id)
}
*/
