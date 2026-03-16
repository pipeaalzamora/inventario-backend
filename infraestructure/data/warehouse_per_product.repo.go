package data

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

// WarehouseProductRepo implementa PortWarehouseProduct.
// Ahora usa store_product_id (producto por tienda) en lugar de product_company_id.
type WarehouseProductRepo struct {
	db *sqlx.DB
}

func NewWarehouseProductRepo(db *sqlx.DB) ports.PortWarehouseProduct {
	return &WarehouseProductRepo{
		db: db,
	}
}

// GetAll implements ports.PortWarehouseProduct.
func (p *WarehouseProductRepo) GetAll() ([]models.ModelProductWarehouse, error) {
	var productWarehouses = make([]entities.EntityWarehousePerProduct, 0)
	query := `
		SELECT wpp.id, wpp.store_product_id, wpp.warehouse_id, wpp.in_stock,
			   COALESCE(0, 0) as in_transit, COALESCE(0, 0) as ordered,
			   pps.store_id
		FROM warehouse_per_product wpp
		JOIN product_per_store pps ON wpp.store_product_id = pps.id
	`
	if err := p.db.Select(&productWarehouses, query); err != nil {
		return nil, types.ThrowData("error al obtener los productos de bodega")
	}

	return p.toModelList(productWarehouses), nil
}

// GetAllByWarehouse implements ports.PortWarehouseProduct.
func (p *WarehouseProductRepo) GetAllByWarehouse(warehouseId string) ([]models.ModelProductWarehouse, error) {
	var productWarehouses = make([]entities.EntityWarehousePerProduct, 0)
	query := `
		SELECT wpp.id, wpp.store_product_id, wpp.warehouse_id, wpp.in_stock,
			   COALESCE(0, 0) as in_transit, COALESCE(0, 0) as ordered,
			   pps.store_id
		FROM warehouse_per_product wpp
		JOIN product_per_store pps ON wpp.store_product_id = pps.id
		WHERE wpp.warehouse_id = $1
	`
	if err := p.db.Select(&productWarehouses, query, warehouseId); err != nil {
		return nil, types.ThrowData("error al obtener los productos de la bodega")
	}

	return p.toModelList(productWarehouses), nil
}

// GetAllByStoreId obtiene todos los productos de bodega por tienda.
// Reemplaza GetAllByCompanyId.
func (p *WarehouseProductRepo) GetAllByStoreId(storeId string) ([]models.ModelProductWarehouse, error) {
	var productWarehouses = make([]entities.EntityWarehousePerProduct, 0)
	query := `
		SELECT wpp.id, wpp.store_product_id, wpp.warehouse_id, wpp.in_stock,
			   COALESCE(0, 0) as in_transit, COALESCE(0, 0) as ordered,
			   pps.store_id
		FROM warehouse_per_product wpp
		JOIN product_per_store pps ON wpp.store_product_id = pps.id
		WHERE pps.store_id = $1
	`
	if err := p.db.Select(&productWarehouses, query, storeId); err != nil {
		return nil, types.ThrowData("error al obtener los productos de la tienda")
	}

	return p.toModelList(productWarehouses), nil
}

// GetById implements ports.PortWarehouseProduct.
func (p *WarehouseProductRepo) GetById(id string) (*models.ModelProductWarehouse, error) {
	var productWarehouse entities.EntityWarehousePerProduct
	query := `
		SELECT wpp.id, wpp.store_product_id, wpp.warehouse_id, wpp.in_stock,
			   COALESCE(0, 0) as in_transit, COALESCE(0, 0) as ordered,
			   pps.store_id
		FROM warehouse_per_product wpp
		JOIN product_per_store pps ON wpp.store_product_id = pps.id
		WHERE wpp.id = $1
	`
	if err := p.db.Get(&productWarehouse, query, id); err != nil {
		return nil, types.ThrowData("error al obtener el producto de bodega")
	}
	return p.toModel(productWarehouse), nil
}

// GetByProductAndWarehouseId obtiene el registro de stock por producto y bodega.
func (p *WarehouseProductRepo) GetByProductAndWarehouseId(storeProductId, warehouseId string) (*models.ModelProductWarehouse, error) {
	var productWarehouse entities.EntityWarehousePerProduct
	query := `
		SELECT wpp.id, wpp.store_product_id, wpp.warehouse_id, wpp.warehouse_id_reference, 
			   wpp.direction, wpp.in_stock, wpp.cost_avg,
			   COALESCE(0, 0) as in_transit, COALESCE(0, 0) as ordered,
			   pps.store_id
		FROM warehouse_per_product wpp
		JOIN product_per_store pps ON wpp.store_product_id = pps.id
		WHERE wpp.store_product_id = $1 AND wpp.warehouse_id = $2
	`
	if err := p.db.Get(&productWarehouse, query, storeProductId, warehouseId); err != nil {
		return nil, types.ThrowData("error al obtener el producto en la bodega especificada")
	}
	return p.toModel(productWarehouse), nil
}

// CreateWPP implements ports.PortWarehouseProduct.
func (p *WarehouseProductRepo) CreateWPP(productWarehouse *models.ModelProductWarehouse) (*models.ModelProductWarehouse, error) {
	query := `
		INSERT INTO warehouse_per_product (
			store_product_id,
			warehouse_id,
			warehouse_id_reference,
			direction,
			in_stock,
			cost_avg
		) VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	var id string
	if err := p.db.QueryRow(query,
		productWarehouse.StoreProductId,
		productWarehouse.WarehouseId,
		productWarehouse.WarehouseIdReference,
		productWarehouse.Direction,
		productWarehouse.InStock,
		productWarehouse.CostAvg,
	).Scan(&id); err != nil {
		return nil, types.ThrowData("error al crear el producto en bodega: " + err.Error())
	}

	productWarehouse.ID = id
	return productWarehouse, nil
}

// UpdateWPP implements ports.PortWarehouseProduct.
func (p *WarehouseProductRepo) UpdateWPP(productWarehouse *models.ModelProductWarehouse) (*models.ModelProductWarehouse, error) {
	query := `
		UPDATE warehouse_per_product
		SET in_stock = $1
		WHERE id = $2
		RETURNING id
	`
	var id string
	if err := p.db.QueryRow(query,
		productWarehouse.InStock,
		productWarehouse.ID,
	).Scan(&id); err != nil {
		return nil, types.ThrowData("error al actualizar el producto en bodega")
	}

	return productWarehouse, nil
}

// MakeProductTransfer implements ports.PortWarehouseProduct.
// storeProductId es el ID de product_per_store.
func (p *WarehouseProductRepo) MakeProductTransfer(storeProductId string, warehouseFrom, warehouseTo *models.ModelProductWarehouse) error {
	tx := p.db.MustBegin()

	fromQuery := `
		UPDATE warehouse_per_product
		SET in_stock = $1
		WHERE store_product_id = $2 AND warehouse_id = $3
	`

	if _, err := tx.Exec(fromQuery,
		warehouseFrom.InStock,
		storeProductId,
		warehouseFrom.WarehouseId,
	); err != nil {
		tx.Rollback()
		return types.ThrowData("error al actualizar el stock de la bodega origen")
	}

	toQuery := `
		UPDATE warehouse_per_product
		SET in_stock = $1
		WHERE store_product_id = $2 AND warehouse_id = $3
	`

	if _, err := tx.Exec(toQuery,
		warehouseTo.InStock,
		storeProductId,
		warehouseTo.WarehouseId,
	); err != nil {
		tx.Rollback()
		return types.ThrowData("error al actualizar el stock de la bodega destino")
	}

	if err := tx.Commit(); err != nil {
		return types.ThrowData("error al confirmar la transacción")
	}

	return nil

}

// MakeProductInput implements ports.PortWarehouseProduct.
// storeProductId es el ID de product_per_store.
func (p *WarehouseProductRepo) MakeProductInput(storeProductId string, warehouse *models.ModelProductWarehouse) error {
	query := `
		UPDATE warehouse_per_product
		SET in_stock = $1
		WHERE store_product_id = $2 AND warehouse_id = $3
	`
	_, err := p.db.Exec(query,
		warehouse.InStock,
		storeProductId,
		warehouse.WarehouseId,
	)
	if err != nil {
		return types.ThrowData("error al actualizar el stock en el ingreso de producto")
	}

	return nil
}

// MakeProductWaste implements ports.PortWarehouseProduct.
// storeProductId es el ID de product_per_store.
func (p *WarehouseProductRepo) MakeProductWaste(storeProductId string, warehouse *models.ModelProductWarehouse) error {
	query := `
		UPDATE warehouse_per_product
		SET in_stock = $1
		WHERE store_product_id = $2 AND warehouse_id = $3
	`
	_, err := p.db.Exec(query,
		warehouse.InStock,
		storeProductId,
		warehouse.WarehouseId,
	)
	if err != nil {
		return types.ThrowData("error al actualizar el stock en el desperdicio de producto")
	}

	return nil
}

func (p *WarehouseProductRepo) toModel(productWarehouse entities.EntityWarehousePerProduct) *models.ModelProductWarehouse {
	return &models.ModelProductWarehouse{
		ID:                   productWarehouse.ID,
		StoreProductId:       productWarehouse.StoreProductId,
		WarehouseId:          productWarehouse.WarehouseId,
		WarehouseIdReference: productWarehouse.WarehouseIdReference,
		Direction:            productWarehouse.Direction,
		InStock:              productWarehouse.InStock,
		CostAvg:              productWarehouse.CostAvg,
		InTransit:            productWarehouse.InTransit,
	}
}

func (p *WarehouseProductRepo) toModelList(entitiesList []entities.EntityWarehousePerProduct) []models.ModelProductWarehouse {
	modelsList := make([]models.ModelProductWarehouse, len(entitiesList))
	for i, entity := range entitiesList {
		modelsList[i] = *p.toModel(entity)
	}
	return modelsList
}
