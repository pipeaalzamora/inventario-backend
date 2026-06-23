package services

import (
	"testing"

	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
)

type fakeRequestRepo struct {
	created *models.ModelRequest
}

func (r *fakeRequestRepo) CreateRequest(request *models.ModelRequest) (*models.ModelRequest, error) {
	r.created = request
	return request, nil
}

type fakeStoreRepo struct {
	stores              map[string]*models.StoreModel
	getByCompanyIDCalls int
}

func (r *fakeStoreRepo) GetStores() ([]models.StoreModel, error) {
	return nil, nil
}

func (r *fakeStoreRepo) GetStoreByID(id string) (*models.StoreModel, error) {
	return r.stores[id], nil
}

func (r *fakeStoreRepo) GetStoreByCompanyID(companyID string) (*models.StoreModel, error) {
	r.getByCompanyIDCalls++
	return nil, nil
}

func (r *fakeStoreRepo) CreateStore(store *models.StoreModel, profiles []models.ProfileAccountModel) (*models.StoreModel, error) {
	return store, nil
}

func (r *fakeStoreRepo) UpdateStore(id string, store *models.StoreModel) (*models.StoreModel, error) {
	return store, nil
}

func (r *fakeStoreRepo) UpdateStoreSuppliers(storeID string, supplierIDs []string) error {
	return nil
}

func (r *fakeStoreRepo) GetStoresByCompanyID(companyID string) ([]models.StoreModel, error) {
	return nil, nil
}

type fakeCompanyRepo struct {
	companies map[string]*models.ModelCompany
}

func (r *fakeCompanyRepo) GetCompanies() ([]models.ModelCompany, error) {
	return nil, nil
}

func (r *fakeCompanyRepo) GetCompanyByID(id string) (*models.ModelCompany, error) {
	return r.companies[id], nil
}

func (r *fakeCompanyRepo) GetCompanyByFiscalIDAndCountry(fiscalID string, countryID int) (*models.ModelCompany, error) {
	return nil, nil
}

func (r *fakeCompanyRepo) GetCompanyByFiscalNameAndCountry(fiscalName string, countryID int) (*models.ModelCompany, error) {
	return nil, nil
}

func (r *fakeCompanyRepo) GetSuppliersByCompanyID(companyID string) ([]models.CompanySupplierModel, error) {
	return nil, nil
}

func (r *fakeCompanyRepo) AssignSuppliersToCompany(companyID string, supplierIDs []string) error {
	return nil
}

func (r *fakeCompanyRepo) UnassignSupplierFromCompany(companyID, supplierID string) error {
	return nil
}

func (r *fakeCompanyRepo) CreateCompany(company *models.ModelCompany, profiles []models.ProfileAccountModel) (*models.ModelCompany, error) {
	return company, nil
}

func (r *fakeCompanyRepo) UpdateCompany(id string, company *models.ModelCompany) (*models.ModelCompany, error) {
	return company, nil
}

func (r *fakeCompanyRepo) AddLogoToCompany(companyID string, url string) error {
	return nil
}

func (r *fakeCompanyRepo) RemoveLogoFromCompany(fileID string) error {
	return nil
}

func TestRequestServiceCreateRequestUsesRequestedStore(t *testing.T) {
	const companyID = "company-1"
	const storeID = "store-2"

	requestRepo := &fakeRequestRepo{}
	storeRepo := &fakeStoreRepo{
		stores: map[string]*models.StoreModel{
			storeID: {ID: storeID, CompanyID: companyID},
		},
	}
	companyRepo := &fakeCompanyRepo{
		companies: map[string]*models.ModelCompany{
			companyID: {ID: companyID},
		},
	}
	service := NewRequestService(requestRepo, storeRepo, companyRepo)
	ctx := serviceTestContext(PowerRequestCreateForStore, PowerPrefixCompany+companyID, PowerPrefixStore+storeID)

	created, err := service.CreateRequest(ctx, &recipe.RecipeNewRequest{
		CompanyID:   companyID,
		StoreID:     storeID,
		RequestType: models.RequestKindStore.ToString(),
	})
	if err != nil {
		t.Fatalf("CreateRequest returned error: %v", err)
	}
	if created == nil {
		t.Fatal("CreateRequest returned nil request")
	}
	if requestRepo.created == nil {
		t.Fatal("CreateRequest did not persist the request")
	}
	if requestRepo.created.StoreId != storeID {
		t.Fatalf("CreateRequest used store %q, want %q", requestRepo.created.StoreId, storeID)
	}
	if storeRepo.getByCompanyIDCalls != 0 {
		t.Fatalf("CreateRequest called GetStoreByCompanyID %d times, want 0", storeRepo.getByCompanyIDCalls)
	}
}

func TestRequestServiceCreateRequestRejectsStoreFromAnotherCompany(t *testing.T) {
	const companyID = "company-1"
	const storeID = "store-2"

	requestRepo := &fakeRequestRepo{}
	storeRepo := &fakeStoreRepo{
		stores: map[string]*models.StoreModel{
			storeID: {ID: storeID, CompanyID: "company-2"},
		},
	}
	companyRepo := &fakeCompanyRepo{
		companies: map[string]*models.ModelCompany{
			companyID: {ID: companyID},
		},
	}
	service := NewRequestService(requestRepo, storeRepo, companyRepo)
	ctx := serviceTestContext(PowerRequestCreateForStore, PowerPrefixCompany+companyID, PowerPrefixStore+storeID)

	created, err := service.CreateRequest(ctx, &recipe.RecipeNewRequest{
		CompanyID:   companyID,
		StoreID:     storeID,
		RequestType: models.RequestKindStore.ToString(),
	})
	if err == nil {
		t.Fatal("CreateRequest returned nil error for a store from another company")
	}
	if created != nil {
		t.Fatalf("CreateRequest returned request for invalid store: %+v", created)
	}
	if requestRepo.created != nil {
		t.Fatalf("CreateRequest persisted request for invalid store: %+v", requestRepo.created)
	}
}
