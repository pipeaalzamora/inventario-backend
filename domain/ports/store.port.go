package ports

import "sofia-backend/domain/models"

type PortStore interface {
	GetStores() ([]models.StoreModel, error)
	GetStoreByID(id string) (*models.StoreModel, error)
	GetStoreByCompanyID(companyID string) (*models.StoreModel, error)
	CreateStore(
		store *models.StoreModel,
		profiles []models.ProfileAccountModel,
	) (*models.StoreModel, error)
	UpdateStore(id string, store *models.StoreModel) (*models.StoreModel, error)
	// Deprecated: use company-level supplier assignment endpoints.
	UpdateStoreSuppliers(storeID string, supplierIDs []string) error
	//DeleteStore(id string) error FALTA CREAR
	GetStoresByCompanyID(companyID string) ([]models.StoreModel, error)
}
