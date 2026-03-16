package ports

import "sofia-backend/domain/models"

type PortSupplierProduct interface {
	GetSupplierProductsByStoreIDAndSupplierIDWithProductCompanyIDs(storeID string, supplierID string, productCompanyIDs []string) ([]models.ModelSupplierProductLegacy, error)
	GetSupplierProductsByStoreIDAndSupplierID(storeID string, supplierID string) ([]models.ModelSupplierProductLegacy, error)
}
