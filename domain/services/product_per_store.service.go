package services

import (
	"context"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
)

// Constantes de permisos para ProductPerStore
const (
	POWER_STORE_PRODUCT_CREATE = "store_product:create"
	POWER_STORE_PRODUCT_UPDATE = "store_product:update"
	POWER_STORE_PRODUCT_DELETE = "store_product:delete"
)

// ProductPerStoreService maneja la lógica de negocio para productos por tienda.
// Reemplaza a ProductCompanyService, ahora con granularidad a nivel de tienda.
type ProductPerStoreService struct {
	PowerChecker
	ProductPerStoreRepo ports.PortProductPerStore
	storeRepo           ports.PortStore
	PriceHistoryRepo    ports.PortPriceHistory
}

// NewProductPerStoreService crea una nueva instancia del servicio.
func NewProductPerStoreService(
	repo ports.PortProductPerStore,
	storeRepo ports.PortStore,
	priceHistoryRepo ports.PortPriceHistory,
) *ProductPerStoreService {
	return &ProductPerStoreService{
		ProductPerStoreRepo: repo,
		storeRepo:           storeRepo,
		PriceHistoryRepo:    priceHistoryRepo,
	}
}

// GetProductsByStore obtiene todos los productos configurados para una tienda.
// Requiere permiso de acceso a la tienda.
func (s *ProductPerStoreService) GetProductsByStore(ctx context.Context, storeID string) ([]models.ModelProductPerStore, error) {

	if ok := s.EveryPower(ctx, PowerPrefixStore+storeID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.ProductPerStoreRepo.GetProductsByStore(storeID)
}

// GetProductPerStoreByID obtiene un producto específico por su ID.
// Valida el permiso de acceso a la tienda del producto.
func (s *ProductPerStoreService) GetProductPerStoreByID(ctx context.Context, storeProductID string) (*models.ModelProductPerStore, error) {
	product, err := s.ProductPerStoreRepo.GetProductPerStoreByID(storeProductID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil
	}

	// Validar acceso a la tienda
	if ok := s.EveryPower(ctx, PowerPrefixStore+product.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}

	return product, nil
}

func (s *ProductPerStoreService) GetProductsWithSuppliersByStore(ctx context.Context, storeID string) ([]models.ModelProductPerStore, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.ProductPerStoreRepo.GetProductsWithSuppliersByStore(storeID)
}

// GetProductsRequestRestriction obtiene todos los productos con información de restricción.
// Requiere permiso de acceso a la tienda.
func (s *ProductPerStoreService) GetProductsRequestRestriction(ctx context.Context, storeId string) ([]models.ModelProductRequestRestriction, int, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeId); !ok {
		return nil, 0, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	products, err := s.ProductPerStoreRepo.GetAllProductRequestRestrictionByStoreId(storeId)
	if err != nil {
		return nil, 0, err
	}
	return products, len(products), nil
}

// GetOnlyRestrictedProducts obtiene solo los productos que tienen restricción.
// Requiere permiso de acceso a la tienda.
func (s *ProductPerStoreService) GetOnlyRestrictedProducts(ctx context.Context, storeId string) ([]models.ModelProductRequestRestriction, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeId); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.ProductPerStoreRepo.GetProductsRequestRetrictedByStoreId(storeId)
}

// CreateStoreProduct crea un nuevo producto para una tienda validando que store pertenece a company.
// Requiere permiso de creación y acceso a la tienda.
func (s *ProductPerStoreService) CreateStoreProduct(ctx context.Context, product *models.ModelProductPerStore) (*models.ModelProductPerStore, error) {
	// Validar permiso de creación y acceso a la tienda
	if ok := s.EveryPower(ctx, POWER_STORE_PRODUCT_CREATE, PowerPrefixStore+product.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear productos en esta tienda")
	}

	// Validar que la tienda existe y obtener su company_id
	store, err := s.storeRepo.GetStoreByID(product.StoreID)
	if err != nil {
		return nil, err
	}
	if store == nil {
		return nil, types.ThrowRecipe("La tienda no existe", "storeId")
	}

	// Validar que el producto plantilla no esté ya asignado a esta tienda
	exists, err := s.ProductPerStoreRepo.ExistsByProductAndStore(product.ProductTemplate.ID, product.StoreID, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, types.ThrowRecipe("Este producto ya está asignado a la tienda", "productTemplate")
	}

	// La validación de que store.CompanyID == input.CompanyID se hace en el facade
	// Aquí solo verificamos que la tienda existe

	return s.ProductPerStoreRepo.CreateProductPerStore(product)
}

// CreateProductPerStore crea un nuevo producto para una tienda.
// Requiere permiso de creación y acceso a la tienda.
func (s *ProductPerStoreService) CreateProductPerStore(ctx context.Context, product *models.ModelProductPerStore) (*models.ModelProductPerStore, error) {
	// Validar permiso de creación y acceso a la tienda
	if ok := s.EveryPower(ctx, POWER_STORE_PRODUCT_CREATE, PowerPrefixStore+product.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear productos en esta tienda")
	}

	// Validar que el producto plantilla no esté ya asignado a esta tienda
	exists, err := s.ProductPerStoreRepo.ExistsByProductAndStore(product.ProductTemplate.ID, product.StoreID, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, types.ThrowRecipe("Este producto ya está asignado a la tienda", "productTemplate")
	}

	return s.ProductPerStoreRepo.CreateProductPerStore(product)
}

// UpdateProductPerStore actualiza un producto existente.
// Requiere permiso de actualización y acceso a la tienda.
func (s *ProductPerStoreService) UpdateProductPerStore(ctx context.Context, product *models.ModelProductPerStore) (*models.ModelProductPerStore, error) {
	// Primero obtenemos el producto actual para validar permisos
	existingProduct, err := s.ProductPerStoreRepo.GetProductPerStoreByID(product.ID)
	if err != nil {
		return nil, err
	}
	if existingProduct == nil {
		return nil, types.ThrowMsg("El producto no existe")
	}

	// Validar permiso de actualización y acceso a la tienda
	if ok := s.EveryPower(ctx, POWER_STORE_PRODUCT_UPDATE, PowerPrefixStore+existingProduct.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar productos en esta tienda")
	}

	// No permitir cambiar la tienda del producto
	product.StoreID = existingProduct.StoreID

	// Si se intenta cambiar el productId, validar que no exista ya en la tienda
	if product.ProductTemplate.ID != existingProduct.ProductTemplate.ID {
		exists, err := s.ProductPerStoreRepo.ExistsByProductAndStore(product.ProductTemplate.ID, product.StoreID, &product.ID)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, types.ThrowRecipe("Este producto ya está asignado a la tienda", "productTemplate")
		}
	}

	return s.ProductPerStoreRepo.UpdateProductPerStore(product)
}

// DeleteProductPerStore elimina un producto de una tienda.
// Requiere permiso de eliminación y acceso a la tienda.
func (s *ProductPerStoreService) DeleteProductPerStore(ctx context.Context, storeProductID string) error {
	// Primero obtenemos el producto para validar permisos
	product, err := s.ProductPerStoreRepo.GetProductPerStoreByID(storeProductID)
	if err != nil {
		return err
	}
	if product == nil {
		return types.ThrowMsg("El producto no existe")
	}

	// Validar permiso de eliminación y acceso a la tienda
	if ok := s.EveryPower(ctx, POWER_STORE_PRODUCT_DELETE, PowerPrefixStore+product.StoreID); !ok {
		return types.ThrowPower("No tienes permiso para eliminar productos en esta tienda")
	}

	return s.ProductPerStoreRepo.DeleteProductPerStore(storeProductID)
}

// GetProductsByStoreAndProductIDs obtiene productos por tienda y lista de IDs de producto base.
// Requiere permiso de acceso a la tienda.
func (s *ProductPerStoreService) GetProductsByStoreAndProductIDs(ctx context.Context, storeID string, productIDs []string) ([]models.ModelProductPerStore, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.ProductPerStoreRepo.GetProductsByStoreAndProductIDs(storeID, productIDs)
}
