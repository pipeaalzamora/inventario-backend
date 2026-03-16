package services

import (
	"context"
	"fmt"
	"mime/multipart"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
)

const (
	POWER_TEMPLATE_PRODUCT_CREATE = "template_product:create"
	POWER_TEMPLATE_PRODUCT_UPDATE = "template_product:update"
)

type ProductService struct {
	PowerChecker

	ProductRepo     ports.PortProduct
	productCodeKind ports.PortProductCode
	productCategory ports.PortCategory
	bucket          ports.PortBucket
}

func NewProductService(repo ports.PortProduct,
	codeKind ports.PortProductCode,
	bucket ports.PortBucket,
	category ports.PortCategory,
) *ProductService {
	return &ProductService{
		ProductRepo:     repo,
		productCodeKind: codeKind,
		productCategory: category,
		bucket:          bucket,
	}
}

func (s *ProductService) GetAllCodes() ([]models.ModelProductCodeKind, error) {
	return s.productCodeKind.GetAll()
}

func (s *ProductService) GetAllCategories() ([]models.ModelProductCategory, error) {
	return s.productCategory.GetAll()
}

func (s *ProductService) GetProductById(id string) (*models.ModelProduct, error) {
	return s.ProductRepo.GetById(id)
}

func (s *ProductService) GetAllProducts() ([]models.ModelProduct, error) {
	return s.ProductRepo.GetAllFull()
}

func (s *ProductService) CreateProduct(ctx context.Context, productInput *recipe.RecipeProductInput) (*models.ModelProduct, error) {
	if ok := s.EveryPower(ctx, PowerProductUpdate); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar productos")
	}

	newProduct := &models.ModelProduct{
		Name:          productInput.Name,
		Description:   productInput.Description,
		CostEstimated: productInput.CostEstimated,
		Image:         nil,
	}

	if productInput.Image != nil && productInput.Image.Size > 0 {
		url, err := s.SaveImageToBucket(ctx, productInput.Image)
		if err != nil {
			return nil, err
		}
		newProduct.Image = url
	}

	productCodes, err := s.validateProductCodes(productInput, nil)
	if err != nil {
		return nil, err
	}

	productCategories, err := s.validateProductCategories(productInput.CategoryIds)
	if err != nil {
		return nil, err
	}

	newProduct.Codes = productCodes
	newProduct.Categories = productCategories

	return s.ProductRepo.Create(newProduct)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id string, productInput *recipe.RecipeProductInput, oldProduct *models.ModelProduct) (*models.ModelProduct, error) {
	if ok := s.EveryPower(ctx, PowerProductUpdate); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar productos")
	}

	updateProduct := &models.ModelProduct{
		ID:            id,
		Name:          productInput.Name,
		Description:   productInput.Description,
		CostEstimated: productInput.CostEstimated,
	}

	if productInput.IsNewImage {
		if productInput.Image != nil && productInput.Image.Size > 0 {
			url, err := s.SaveImageToBucket(ctx, productInput.Image)
			if err != nil {
				return nil, err
			}
			updateProduct.Image = url
		}
	} else {
		updateProduct.Image = oldProduct.Image
	}

	productCodes, err := s.validateProductCodes(productInput, &id)
	if err != nil {
		return nil, err
	}

	productCategories, err := s.validateProductCategories(productInput.CategoryIds)
	if err != nil {
		return nil, err
	}

	updateProduct.Codes = productCodes
	updateProduct.Categories = productCategories

	return s.ProductRepo.Update(updateProduct, oldProduct)
}

func (s *ProductService) validateProductCodes(productInput *recipe.RecipeProductInput, productID *string) ([]models.ModelProductCode, error) {
	codes, err := s.productCodeKind.GetAll()
	if err != nil {
		return nil, err
	}

	codeMap := make(map[int]models.ModelProductCodeKind)
	for _, code := range codes {
		codeMap[code.ID] = code
	}

	_codesList := map[int]recipe.RecipeProductCode{}
	for _, codeRecipe := range productInput.CodesList {
		_codesList[codeRecipe.ID] = codeRecipe
	}

	if len(_codesList) != len(productInput.CodesList) {
		return nil, types.ThrowRecipe("No pueden haber dos codigos del mismo tipo", "codes")
	}

	for _, codeRecipe := range productInput.CodesList {
		if _, exists := codeMap[codeRecipe.ID]; !exists {
			return nil, types.ThrowRecipe("Hay códigos erroneos o no encontrados en la base de datos", "codes")
		}
	}

	existAny, err := s.ProductRepo.CheckCodeExists(productID, func() map[int]string {
		m := make(map[int]string)
		for _, codeRecipe := range productInput.CodesList {
			m[codeRecipe.ID] = codeRecipe.Value
		}
		return m
	}())

	if err != nil {
		return nil, err
	}

	_stringError := ""
	for existsId, existValue := range existAny {
		_stringError += fmt.Sprintf("El código %s : %s ya existe en el sistema ", codeMap[existsId].Name, existValue)
	}

	if _stringError != "" {
		return nil, types.ThrowRecipe(_stringError, "codes")
	}

	pCodes := make([]models.ModelProductCode, 0)
	for _, codeRecipe := range productInput.CodesList {
		if _, exists := codeMap[codeRecipe.ID]; !exists {
			return nil, types.ThrowRecipe("Hay códigos erroneos o no encontrados en la base de datos", "codes")
		}

		codeModel := models.ModelProductCode{
			Kind: models.ModelProductCodeKind{
				ID:          codeRecipe.ID,
				Name:        codeMap[codeRecipe.ID].Name,
				Description: codeMap[codeRecipe.ID].Description,
			},
			Value: codeRecipe.Value,
		}
		pCodes = append(pCodes, codeModel)
	}

	return pCodes, nil
}

func (s *ProductService) validateProductCategories(categoryIds []int64) ([]models.ModelProductCategory, error) {
	productCategories := make([]models.ModelProductCategory, len(categoryIds))

	if len(categoryIds) == 0 {
		return productCategories, nil
	}

	categories, err := s.productCategory.GetAll()
	if err != nil {
		return nil, err
	}

	for i, catId := range categoryIds {
		if catId == 0 {
			continue
		}
		found := false
		for _, category := range categories {
			if category.ID == catId {
				productCategories[i] = models.ModelProductCategory{
					ID:          category.ID,
					Name:        category.Name,
					Description: category.Description,
					Available:   category.Available,
				}
				found = true
				break
			}
		}
		if !found {
			return nil, types.ThrowRecipe("Hay categorías erróneas o no encontradas en la base de datos", "categories")
		}
	}

	return productCategories, nil
}

func (s *ProductService) CreateCategory(ctx context.Context, recipe *recipe.RecipeProductCategory) (*models.ModelProductCategory, error) {
	// Validar permiso usando PowerProductCreate
	if ok := s.EveryPower(ctx, PowerProductCreate); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear categorías")
	}

	// Preparar modelo
	category := &models.ModelProductCategory{
		Name:        recipe.Name,
		Description: recipe.Description,
		Available:   true,
	}

	// Llamar al repositorio
	return s.productCategory.Create(category)
}

func (s *ProductService) GetCategoryByName(name string) (*models.ModelProductCategory, error) {
	return s.productCategory.GetByName(name)
}

/*
func (s *ProductService) DeleteProduct(id string) error {
	return s.productRepo.Delete(id)
}
*/

func (s *ProductService) SaveImageToBucket(ctx context.Context, file *multipart.FileHeader) (*string, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Generate a unique name for the file
	url, err := s.bucket.UploadFile(ctx, f, file.Filename)
	if err != nil {
		return nil, err
	}

	return &url, nil
}
