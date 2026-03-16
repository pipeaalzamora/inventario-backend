package data

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type SupplierProductRepo struct {
	db *sqlx.DB
}

func NewSupplierProductRepo(db *sqlx.DB) ports.PortSupplierProduct {
	return &SupplierProductRepo{
		db: db,
	}
}

func (r *SupplierProductRepo) GetSupplierProductsByStoreIDAndSupplierIDWithProductCompanyIDs(storeID string, supplierID string, productCompanyIDs []string) ([]models.ModelSupplierProductLegacy, error) {
	// Implementation for fetching supplier products by store ID, supplier ID, and product company IDs
	// Uses the new schema: supplier_per_store_product + product_per_store + supplier_product
	var supplierProducts []entities.EntitySupplierProductLegacy
	query := `
		SELECT
			pps.product_id AS product_id,
			p.product_name AS product_name,
			p.description AS product_description,
			p.image AS product_image,
			st.company_id AS company_id,
			st.id AS store_id,
			spps.supplier_id AS supplier_id,
			pps.id AS product_company_id,
			p.sku AS product_company_sku,
			sp.id AS supplier_product_id,
			sp.unit_price AS supplier_product_price,
			1 AS supplier_product_min_quantity,
			mu.abbreviation AS unit_purchase
		FROM supplier_per_store_product spps
		JOIN product_per_store pps ON pps.id = spps.store_product_id
		JOIN store st ON st.id = pps.store_id
		JOIN product p ON p.id = pps.product_id
		JOIN supplier_product sp 
			ON sp.supplier_id = spps.supplier_id 
			AND sp.product_id = pps.product_id
		JOIN measurement_unit mu ON mu.id = sp.purchase_unit_id
		WHERE pps.store_id = $1
		  AND spps.supplier_id = $2
		  AND pps.id = ANY($3)
	`

	err := r.db.Select(&supplierProducts, query, storeID, supplierID, pq.Array(productCompanyIDs))
	if err != nil {
		return nil, types.ThrowData("Error al obtener los productos del proveedor")
	}

	// Convert to domain models
	result := make([]models.ModelSupplierProductLegacy, len(supplierProducts))
	for i, sp := range supplierProducts {
		result[i] = *r.toModel(&sp)
	}

	return result, nil
}

func (r *SupplierProductRepo) GetSupplierProductsByStoreIDAndSupplierID(storeID string, supplierID string) ([]models.ModelSupplierProductLegacy, error) {
	// Fetch supplier products by store ID and supplier ID using the new schema
	var supplierProducts []entities.EntitySupplierProductLegacy
	query := `
		SELECT
			pps.product_id AS product_id,
			p.product_name AS product_name,
			p.description AS product_description,
			p.image AS product_image,
			st.company_id AS company_id,
			st.id AS store_id,
			sp.supplier_id AS supplier_id,
			pps.id AS product_company_id,
			p.sku AS product_company_sku,
			sp.id AS supplier_product_id,
			sp.unit_price AS supplier_product_price,
			1 AS supplier_product_min_quantity,
			mu.abbreviation AS unit_purchase
		FROM supplier_per_store_product spps
		JOIN product_per_store pps ON pps.id = spps.store_product_id
		JOIN store st ON st.id = pps.store_id
		JOIN product p ON p.id = pps.product_id
		JOIN supplier_product sp 
			ON sp.supplier_id = spps.supplier_id 
			AND sp.product_id = pps.product_id
		JOIN measurement_unit mu ON mu.id = sp.purchase_unit_id
		WHERE pps.store_id = $1
		  AND spps.supplier_id = $2
	`

	err := r.db.Select(&supplierProducts, query, storeID, supplierID)
	if err != nil {
		return nil, types.ThrowData("Error al obtener los productos del proveedor: " + err.Error())
	}

	// Convert to domain models
	result := make([]models.ModelSupplierProductLegacy, len(supplierProducts))
	for i, sp := range supplierProducts {
		result[i] = *r.toModel(&sp)
	}

	return result, nil
}

func (r *SupplierProductRepo) toModel(entity *entities.EntitySupplierProductLegacy) *models.ModelSupplierProductLegacy {
	if entity == nil {
		return nil
	}
	return &models.ModelSupplierProductLegacy{
		ProductID:                  entity.ProductID,
		ProductName:                entity.ProductName,
		ProductDescription:         entity.ProductDescription,
		ProductImage:               entity.ProductImage,
		CompanyID:                  entity.CompanyID,
		StoreID:                    entity.StoreID,
		SupplierID:                 entity.SupplierID,
		ProductCompanyID:           entity.ProductCompanyID,
		ProductCompanySKU:          entity.ProductCompanySKU,
		SupplierProductID:          entity.SupplierProductID,
		SupplierProductPrice:       entity.SupplierProductPrice,
		SupplierProductMinQuantity: entity.SupplierProductMinQuantity,
		UnitPurchase:               entity.UnitPurchase,
	}
}
