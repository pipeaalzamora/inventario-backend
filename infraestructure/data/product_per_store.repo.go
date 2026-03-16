package data

import (
	"database/sql"
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

// ProductPerStoreRepo implementa PortProductPerStore.
// Reemplaza a ProductCompanyRepo con granularidad a nivel de tienda.
type ProductPerStoreRepo struct {
	db              *sqlx.DB
	productRepo     ports.PortProduct
	measurementRepo ports.PortMeasurement
}

// NewProductPerStoreRepo crea una nueva instancia de ProductPerStoreRepo.
func NewProductPerStoreRepo(
	db *sqlx.DB,
	productRepo ports.PortProduct,
	measurementRepo ports.PortMeasurement,
) ports.PortProductPerStore {
	return &ProductPerStoreRepo{
		db:              db,
		productRepo:     productRepo,
		measurementRepo: measurementRepo,
	}
}

// GetProductsByStore obtiene todos los productos configurados para una tienda.
func (r *ProductPerStoreRepo) GetProductsByStore(storeID string) ([]models.ModelProductPerStore, error) {
	query := `
		SELECT 
			pps.id,
			pps.store_id,
			pps.product_id,
			pps.tag_id,
			p.sku,
			pps.product_name,
			p.image,
			pps.item_sale,
			pps.use_recipe,
			pps.unit_inventory_id,
			pps.description,
			p.cost_estimated,
			pps.cost_avg,
			pps.minimal_stock,
			pps.maximal_stock,
			pps.max_quantity,
			pps.created_at,
			pps.updated_at
		FROM product_per_store pps
		JOIN product p ON pps.product_id = p.id
		WHERE pps.store_id = $1
	`

	var products []entities.EntityProductPerStore
	err := r.db.Select(&products, query, storeID)
	if err != nil {
		fmt.Println(err)
		return nil, types.ThrowData("Error al obtener los productos de la tienda")
	}

	return r.toMapList(products), nil
}

// GetProductPerStoreByID obtiene un producto específico por su ID.
func (r *ProductPerStoreRepo) GetProductPerStoreByID(storeProductID string) (*models.ModelProductPerStore, error) {
	query := `
		SELECT 
			pps.id,
			pps.store_id,
			pps.product_id,
			pps.tag_id,
			p.sku,
			pps.product_name,
			p.image,
			pps.item_sale,
			pps.use_recipe,
			pps.unit_inventory_id,
			pps.description,
			p.cost_estimated,
			pps.cost_avg,
			pps.minimal_stock,
			pps.maximal_stock,
			pps.max_quantity,
			pps.created_at,
			pps.updated_at
		FROM product_per_store pps
		JOIN product p ON pps.product_id = p.id
		WHERE pps.id = $1
	`
	var product entities.EntityProductPerStore
	err := r.db.Get(&product, query, storeProductID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.ThrowData(fmt.Sprintf("No se encontró el producto de la tienda con ID: %s", storeProductID))
		}
		return nil, types.ThrowData(fmt.Sprintf("Error al obtener el producto de la tienda por ID: %v", err))
	}
	return r.toModel(&product, true)
}

func (r *ProductPerStoreRepo) GetProductsWithSuppliersByStore(storeID string) ([]models.ModelProductPerStore, error) {
	query := `
		SELECT 
			pps.id,
			pps.store_id,
			pps.product_id,
			pps.tag_id,
			p.sku,
			pps.product_name,
			p.image,
			pps.item_sale,
			pps.use_recipe,
			pps.unit_inventory_id,
			pps.description,
			p.cost_estimated,
			pps.cost_avg,
			pps.minimal_stock,
			pps.maximal_stock,
			pps.max_quantity,
			pps.created_at,
			pps.updated_at
		FROM product_per_store pps
		JOIN product p ON pps.product_id = p.id
		WHERE pps.store_id = $1
		  AND EXISTS (
			SELECT 1 FROM supplier_per_store_product spsp
			WHERE spsp.store_product_id = pps.id
		  )
	`
	var products []entities.EntityProductPerStore
	err := r.db.Select(&products, query, storeID)
	if err != nil {
		return nil, types.ThrowData("Error al obtener los productos con proveedores de la tienda")
	}

	return r.toMapList(products), nil
}

// GetAllProductRequestRestrictionByStoreId obtiene todos los productos con información de restricción.
// Trae todos los productos independiente de si tienen restriccion o no.
func (r *ProductPerStoreRepo) GetAllProductRequestRestrictionByStoreId(storeId string) ([]models.ModelProductRequestRestriction, error) {
	/*
		query := `
			SELECT
				pps.id AS id,
				pps.product_name AS product_name,
				pps.description AS description,
				pr.image AS image,
				pps.max_quantity,
				mu.abbreviation AS unit_inventory
			FROM product_per_store pps
			JOIN measurement_unit mu ON pps.unit_inventory_id = mu.id
			JOIN product pr ON pr.id = pps.product_id
			WHERE pps.store_id = $1
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
	*/
	return nil, types.ThrowData("Funcionalidad de obtención de restricciones de solicitud no implementada")
}

// GetProductsRequestRetrictedByStoreId obtiene solo los productos que tienen restricción.
// Trae solo los productos que tienen max_quantity definido.
func (r *ProductPerStoreRepo) GetProductsRequestRetrictedByStoreId(storeId string) ([]models.ModelProductRequestRestriction, error) {

	query := `
			SELECT
				pps.id AS id,
				pps.product_name AS product_name,
				pps.description AS description,
				pr.image AS image,
				pps.max_quantity AS max_quantity,
				mu.abbreviation AS unit_inventory
			FROM product_per_store pps
			JOIN measurement_unit mu ON pps.unit_inventory_id = mu.id
			JOIN product pr ON pr.id = pps.product_id
			WHERE pps.store_id = $1 AND pps.max_quantity IS NOT NULL
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

// CreateProductPerStore crea un nuevo producto para una tienda.
func (r *ProductPerStoreRepo) CreateProductPerStore(product *models.ModelProductPerStore) (*models.ModelProductPerStore, error) {
	// Iniciar transacción
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción")
	}
	defer tx.Rollback()

	// 1. Insertar producto en product_per_store
	insertQuery := `
		INSERT INTO product_per_store (
			store_id, product_id, tag_id, product_name,
			item_sale, use_recipe, unit_inventory_id,
			description, cost_avg,
			minimal_stock, maximal_stock, max_quantity
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 0, $9, $10, $11)
		RETURNING id, created_at, updated_at
	`

	tagID := product.TagID
	if tagID == 0 {
		tagID = 1
	}

	var newID, createdAt, updatedAt string
	err = tx.QueryRow(
		insertQuery,
		product.StoreID,
		product.ProductTemplate.ID,
		tagID,
		product.ProductName,
		product.ItemSale,
		product.UseRecipe,
		product.UnitInventory.ID,
		product.Description,
		product.Quantities.MinimalStock,
		product.Quantities.MaximalStock,
		product.Quantities.MaxQuantity,
	).Scan(&newID, &createdAt, &updatedAt)

	if err != nil {
		// Capturar violación de unique constraint en SKU
		if err.Error() == `pq: duplicate key value violates unique constraint "product_per_store_store_id_sku_key"` {
			return nil, types.ThrowData("El SKU ya existe para esta tienda")
		}
		return nil, types.ThrowData("Error al crear el producto de tienda")
	}

	// 2. Insertar conversiones de unidades en product_per_store_per_measurement_unit
	if len(product.UnitMatrix) > 0 {
		insertConversionQuery := `
			INSERT INTO product_per_store_per_measurement_unit (
				store_product_id, measurement_unit_id, conversion_factor
			) VALUES ($1, $2, $3)
		`

		for _, unit := range product.UnitMatrix {
			_, err := tx.Exec(insertConversionQuery, newID, unit.ID, unit.Factor)
			if err != nil {
				return nil, types.ThrowData("Error al insertar conversiones de unidades")
			}
		}
	}

	// 3. Insertar proveedores en supplier_per_store_product
	if len(product.Suppliers) > 0 {
		insertSupplierQuery := `
			INSERT INTO supplier_per_store_product (
				store_product_id, supplier_id, priority
			) VALUES ($1, $2, $3)
		`

		for _, supplier := range product.Suppliers {
			_, err := tx.Exec(insertSupplierQuery, newID, supplier.SupplierID, supplier.Priority)
			if err != nil {
				// Capturar violación de unique constraint en priority
				if err.Error() == `pq: duplicate key value violates unique constraint "supplier_per_store_product_store_product_id_priority_key"` {
					return nil, types.ThrowData("No pueden existir dos proveedores con la misma prioridad")
				}
				return nil, types.ThrowData("Error al insertar proveedores del producto")
			}
		}
	}

	// Commit de la transacción
	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	// Obtener el producto completo con todas sus relaciones
	return r.GetProductPerStoreByID(newID)
}

// ExistsBySKU verifica si existe un producto con el SKU dado en la tienda.
func (r *ProductPerStoreRepo) ExistsBySKU(storeID, sku string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM product_per_store WHERE store_id = $1 AND sku = $2)`
	if err := r.db.QueryRow(query, storeID, sku).Scan(&exists); err != nil {
		return false, types.ThrowData("Error al verificar el SKU")
	}
	return exists, nil
}

// ExistsByProductAndStore verifica si existe un producto plantilla asignado a una tienda.
func (r *ProductPerStoreRepo) ExistsByProductAndStore(productID, storeID string, excludeID *string) (bool, error) {
	var exists bool
	var query string
	
	if excludeID != nil && *excludeID != "" {
		// Excluir el ID actual (para updates)
		query = `SELECT EXISTS(SELECT 1 FROM product_per_store WHERE product_id = $1 AND store_id = $2 AND id != $3)`
		if err := r.db.QueryRow(query, productID, storeID, *excludeID).Scan(&exists); err != nil {
			return false, types.ThrowData("Error al verificar el producto en la tienda")
		}
	} else {
		// Sin exclusión (para creates)
		query = `SELECT EXISTS(SELECT 1 FROM product_per_store WHERE product_id = $1 AND store_id = $2)`
		if err := r.db.QueryRow(query, productID, storeID).Scan(&exists); err != nil {
			return false, types.ThrowData("Error al verificar el producto en la tienda")
		}
	}
	
	return exists, nil
}

// UpdateProductPerStore actualiza un producto existente.
func (r *ProductPerStoreRepo) UpdateProductPerStore(product *models.ModelProductPerStore) (*models.ModelProductPerStore, error) {
	// Iniciar transacción
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción")
	}
	defer tx.Rollback()

	// 1. Actualizar tabla principal product_per_store
	query := `
		UPDATE product_per_store SET
			tag_id = $1,
			product_name = $2,
			item_sale = $3,
			use_recipe = $4,
			unit_inventory_id = $5,
			description = $6,
			minimal_stock = $7,
			maximal_stock = $8,
			max_quantity = $9,
			updated_at = NOW()
		WHERE id = $10
		RETURNING *
	`

	var productEntity entities.EntityProductPerStore
	err = tx.QueryRowx(
		query,
		product.TagID,
		product.ProductName,
		product.ItemSale,
		product.UseRecipe,
		product.UnitInventory.ID,
		product.Description,
		product.Quantities.MinimalStock,
		product.Quantities.MaximalStock,
		product.Quantities.MaxQuantity,
		product.ID,
	).StructScan(&productEntity)

	if err != nil {
		return nil, types.ThrowData("Error al actualizar el producto de la tienda: " + err.Error())
	}

	// 2. Actualizar conversiones de unidades (UnitMatrix)
	// Eliminar las existentes y crear las nuevas
	if len(product.UnitMatrix) > 0 {
		// Eliminar conversiones existentes
		deleteConversionsQuery := `
			DELETE FROM product_per_store_per_measurement_unit 
			WHERE store_product_id = $1
		`
		_, err := tx.Exec(deleteConversionsQuery, product.ID)
		if err != nil {
			return nil, types.ThrowData("Error al eliminar conversiones de unidades existentes")
		}

		// Insertar nuevas conversiones
		insertConversionQuery := `
			INSERT INTO product_per_store_per_measurement_unit (
				store_product_id, measurement_unit_id, conversion_factor
			) VALUES ($1, $2, $3)
		`

		for _, unit := range product.UnitMatrix {
			_, err := tx.Exec(insertConversionQuery, product.ID, unit.ID, unit.Factor)
			if err != nil {
				return nil, types.ThrowData("Error al insertar conversiones de unidades")
			}
		}
	}

	// 3. Actualizar proveedores (Suppliers) de manera eficiente
	// Solo se actualizan si se envían en el request (partial update)
	// nil = campo omitido (no tocar), [] = eliminar todos, [items] = actualizar
	if product.Suppliers != nil {
		// Caso 1: Lista vacía = eliminar todos los proveedores
		if len(product.Suppliers) == 0 {
			deleteAllSuppliersQuery := `
				DELETE FROM supplier_per_store_product 
				WHERE store_product_id = $1
			`
			_, err := tx.Exec(deleteAllSuppliersQuery, product.ID)
			if err != nil {
				return nil, types.ThrowData("Error al eliminar todos los proveedores")
			}
		} else {
			// Caso 2: Lista con elementos = UPSERT y eliminar los que no están
			// Estrategia eficiente: usar UPSERT y eliminar solo los que no están en la nueva lista

			// 3.1. Obtener IDs de proveedores actuales
			currentSupplierIDs := []string{}
			getCurrentSuppliersQuery := `
				SELECT supplier_id 
				FROM supplier_per_store_product 
				WHERE store_product_id = $1
			`
			err := tx.Select(&currentSupplierIDs, getCurrentSuppliersQuery, product.ID)
			if err != nil {
				return nil, types.ThrowData("Error al obtener proveedores actuales")
			}

			// 3.2. Construir lista de nuevos supplier IDs
			newSupplierIDs := make([]string, len(product.Suppliers))
			for i, s := range product.Suppliers {
				newSupplierIDs[i] = s.SupplierID
			}

			// 3.3. UPSERT: Insertar o actualizar proveedores
			// PostgreSQL permite ON CONFLICT para hacer UPSERT eficiente
			upsertSupplierQuery := `
				INSERT INTO supplier_per_store_product (
					store_product_id, supplier_id, priority
				) VALUES ($1, $2, $3)
				ON CONFLICT (store_product_id, supplier_id) 
				DO UPDATE SET priority = EXCLUDED.priority
			`

			for _, supplier := range product.Suppliers {
				_, err := tx.Exec(upsertSupplierQuery, product.ID, supplier.SupplierID, supplier.Priority)
				if err != nil {
					// Capturar violación de unique constraint en priority
					if err.Error() == `pq: duplicate key value violates unique constraint "supplier_per_store_product_store_product_id_priority_key"` {
						return nil, types.ThrowData("No pueden existir dos proveedores con la misma prioridad")
					}
					return nil, types.ThrowData("Error al actualizar proveedores del producto: " + err.Error())
				}
			}

			// 3.4. Eliminar proveedores que ya no están en la lista (solo los que se removieron)
			suppliersToDelete := []string{}
			for _, currentID := range currentSupplierIDs {
				found := false
				for _, newID := range newSupplierIDs {
					if currentID == newID {
						found = true
						break
					}
				}
				if !found {
					suppliersToDelete = append(suppliersToDelete, currentID)
				}
			}

			if len(suppliersToDelete) > 0 {
				deleteSupplierQuery := `
					DELETE FROM supplier_per_store_product 
					WHERE store_product_id = $1 AND supplier_id = ANY($2)
				`
				_, err := tx.Exec(deleteSupplierQuery, product.ID, suppliersToDelete)
				if err != nil {
					return nil, types.ThrowData("Error al eliminar proveedores obsoletos")
				}
			}
		}
	}

	// Commit de la transacción
	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}
	// Retornar el producto actualizado
	updatedProduct, err := r.toModel(&productEntity, true)
	return updatedProduct, err
}

// DeleteProductPerStore elimina un producto de una tienda.
func (r *ProductPerStoreRepo) DeleteProductPerStore(storeProductID string) error {
	query := `DELETE FROM product_per_store WHERE id = $1`

	result, err := r.db.Exec(query, storeProductID)
	if err != nil {
		return types.ThrowData("Error al eliminar el producto de la tienda")
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return types.ThrowData("El producto no existe en la tienda")
	}

	return nil
}

// GetProductsByStoreAndProductIDs obtiene productos por tienda y lista de IDs de producto base.
func (r *ProductPerStoreRepo) GetProductsByStoreAndProductIDs(storeID string, productIDs []string) ([]models.ModelProductPerStore, error) {
	if len(productIDs) == 0 {
		return []models.ModelProductPerStore{}, nil
	}

	query := `
		SELECT 
			pps.id,
			pps.store_id,
			pps.product_id,
			pps.tag_id,
			p.sku,
			pps.product_name,
			p.image,
			pps.item_sale,
			pps.use_recipe,
			pps.description,
			p.cost_estimated,
			pps.cost_avg,
			pps.minimal_stock,
			pps.maximal_stock,
			pps.max_quantity,
			pps.created_at,
			pps.updated_at,
			pps.unit_inventory_id,
			mu.abbreviation AS unit_inventory
		FROM product_per_store pps
		JOIN measurement_unit mu ON pps.unit_inventory_id = mu.id
		JOIN product p ON pps.product_id = p.id
		WHERE pps.store_id = $1 AND pps.product_id = ANY($2)
	`

	var products []entities.EntityProductPerStore
	err := r.db.Select(&products, query, storeID, productIDs)
	if err != nil {
		fmt.Println(err)
		return nil, types.ThrowData("Error al obtener los productos de la tienda por IDs")
	}

	return r.toMapList(products), nil
}

// getImageValue retorna el valor de la imagen, priorizando la imagen del producto plantilla.
// Retorna nil si ambas imágenes son nil.
func getImageValue(productTemplateImage *string, entityImage *string) *string {
	if entityImage != nil && *entityImage != "" {
		return entityImage
	}
	if productTemplateImage != nil {
		return productTemplateImage
	}
	return nil
}

// toModel convierte una entity a modelo.
func (r *ProductPerStoreRepo) toModel(
	entity *entities.EntityProductPerStore,
	withMatrix bool,
) (*models.ModelProductPerStore, error) {
	if entity == nil {
		return nil, types.ThrowData("No hay una entidad producto que procesar")
	}

	_tag := 0
	if entity.TagID != nil {
		_tag = *entity.TagID
	}

	_productTemplate, err := r.productRepo.GetById(entity.ProductID)
	if err != nil {
		return nil, err
	}
	// Actualizar el CostEstimated del template con el valor del JOIN
	_productTemplate.CostEstimated = entity.CostEstimated
	
	_measurementUnit, err := r.measurementRepo.GetById(entity.UnitInventoryID)
	if err != nil {
		return nil, err
	}

	_matrix := make([]models.ModelMeasurementConversionUnit, 0)
	if withMatrix {
		matrix, err := r.getMatrix(entity.ID)
		if err != nil {
			return nil, err
		}
		_matrix = matrix
	}

	// Obtener proveedores cuando withMatrix es true (usado en GetProductPerStoreByID)
	_suppliers := make([]models.ModelSupplierStoreProduct, 0)
	if withMatrix {
		suppliers, err := r.getSuppliers(entity.ID)
		if err != nil {
			return nil, err
		}
		_suppliers = suppliers
	}

	return &models.ModelProductPerStore{
		ID:              entity.ID,
		StoreID:         entity.StoreID,
		ProductTemplate: *_productTemplate,
		TagID:           _tag,
		ProductName:     entity.ProductName,
		Image:           getImageValue(_productTemplate.Image, entity.Image),
		// DEPRECATED: ItemPurchase:  entity.ItemPurchase,
		ItemSale: entity.ItemSale,
		// DEPRECATED: ItemInventory: entity.ItemInventory,
		// DEPRECATED: IsFrozen:      entity.IsFrozen,
		UseRecipe: entity.UseRecipe,
		// DEPRECATED: UnitPurchase: models.ModelMeasurementUnit{
		// 	ID:           entity.UnitPurchaseID,
		// 	Abbreviation: entity.UnitPurchase,
		// },

		UnitInventory: *_measurementUnit,
		UnitMatrix:    _matrix,

		Description: entity.Description,
		// DEPRECATED: CostLast:      entity.CostLast,
		Costs: models.ModelStoreProductCost{
			// DEPRECATED: CostEstimated ahora está en ProductTemplate
			//CostAvg:       entity.CostAvg,
		},
		Quantities: models.ModelStoreProductQuantities{
			MinimalStock: entity.MinimalStock,
			MaximalStock: entity.MaximalStock,
			MaxQuantity:  entity.MaxQuantity,
		},
		Suppliers: _suppliers,
		// DEPRECATED: MinimalOrder:  entity.MinimalOrder,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

// toMapList convierte una lista de entities a modelos.
func (r *ProductPerStoreRepo) toMapList(entitiesList []entities.EntityProductPerStore) []models.ModelProductPerStore {
	modelsCollection := make([]models.ModelProductPerStore, len(entitiesList))
	for index, entity := range entitiesList {
		_model, err := r.toModel(&entity, false)
		if err != nil {
			continue
		}
		modelsCollection[index] = *_model
	}
	return modelsCollection
}

func (r *ProductPerStoreRepo) getMatrix(productId string) ([]models.ModelMeasurementConversionUnit, error) {
	var entityList []entities.EntityMeasurementUniqueWithFactor
	query := `
		SELECT 
			mu.id, 
			mu.unit_name, 
			mu.abbreviation, 
			mu.description, 
			mu.basic_unit, 
			ppsm.conversion_factor
		FROM product_per_store_per_measurement_unit ppsm
		JOIN measurement_unit mu ON ppsm.measurement_unit_id = mu.id
		WHERE ppsm.store_product_id = $1
	`
	err := r.db.Select(&entityList, query, productId)
	if err != nil {
		return nil, types.ThrowData("ocurrió un error al traer la matrix de unidades del producto por tienda")
	}

	matrix := make([]models.ModelMeasurementConversionUnit, len(entityList))
	for i, entity := range entityList {
		matrix[i] = models.ModelMeasurementConversionUnit{
			ID:           entity.ID,
			Name:         entity.Name,
			Abbreviation: entity.Abbreviation,
			Description:  entity.Description,
			Factor:       float32(entity.ConversionFactor),
		}
	}

	return matrix, nil
}

// getSuppliers obtiene los proveedores asignados a un producto de tienda.
func (r *ProductPerStoreRepo) getSuppliers(storeProductID string) ([]models.ModelSupplierStoreProduct, error) {
	type supplierRow struct {
		ID           string `db:"id"`
		SupplierID   string `db:"supplier_id"`
		SupplierName string `db:"supplier_name"`
		ProductID    string `db:"product_id"`
		ProductName  string `db:"product_name"`
		Priority     int    `db:"priority"`
	}

	query := `
		SELECT 
			spp.id,
			spp.supplier_id,
			s.supplier_name,
			pps.product_id,
			pps.product_name,
			spp.priority
		FROM supplier_per_store_product spp
		JOIN supplier s ON s.id = spp.supplier_id
		JOIN product_per_store pps ON pps.id = spp.store_product_id
		WHERE spp.store_product_id = $1
		ORDER BY spp.priority ASC
	`

	var rows []supplierRow
	err := r.db.Select(&rows, query, storeProductID)
	if err != nil {
		return nil, types.ThrowData("Error al obtener los proveedores del producto de tienda")
	}

	suppliers := make([]models.ModelSupplierStoreProduct, len(rows))
	for i, row := range rows {
		suppliers[i] = models.ModelSupplierStoreProduct{
			ID:           row.ID,
			SupplierID:   row.SupplierID,
			SupplierName: row.SupplierName,
			ProductID:    row.ProductID,
			ProductName:  row.ProductName,
			Priority:     row.Priority,
		}
	}

	return suppliers, nil
}

// toMatrixModel convierte la matrix de unidades de entity a modelo.
/*
func (r *ProductPerStoreRepo) toMatrixModel(entitiesMap map[int]entities.EntityConversionUnit) []models.ModelMeasurementConversionUnit {
	var matrix []models.ModelMeasurementConversionUnit

	for k, v := range entitiesMap {
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

// toEntityMatrix convierte la matrix de modelos a entity para guardar en DB.
/*
func (r *ProductPerStoreRepo) toEntityMatrix(matrix []models.ModelMeasurementConversionUnit) entities.EntityUnitMatrix {
	entityMatrix := make(entities.EntityUnitMatrix)

	for _, unit := range matrix {
		entityMatrix[unit.ID] = entities.EntityConversionUnit{
			Name:         unit.Name,
			Abbreviation: unit.Abbreviation,
			Description:  unit.Description,
			Factor:       unit.Factor,
		}
	}

	return entityMatrix
}
*/
func (r *ProductPerStoreRepo) handleSQLNotFoundError(err error) bool {
	return err != nil && err.Error() == "sql: no rows in result set"
}
