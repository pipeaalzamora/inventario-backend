package services

import (
	"testing"

	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
)

type fakeCompanyBrandingRepo struct {
	saved *models.ModelCompanyBranding
}

func (r *fakeCompanyBrandingRepo) GetByCompanyID(companyID string) (*models.ModelCompanyBranding, error) {
	return &models.ModelCompanyBranding{
		CompanyID:     companyID,
		AppName:       "Marca Test",
		PrimaryColor:  "#2563eb",
		AccentColor:   "#16a34a",
		EmailFromName: "Marca Test",
	}, nil
}

func (r *fakeCompanyBrandingRepo) Upsert(branding *models.ModelCompanyBranding) (*models.ModelCompanyBranding, error) {
	r.saved = branding
	return branding, nil
}

func (r *fakeCompanyBrandingRepo) AddAsset(asset *models.ModelCompanyBrandingAsset) (*models.ModelCompanyBrandingAsset, error) {
	return asset, nil
}

func (r *fakeCompanyBrandingRepo) GetAssetByID(assetID string) (*models.ModelCompanyBrandingAsset, error) {
	return &models.ModelCompanyBrandingAsset{ID: assetID, CompanyID: "company-1"}, nil
}

func (r *fakeCompanyBrandingRepo) ListAssets(companyID string) ([]models.ModelCompanyBrandingAsset, error) {
	return nil, nil
}

func (r *fakeCompanyBrandingRepo) DeleteAsset(assetID string) error {
	return nil
}

func (r *fakeCompanyBrandingRepo) ListEmailTemplates(companyID string) ([]models.ModelCompanyEmailTemplate, error) {
	return nil, nil
}

func (r *fakeCompanyBrandingRepo) UpsertEmailTemplate(template *models.ModelCompanyEmailTemplate) (*models.ModelCompanyEmailTemplate, error) {
	return template, nil
}

func (r *fakeCompanyBrandingRepo) CreateImportJob(job *models.ModelCompanyImportJob) (*models.ModelCompanyImportJob, error) {
	return job, nil
}

func (r *fakeCompanyBrandingRepo) GetImportJobByID(companyID string, importJobID string) (*models.ModelCompanyImportJob, error) {
	return &models.ModelCompanyImportJob{ID: importJobID, CompanyID: companyID}, nil
}

func (r *fakeCompanyBrandingRepo) ListImportJobs(companyID string) ([]models.ModelCompanyImportJob, error) {
	return nil, nil
}

func (r *fakeCompanyBrandingRepo) UpdateImportJobExecution(job *models.ModelCompanyImportJob) (*models.ModelCompanyImportJob, error) {
	return job, nil
}

func (r *fakeCompanyBrandingRepo) UpsertImportedProducts(rows []models.ModelProductImportItem) (int, []string, error) {
	return len(rows), nil, nil
}

func (r *fakeCompanyBrandingRepo) UpsertImportedSuppliers(companyID string, rows []models.ModelSupplierImportItem) (int, []string, error) {
	return len(rows), nil, nil
}

func (r *fakeCompanyBrandingRepo) UpsertImportedInventory(companyID string, rows []models.ModelInventoryImportItem) (int, []string, error) {
	return len(rows), nil, nil
}

func TestCompanyBrandingServiceUpsertRequiresCompanyPower(t *testing.T) {
	repo := &fakeCompanyBrandingRepo{}
	service := NewCompanyBrandingService(repo, nil)
	ctx := serviceTestContext(PowerPrefixCompany + "other-company")

	created, err := service.Upsert(ctx, "company-1", validCompanyBrandingRecipe())
	if err == nil {
		t.Fatal("Upsert returned nil error without company ownership power")
	}
	if created != nil {
		t.Fatalf("Upsert returned branding without permission: %+v", created)
	}
	if repo.saved != nil {
		t.Fatalf("Upsert persisted branding without permission: %+v", repo.saved)
	}
}

func TestCompanyBrandingServiceUpsertValidatesHexColors(t *testing.T) {
	repo := &fakeCompanyBrandingRepo{}
	service := NewCompanyBrandingService(repo, nil)
	ctx := serviceTestContext(PowerPrefixCompany + "company-1")
	input := validCompanyBrandingRecipe()
	input.PrimaryColor = "blue"

	created, err := service.Upsert(ctx, "company-1", input)
	if err == nil {
		t.Fatal("Upsert returned nil error for invalid primary color")
	}
	if created != nil {
		t.Fatalf("Upsert returned branding for invalid color: %+v", created)
	}
	if repo.saved != nil {
		t.Fatalf("Upsert persisted branding for invalid color: %+v", repo.saved)
	}
}

func TestCompanyBrandingServiceUpsertSavesValidBranding(t *testing.T) {
	repo := &fakeCompanyBrandingRepo{}
	service := NewCompanyBrandingService(repo, nil)
	ctx := serviceTestContext(PowerPrefixCompany + "company-1")

	created, err := service.Upsert(ctx, "company-1", validCompanyBrandingRecipe())
	if err != nil {
		t.Fatalf("Upsert returned error: %v", err)
	}
	if created == nil {
		t.Fatal("Upsert returned nil branding")
	}
	if repo.saved == nil {
		t.Fatal("Upsert did not persist branding")
	}
	if repo.saved.PrimaryColor != "#123456" {
		t.Fatalf("PrimaryColor = %q, want #123456", repo.saved.PrimaryColor)
	}
}

func validCompanyBrandingRecipe() *recipe.CompanyBrandingRecipe {
	return &recipe.CompanyBrandingRecipe{
		AppName:       "Mi ERP",
		PrimaryColor:  "#123456",
		AccentColor:   "#abcdef",
		EmailFromName: "Mi ERP",
	}
}
