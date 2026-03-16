package facades

import (
	"context"
	"sofia-backend/api/v1/dto"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/shared"
	"sofia-backend/types"
)

type ProductFacade struct {
	appService *services.ServiceContainer
}

func NewProductFacade(appService *services.ServiceContainer) *ProductFacade {
	return &ProductFacade{
		appService: appService,
	}
}

// /////// Products /////////
func (f *ProductFacade) GetAllCodes() ([]models.ModelProductCodeKind, error) {
	return f.appService.ProductService.GetAllCodes()
}

func (f *ProductFacade) GetAllCategories() ([]models.ModelProductCategory, error) {
	return f.appService.ProductService.GetAllCategories()
}

func (f *ProductFacade) CreateCategory(ctx context.Context, recipe *recipe.RecipeProductCategory) (*models.ModelProductCategory, error) {
	// Validar si ya existe una categoría con ese nombre
	if _, err := f.appService.ProductService.GetCategoryByName(recipe.Name); err == nil {
		return nil, types.ThrowRecipe("Ya existe una categoría con ese nombre", "name")
	}

	return f.appService.ProductService.CreateCategory(ctx, recipe)
}

func (f *ProductFacade) GetProductById(id string) (*models.ModelProduct, error) {
	return f.appService.ProductService.GetProductById(id)
}

func (f *ProductFacade) GetProductWithSuppliersById(companyId, id string) (*models.ModelProductWithSuppliers, error) {
	_product, err := f.appService.ProductService.GetProductById(id)
	if err != nil {
		return nil, err
	}
	if _product == nil {
		return nil, nil
	}

	_suppliers, err := f.appService.SupplierService.SupplierRepo.GetSuppliersByTemplateProductId(
		companyId, id,
	)
	if err != nil {
		return nil, err
	}

	_productWithSuppliers := &models.ModelProductWithSuppliers{
		ModelProduct: *_product,
		Suppliers:    make([]models.ModelSupplier, 0),
	}

	_productWithSuppliers.Suppliers = _suppliers

	return _productWithSuppliers, nil
}

func (f *ProductFacade) GetAllProducts() (shared.PaginationResponse[dto.DtoSimpleProduct], error) {
	products, err := f.appService.ProductService.GetAllProducts()
	if err != nil {
		return shared.PaginationResponse[dto.DtoSimpleProduct]{}, err
	}

	dtoProducts := f.toDtoSimpleList(products)

	return shared.NewPagination(dtoProducts, len(dtoProducts), 1, len(dtoProducts)), nil
}

func (f *ProductFacade) GetProductsRequestRestriction(ctx context.Context, storeId string) (shared.PaginationResponse[models.ModelProductRequestRestriction], error) {
	products, _, err := f.appService.ProductPerStoreService.GetProductsRequestRestriction(ctx, storeId)
	if err != nil {
		return shared.PaginationResponse[models.ModelProductRequestRestriction]{}, err
	}
	return shared.NewPagination(products, 1, 1, len(products)), nil
}

func (f *ProductFacade) CreateProduct(ctx context.Context, productInput *recipe.RecipeProductInput) (*dto.DtoSimpleProduct, error) {

	if productInput.Description == nil {
		defaultDesc := ""
		productInput.Description = &defaultDesc
	}

	if productInput.CategoryIds == nil {
		productInput.CategoryIds = []int64{}
	}

	if productInput.CodesList == nil {
		productInput.CodesList = []recipe.RecipeProductCode{}
	}

	createdProduct, err := f.appService.ProductService.CreateProduct(ctx, productInput)
	if err != nil {
		return nil, err
	}

	return f.toDtoSimpleProduct(createdProduct), nil
}

func (f *ProductFacade) UpdateProduct(ctx context.Context, id string, productInput *recipe.RecipeProductInput) (*dto.DtoSimpleProduct, error) {
	ogProduct, err := f.appService.ProductService.GetProductById(id)
	if err != nil {
		return nil, types.ThrowMsg("No se encontró el producto para la actualización")
	}

	if productInput.Description == nil {
		defaultDesc := ""
		productInput.Description = &defaultDesc
	}

	if productInput.CategoryIds == nil {
		productInput.CategoryIds = []int64{}
	}

	if productInput.CodesList == nil {
		productInput.CodesList = []recipe.RecipeProductCode{}
	}

	updatedProduct, err := f.appService.ProductService.UpdateProduct(ctx, id, productInput, ogProduct)
	if err != nil {
		return nil, err
	}

	return f.toDtoSimpleProduct(updatedProduct), nil
}

/*
func (f *ProductFacade) DeleteProduct(id string) error {
	return f.appService.ProductService.DeleteProduct(id)
}
*/

// /////// Store Products (Productos por Tienda) /////////

// GetStoreProductsByStoreId obtiene los productos configurados para una tienda.
// Reemplaza GetCompanyProductsByCompanyId.
func (f *ProductFacade) GetStoreProductsByStoreId(ctx context.Context, storeId string) (shared.PaginationResponse[models.ModelProductPerStore], error) {
	products, err := f.appService.ProductPerStoreService.GetProductsByStore(ctx, storeId)
	if err != nil {
		return shared.PaginationResponse[models.ModelProductPerStore]{}, err
	}

	return shared.NewPagination(products, 1, 1, len(products)), nil
}

// DEPRECATED: GetCompanyProductsByCompanyId - usar GetStoreProductsByStoreId
/*
func (f *ProductFacade) GetCompanyProductsByCompanyId(ctx context.Context, companyId string) (shared.PaginationResponse[models.ModelProductCompany], error) {
	products, err := f.appService.ProductCompanyService.GetProductsByCompany(ctx, companyId)
	if err != nil {
		return shared.PaginationResponse[models.ModelProductCompany]{}, err
	}

	return shared.NewPagination(products, 1, 1, len(products)), nil
}
*/

/////////////////// Private Methods ///////////////////

func (f *ProductFacade) toDtoSimpleList(products []models.ModelProduct) []dto.DtoSimpleProduct {
	dtoProducts := make([]dto.DtoSimpleProduct, 0, len(products))
	for _, product := range products {
		dtoProducts = append(dtoProducts, *f.toDtoSimpleProduct(&product))
	}

	return dtoProducts
}

func (f *ProductFacade) toDtoSimpleProduct(product *models.ModelProduct) *dto.DtoSimpleProduct {
	dtoCategories := make([]dto.DtoProductCategory, 0, len(product.Categories))
	for _, category := range product.Categories {
		dtoCategories = append(dtoCategories, *f.toDtoCategory(&category))
	}

	dtoCodes := make([]dto.DtoProductCode, 0, len(product.Codes))
	for _, code := range product.Codes {
		dtoCodes = append(dtoCodes, *f.toDtoCode(&code))
	}

	// Obtener precio anterior del historial
	previousPrice, _ := f.appService.PriceHistoryService.GetPreviousPrice(product.ID)

	return &dto.DtoSimpleProduct{
		ID:            product.ID,
		Name:          product.Name,
		SKU:           product.SKU,
		CostEstimated: product.CostEstimated,
		PreviousPrice: previousPrice,
		Description:   product.Description,
		Image:         product.Image,
		Categories:    dtoCategories,
		Codes:         dtoCodes,
	}
}

func (f *ProductFacade) toDtoCategory(category *models.ModelProductCategory) *dto.DtoProductCategory {
	return &dto.DtoProductCategory{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		Available:   category.Available,
	}
}

func (f *ProductFacade) toDtoCode(code *models.ModelProductCode) *dto.DtoProductCode {
	return &dto.DtoProductCode{
		ID:    code.ID,
		Kind:  *f.toDtoCodeKind(&code.Kind),
		Value: code.Value,
	}
}

func (f *ProductFacade) toDtoCodeKind(codeKind *models.ModelProductCodeKind) *dto.DtoProductCodeKind {
	return &dto.DtoProductCodeKind{
		ID:          codeKind.ID,
		Name:        codeKind.Name,
		Description: codeKind.Description,
	}
}
