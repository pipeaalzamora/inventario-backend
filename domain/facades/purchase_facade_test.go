package facades

import (
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"sofia-backend/api/v1/recipe"
	"sofia-backend/config"
	"sofia-backend/domain/external"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

func TestPurchaseFacadeCreatePurchaseOrderCreatesRequestTokenAndSupplierEmail(t *testing.T) {
	purchaseRepo := &purchaseFacadePurchaseRepo{purchases: map[string]*models.ModelPurchase{}}
	supplierProductRepo := &purchaseFacadeSupplierProductRepo{
		products: []models.ModelSupplierProductLegacy{
			{
				StoreID:              "store-1",
				SupplierID:           "supplier-1",
				ProductCompanyID:     "store-product-1",
				SupplierProductPrice: 1000,
				UnitPurchase:         "UN",
			},
		},
	}
	supplierRepo := &purchaseFacadeSupplierRepo{email: "compras@proveedor.test"}
	supplierOCRepo := &purchaseFacadeSupplierOCRepo{}
	mailer := &purchaseFacadeMailer{}

	cfg := &config.Config{Debug: true, FrontUrl: "https://erp.test"}
	appServices := &services.ServiceContainer{
		AuthService:            services.NewAuthService(nil, cfg),
		PurchaseService:        services.NewPurchaseService(purchaseRepo),
		SupplierProductService: services.NewSupplierProductService(supplierProductRepo),
		SupplierService:        services.NewSupplierService(supplierRepo),
		SupplierOCService:      services.NewSupplierOCService(supplierOCRepo),
	}
	externalServices := &external.ServiceContainer{
		EmailService: external.NewEmailService(mailer, &purchaseFacadeRender{}, cfg),
	}
	facade := NewPurchaseFacade(appServices, externalServices, cfg)

	ctx := facadeTestContext(
		services.PowerPurchaseCreate,
		services.PowerPrefixCompany+"company-1",
		services.PowerPrefixStore+"store-1",
	)

	created, err := facade.CreatePurchaseOrder(ctx, &recipe.RecipePurchase{
		Description: "Compra semanal",
		SupplierID:  "supplier-1",
		CompanyID:   "company-1",
		StoreID:     "store-1",
		WarehouseID: "warehouse-1",
		Items: []recipe.RecipePurchaseItem{
			{
				StoreProductID: "store-product-1",
				Quantity:       2,
				PurchaseUnit:   "UN",
				UnitPrice:      1000,
			},
		},
	})
	if err != nil {
		t.Fatalf("CreatePurchaseOrder returned error: %v", err)
	}
	if created == nil || created.ID != "purchase-1" {
		t.Fatalf("CreatePurchaseOrder returned unexpected purchase: %+v", created)
	}

	if purchaseRepo.createdRequest == nil {
		t.Fatal("CreatePurchaseOrder did not create the linked inventory request")
	}
	if purchaseRepo.createdRequest.RequesterID != "user-1" {
		t.Fatalf("linked request requester = %q, want user-1", purchaseRepo.createdRequest.RequesterID)
	}
	if purchaseRepo.createdRequest.Status != entities.RequestStatusApproved {
		t.Fatalf("linked request status = %q, want approved", purchaseRepo.createdRequest.Status)
	}

	if supplierProductRepo.gotSupplierID != "supplier-1" || supplierProductRepo.gotStoreID != "store-1" {
		t.Fatalf("supplier products looked up with store=%q supplier=%q", supplierProductRepo.gotStoreID, supplierProductRepo.gotSupplierID)
	}
	if len(supplierProductRepo.gotProductIDs) != 1 || supplierProductRepo.gotProductIDs[0] != "store-product-1" {
		t.Fatalf("supplier products looked up with product ids: %+v", supplierProductRepo.gotProductIDs)
	}

	if supplierOCRepo.created == nil {
		t.Fatal("CreatePurchaseOrder did not create supplier token")
	}
	if supplierOCRepo.created.PurchaseID != "purchase-1" {
		t.Fatalf("supplier token purchase id = %q, want purchase-1", supplierOCRepo.created.PurchaseID)
	}
	if supplierOCRepo.created.Exp == nil || !supplierOCRepo.created.Exp.After(time.Now()) {
		t.Fatalf("supplier token expiration is not in the future: %+v", supplierOCRepo.created.Exp)
	}

	if len(mailer.to) != 1 || mailer.to[0] != "compras@proveedor.test" {
		t.Fatalf("supplier email recipients = %+v", mailer.to)
	}
	if mailer.subject != "Aprobar solicitud de compra" {
		t.Fatalf("supplier email subject = %q", mailer.subject)
	}
	if !strings.Contains(mailer.body, "supplier-manage-oc?token=token-OC-001") {
		t.Fatalf("supplier email body does not include debug token URL: %q", mailer.body)
	}
}

func facadeTestContext(powers ...string) *gin.Context {
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	user := &models.UserAccountModel{
		ID:        "user-1",
		UserName:  "Usuario Test",
		UserEmail: "test@example.com",
	}

	ctx.Set(shared.UserIdKey(), user.ID)
	ctx.Set(shared.UserKey(), user)
	ctx.Set(shared.UserPowersKeys(), powers)

	return ctx
}

type purchaseFacadePurchaseRepo struct {
	createdPurchase *models.ModelPurchase
	createdRequest  *models.ModelInventoryRequest
	purchases       map[string]*models.ModelPurchase
}

func (r *purchaseFacadePurchaseRepo) CreatePurchaseOrder(purchase *models.ModelPurchase) (*models.ModelPurchase, error) {
	r.createdPurchase = purchase
	return r.persistPurchase(purchase), nil
}

func (r *purchaseFacadePurchaseRepo) CreatePurchaseOrderApproved(purchase *models.ModelPurchase) (*models.ModelPurchase, error) {
	r.createdPurchase = purchase
	return r.persistPurchase(purchase), nil
}

func (r *purchaseFacadePurchaseRepo) GetAllPurchase(storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelPurchase, int, error) {
	return nil, 0, nil
}

func (r *purchaseFacadePurchaseRepo) GetPurchaseByID(purchaseID string) (*models.ModelPurchase, error) {
	purchase := r.purchases[purchaseID]
	if purchase == nil {
		return nil, fmt.Errorf("purchase %s not found", purchaseID)
	}
	return purchase, nil
}

func (r *purchaseFacadePurchaseRepo) GetPurchasesByInventoryRequestID(inventoryRequestID string) ([]models.ModelPurchase, error) {
	return nil, nil
}

func (r *purchaseFacadePurchaseRepo) CancelPurchase(purchaseID string, observation string) error {
	return nil
}

func (r *purchaseFacadePurchaseRepo) CreatePurchaseOrderWithInventoryRequest(purchase *models.ModelPurchase, request *models.ModelInventoryRequest) (*models.ModelPurchase, error) {
	r.createdPurchase = purchase
	r.createdRequest = request
	return r.persistPurchase(purchase), nil
}

func (r *purchaseFacadePurchaseRepo) AddDeliveryNoteIdAndSetArrivedStatus(purchaseID string, deliveryNoteID string) error {
	return nil
}

func (r *purchaseFacadePurchaseRepo) AddSonOCToPurchase(purchaseID string, sonDisplayID string) error {
	return nil
}

func (r *purchaseFacadePurchaseRepo) UpdatePurchaseState(purchaseID string, state entities.PurchaseStatus, observation string) error {
	return nil
}

func (r *purchaseFacadePurchaseRepo) UpdatePurchaseItemsStatus(purchaseID string, items []models.ModelPurchaseItem) error {
	return nil
}

func (r *purchaseFacadePurchaseRepo) GetPurchasesByDeliveryPurchaseNote(id string) ([]models.ModelPurchase, error) {
	return nil, nil
}

func (r *purchaseFacadePurchaseRepo) ApprovePurchase(purchaseID string) error {
	return nil
}

func (r *purchaseFacadePurchaseRepo) persistPurchase(purchase *models.ModelPurchase) *models.ModelPurchase {
	if r.purchases == nil {
		r.purchases = map[string]*models.ModelPurchase{}
	}

	created := *purchase
	created.ID = "purchase-1"
	created.DisplayID = "OC-001"
	created.Status = entities.PurchaseStatusPending
	created.CreatedAt = time.Now()
	created.UpdatedAt = created.CreatedAt
	created.Items = make([]models.ModelPurchaseItem, len(purchase.Items))
	for i, item := range purchase.Items {
		created.Items[i] = item
		created.Items[i].ID = fmt.Sprintf("purchase-item-%d", i+1)
		created.Items[i].PurchaseID = created.ID
		created.Items[i].Subtotal = item.Quantity * item.UnitPrice
		created.Items[i].Status = entities.ItemPurchaseStatusPending
	}
	r.purchases[created.ID] = &created
	return &created
}

type purchaseFacadeSupplierProductRepo struct {
	products      []models.ModelSupplierProductLegacy
	gotStoreID    string
	gotSupplierID string
	gotProductIDs []string
}

func (r *purchaseFacadeSupplierProductRepo) GetSupplierProductsByStoreIDAndSupplierIDWithProductCompanyIDs(storeID string, supplierID string, productCompanyIDs []string) ([]models.ModelSupplierProductLegacy, error) {
	r.gotStoreID = storeID
	r.gotSupplierID = supplierID
	r.gotProductIDs = append([]string(nil), productCompanyIDs...)
	return r.products, nil
}

func (r *purchaseFacadeSupplierProductRepo) GetSupplierProductsByStoreIDAndSupplierID(storeID string, supplierID string) ([]models.ModelSupplierProductLegacy, error) {
	return r.products, nil
}

type purchaseFacadeSupplierRepo struct {
	email string
}

func (r *purchaseFacadeSupplierRepo) GetAllSuppliers() ([]models.ModelSupplier, error) {
	return nil, nil
}

func (r *purchaseFacadeSupplierRepo) GetSupplierByID(id string) (*models.ModelSupplier, error) {
	return &models.ModelSupplier{ID: id, FiscalData: models.ModelFiscalData{ID: "fiscal-1"}}, nil
}

func (r *purchaseFacadeSupplierRepo) GetSupplierByFiscalIDAndCountry(fiscalID string, countryID int) (*models.ModelSupplier, error) {
	return nil, nil
}

func (r *purchaseFacadeSupplierRepo) GetCompanyByFiscalNameAndCountry(fiscalName string, countryID int) (*models.ModelSupplier, error) {
	return nil, nil
}

func (r *purchaseFacadeSupplierRepo) CreateSupplier(supplier *models.ModelSupplier) (*models.ModelSupplier, error) {
	return supplier, nil
}

func (r *purchaseFacadeSupplierRepo) UpdateSupplier(supplier *models.ModelSupplier, ogSupplier *models.ModelSupplier) (*models.ModelSupplier, error) {
	return supplier, nil
}

func (r *purchaseFacadeSupplierRepo) DeleteSupplier(id string) error {
	return nil
}

func (r *purchaseFacadeSupplierRepo) GetSupplierProducts(supplierID string) ([]models.ModelSupplierProduct, error) {
	return nil, nil
}

func (r *purchaseFacadeSupplierRepo) GetSupplierProductById(supplierID, productID string) (*models.ModelSupplierProduct, error) {
	return nil, nil
}

func (r *purchaseFacadeSupplierRepo) GetSupplierProductBySku(supplierID, sku string) (*models.ModelSupplierProduct, error) {
	return nil, nil
}

func (r *purchaseFacadeSupplierRepo) AddProductToSupplier(supplierID string, product *models.ModelSupplierProduct) (*models.ModelSupplierProduct, error) {
	return product, nil
}

func (r *purchaseFacadeSupplierRepo) UpdateSupplierProductsPrices(supplierID string, products []models.ModelSupplierProduct) ([]models.ModelSupplierProduct, error) {
	return products, nil
}

func (r *purchaseFacadeSupplierRepo) UpdateSupplierProduct(supplierID string, product *models.ModelSupplierProduct) (*models.ModelSupplierProduct, error) {
	return product, nil
}

func (r *purchaseFacadeSupplierRepo) DeleteSupplierProduct(supplierID, productID string) (*models.ModelSupplierProduct, error) {
	return nil, nil
}

func (r *purchaseFacadeSupplierRepo) GetSuppliersByStoreProductId(storeID string, productIDs []string) ([]models.ModelSupplierStoreProduct, error) {
	return nil, nil
}

func (r *purchaseFacadeSupplierRepo) GetSupplierProductsByProductID(productID string) ([]models.ModelSupplierProduct, error) {
	return nil, nil
}

func (r *purchaseFacadeSupplierRepo) UpsertSupplierProductPerStore(storeID string, suppliers []models.ModelSupplierStoreProduct) error {
	return nil
}

func (r *purchaseFacadeSupplierRepo) DeleteSupplierProductPerStoreByIDs(storeID string, supplierProductIDs []string) error {
	return nil
}

func (r *purchaseFacadeSupplierRepo) EnableDisableSupplier(id string, available bool) error {
	return nil
}

func (r *purchaseFacadeSupplierRepo) EnableDisableSupplierStore(supplierID, storeID string, available bool) error {
	return nil
}

func (r *purchaseFacadeSupplierRepo) GetSupplierEmail(fiscalDataID string) (string, error) {
	return r.email, nil
}

func (r *purchaseFacadeSupplierRepo) ExistsSupplierInCompany(supplierID, companyID string) (bool, error) {
	return true, nil
}

func (r *purchaseFacadeSupplierRepo) GetSuppliersByTemplateProductId(companyId, templateProductId string) ([]models.ModelSupplier, error) {
	return nil, nil
}

type purchaseFacadeSupplierOCRepo struct {
	created *models.ModelSupplierToken
}

func (r *purchaseFacadeSupplierOCRepo) GetSupplierOC(hash string) (*models.ModelSupplierToken, error) {
	return nil, nil
}

func (r *purchaseFacadeSupplierOCRepo) CreateSupplierOC(supplierOC *models.ModelSupplierToken) (*models.ModelSupplierToken, error) {
	r.created = supplierOC
	return supplierOC, nil
}

func (r *purchaseFacadeSupplierOCRepo) UpdateSupplierOC(supplier *models.ModelSupplierToken) (*models.ModelSupplierToken, error) {
	return supplier, nil
}

type purchaseFacadeRender struct{}

func (r *purchaseFacadeRender) Render(filePath string, data any) (string, error) {
	values, _ := data.(map[string]interface{})
	return fmt.Sprintf("%s %v", filePath, values["Url"]), nil
}

type purchaseFacadeMailer struct {
	to      []string
	subject string
	body    string
}

func (m *purchaseFacadeMailer) Send(to []string, subject string, body string) error {
	m.to = append([]string(nil), to...)
	m.subject = subject
	m.body = body
	return nil
}
