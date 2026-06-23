package services

import (
	"context"
	"mime/multipart"
	"testing"

	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
)

type fakeProductRepo struct {
	created *models.ModelProduct
}

func (r *fakeProductRepo) GetById(id string) (*models.ModelProduct, error) {
	return nil, nil
}

func (r *fakeProductRepo) GetAll() ([]models.ModelProduct, error) {
	return nil, nil
}

func (r *fakeProductRepo) GetAllFull() ([]models.ModelProduct, error) {
	return nil, nil
}

func (r *fakeProductRepo) GetByCategory(categoryID int) ([]models.ModelProduct, error) {
	return nil, nil
}

func (r *fakeProductRepo) Create(product *models.ModelProduct) (*models.ModelProduct, error) {
	r.created = product
	if product.ID == "" {
		product.ID = "product-1"
	}
	return product, nil
}

func (r *fakeProductRepo) Update(product *models.ModelProduct, oldProduct *models.ModelProduct) (*models.ModelProduct, error) {
	return product, nil
}

func (r *fakeProductRepo) Delete(id string) error {
	return nil
}

func (r *fakeProductRepo) CheckCodeExists(productID *string, codes map[int]string) (map[int]string, error) {
	return map[int]string{}, nil
}

type fakeProductCodeRepo struct {
	codes []models.ModelProductCodeKind
}

func (r *fakeProductCodeRepo) GetAll() ([]models.ModelProductCodeKind, error) {
	return r.codes, nil
}

func (r *fakeProductCodeRepo) GetAllByProductId(productId string) ([]models.ModelProductCode, error) {
	return nil, nil
}

type fakeProductCategoryRepo struct {
	categories []models.ModelProductCategory
}

func (r *fakeProductCategoryRepo) GetById(id string) (*models.ModelProductCategory, error) {
	return nil, nil
}

func (r *fakeProductCategoryRepo) GetByName(name string) (*models.ModelProductCategory, error) {
	return nil, nil
}

func (r *fakeProductCategoryRepo) GetAll() ([]models.ModelProductCategory, error) {
	return r.categories, nil
}

func (r *fakeProductCategoryRepo) Create(category *models.ModelProductCategory) (*models.ModelProductCategory, error) {
	return category, nil
}

func (r *fakeProductCategoryRepo) Update(category *models.ModelProductCategory) (*models.ModelProductCategory, error) {
	return category, nil
}

func (r *fakeProductCategoryRepo) EnableDisable(category *models.ModelProductCategory, value bool) (*models.ModelProductCategory, error) {
	category.Available = value
	return category, nil
}

func (r *fakeProductCategoryRepo) GetAllByProductId(productId string) ([]models.ModelProductCategory, error) {
	return nil, nil
}

type fakeProductBucket struct{}

func (b *fakeProductBucket) UploadFile(ctx context.Context, file multipart.File, fileName string) (string, error) {
	return "https://example.com/product.png", nil
}

func (b *fakeProductBucket) DeleteFile(ctx context.Context, fileURL string) error {
	return nil
}

func newProductInput() *recipe.RecipeProductInput {
	description := "Producto de prueba"

	return &recipe.RecipeProductInput{
		Name:          "Producto Test",
		Description:   &description,
		CostEstimated: 1250,
		CodesList: []recipe.RecipeProductCode{
			{ID: 1, Value: "SKU-001"},
		},
	}
}

func newProductService(repo *fakeProductRepo) *ProductService {
	return NewProductService(
		repo,
		&fakeProductCodeRepo{
			codes: []models.ModelProductCodeKind{
				{ID: 1, Name: "SKU", Description: "Codigo interno"},
			},
		},
		&fakeProductBucket{},
		&fakeProductCategoryRepo{},
	)
}

func TestProductServiceCreateProductRequiresCreatePower(t *testing.T) {
	productRepo := &fakeProductRepo{}
	service := newProductService(productRepo)
	ctx := serviceTestContext(PowerProductUpdate)

	created, err := service.CreateProduct(ctx, newProductInput())
	if err == nil {
		t.Fatal("CreateProduct returned nil error without product:create power")
	}
	if created != nil {
		t.Fatalf("CreateProduct returned product without permission: %+v", created)
	}
	if productRepo.created != nil {
		t.Fatalf("CreateProduct persisted product without permission: %+v", productRepo.created)
	}
}

func TestProductServiceCreateProductAllowsCreatePower(t *testing.T) {
	productRepo := &fakeProductRepo{}
	service := newProductService(productRepo)
	ctx := serviceTestContext(PowerProductCreate)

	created, err := service.CreateProduct(ctx, newProductInput())
	if err != nil {
		t.Fatalf("CreateProduct returned error: %v", err)
	}
	if created == nil {
		t.Fatal("CreateProduct returned nil product")
	}
	if productRepo.created == nil {
		t.Fatal("CreateProduct did not persist the product")
	}
	if created.Name != "Producto Test" {
		t.Fatalf("CreateProduct name = %q, want %q", created.Name, "Producto Test")
	}
	if len(created.Codes) != 1 || created.Codes[0].Value != "SKU-001" {
		t.Fatalf("CreateProduct codes = %+v, want one SKU-001 code", created.Codes)
	}
}
