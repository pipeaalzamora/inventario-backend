package ports

import (
	"sofia-backend/domain/models"
)

type PortSupplierOC interface {
	GetSupplierOC(hash string) (*models.ModelSupplierToken, error)
	CreateSupplierOC(supplierOC *models.ModelSupplierToken) (*models.ModelSupplierToken, error)
	UpdateSupplierOC(supplier *models.ModelSupplierToken) (*models.ModelSupplierToken, error)
}
