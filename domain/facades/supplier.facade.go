package facades

import (
	"context"
	"sofia-backend/api/v1/dto"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/shared"
	"sofia-backend/types"
	"strings"
)

type SupplierFacade struct {
	appService *services.ServiceContainer
}

func NewSupplierFacade(appService *services.ServiceContainer) *SupplierFacade {
	return &SupplierFacade{appService: appService}
}

func (f *SupplierFacade) GetAllSuppliers() (shared.PaginationResponse[dto.DtoSupplierSimple], error) {
	suppliers, err := f.appService.SupplierService.GetSuppliers()
	if err != nil {
		return shared.PaginationResponse[dto.DtoSupplierSimple]{}, err
	}

	dtoSuppliers := f.toDtoSupplierSimpleList(suppliers)

	return shared.NewPagination(dtoSuppliers, 1, 1, len(dtoSuppliers)), nil
}

func (f *SupplierFacade) GetSupplierById(ctx context.Context, id string) (*dto.DtoSupplier, error) {
	supplier, err := f.appService.SupplierService.GetSupplierByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return f.toDtoSupplier(supplier)
}

func (f *SupplierFacade) GetCompanyProductsBySupplierId(storeId, supplierId string) (shared.PaginationResponse[models.ModelSupplierProductLegacy], error) {
	products, err := f.appService.SupplierProductService.GetSupplierProductsByStoreIDAndSupplierID(storeId, supplierId)
	if err != nil {
		return shared.PaginationResponse[models.ModelSupplierProductLegacy]{}, err
	}
	return shared.NewPagination(products, 1, 1, len(products)), nil
}

func (f *SupplierFacade) CreateSupplier(ctx context.Context, supRecipe recipe.RecipeCreateSupplier) (*dto.DtoSupplier, error) {

	if !strings.HasPrefix(supRecipe.IDFiscal, "CL-") {
		supRecipe.IDFiscal = "CL-" + supRecipe.IDFiscal
	}

	rawIDFiscal := strings.TrimPrefix(supRecipe.IDFiscal, "CL-")
	if !shared.IsValidRUT(rawIDFiscal) {
		return nil, types.ThrowMsg("El RUT no es válido")
	}

	if supRecipe.Description == nil {
		supRecipe.Description = new(string)
	}

	if supRecipe.Contacts == nil {
		supRecipe.Contacts = make([]recipe.RecipeSupplierContact, 0)
	}

	for i, contact := range supRecipe.Contacts {

		// en el futuro se debe validar con los telefonos extranjeros segun el countryId
		if !strings.HasPrefix(contact.Phone, "+569") {
			supRecipe.Contacts[i].Phone = "+569" + contact.Phone
		}

		trimed := strings.TrimPrefix(supRecipe.Contacts[i].Phone, "+569")
		if len(trimed) != 8 {
			return nil, types.ThrowMsg("El número de contacto " + supRecipe.Contacts[i].Phone + " no es válido.")
		}
	}

	createdSupplier, err := f.appService.SupplierService.CreateSupplier(ctx, supRecipe)
	if err != nil {
		return nil, err
	}

	return f.toDtoSupplier(createdSupplier)
}

func (f *SupplierFacade) UpdateSupplier(ctx context.Context, id string, supRecipe *recipe.RecipeCreateSupplier) (*dto.DtoSupplier, error) {

	if _, err := f.appService.SupplierService.GetSupplierByFiscalIDAndCountry(supRecipe.IDFiscal, 1); err == nil {
		return nil, types.ThrowRecipe("Ya existe un proveedor con el mismo RUT", "idFiscal")
	}

	if _, err := f.appService.SupplierService.GetSupplierByFiscalNameAndCountry(supRecipe.FiscalName, 1); err == nil {
		return nil, types.ThrowRecipe("Ya existe un proveedor con el mismo nombre fiscal", "fiscalName")
	}

	ogSupplier, err := f.appService.SupplierService.GetSupplierByID(ctx, id)
	if err != nil {
		return nil, types.ThrowMsg("No se encontró el proveedor a actualizar")
	}

	if !strings.HasPrefix(supRecipe.IDFiscal, "CL-") {
		supRecipe.IDFiscal = "CL-" + supRecipe.IDFiscal
	}

	if supRecipe.IDFiscal != ogSupplier.FiscalData.IDFiscal {
		return nil, types.ThrowRecipe("No se puede cambiar el id fiscal", "idFiscal")
	}

	if supRecipe.Description == nil {
		supRecipe.Description = new(string)
	}

	if supRecipe.Contacts == nil {
		supRecipe.Contacts = make([]recipe.RecipeSupplierContact, 0)
	}

	contacts := make([]models.ModelSupplierContact, len(supRecipe.Contacts))
	for i, contact := range supRecipe.Contacts {
		// en el futuro se debe validar con los telefonos extranjeros segun el countryId
		if !strings.HasPrefix(contact.Phone, "+569") {
			supRecipe.Contacts[i].Phone = "+569" + contact.Phone
		}

		trimed := strings.TrimPrefix(supRecipe.Contacts[i].Phone, "+569")
		if len(trimed) != 8 {
			return nil, types.ThrowMsg("El número de contacto " + supRecipe.Contacts[i].Phone + " no es válido.")
		}

		contacts[i] = models.ModelSupplierContact{
			Name:        contact.Name,
			Description: contact.Description,
			Email:       contact.Email,
			Phone:       contact.Phone,
		}
	}

	modelSupplier := &models.ModelSupplier{
		ID:           ogSupplier.ID,
		CountryID:    ogSupplier.CountryID,
		SupplierName: supRecipe.Name,
		Description:  *supRecipe.Description,
		Available:    supRecipe.Available,
		FiscalData: models.ModelFiscalData{
			ID:            ogSupplier.FiscalData.ID,
			IDFiscal:      ogSupplier.FiscalData.IDFiscal,
			RawFiscalID:   ogSupplier.FiscalData.RawFiscalID,
			FiscalName:    supRecipe.FiscalName,
			FiscalAddress: supRecipe.FiscalAddress,
			FiscalState:   supRecipe.FiscalState,
			FiscalCity:    supRecipe.FiscalCity,
			Email:         supRecipe.Email,
		},
		Contacts: contacts,
	}

	updatedSupp, err := f.appService.SupplierService.UpdateSupplier(ctx, modelSupplier, ogSupplier)
	if err != nil {
		return nil, err
	}

	return f.toDtoSupplier(updatedSupp)
}

func (f *SupplierFacade) EnableDisableSupplier(id string, available *bool) error {

	if available == nil {
		return types.ThrowRecipe("Available es obligatorio", "available")
	}

	return f.appService.SupplierService.EnableDisableSupplier(id, *available)
}

///////////////// SUPPLIER PRODUCTS CRUD ///////////////////

func (f *SupplierFacade) AddProductToSupplier(ctx context.Context, id string, supProdRecipe recipe.RecipeSupplierProductCreate) (*dto.DtoSupplier, error) {
	_, err := f.appService.SupplierService.GetSupplierByID(ctx, id)
	if err != nil {
		return nil, types.ThrowMsg("Id de proveedor inválido")
	}

	_, err = f.appService.ProductService.GetProductById(supProdRecipe.ProductID)
	if err != nil {
		return nil, types.ThrowMsg("Id de producto inválido")
	}

	if supProdRecipe.Available == nil {
		available := true
		supProdRecipe.Available = &available
	}

	if supProdRecipe.Description == nil {
		supProdRecipe.Description = new(string)
	}

	if _, err := f.appService.SupplierService.GetSupplierProductBySku(id, supProdRecipe.SKU); err == nil {
		return nil, types.ThrowRecipe("Ya existe un producto con el mismo SKU para este proveedor", "sku")
	}

	productToAdd := &models.ModelSupplierProduct{
		SupplierID:  id,
		ProductID:   supProdRecipe.ProductID,
		Name:        supProdRecipe.Name,
		Description: supProdRecipe.Description,
		SKU:         supProdRecipe.SKU,
		Price:       supProdRecipe.Price,
		PurchaseUnit: models.ModelProductPurchaseUnit{
			UnitID: supProdRecipe.UnitId,
		},
		Available: *supProdRecipe.Available,
	}

	_, err = f.appService.SupplierService.AddProductToSupplier(id, productToAdd)
	if err != nil {
		return nil, err
	}

	return f.GetSupplierById(ctx, id)
}

func (f *SupplierFacade) UpdateSupplierProductPrices(ctx context.Context, supplierID string, products []recipe.RecipeSuplierProductPriceUpdate) ([]dto.DtoSupplierProduct, error) {
	_, err := f.appService.SupplierService.GetSupplierByID(ctx, supplierID)
	if err != nil {
		return nil, types.ThrowMsg("Id de proveedor inválido")
	}

	supplierProducts, err := f.appService.SupplierService.GetSupplierProducts(supplierID)
	if err != nil {
		return nil, types.ThrowMsg("No se encontraron productos para actualizar en este proveedor")
	}

	productsToUpdate := make([]models.ModelSupplierProduct, 0)

	for _, product := range products {
		for _, supProd := range supplierProducts {
			if supProd.ProductID == product.ProductID {
				supProd.Price = product.Price
				productsToUpdate = append(productsToUpdate, supProd)
				break
			}
		}
	}

	_, err = f.appService.SupplierService.UpdateSupplierProductsPrice(ctx, supplierID, productsToUpdate)
	if err != nil {
		return nil, err
	}

	productsUpdated, err := f.appService.SupplierService.GetSupplierProducts(supplierID)
	if err != nil {
		return nil, types.ThrowMsg("Error al obtener los productos actualizados")
	}

	return f.toDtoSupplierProductList(productsUpdated), nil
}

func (f *SupplierFacade) UpdateSupplierProduct(ctx context.Context, supplierID string, supProdRecipe recipe.RecipeSupplierProductCreate) (*dto.DtoSupplierProduct, error) {
	_, err := f.appService.SupplierService.GetSupplierByID(ctx, supplierID)
	if err != nil {
		return nil, types.ThrowMsg("Id de proveedor inválido")
	}

	productToUpdate, err := f.appService.SupplierService.GetSupplierProductById(supplierID, supProdRecipe.ProductID)
	if err != nil {
		return nil, types.ThrowMsg("Id de producto inválido")
	}

	if supplierProduct, err := f.appService.SupplierService.GetSupplierProductBySku(supplierID, supProdRecipe.SKU); err == nil && supplierProduct.ID != productToUpdate.ID {
		return nil, types.ThrowRecipe("Ya existe un producto con el mismo SKU para este proveedor", "sku")
	}

	productToUpdate.Name = supProdRecipe.Name
	productToUpdate.Description = supProdRecipe.Description
	productToUpdate.SKU = supProdRecipe.SKU
	productToUpdate.Price = supProdRecipe.Price
	productToUpdate.PurchaseUnit = models.ModelProductPurchaseUnit{
		UnitID: supProdRecipe.UnitId,
	}

	if supProdRecipe.Available != nil {
		productToUpdate.Available = *supProdRecipe.Available
	}

	_, err = f.appService.SupplierService.UpdateSupplierProduct(ctx, supplierID, productToUpdate)
	if err != nil {
		return nil, err
	}

	productUpdated, err := f.appService.SupplierService.GetSupplierProductById(supplierID, productToUpdate.ID)
	if err != nil {
		return nil, err
	}

	return f.toDtoSupplierProduct(*productUpdated), nil
}

func (f *SupplierFacade) DeleteSupplierProduct(ctx context.Context, supplierID, productID string) (*dto.DtoSupplierProduct, error) {
	_, err := f.appService.SupplierService.GetSupplierByID(ctx, supplierID)
	if err != nil {
		return nil, types.ThrowMsg("Id de proveedor inválido")
	}

	_, err = f.appService.SupplierService.GetSupplierProductById(supplierID, productID)
	if err != nil {
		return nil, types.ThrowMsg("Id de producto inválido")
	}

	productDeleted, err := f.appService.SupplierService.DeleteSupplierProduct(ctx, supplierID, productID)
	if err != nil {
		return nil, err
	}

	return f.toDtoSupplierProduct(*productDeleted), nil

}

///////////////// PRIVATE METHODS ///////////////////

func (f *SupplierFacade) toDtoSupplierSimple(supplier models.ModelSupplier) dto.DtoSupplierSimple {
	return dto.DtoSupplierSimple{
		ID:           supplier.ID,
		IDFiscal:     supplier.FiscalData.IDFiscal,
		RawFiscalID:  supplier.FiscalData.RawFiscalID,
		CountryID:    supplier.CountryID,
		SupplierName: supplier.SupplierName,
		Description:  supplier.Description,
		Available:    supplier.Available,
	}
}

func (f *SupplierFacade) toDtoSupplierSimpleList(suppliers []models.ModelSupplier) []dto.DtoSupplierSimple {
	dtoSuppliers := make([]dto.DtoSupplierSimple, len(suppliers))
	for i, supplier := range suppliers {
		dtoSuppliers[i] = f.toDtoSupplierSimple(supplier)
	}
	return dtoSuppliers
}

func (f *SupplierFacade) toDtoSupplier(supplier *models.ModelSupplier) (*dto.DtoSupplier, error) {

	if supplier == nil {
		return nil, types.ThrowMsg("Error al enviar los datos del proveedor")
	}

	contacts := make([]dto.DtoSupplierContact, 0)
	for _, contact := range supplier.Contacts {
		contacts = append(contacts, dto.DtoSupplierContact{
			ID:          contact.ID,
			Name:        contact.Name,
			Description: contact.Description,
			Email:       contact.Email,
			Phone:       contact.Phone,
		})
	}

	return &dto.DtoSupplier{
		ID:          supplier.ID,
		Name:        supplier.SupplierName,
		Description: supplier.Description,
		Available:   supplier.Available,
		CountryID:   supplier.CountryID,
		FiscalData: dto.DtoSupplierFiscalData{
			ID:            supplier.FiscalData.ID,
			IDFiscal:      supplier.FiscalData.IDFiscal,
			RawFiscalID:   supplier.FiscalData.RawFiscalID,
			FiscalName:    supplier.FiscalData.FiscalName,
			FiscalAddress: supplier.FiscalData.FiscalAddress,
			FiscalState:   supplier.FiscalData.FiscalState,
			FiscalCity:    supplier.FiscalData.FiscalCity,
			Email:         supplier.FiscalData.Email,
		},
		Contacts: contacts,
		Products: f.toDtoSupplierProductList(supplier.Products),
	}, nil
}

func (f *SupplierFacade) toDtoSupplierProduct(supplierProduct models.ModelSupplierProduct) *dto.DtoSupplierProduct {
	return &dto.DtoSupplierProduct{
		ID:          supplierProduct.ID,
		SupplierID:  supplierProduct.SupplierID,
		ProductID:   supplierProduct.ProductID,
		Name:        supplierProduct.Name,
		Description: supplierProduct.Description,
		SKU:         supplierProduct.SKU,
		Price:       supplierProduct.Price,
		Unit:        supplierProduct.PurchaseUnit.UnitID,
		UnitAbv:     supplierProduct.PurchaseUnit.UnitAbv,
		UnitName:    supplierProduct.PurchaseUnit.UnitName,
		Available:   supplierProduct.Available,
		CreatedAt:   supplierProduct.CreatedAt,
		UpdatedAt:   supplierProduct.UpdatedAt,
	}
}

func (f *SupplierFacade) toDtoSupplierProductList(supplierProducts []models.ModelSupplierProduct) []dto.DtoSupplierProduct {
	dtoSupplierProducts := make([]dto.DtoSupplierProduct, len(supplierProducts))
	for i, supplierProduct := range supplierProducts {
		dtoSupplierProducts[i] = *f.toDtoSupplierProduct(supplierProduct)
	}
	return dtoSupplierProducts
}
