package facades

import (
	"context"
	"sofia-backend/api/v1/dto"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/types"
)

// StoreProductFacade maneja la lógica de negocio para productos de tienda.
type StoreProductFacade struct {
	appServices *services.ServiceContainer
}

// NewStoreProductFacade crea una nueva instancia del facade.
func NewStoreProductFacade(appServices *services.ServiceContainer) *StoreProductFacade {
	return &StoreProductFacade{
		appServices: appServices,
	}
}

// GetStoreProductByID obtiene un producto de tienda por su ID.
func (f *StoreProductFacade) GetStoreProductByID(ctx context.Context, storeProductID string) (*dto.DtoStoreProduct, error) {
	product, err := f.appServices.ProductPerStoreService.GetProductPerStoreByID(ctx, storeProductID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, types.ThrowMsg("El producto de tienda no existe")
	}

	// Enriquecer con proveedores
	dtoProduct := dto.NewDtoStoreProduct(product)

	enrichedSuppliers, err := f.enrichSuppliersWithFiscalData(product.Suppliers)
	if err != nil {
		return nil, err
	}
	dtoProduct.Suppliers = enrichedSuppliers

	return dtoProduct, nil
}

// GetProductsByStore obtiene todos los productos de una tienda.
func (f *StoreProductFacade) GetProductsByStore(ctx context.Context, storeID string) ([]models.ModelProductPerStore, error) {
	products, err := f.appServices.ProductPerStoreService.GetProductsByStore(ctx, storeID)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return []models.ModelProductPerStore{}, nil
	}

	return products, nil

	/* // Preparar IDs de producto plantilla para proveedores
	productIDs := make([]string, len(products))
	for i := range products {
		productIDs[i] = products[i].ProductTemplate.ID
	}

	// Obtener proveedores en lote (por producto plantilla)
	suppliers, err := f.appServices.SupplierService.SupplierRepo.GetSuppliersByStoreProductId(storeID, productIDs)
	if err != nil {
		return nil, err
	}

	// Agrupar proveedores por productID (plantilla)
	supByProduct := make(map[string][]models.ModelSupplierStoreProduct)
	for _, s := range suppliers {
		supByProduct[s.ProductID] = append(supByProduct[s.ProductID], s)
	}

	// Mapear a DTO
	result := make([]dto.DtoStoreProduct, len(products))
	for i := range products {
		dtoP := dto.NewDtoStoreProduct(&products[i])
		if list, ok := supByProduct[products[i].ProductTemplate.ID]; ok {
			enrichedSuppliers, err := f.enrichSuppliersWithFiscalData(list)
			if err != nil {
				return nil, err
			}
			dtoP.Suppliers = enrichedSuppliers
		}
		result[i] = *dtoP
	}

	return result, nil */
}

func (f *StoreProductFacade) GetProductsWithSuppliersByStore(ctx context.Context, storeID string) ([]dto.DtoStoreProduct, error) {
	products, err := f.appServices.ProductPerStoreService.GetProductsWithSuppliersByStore(ctx, storeID)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return []dto.DtoStoreProduct{}, nil
	}

	// Mapear a DTO y enriquecer proveedores
	result := make([]dto.DtoStoreProduct, len(products))
	for i := range products {
		dtoP := dto.NewDtoStoreProduct(&products[i])
		enrichedSuppliers, err := f.enrichSuppliersWithFiscalData(products[i].Suppliers)
		if err != nil {
			return nil, err
		}
		dtoP.Suppliers = enrichedSuppliers
		result[i] = *dtoP
	}

	return result, nil
}

// GetSuppliersForStoreProduct obtiene los proveedores asignados a un producto de tienda.
func (f *StoreProductFacade) GetSuppliersForStoreProduct(ctx context.Context, storeProductID string) ([]dto.DtoStoreProductSupplier, error) {
	return nil, types.ThrowMsg("no implementado aun")
	/*
		// Primero obtenemos el producto de tienda para saber su productID y storeID
		product, err := f.appServices.ProductPerStoreService.GetProductPerStoreByID(ctx, storeProductID)
		if err != nil {
			return nil, err
		}
		if product == nil {
			return nil, types.ThrowMsg("El producto de tienda no existe")
		}

		// Obtener proveedores con información completa directamente del repo
		suppliersComplete, err := f.appServices.SupplierService.SupplierRepo.GetSuppliersByStoreProductId(product.StoreID, []string{product.ProductID})
		if err != nil {
			return nil, err
		}

		result := make([]dto.DtoStoreProductSupplier, len(suppliersComplete))
		for i, sup := range suppliersComplete {
			result[i] = dto.DtoStoreProductSupplier{
				SupplierProductID: sup.SupplierProductID,
				SupplierID:        sup.SupplierID,
				SupplierName:      sup.SupplierName,
				ProductID:         sup.ProductID,
				ProductName:       sup.ProductName,
				Priority:          sup.Priority,
				Preferred:         sup.Preferred,
				Price:             sup.Price,
				MinOrderQuantity:  sup.MinOrderQuantity,
				PaymentTerms:      sup.PaymentTerms,
				LeadTimeDays:      sup.LeadTimeDays,
				Available:         sup.Available,
			}
		}

		return result, nil
	*/
}

// CreateStoreProduct crea un nuevo producto de tienda con sus proveedores.
func (f *StoreProductFacade) CreateStoreProduct(ctx context.Context, input *recipe.RecipeCreateStoreProduct) (*dto.DtoStoreProduct, error) {
	// 0. Validar que la tienda existe y pertenece a la compañía
	store, err := f.appServices.StoreService.GetStoreByID(ctx, input.StoreID)
	if err != nil {
		return nil, err
	}
	if store == nil {
		return nil, types.ThrowRecipe("La tienda no existe", "storeId")
	}
	if store.CompanyID != input.CompanyID {
		return nil, types.ThrowRecipe("La tienda no pertenece a la compañía especificada", "storeId")
	}

	// 1. Validar que el producto base existe
	product, err := f.appServices.ProductService.ProductRepo.GetById(input.ProductID)
	if err != nil {
		return nil, types.ThrowRecipe("No se pudo obtener la plantilla", "productId")
	}
	if product == nil {
		return nil, types.ThrowRecipe("La plantilla no existe", "productId")
	}

	// 2. Validar que la unidad de inventario existe
	unitExists, err := f.appServices.MeasurementService.MeasurementRepo.ExistsMeasurementById(input.UnitInventoryID)
	if err != nil {
		return nil, err
	}
	if !unitExists {
		return nil, types.ThrowRecipe("La unidad de inventario no existe", "unitInventoryId")
	}

	// 3. Validar que la unidad de inventario no esté en las conversiones
	/* for _, unit := range input.UnitMatrix {
		if unit.UnitID == input.UnitInventoryID {
			return nil, types.ThrowRecipe("La unidad de inventario no puede estar en las conversiones", "unitMatrix")
		}
	} */

	// 4. Validar suppliers
	if err := f.validateSuppliers(input.CompanyID, input.Suppliers); err != nil {
		return nil, err
	}

	// 6. Construir modelo para crear
	storeProduct := &models.ModelProductPerStore{
		StoreID:         input.StoreID,
		ProductTemplate: *product,
		TagID:           input.TagID,
		ProductName:     input.ProductName,
		Image:           product.Image,
		ItemSale:        input.IsSellable,
		UseRecipe:       input.UseRecipe,
		UnitInventory: models.ModelMeasurementUnique{
			ID: input.UnitInventoryID,
		},
		UnitMatrix:  f.recipeToModelUnitMatrix(input.UnitMatrix),
		Description: input.Description,
		Costs: models.ModelStoreProductCost{
			// DEPRECATED: CostEstimated ahora está en ProductTemplate
			// CostAvg:       input.Costs.CostAvg,
			// CostAvg:       0,
		},
		Quantities: models.ModelStoreProductQuantities{
			MinimalStock: input.Quantities.MinimalStock,
			MaximalStock: input.Quantities.MaximalStock,
			MaxQuantity:  input.Quantities.MaxQuantity,
		},
	}

	// 7. Agregar suppliers al modelo
	storeProduct.Suppliers = f.recipeToModelSuppliersForCreate(input.Suppliers)

	// 8. Llamar al servicio para crear
	created, err := f.appServices.ProductPerStoreService.CreateStoreProduct(ctx, storeProduct)
	if err != nil {
		return nil, err
	}

	// 9. Enriquecer datos de unidades (nombre y abreviación) desde la BD
	if err := f.enrichUnitInventoryAndMatrix(created); err != nil {
		return nil, err
	}

	// 10. Enriquecer proveedores con datos fiscales
	enrichedSuppliers, err := f.enrichSuppliersWithFiscalData(created.Suppliers)
	if err != nil {
		return nil, err
	}

	// 11. Convertir a DTO y asignar suppliers enriquecidos
	dtoProduct := dto.NewDtoStoreProduct(created)
	dtoProduct.Suppliers = enrichedSuppliers

	return dtoProduct, nil
}

// UpdateStoreProduct actualiza un producto de tienda con lógica de upsert para proveedores.
func (f *StoreProductFacade) UpdateStoreProduct(ctx context.Context, storeProductID string, input *recipe.RecipeUpdateStoreProduct) (*dto.DtoStoreProduct, error) {
	// 1. Obtener producto existente
	existingProduct, err := f.appServices.ProductPerStoreService.GetProductPerStoreByID(ctx, storeProductID)
	if err != nil {
		return nil, err
	}
	if existingProduct == nil {
		return nil, types.ThrowMsg("El producto de tienda no existe")
	}

	// 2. Obtener información de la tienda para validaciones
	store, err := f.appServices.StoreService.GetStoreByID(ctx, existingProduct.StoreID)
	if err != nil {
		return nil, err
	}
	if store == nil {
		return nil, types.ThrowMsg("La tienda no existe")
	}

	// 3. Actualizar campos solo si están presentes en el input (partial update)
	if input.ProductName != nil {
		existingProduct.ProductName = *input.ProductName
	}
	if input.TagID != nil {
		existingProduct.TagID = *input.TagID
	}
	if input.IsSellable != nil {
		existingProduct.ItemSale = *input.IsSellable
	}
	if input.UseRecipe != nil {
		existingProduct.UseRecipe = *input.UseRecipe
	}
	if input.UnitInventoryID != nil {
		// Validar que la unidad de inventario existe
		unitExists, err := f.appServices.MeasurementService.MeasurementRepo.ExistsMeasurementById(*input.UnitInventoryID)
		if err != nil {
			return nil, err
		}
		if !unitExists {
			return nil, types.ThrowRecipe("La unidad de inventario no existe", "unitInventoryId")
		}
		existingProduct.UnitInventory.ID = *input.UnitInventoryID
	}
	if len(input.UnitMatrix) > 0 {
		existingProduct.UnitMatrix = f.recipeToModelUnitMatrix(input.UnitMatrix)
	}
	if input.Description != nil {
		existingProduct.Description = *input.Description
	}

	// DEPRECATED: CostEstimated ahora está en ProductTemplate, no en Costs
	// existingProduct.Costs = models.ModelStoreProductCost{
	// 	// CostAvg:       input.Costs.CostAvg,
	// 	// CostAvg:       existingProduct.Costs.CostAvg,
	// }

	if input.Quantities != nil {
		existingProduct.Quantities = models.ModelStoreProductQuantities{
			MinimalStock: input.Quantities.MinimalStock,
			MaximalStock: input.Quantities.MaximalStock,
			MaxQuantity:  input.Quantities.MaxQuantity,
		}
	}
	if len(input.Suppliers) > 0 {
		// Validar suppliers si se envían
		if err := f.validateSuppliers(store.CompanyID, input.Suppliers); err != nil {
			return nil, err
		}
		existingProduct.Suppliers = f.recipeToModelSuppliersForCreate(input.Suppliers)
	}

	// 4. Llamar al servicio para actualizar
	updated, err := f.appServices.ProductPerStoreService.UpdateProductPerStore(ctx, existingProduct)
	if err != nil {
		return nil, err
	}

	// 5. Enriquecer con proveedores (incluyendo rawFiscalId)
	suppliers, err := f.appServices.SupplierService.SupplierRepo.GetSuppliersByStoreProductId(updated.StoreID, []string{updated.ProductTemplate.ID})
	if err != nil {
		return nil, err
	}

	dtoProduct := dto.NewDtoStoreProduct(updated)
	if len(suppliers) > 0 {
		enrichedSuppliers, err := f.enrichSuppliersWithFiscalData(suppliers)
		if err != nil {
			return nil, err
		}
		dtoProduct.Suppliers = enrichedSuppliers
	}

	return dtoProduct, nil
}

// recipeToModelUnitMatrix convierte el recipe de unit matrix a modelo.
func (f *StoreProductFacade) recipeToModelUnitMatrix(recipe []recipe.RecipeUnitConversion) []models.ModelMeasurementConversionUnit {

	result := make([]models.ModelMeasurementConversionUnit, len(recipe))
	for i, r := range recipe {
		result[i] = models.ModelMeasurementConversionUnit{
			ID:     r.UnitID,
			Factor: r.Factor,
		}
	}
	return result
}

// recipeToModelSuppliers convierte el recipe de suppliers a modelo.
/* func (f *StoreProductFacade) recipeToModelSuppliers(recipe []recipe.RecipeStoreProductSupplier) []models.ModelSupplierStoreProduct {
	result := make([]models.ModelSupplierStoreProduct, len(recipe))
	for i, r := range recipe {
		result[i] = models.ModelSupplierStoreProduct{
			SupplierID: r.SupplierID,
			Priority:   r.Priority,
		}
	}
	return result
} */

// recipeToModelSuppliersForCreate convierte el recipe de suppliers a modelo para crear storeproduct.
func (f *StoreProductFacade) recipeToModelSuppliersForCreate(recipe []recipe.RecipeStoreProductSupplier) []models.ModelSupplierStoreProduct {
	result := make([]models.ModelSupplierStoreProduct, len(recipe))
	for i, r := range recipe {
		result[i] = models.ModelSupplierStoreProduct{
			SupplierID: r.SupplierID,
			Priority:   r.Priority,
		}
	}
	return result
}

// enrichUnitInventoryAndMatrix enriquece los datos de unidades con nombre y abreviación desde la BD.
// Esto es necesario después de crear un producto ya que el modelo solo contiene IDs.
func (f *StoreProductFacade) enrichUnitInventoryAndMatrix(product *models.ModelProductPerStore) error {
	// 1. Enriquecer UnitInventory con datos completos de la unidad
	enrichedInventory, err := f.appServices.MeasurementService.MeasurementRepo.GetById(product.UnitInventory.ID)
	if err != nil {
		return err
	}
	if enrichedInventory != nil {
		product.UnitInventory = *enrichedInventory
	}

	// 2. Enriquecer UnitMatrix con datos completos de cada unidad de conversión
	if len(product.UnitMatrix) > 0 {
		for i, unit := range product.UnitMatrix {
			enrichedUnit, err := f.appServices.MeasurementService.MeasurementRepo.GetById(unit.ID)
			if err != nil {
				return err
			}
			if enrichedUnit != nil {
				// Preservar el factor de conversión que ya está en el modelo
				product.UnitMatrix[i] = models.ModelMeasurementConversionUnit{
					ID:           enrichedUnit.ID,
					Name:         enrichedUnit.Name,
					Abbreviation: enrichedUnit.Abbreviation,
					Description:  enrichedUnit.Description,
					Factor:       unit.Factor,
				}
			}
		}
	}

	return nil
}

// enrichSuppliersWithFiscalData enriquece los proveedores con datos fiscales.
func (f *StoreProductFacade) enrichSuppliersWithFiscalData(suppliers []models.ModelSupplierStoreProduct) ([]dto.DtoStoreProductSupplier, error) {
	if len(suppliers) == 0 {
		return []dto.DtoStoreProductSupplier{}, nil
	}

	dtoSuppliers := make([]dto.DtoStoreProductSupplier, 0)
	for _, sup := range suppliers {
		// Obtener datos completos del proveedor incluyendo fiscal data
		fullSupplier, err := f.appServices.SupplierService.SupplierRepo.GetSupplierByID(sup.SupplierID)
		if err != nil {

			return nil, err
		}
		if fullSupplier != nil {
			dtoSuppliers = append(dtoSuppliers, dto.DtoStoreProductSupplier{
				ID:          sup.SupplierID,
				Name:        fullSupplier.SupplierName,
				Priority:    sup.Priority,
				RawFiscalId: fullSupplier.FiscalData.RawFiscalID,
			})
		}
	}

	return dtoSuppliers, nil
}

// validateSuppliers valida que todos los suppliers existen y pertenecen a la compañía.
// También valida que las prioridades sean >= 0 y no haya duplicados.
func (f *StoreProductFacade) validateSuppliers(companyID string, suppliers []recipe.RecipeStoreProductSupplier) error {
	if len(suppliers) == 0 {
		return nil // No hay suppliers que validar
	}

	// Mapa para detectar prioridades duplicadas
	priorityMap := make(map[int]bool)
	count := 0
	for _, sup := range suppliers {
		count++
		// 1. Validar que priority >= 0
		if sup.Priority < 0 {
			return types.ThrowRecipe("La prioridad no puede ser negativa", "suppliers")
		}

		// 2. Validar que no hay prioridades duplicadas
		if priorityMap[sup.Priority] {
			return types.ThrowRecipe("No pueden existir dos proveedores con la misma prioridad", "suppliers")
		}
		priorityMap[sup.Priority] = true

		// 3. Validar que el supplier existe
		supplier, err := f.appServices.SupplierService.SupplierRepo.GetSupplierByID(sup.SupplierID)
		if err != nil {
			return err
		}
		if supplier == nil {
			return types.ThrowRecipe("El proveedor con ID "+sup.SupplierID+" no existe", "suppliers")
		}

		// 4. Validar que el supplier pertenece a la compañía
		exists, err := f.appServices.SupplierService.SupplierRepo.ExistsSupplierInCompany(sup.SupplierID, companyID)
		if err != nil {
			return err
		}
		if !exists {
			return types.ThrowRecipe("El proveedor con ID "+sup.SupplierID+" no pertenece a la compañía", "suppliers")
		}
	}

	return nil
}
