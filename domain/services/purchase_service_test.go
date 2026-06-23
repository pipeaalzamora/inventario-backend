package services

import (
	"testing"

	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/infraestructure/entities"
)

type fakePurchaseRepo struct {
	created *models.ModelPurchase
}

func (r *fakePurchaseRepo) CreatePurchaseOrder(purchase *models.ModelPurchase) (*models.ModelPurchase, error) {
	r.created = purchase
	return purchase, nil
}

func (r *fakePurchaseRepo) CreatePurchaseOrderApproved(purchase *models.ModelPurchase) (*models.ModelPurchase, error) {
	r.created = purchase
	return purchase, nil
}

func (r *fakePurchaseRepo) GetAllPurchase(storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelPurchase, int, error) {
	return nil, 0, nil
}

func (r *fakePurchaseRepo) GetPurchaseByID(purchaseID string) (*models.ModelPurchase, error) {
	return &models.ModelPurchase{ID: purchaseID, CompanyID: "company-1", StoreID: "store-1"}, nil
}

func (r *fakePurchaseRepo) GetPurchasesByInventoryRequestID(inventoryRequestID string) ([]models.ModelPurchase, error) {
	return nil, nil
}

func (r *fakePurchaseRepo) CancelPurchase(purchaseID string, observation string) error {
	return nil
}

func (r *fakePurchaseRepo) CreatePurchaseOrderWithInventoryRequest(purchase *models.ModelPurchase, request *models.ModelInventoryRequest) (*models.ModelPurchase, error) {
	r.created = purchase
	return purchase, nil
}

func (r *fakePurchaseRepo) AddDeliveryNoteIdAndSetArrivedStatus(purchaseID string, deliveryNoteID string) error {
	return nil
}

func (r *fakePurchaseRepo) AddSonOCToPurchase(purchaseID string, sonDisplayID string) error {
	return nil
}

func (r *fakePurchaseRepo) UpdatePurchaseState(purchaseID string, state entities.PurchaseStatus, observation string) error {
	return nil
}

func (r *fakePurchaseRepo) UpdatePurchaseItemsStatus(purchaseID string, items []models.ModelPurchaseItem) error {
	return nil
}

func (r *fakePurchaseRepo) GetPurchasesByDeliveryPurchaseNote(id string) ([]models.ModelPurchase, error) {
	return nil, nil
}

func (r *fakePurchaseRepo) ApprovePurchase(purchaseID string) error {
	return nil
}

func TestPurchaseServiceCreatePurchaseRequiresCreatePower(t *testing.T) {
	repo := &fakePurchaseRepo{}
	service := NewPurchaseService(repo)
	ctx := serviceTestContext(PowerPrefixCompany+"company-1", PowerPrefixStore+"store-1")

	created, err := service.CreatePurchaseOrder(ctx, validPurchaseRecipe())
	if err == nil {
		t.Fatal("CreatePurchaseOrder returned nil error without purchase:create power")
	}
	if created != nil {
		t.Fatalf("CreatePurchaseOrder returned purchase without permission: %+v", created)
	}
	if repo.created != nil {
		t.Fatalf("CreatePurchaseOrder persisted purchase without permission: %+v", repo.created)
	}
}

func TestPurchaseServiceCreatePurchaseRequiresStoreOwnership(t *testing.T) {
	repo := &fakePurchaseRepo{}
	service := NewPurchaseService(repo)
	ctx := serviceTestContext(PowerPurchaseCreate, PowerPrefixCompany+"company-1", PowerPrefixStore+"store-2")

	created, err := service.CreatePurchaseOrder(ctx, validPurchaseRecipe())
	if err == nil {
		t.Fatal("CreatePurchaseOrder returned nil error without target store ownership")
	}
	if created != nil {
		t.Fatalf("CreatePurchaseOrder returned purchase without store ownership: %+v", created)
	}
	if repo.created != nil {
		t.Fatalf("CreatePurchaseOrder persisted purchase without store ownership: %+v", repo.created)
	}
}

func TestPurchaseServiceCreatePurchaseAllowsValidPowers(t *testing.T) {
	repo := &fakePurchaseRepo{}
	service := NewPurchaseService(repo)
	ctx := serviceTestContext(PowerPurchaseCreate, PowerPrefixCompany+"company-1", PowerPrefixStore+"store-1")

	created, err := service.CreatePurchaseOrder(ctx, validPurchaseRecipe())
	if err != nil {
		t.Fatalf("CreatePurchaseOrder returned error: %v", err)
	}
	if created == nil {
		t.Fatal("CreatePurchaseOrder returned nil purchase")
	}
	if repo.created == nil {
		t.Fatal("CreatePurchaseOrder did not persist purchase")
	}
}

func validPurchaseRecipe() *recipe.RecipePurchase {
	return &recipe.RecipePurchase{
		SupplierID:         "supplier-1",
		CompanyID:          "company-1",
		StoreID:            "store-1",
		WarehouseID:        "warehouse-1",
		InventoryRequestID: "request-1",
		Items: []recipe.RecipePurchaseItem{
			{StoreProductID: "store-product-1", Quantity: 2, UnitPrice: 1000},
		},
	}
}
