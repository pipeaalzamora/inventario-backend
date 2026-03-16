package data

// DEPRECATED: ProductCompanyRepo ha sido reemplazado por ProductPerStoreRepo
// Este repositorio se mantiene comentado para referencia durante la migración.
// Ver: product_per_store.repo.go
/*
import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

type ProductCompanyRepo struct {
	db *sqlx.DB
}

func NewProductCompanyRepo(db *sqlx.DB) ports.PortProductCompany {
	return &ProductCompanyRepo{
		db: db,
	}
}

func (r *ProductCompanyRepo) GetProductsByCompany(companyID string) ([]models.ModelProductCompany, error) {

	query := `
		SELECT
			pc.id,
			pc.company_id,
			pc.product_id,
			pc.tag_id,
			pc.sku,
			pc.product_name,
			pc.item_purchase,
			pc.item_sale,
			pc.item_inventory,
			pc.is_frozen,
			pc.use_recipe,
			pc.cost_last,
			pc.description,
			pc.cost_estimated,
			pc.cost_avg,
			pc.minimal_stock,
			pc.maximal_stock,
			pc.minimal_order,
			pc.created_at,
			pc.updated_at,
			mu.abbreviation AS unit_inventory,
			mu2.abbreviation AS unit_purchase,
			pc.unit_matrix
		FROM product_company pc
		JOIN measurement_unit mu ON pc.unit_inventory_id = mu.id
		JOIN measurement_unit mu2 ON pc.unit_purchase_id = mu2.id
		WHERE company_id = $1
	`

	var products []entities.EntityProductCompany
	err := r.db.Select(&products, query, companyID)
	if err != nil {
		return nil, err
	}

	return r.toMapList(products), nil
}

func (r *ProductCompanyRepo) GetProductCompanyByID(productCompanyID string) (*models.ModelProductCompany, error) {

	query := `
		SELECT
			pc.id,
			pc.company_id,
			pc.product_id,
			pc.tag_id,
			pc.sku,
			pc.product_name,
			pc.item_purchase,
			pc.item_sale,
			pc.item_inventory,
			pc.is_frozen,
			pc.use_recipe,
			pc.cost_last,
			pc.description,
			pc.cost_estimated,
			pc.cost_avg,
			pc.minimal_stock,
			pc.maximal_stock,
			pc.minimal_order,
			pc.created_at,
			pc.updated_at,
			mu.id AS unit_inventory_id,
			mu.abbreviation AS unit_inventory,
			mu2.id AS unit_purchase_id,
			mu2.abbreviation AS unit_purchase,
			pc.unit_matrix
		FROM product_company pc
		JOIN measurement_unit mu ON pc.unit_inventory_id = mu.id
		JOIN measurement_unit mu2 ON pc.unit_purchase_id = mu2.id
		WHERE pc.id = $1
	`
	var product entities.EntityProductCompany
	err := r.db.Get(&product, query, productCompanyID)
	if err != nil {
		return nil, types.ThrowData("Error al obtener el producto de la empresa por ID")
	}
	return r.toModel(&product), nil
}

func (r *ProductCompanyRepo) GetAllProductRequestRestrictionByStoreId(storeId string) ([]models.ModelProductRequestRestriction, error) {
	query := `
		SELECT
			pc.id AS id,
			pc.product_name AS product_name,
			pc.description AS description,
			pr.image AS image,
			rr.max_quantity,
			mu.abbreviation AS unit_inventory
		FROM product_company pc
		JOIN measurement_unit mu ON pc.unit_inventory_id = mu.id
		JOIN store st ON st.id = $1 AND pc.company_id = st.company_id
		JOIN product pr ON pr.id = pc.product_id
		LEFT JOIN request_restriction rr
			ON rr.product_company_id = pc.id
			AND rr.store_id = st.id
	`

	var restrictions []entities.EntityProductRequestRestriction
	if err := r.db.Select(&restrictions, query, storeId); err != nil {
		return nil, types.ThrowData("Error al obtener las restricciones de solicitud de productos")
	}

	var modelsRestrictions []models.ModelProductRequestRestriction
	for _, restriction := range restrictions {
		modelsRestrictions = append(modelsRestrictions, models.ModelProductRequestRestriction{
			ID:          restriction.ID,
			Name:        restriction.Name,
			Description: restriction.Description,
			Image:       restriction.Image,
			MaxQuantity: restriction.MaxQuantity,
			Unit:        restriction.Unit,
		})
	}

	return modelsRestrictions, nil
}

func (r *ProductCompanyRepo) GetProductsRequestRetrictedByStoreId(storeId string) ([]models.ModelProductRequestRestriction, error) {
	query := `
		SELECT
			pc.id as id,
			pc.product_name as product_name,
			pc.description as description,
			pr.image as image,
			rr.max_quantity as max_quantity,
			mu.abbreviation AS unit_inventory
		FROM product_company pc
		JOIN measurement_unit mu ON pc.unit_inventory_id = mu.id
		JOIN store st ON st.id = $1 AND pc.company_id = st.company_id
		JOIN product pr ON pr.id = pc.product_id
		INNER JOIN request_restriction rr ON rr.product_company_id = pc.id AND rr.store_id = st.id
	`
	var restrictions []entities.EntityProductRequestRestriction
	if err := r.db.Select(&restrictions, query, storeId); err != nil {
		return nil, types.ThrowData("Error al obtener las restricciones de solicitud de productos")
	}

	var modelsRestrictions []models.ModelProductRequestRestriction
	for _, restriction := range restrictions {
		modelsRestrictions = append(modelsRestrictions, models.ModelProductRequestRestriction{
			ID:          restriction.ID,
			Name:        restriction.Name,
			Description: restriction.Description,
			Image:       restriction.Image,
			MaxQuantity: restriction.MaxQuantity,
			Unit:        restriction.Unit,
		})
	}

	return modelsRestrictions, nil
}

func (r *ProductCompanyRepo) toModel(entity *entities.EntityProductCompany) *models.ModelProductCompany {
	if entity == nil {
		return nil
	}

	// map every field
	return &models.ModelProductCompany{
		ID:            entity.ID,
		ProductID:     entity.ProductID,
		CompanyID:     entity.CompanyID,
		TagID:         entity.TagID,
		SKU:           entity.SKU,
		ProductName:   entity.ProductName,
		ItemPurchase:  entity.ItemPurchase,
		ItemSale:      entity.ItemSale,
		ItemInventory: entity.ItemInventory,
		IsFrozen:      entity.IsFrozen,
		UseRecipe:     entity.UseRecipe,
		UnitPurchase: models.ModelMeasurementUnit{
			ID:           entity.UnitPurchaseID,
			Abbreviation: entity.UnitPurchase,
		},
		UnitInventory: models.ModelMeasurementUnit{
			ID:           entity.UnitInventoryID,
			Abbreviation: entity.UnitInventory,
		},
		UnitMatrix:    r.toMatrixModel(entity.UnitMatrix),
		CostLast:      entity.CostLast,
		Description:   entity.Description,
		CostEstimated: entity.CostEstimated,
		CostAvg:       entity.CostAvg,
		MinimalStock:  entity.MinimalStock,
		MaximalStock:  entity.MaximalStock,
		MinimalOrder:  entity.MinimalOrder,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}
}

func (r *ProductCompanyRepo) toMapList(entities []entities.EntityProductCompany) []models.ModelProductCompany {
	modelsCollection := make([]models.ModelProductCompany, len(entities))
	for index, entity := range entities {
		modelsCollection[index] = *r.toModel(&entity)
	}
	return modelsCollection
}

func (r *ProductCompanyRepo) toMatrixModel(entities map[int]entities.EntityConversionUnit) []models.ModelMeasurementConversionUnit {
	var matrix []models.ModelMeasurementConversionUnit

	for k, v := range entities {
		unitFactor := models.ModelMeasurementConversionUnit{
			ID:           k,
			Name:         v.Name,
			Abbreviation: v.Abbreviation,
			Description:  v.Description,
			Factor:       v.Factor,
		}

		matrix = append(matrix, unitFactor)
	}

	return matrix
}
*/
