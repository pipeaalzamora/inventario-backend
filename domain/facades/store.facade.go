package facades

import (
	"context"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/types"
)

type StoreFacade struct {
	appService *services.ServiceContainer
}

func NewStoreFacade(appService *services.ServiceContainer) *StoreFacade {
	return &StoreFacade{
		appService: appService,
	}
}

func (f *StoreFacade) GetStores(ctx context.Context) ([]models.StoreModel, error) {
	return f.appService.StoreService.GetStores(ctx)
}

func (f *StoreFacade) GetStoresByCompanyID(ctx context.Context, companyId string) ([]models.StoreModel, error) {
	return f.appService.StoreService.GetStoresByCompanyID(ctx, companyId)
}

func (f *StoreFacade) GetStoreByID(ctx context.Context, id string) (*models.StoreModel, error) {
	return f.appService.StoreService.GetStoreByID(ctx, id)
}

func (f *StoreFacade) CreateStore(ctx context.Context, recipe *recipe.StoreRecipe) (*models.StoreModel, error) {
	return f.appService.StoreService.CreateStore(ctx, recipe)
}

func (f *StoreFacade) UpdateStore(ctx context.Context, id string, recipe *recipe.UpdateStoreRecipe) (*models.StoreModel, error) {
	return f.appService.StoreService.UpdateStore(ctx, id, recipe)
}

// Deprecated: use CompanyFacade.AssignSuppliersToCompany en su lugar.
func (f *StoreFacade) UpdateStoreSuppliers(ctx context.Context, id string, recipe *recipe.UpdateStoreSuppliersRecipe) (*models.StoreModel, error) {
	return nil, types.ThrowMsg("Endpoint deprecado. Use POST /companies/:id/suppliers")
}
