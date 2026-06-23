package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"sofia-backend/domain/facades"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/shared"

	"github.com/gin-gonic/gin"
)

func TestCompanyControllerUpsertCompanyBranding(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &companyBrandingControllerRepo{}
	companyFacade := facades.NewCompanyFacade(&services.ServiceContainer{
		CompanyBrandingService: services.NewCompanyBrandingService(repo, nil),
	})
	controller := NewCompanyController(companyFacade)

	router := gin.New()
	router.Use(func(ctx *gin.Context) {
		ctx.Set(shared.UserPowersKeys(), []string{services.PowerPrefixCompany + "company-1"})
		ctx.Next()
	})
	controller.RegisterRoutes(&router.RouterGroup)

	body := bytes.NewBufferString(`{
		"appName":"Inventario Cliente",
		"logoUrl":"https://cdn.test/logo.png",
		"faviconUrl":"https://cdn.test/favicon.ico",
		"primaryColor":"#123456",
		"accentColor":"#abcdef",
		"emailFromName":"Inventario Cliente",
		"supportEmail":"soporte@cliente.test",
		"customDomain":"inventario.cliente.test",
		"subdomain":"cliente",
		"welcomeText":"Bienvenido"
	}`)

	request := httptest.NewRequest(http.MethodPut, "/companies/company-1/branding", body)
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("PUT /companies/company-1/branding returned %d with body %s", response.Code, response.Body.String())
	}
	if repo.saved == nil {
		t.Fatal("branding controller did not save branding")
	}
	if repo.saved.CompanyID != "company-1" {
		t.Fatalf("saved company id = %q, want company-1", repo.saved.CompanyID)
	}
	if repo.saved.AppName != "Inventario Cliente" {
		t.Fatalf("saved app name = %q", repo.saved.AppName)
	}
	if repo.saved.LogoURL == nil || *repo.saved.LogoURL != "https://cdn.test/logo.png" {
		t.Fatalf("saved logo url = %+v", repo.saved.LogoURL)
	}
}

type companyBrandingControllerRepo struct {
	saved *models.ModelCompanyBranding
}

func (r *companyBrandingControllerRepo) GetByCompanyID(companyID string) (*models.ModelCompanyBranding, error) {
	return &models.ModelCompanyBranding{
		CompanyID:     companyID,
		AppName:       "Inventario",
		PrimaryColor:  "#2563eb",
		AccentColor:   "#16a34a",
		EmailFromName: "Inventario",
		UpdatedAt:     time.Now(),
	}, nil
}

func (r *companyBrandingControllerRepo) Upsert(branding *models.ModelCompanyBranding) (*models.ModelCompanyBranding, error) {
	branding.UpdatedAt = time.Now()
	r.saved = branding
	return branding, nil
}

func (r *companyBrandingControllerRepo) AddAsset(asset *models.ModelCompanyBrandingAsset) (*models.ModelCompanyBrandingAsset, error) {
	return asset, nil
}

func (r *companyBrandingControllerRepo) GetAssetByID(assetID string) (*models.ModelCompanyBrandingAsset, error) {
	return &models.ModelCompanyBrandingAsset{ID: assetID, CompanyID: "company-1"}, nil
}

func (r *companyBrandingControllerRepo) ListAssets(companyID string) ([]models.ModelCompanyBrandingAsset, error) {
	return nil, nil
}

func (r *companyBrandingControllerRepo) DeleteAsset(assetID string) error {
	return nil
}

func (r *companyBrandingControllerRepo) ListEmailTemplates(companyID string) ([]models.ModelCompanyEmailTemplate, error) {
	return nil, nil
}

func (r *companyBrandingControllerRepo) UpsertEmailTemplate(template *models.ModelCompanyEmailTemplate) (*models.ModelCompanyEmailTemplate, error) {
	return template, nil
}

func (r *companyBrandingControllerRepo) CreateImportJob(job *models.ModelCompanyImportJob) (*models.ModelCompanyImportJob, error) {
	return job, nil
}

func (r *companyBrandingControllerRepo) GetImportJobByID(companyID string, importJobID string) (*models.ModelCompanyImportJob, error) {
	return &models.ModelCompanyImportJob{ID: importJobID, CompanyID: companyID}, nil
}

func (r *companyBrandingControllerRepo) ListImportJobs(companyID string) ([]models.ModelCompanyImportJob, error) {
	return nil, nil
}

func (r *companyBrandingControllerRepo) UpdateImportJobExecution(job *models.ModelCompanyImportJob) (*models.ModelCompanyImportJob, error) {
	return job, nil
}

func (r *companyBrandingControllerRepo) UpsertImportedProducts(rows []models.ModelProductImportItem) (int, []string, error) {
	return len(rows), nil, nil
}

func (r *companyBrandingControllerRepo) UpsertImportedSuppliers(companyID string, rows []models.ModelSupplierImportItem) (int, []string, error) {
	return len(rows), nil, nil
}

func (r *companyBrandingControllerRepo) UpsertImportedInventory(companyID string, rows []models.ModelInventoryImportItem) (int, []string, error) {
	return len(rows), nil, nil
}
