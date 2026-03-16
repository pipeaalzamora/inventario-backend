package ports

import (
	"sofia-backend/domain/models"
)

type PortSupplier interface {
	/// Supplier CRUD (taxMapia) ///
	GetAllSuppliers() ([]models.ModelSupplier, error)
	GetSupplierByID(id string) (*models.ModelSupplier, error)

	GetSupplierByFiscalIDAndCountry(fiscalID string, countryID int) (*models.ModelSupplier, error)
	GetCompanyByFiscalNameAndCountry(fiscalName string, countryID int) (*models.ModelSupplier, error)

	CreateSupplier(supplier *models.ModelSupplier) (*models.ModelSupplier, error)
	UpdateSupplier(supplier *models.ModelSupplier, ogSupplier *models.ModelSupplier) (*models.ModelSupplier, error)
	DeleteSupplier(id string) error

	/// Supplier Products CRUD (taxMapia) ///
	GetSupplierProducts(supplierID string) ([]models.ModelSupplierProduct, error)
	GetSupplierProductById(supplierID, productID string) (*models.ModelSupplierProduct, error)
	GetSupplierProductBySku(supplierID, sku string) (*models.ModelSupplierProduct, error)

	AddProductToSupplier(supplierID string, product *models.ModelSupplierProduct) (*models.ModelSupplierProduct, error)
	UpdateSupplierProductsPrices(supplierID string, products []models.ModelSupplierProduct) ([]models.ModelSupplierProduct, error)
	UpdateSupplierProduct(supplierID string, product *models.ModelSupplierProduct) (*models.ModelSupplierProduct, error)
	DeleteSupplierProduct(supplierID, productID string) (*models.ModelSupplierProduct, error)

	/// Supplier Product Per Store (Store Product Suppliers) ///
	GetSuppliersByStoreProductId(storeID string, productIDs []string) ([]models.ModelSupplierStoreProduct, error)
	GetSupplierProductsByProductID(productID string) ([]models.ModelSupplierProduct, error)
	UpsertSupplierProductPerStore(storeID string, suppliers []models.ModelSupplierStoreProduct) error
	DeleteSupplierProductPerStoreByIDs(storeID string, supplierProductIDs []string) error

	/// Other Methods ///
	EnableDisableSupplier(id string, available bool) error
	// Deprecated: use company-level supplier assignment.
	EnableDisableSupplierStore(supplierID, storeID string, available bool) error
	GetSupplierEmail(fiscalDataID string) (string, error)
	ExistsSupplierInCompany(supplierID, companyID string) (bool, error)

	GetSuppliersByTemplateProductId(companyId, templateProductId string) ([]models.ModelSupplier, error)
}
