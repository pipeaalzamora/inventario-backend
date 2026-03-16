package ports

import (
	"sofia-backend/domain/models"
)

// PortProductPerStore define las operaciones para gestionar productos por tienda.
// Reemplaza a PortProductCompany, ahora con granularidad a nivel de tienda.
type PortProductPerStore interface {
	// GetProductsByStore obtiene todos los productos configurados para una tienda
	GetProductsByStore(storeID string) ([]models.ModelProductPerStore, error)

	// GetProductPerStoreByID obtiene un producto específico por su ID
	GetProductPerStoreByID(storeProductID string) (*models.ModelProductPerStore, error)

	// GetProductsWithSuppliersByStore obtiene todos los productos de una tienda que tienen proveedores asignados
	GetProductsWithSuppliersByStore(storeID string) ([]models.ModelProductPerStore, error)

	// GetAllProductRequestRestrictionByStoreId obtiene todos los productos con información de restricción
	// Trae todos los productos independiente de si tienen restriccion o no
	GetAllProductRequestRestrictionByStoreId(storeId string) ([]models.ModelProductRequestRestriction, error)

	// GetProductsRequestRetrictedByStoreId obtiene solo los productos que tienen restricción
	// Trae solo los productos que tienen max_quantity definido
	GetProductsRequestRetrictedByStoreId(storeId string) ([]models.ModelProductRequestRestriction, error)

	// CreateProductPerStore crea un nuevo producto para una tienda
	CreateProductPerStore(product *models.ModelProductPerStore) (*models.ModelProductPerStore, error)

	// UpdateProductPerStore actualiza un producto existente
	UpdateProductPerStore(product *models.ModelProductPerStore) (*models.ModelProductPerStore, error)

	// DeleteProductPerStore elimina un producto de una tienda
	DeleteProductPerStore(storeProductID string) error

	// ExistsBySKU verifica si existe un producto con el SKU dado en la tienda
	ExistsBySKU(storeID, sku string) (bool, error)

	// ExistsByProductAndStore verifica si existe un producto plantilla ya asignado a la tienda
	ExistsByProductAndStore(productID, storeID string, excludeID *string) (bool, error)

	// GetProductsByStoreAndProductIDs obtiene productos por tienda y lista de IDs de producto base
	GetProductsByStoreAndProductIDs(storeID string, productIDs []string) ([]models.ModelProductPerStore, error)
}
