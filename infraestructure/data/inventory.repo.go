package data

import (
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"

	"github.com/lib/pq" // <--
)

type inventoryRepo struct {
	db *sqlx.DB
}

func NewInventoryRepo(db *sqlx.DB) ports.PortInventory {
	return &inventoryRepo{
		db: db,
	}
}

func (r *inventoryRepo) GetProductStoreInventoryByStoreID(storeID string, warehouseIDs []string) ([]models.ModelWarehouseProductStock, error) {
	query := `
		SELECT 
			wpp.store_product_id,
			wpp.warehouse_id,
			w.warehouse_name,
			wpp.in_stock as current_stock,
			wpp.cost_avg as avg_cost
			-- pps.cost_estimated as avg_estimated_cost
		FROM warehouse_per_product wpp
		INNER JOIN product_per_store pps ON pps.id = wpp.store_product_id
		INNER JOIN warehouse w ON w.id = wpp.warehouse_id
		WHERE pps.store_id = $1
	`

	args := []interface{}{storeID}

	// Si se especifican IDs de bodegas, filtrar por ellas
	if len(warehouseIDs) > 0 {
		query += ` AND wpp.warehouse_id = ANY($2)`
		args = append(args, pq.Array(warehouseIDs))
	}

	var entities []entities.EntityWarehouseProductStock
	if err := r.db.Select(&entities, query, args...); err != nil {
		fmt.Println("Error al obtener el inventario:", err)
		return nil, types.ThrowData("Error al obtener el inventario de la tienda")
	}

	return r.toModelList(entities), nil
}

func (r *inventoryRepo) GetProductTransitByReferences(storeID string, warehouseIDs []string) ([]models.ModelWarehouseProductTransit, error) {
	if len(warehouseIDs) == 0 {
		return []models.ModelWarehouseProductTransit{}, nil
	}

	query := `
		SELECT 
			wpp.store_product_id,
			wpp.warehouse_id_reference,
			wpp.direction,
			wpp.in_stock,
			wpp.cost_avg as avg_cost
		FROM warehouse_per_product wpp
		INNER JOIN product_per_store pps ON pps.id = wpp.store_product_id
		INNER JOIN warehouse w ON w.id = wpp.warehouse_id
		WHERE pps.store_id = $1
			AND w.is_momevent_warehouse = true
			AND wpp.warehouse_id_reference = ANY($2)
			AND wpp.direction IS NOT NULL
	`

	var transitEntities []entities.EntityWarehouseProductTransit
	if err := r.db.Select(&transitEntities, query, storeID, pq.Array(warehouseIDs)); err != nil {
		fmt.Println("Error al obtener el inventario en tránsito:", err)
		return nil, types.ThrowData("Error al obtener el inventario en tránsito")
	}

	return r.toTransitModelList(transitEntities), nil
}

func (r *inventoryRepo) GetSingleProductStock(storeID string, warehouseID string, storeProductID string) (*models.ModelWarehouseProductStock, error) {
	query := `
		SELECT 
			wpp.store_product_id,
			wpp.warehouse_id,
			w.warehouse_name,
			wpp.in_stock as current_stock,
			wpp.cost_avg as avg_cost
		FROM warehouse_per_product wpp
		INNER JOIN product_per_store pps ON pps.id = wpp.store_product_id
		INNER JOIN warehouse w ON w.id = wpp.warehouse_id
		WHERE pps.store_id = $1
			AND wpp.warehouse_id = $2
			AND wpp.store_product_id = $3
	`

	var entity entities.EntityWarehouseProductStock
	if err := r.db.Get(&entity, query, storeID, warehouseID, storeProductID); err != nil {
		// capturar el error sql: no rows in result set
		if err.Error() == "sql: no rows in result set" {
			return r.getEmptyModel(storeID, warehouseID, storeProductID), nil
		}
		return nil, types.ThrowData("Error al obtener el stock del producto")
	}

	return r.toModel(&entity), nil
}

func (r *inventoryRepo) GetSingleProductTransit(storeID string, warehouseID string, storeProductID string) ([]models.ModelWarehouseProductTransit, error) {
	query := `
		SELECT 
			wpp.store_product_id,
			wpp.warehouse_id_reference,
			wpp.direction,
			wpp.in_stock,
			wpp.cost_avg as avg_cost
		FROM warehouse_per_product wpp
		INNER JOIN product_per_store pps ON pps.id = wpp.store_product_id
		INNER JOIN warehouse w ON w.id = wpp.warehouse_id
		WHERE pps.store_id = $1
			AND w.is_momevent_warehouse = true
			AND wpp.warehouse_id_reference = $2
			AND wpp.store_product_id = $3
			AND wpp.direction IS NOT NULL
	`

	var transitEntities []entities.EntityWarehouseProductTransit
	if err := r.db.Select(&transitEntities, query, storeID, warehouseID, storeProductID); err != nil {
		fmt.Println("Error al obtener el inventario en tránsito del producto:", err)
		return nil, types.ThrowData("Error al obtener el inventario en tránsito del producto")
	}

	return r.toTransitModelList(transitEntities), nil
}

func (r *inventoryRepo) GetCurrentStock(companyId, storeId, warehouseId string) (*models.ModelWarehouseProductStock, error) {
	query := `
		SELECT 
			wpp.store_product_id,
			wpp.warehouse_id,
			w.warehouse_name,
			wpp.in_stock as current_stock,
			wpp.cost_avg as avg_cost
		FROM warehouse_per_product wpp
		INNER JOIN product_per_store pps ON pps.id = wpp.store_product_id
		INNER JOIN store s ON s.id = pps.store_id
		INNER JOIN warehouse w ON w.id = wpp.warehouse_id
		WHERE s.company_id = $1
			AND s.id = $2
			AND wpp.warehouse_id = $3
	`

	var entity entities.EntityWarehouseProductStock
	if err := r.db.Get(&entity, query, companyId, storeId, warehouseId); err != nil {
		// capturar el error sql: no rows in result set
		if err.Error() == "sql: no rows in result set" {
			return r.getEmptyModel(storeId, warehouseId, ""), nil
		}
		return nil, types.ThrowData("Error al obtener el stock actual")
	}

	return r.toModel(&entity), nil
}

func (r *inventoryRepo) getEmptyModel(storeId, warehouseId, storeProductId string) *models.ModelWarehouseProductStock {
	query := `
		SELECT
			warehouse_name
		FROM warehouse
		JOIN store ON store.id = warehouse.store_id
		JOIN product_per_store pps ON pps.store_id = store.id
		WHERE store.id = $1 AND pps.id = $2 AND warehouse.id = $3
		RETURNING warehouse_name
	`
	var warehouseName string
	if err := r.db.Get(&warehouseName, query, storeId, storeProductId, warehouseId); err != nil {
		warehouseName = ""
	}

	return &models.ModelWarehouseProductStock{
		StoreProductId: storeProductId,
		WarehouseID:    warehouseId,
		WarehouseName:  warehouseName,
		CurrentStock:   0,
		AvgCost:        0,
	}

}

func (r *inventoryRepo) toTransitModel(entity *entities.EntityWarehouseProductTransit) *models.ModelWarehouseProductTransit {
	return &models.ModelWarehouseProductTransit{
		StoreProductId:       entity.StoreProductId,
		WarehouseIDReference: entity.WarehouseIDReference,
		Direction:            entity.Direction,
		InStock:              entity.InStock,
		AvgCost:              entity.AvgCost,
	}
}

func (r *inventoryRepo) toTransitModelList(transitEntities []entities.EntityWarehouseProductTransit) []models.ModelWarehouseProductTransit {
	result := make([]models.ModelWarehouseProductTransit, len(transitEntities))
	for i, entity := range transitEntities {
		result[i] = *r.toTransitModel(&entity)
	}
	return result
}

func (r *inventoryRepo) toModel(entity *entities.EntityWarehouseProductStock) *models.ModelWarehouseProductStock {
	return &models.ModelWarehouseProductStock{
		StoreProductId: entity.StoreProductId,
		WarehouseID:    entity.WarehouseID,
		WarehouseName:  entity.WarehouseName,
		CurrentStock:   entity.CurrentStock,
		AvgCost:        entity.AvgCost,
	}
}

func (r *inventoryRepo) toModelList(entities []entities.EntityWarehouseProductStock) []models.ModelWarehouseProductStock {
	models := make([]models.ModelWarehouseProductStock, len(entities))
	for i, entity := range entities {
		models[i] = *r.toModel(&entity)
	}
	return models
}
