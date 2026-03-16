package ports

// DEPRECATED: PortProductCompany ha sido reemplazado por PortProductPerStore
// Esta interfaz se mantiene comentada para referencia durante la migración.
// Ver: product_per_store.port.go
/*
import "sofia-backend/domain/models"

type PortProductCompany interface {
	GetProductsByCompany(companyID string) ([]models.ModelProductCompany, error)
	// Trae todos los productos independiente de si tienen restriccion o no
	GetAllProductRequestRestrictionByStoreId(storeId string) ([]models.ModelProductRequestRestriction, error)

	// Trae solo los productos que tienen restriccion
	GetProductsRequestRetrictedByStoreId(storeId string) ([]models.ModelProductRequestRestriction, error)
	GetProductCompanyByID(productCompanyID string) (*models.ModelProductCompany, error)
}
*/
