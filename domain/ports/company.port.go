package ports

import "sofia-backend/domain/models"

type PortCompany interface {
	GetCompanies() ([]models.ModelCompany, error)
	GetCompanyByID(id string) (*models.ModelCompany, error)
	GetCompanyByFiscalIDAndCountry(fiscalID string, countryID int) (*models.ModelCompany, error)
	GetCompanyByFiscalNameAndCountry(fiscalName string, countryID int) (*models.ModelCompany, error)

	GetSuppliersByCompanyID(companyID string) ([]models.CompanySupplierModel, error)
	AssignSuppliersToCompany(companyID string, supplierIDs []string) error
	UnassignSupplierFromCompany(companyID, supplierID string) error

	CreateCompany(company *models.ModelCompany, profiles []models.ProfileAccountModel) (*models.ModelCompany, error)
	UpdateCompany(id string, company *models.ModelCompany) (*models.ModelCompany, error)

	AddLogoToCompany(companyID string, url string) error
	RemoveLogoFromCompany(fileID string) error
}
