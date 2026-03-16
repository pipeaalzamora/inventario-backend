package services

import (
	"context"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
	"strings"
	"time"
)

type StoreService struct {
	PowerChecker
	storeRepo    ports.PortStore
	profileRepo  ports.PortProfile
	userRepo     ports.PortUser
	cacheService ports.PortCache
	powerRepo    ports.PortPower
}

func NewStoreService(
	storeRepo ports.PortStore,
	profileRepo ports.PortProfile,
	userRepo ports.PortUser,
	cacheService ports.PortCache,
	powerRepo ports.PortPower,

) *StoreService {
	return &StoreService{
		storeRepo:    storeRepo,
		profileRepo:  profileRepo,
		userRepo:     userRepo,
		cacheService: cacheService,
		powerRepo:    powerRepo,
	}
}

func (s *StoreService) GetStores(ctx context.Context) ([]models.StoreModel, error) {
	// get ownable powers
	powers, ok := s.GetPowersFromContext(ctx)
	if !ok {
		return nil, nil
	}

	// if the user has the "store:" prefix power, return all stores
	filteredStoreID := map[string]bool{}
	for _, p := range powers {
		if strings.HasPrefix(p, PowerPrefixStore) {
			storeId := p[len(PowerPrefixStore):]
			filteredStoreID[storeId] = true
		}
	}

	allStores, err := s.storeRepo.GetStores()
	if err != nil {
		return nil, err
	}

	result := []models.StoreModel{}
	for _, store := range allStores {
		if filteredStoreID[store.ID] {
			result = append(result, store)
		}
	}

	return result, nil
}

func (s *StoreService) GetStoresByCompanyID(ctx context.Context, companyId string) ([]models.StoreModel, error) {
	// TODO: Descomentar cuando los permisos estén listos

	// check if user has access to the company
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyId); !ok {
		return make([]models.StoreModel, 0), nil
	}

	// get ownable powers
	powers, ok := s.GetPowersFromContext(ctx)
	if !ok {
		return make([]models.StoreModel, 0), nil
	}

	// if the user has the "store:" prefix power, return all stores
	filteredStoreID := map[string]bool{}
	for _, p := range powers {
		if strings.HasPrefix(p, "store:") {
			storeId := p[len("store:"):]
			filteredStoreID[storeId] = true
		}
	}

	allStores, err := s.storeRepo.GetStoresByCompanyID(companyId)
	if err != nil {
		return nil, err
	}

	result := []models.StoreModel{}
	for _, store := range allStores {
		if filteredStoreID[store.ID] {
			result = append(result, store)
		}
	}

	// FALTAN LOS PERMISOS PARA PODER FILTRAR
	allStores, err = s.storeRepo.GetStoresByCompanyID(companyId)
	if err != nil {
		return nil, err
	}

	return allStores, nil
}

func (s *StoreService) GetStoreByID(ctx context.Context, id string) (*models.StoreModel, error) {

	powers, ok := s.GetPowersFromContext(ctx)
	if !ok {
		return nil, nil
	}

	filteredStoreID := map[string]bool{}
	for _, p := range powers {
		if strings.HasPrefix(p, PowerPrefixStore) {
			storeId := p[len(PowerPrefixStore):]
			filteredStoreID[storeId] = true
		}
	}

	if !filteredStoreID[id] {
		return nil, nil
	}

	return s.storeRepo.GetStoreByID(id)
}

func (s *StoreService) CreateStore(ctx context.Context, recipe *recipe.StoreRecipe) (*models.StoreModel, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+recipe.CompanyId); !ok {
		return nil, types.ThrowPower("No tienes permiso para trabajar en esta empresa")
	}
	if ok := s.EveryPower(ctx, PowerStoreCreate); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear tiendas")
	}

	userContext, ok := s.GetUserFromContext(ctx)
	if !ok {
		return nil, types.ThrowData("No se pudo obtener el usuario del contexto")
	}

	myProfiles, err := s.profileRepo.GetProfilesByUserID(userContext.ID)
	if err != nil {
		return nil, types.ThrowData("Ocurrió un error al extraer los perfiles del usuario")
	}

	profilesWithCreatePermission := make([]models.ProfileAccountModel, 0)
	for _, profile := range myProfiles {
		powers, err := s.powerRepo.GetPowersByProfile(profile.ID)
		if err != nil {
			continue
		}
		for _, p := range powers {
			if p.PowerName == PowerStoreCreate {
				profilesWithCreatePermission = append(profilesWithCreatePermission, profile)
				break
			}
		}
	}

	newStore := &models.StoreModel{
		ID:           "",
		CompanyID:    recipe.CompanyId,
		StoreName:    recipe.Name,
		Description:  recipe.Description,
		StoreAddress: recipe.Address,
		IDCostCenter: recipe.CostCenter,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		WareHouses:   nil,
		//SupplierApplied: nil,
	}

	createdStore, err := s.storeRepo.CreateStore(newStore, profilesWithCreatePermission)
	if err != nil {
		return nil, err
	}

	s.cacheService.DeleteByKey("POWERS:" + userContext.ID)

	return createdStore, nil
}

func (s *StoreService) UpdateStore(ctx context.Context, id string, recipe *recipe.UpdateStoreRecipe) (*models.StoreModel, error) {
	if ok := s.EveryPower(ctx, PowerStoreUpdate); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar tiendas")
	}
	if ok := s.EveryPower(ctx, PowerPrefixStore+id); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar esta tienda")
	}
	currentStore, err := s.storeRepo.GetStoreByID(id)
	if err != nil || currentStore == nil {
		return nil, types.ThrowData("No se pudo obtener la tienda")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+currentStore.CompanyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta compañía")
	}

	currentStore.StoreName = recipe.Name
	currentStore.Description = recipe.Description
	currentStore.StoreAddress = recipe.Address
	currentStore.IDCostCenter = recipe.CostCenter
	currentStore.UpdatedAt = time.Now()

	updatedStore, err := s.storeRepo.UpdateStore(id, currentStore)
	if err != nil {
		return nil, err
	}

	return updatedStore, nil
}

// Deprecated: use CompanyService.AssignSuppliersToCompany en su lugar.
func (s *StoreService) UpdateStoreSuppliers(ctx context.Context, id string, recipe *recipe.UpdateStoreSuppliersRecipe) (*models.StoreModel, error) {
	return nil, types.ThrowMsg("Endpoint deprecado. Use POST /companies/:id/suppliers")
}
