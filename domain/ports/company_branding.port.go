package ports

import "sofia-backend/domain/models"

type PortCompanyBranding interface {
	GetByCompanyID(companyID string) (*models.ModelCompanyBranding, error)
	Upsert(branding *models.ModelCompanyBranding) (*models.ModelCompanyBranding, error)
	AddAsset(asset *models.ModelCompanyBrandingAsset) (*models.ModelCompanyBrandingAsset, error)
	GetAssetByID(assetID string) (*models.ModelCompanyBrandingAsset, error)
	ListAssets(companyID string) ([]models.ModelCompanyBrandingAsset, error)
	DeleteAsset(assetID string) error
	ListEmailTemplates(companyID string) ([]models.ModelCompanyEmailTemplate, error)
	UpsertEmailTemplate(template *models.ModelCompanyEmailTemplate) (*models.ModelCompanyEmailTemplate, error)
	CreateImportJob(job *models.ModelCompanyImportJob) (*models.ModelCompanyImportJob, error)
	GetImportJobByID(companyID string, importJobID string) (*models.ModelCompanyImportJob, error)
	ListImportJobs(companyID string) ([]models.ModelCompanyImportJob, error)
	UpdateImportJobExecution(job *models.ModelCompanyImportJob) (*models.ModelCompanyImportJob, error)
	UpsertImportedProducts(rows []models.ModelProductImportItem) (int, []string, error)
	UpsertImportedSuppliers(companyID string, rows []models.ModelSupplierImportItem) (int, []string, error)
	UpsertImportedInventory(companyID string, rows []models.ModelInventoryImportItem) (int, []string, error)
}
