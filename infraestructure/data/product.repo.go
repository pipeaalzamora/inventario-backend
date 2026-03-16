package data

import (
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/shared"
	"sofia-backend/types"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB

	categoryPort    ports.PortCategory
	productCodePort ports.PortProductCode
}

func NewProductRepo(
	db *sqlx.DB,

	categoryPort ports.PortCategory,
	productCodePort ports.PortProductCode,

) ports.PortProduct {
	return &productRepo{
		db:              db,
		categoryPort:    categoryPort,
		productCodePort: productCodePort,
	}
}

func (p *productRepo) GetById(id string) (*models.ModelProduct, error) {
	var productEntity entities.EntityProduct

	query := "SELECT * FROM product WHERE id = $1"
	if err := p.db.Get(&productEntity, query, id); err != nil {
		return nil, err
	}
	return p.toModelFull(&productEntity), nil
}

func (p *productRepo) GetByCategory(categoryID int) ([]models.ModelProduct, error) {
	var productEntities []entities.EntityProduct

	query := `
		SELECT product.* FROM product
		JOIN product_per_category ON product_per_category.product_id = product.id
		WHERE product_per_category.category_id = $1
	`
	if err := p.db.Select(&productEntities, query, categoryID); err != nil {
		return nil, err
	}
	return p.toModelList(productEntities), nil
}

func (p *productRepo) GetAll() ([]models.ModelProduct, error) {
	var productEntities []entities.EntityProduct

	query := "SELECT * FROM product"
	if err := p.db.Select(&productEntities, query); err != nil {
		return nil, types.ThrowData("Error al obtener los productos")
	}

	return p.toModelList(productEntities), nil
}

func (p *productRepo) GetAllFull() ([]models.ModelProduct, error) {
	var productEntities []entities.EntityProduct

	query := "SELECT * FROM product"
	if err := p.db.Select(&productEntities, query); err != nil {
		return nil, types.ThrowData("Error al obtener los productos")
	}

	return p.toModelFullList(productEntities), nil
}

func (p *productRepo) Create(product *models.ModelProduct) (*models.ModelProduct, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Ocurrio un error al iniciar la creación del producto")
	}
	defer tx.Rollback()

	// Generar SKU a partir del primer código
	sku := ""
	if len(product.Codes) > 0 {
		sku = "SKU-" + product.Codes[0].Value
	}

	query := `
		INSERT INTO product (product_name, description, image, cost_estimated, sku)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING *
	`
	var productEntity entities.EntityProduct
	if err := tx.QueryRowx(
		query,
		product.Name,
		product.Description,
		product.Image,
		product.CostEstimated,
		sku,
	).StructScan(&productEntity); err != nil {
		return nil, types.ThrowData("Error al crear nuevo producto")
	}

	queryCodes := `
		INSERT INTO product_code (product_id, kind_id, code_value)
		VALUES ($1, $2, $3)
	`

	for _, code := range product.Codes {
		_, err := tx.Exec(
			queryCodes,
			productEntity.ID,
			code.Kind.ID,
			code.Value,
		)
		if err != nil {
			return nil, types.ThrowData("Error asignando códigos al producto")
		}
	}

	queryCategories := `
		INSERT INTO product_per_category (product_id, category_id)
		VALUES ($1, $2)
	`
	for _, category := range product.Categories {
		_, err := tx.Exec(
			queryCategories,
			productEntity.ID,
			category.ID,
		)
		if err != nil {
			return nil, types.ThrowData("Error asignando categorías al producto")
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return p.toModelFull(&productEntity), nil
}

func (p *productRepo) Update(product *models.ModelProduct, oldProduct *models.ModelProduct) (*models.ModelProduct, error) {
	tx, err := p.db.Beginx()
	if err != nil {
		return nil, types.ThrowData(fmt.Sprintf("Error iniciando transacción: %v", err))
	}
	defer tx.Rollback()

	// 1. Actualizar datos base del producto
	queryUpdate := `
		UPDATE product
		SET
			product_name   = $1,
			description    = $2,
			image          = $3,
			cost_estimated = $4,
			updated_at     = NOW()
		WHERE id = $5
		RETURNING *
	`

	var productEntity entities.EntityProduct
	if err := tx.QueryRowx(
		queryUpdate,
		product.Name,
		product.Description,
		product.Image,
		product.CostEstimated,
		product.ID,
	).StructScan(&productEntity); err != nil {
		return nil, types.ThrowData("Error al actualizar producto")
	}

	//Aca debemos actualizar los codigos por diferencia
	//actualizamos codigos por diferencia
	arrayStringCodesNew := make([]string, 0)
	for _, code := range product.Codes {
		arrayStringCodesNew = append(arrayStringCodesNew, strconv.Itoa(code.Kind.ID))
	}

	arrayStringCodesOld := make([]string, 0)
	for _, code := range oldProduct.Codes {
		arrayStringCodesOld = append(arrayStringCodesOld, strconv.Itoa(code.Kind.ID))
	}

	sliceUtil := shared.NewSliceUtils()
	codesToAdd := sliceUtil.DifferenceString(arrayStringCodesNew, arrayStringCodesOld)
	codesToRemove := sliceUtil.DifferenceString(arrayStringCodesOld, arrayStringCodesNew)

	for _, kindID := range codesToRemove {
		_, err := tx.Exec(
			`DELETE FROM product_code WHERE product_id = $1 AND kind_id = $2`,
			productEntity.ID,
			kindID,
		)
		if err != nil {
			return nil, types.ThrowData("Error al eliminar códigos del producto")
		}
	}

	for _, kindID := range codesToAdd {

		var value string
		for _, code := range product.Codes {
			if strconv.Itoa(code.Kind.ID) == kindID {
				value = code.Value
			}
		}

		_, err := tx.Exec(
			`INSERT INTO product_code (product_id, kind_id, code_value) VALUES ($1, $2, $3)`,
			productEntity.ID,
			kindID,
			value,
		)
		if err != nil {
			return nil, types.ThrowData("Error al agregar códigos al producto")
		}
	}

	//actualizamos categorias por diferencia
	arrayStringCategoriesNew := make([]string, 0)
	for _, category := range product.Categories {
		arrayStringCategoriesNew = append(arrayStringCategoriesNew, strconv.FormatInt(category.ID, 10))
	}
	arrayStringCategoriesOld := make([]string, 0)
	for _, category := range oldProduct.Categories {
		arrayStringCategoriesOld = append(arrayStringCategoriesOld, strconv.FormatInt(category.ID, 10))
	}

	categoriesToAdd := sliceUtil.DifferenceString(arrayStringCategoriesNew, arrayStringCategoriesOld)
	categoriesToRemove := sliceUtil.DifferenceString(arrayStringCategoriesOld, arrayStringCategoriesNew)

	for _, categoryID := range categoriesToRemove {
		_, err := tx.Exec(
			`DELETE FROM product_per_category WHERE product_id = $1 AND category_id = $2`,
			productEntity.ID,
			categoryID,
		)
		if err != nil {
			return nil, types.ThrowData("Error al eliminar categorías del producto")
		}
	}

	for _, categoryID := range categoriesToAdd {
		_, err := tx.Exec(
			`INSERT INTO product_per_category (product_id, category_id) VALUES ($1, $2)`,
			productEntity.ID,
			categoryID,
		)
		if err != nil {
			return nil, types.ThrowData("Error al agregar categorías al producto")
		}
	}

	// 6. Confirmar transacción
	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return p.toModelFull(&productEntity), nil
}

func (p *productRepo) Delete(id string) error {
	res, err := p.db.Exec(
		`DELETE FROM product WHERE id = $1`,
		id,
	)
	if err != nil {
		return types.ThrowData("Error al eliminar producto")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return types.ThrowData("Error al verificar eliminación del producto")
	}
	if rows == 0 {
		return types.ThrowData("No se encontró el producto a eliminar")
	}

	return nil
}

// ////////////////////// Private Methods /////////////////////////////
func (p *productRepo) toModelFull(entity *entities.EntityProduct) *models.ModelProduct {
	cats, _ := p.categoryPort.GetAllByProductId(entity.ID)
	codes, _ := p.productCodePort.GetAllByProductId(entity.ID)

	return &models.ModelProduct{
		ID:            entity.ID,
		Name:          entity.Name,
		SKU:           entity.SKU,
		CostEstimated: entity.CostEstimated,
		Description:   entity.Description,
		Image:         entity.Image,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
		Categories:    cats,
		Codes:         codes,
	}
}

func (p *productRepo) toModel(entity *entities.EntityProduct) *models.ModelProduct {
	return &models.ModelProduct{
		ID:            entity.ID,
		Name:          entity.Name,
		SKU:           entity.SKU,
		CostEstimated: entity.CostEstimated,
		Description:   entity.Description,
		Image:         entity.Image,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
		Categories:    make([]models.ModelProductCategory, 0),
		Codes:         make([]models.ModelProductCode, 0),
	}
}

func (p *productRepo) toModelList(entities []entities.EntityProduct) []models.ModelProduct {
	models := make([]models.ModelProduct, len(entities))
	for i, e := range entities {
		models[i] = *p.toModel(&e)
	}
	return models
}

func (p *productRepo) toModelFullList(entities []entities.EntityProduct) []models.ModelProduct {
	models := make([]models.ModelProduct, len(entities))
	for i, e := range entities {
		models[i] = *p.toModelFull(&e)
	}
	return models
}

func (p *productRepo) CheckCodeExists(productID *string, codes map[int]string) (map[int]string, error) {
	if len(codes) == 0 {
		return make(map[int]string), nil
	}

	tx, err := p.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	duplicates := make(map[int]string)
	baseQuery := "SELECT 1 FROM product_code WHERE kind_id = $1 AND code_value = $2"
	queryWithProduct := baseQuery + " AND product_id <> $3 LIMIT 1"
	queryWithoutProduct := baseQuery + " LIMIT 1"

	for kindID, codeValue := range codes {
		var exists int
		if productID != nil {
			err = tx.Get(&exists, queryWithProduct, kindID, codeValue, *productID)
		} else {
			err = tx.Get(&exists, queryWithoutProduct, kindID, codeValue)
		}
		if err == nil {
			// El código existe, agregarlo al map de duplicados
			duplicates[kindID] = codeValue
		}
		// Si hay error (no existe), simplemente continuamos con el siguiente
	}

	return duplicates, nil
}
