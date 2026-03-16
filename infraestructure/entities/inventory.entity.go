package entities

// EntityWarehouseProductStock representa el stock de un producto en una bodega específica.
// Combina datos de warehouse_per_product, product_per_store y warehouse.
type EntityWarehouseProductStock struct {
	StoreProductId string  `db:"store_product_id"`
	WarehouseID    string  `db:"warehouse_id"`
	WarehouseName  string  `db:"warehouse_name"`
	CurrentStock   float32 `db:"current_stock"`
	AvgCost        float32 `db:"avg_cost"`
}

// EntityWarehouseProductTransit representa stock en tránsito en bodegas de traspaso.
// warehouse_id_reference indica la bodega destino/origen del traspaso.
// direction puede ser 'IN' (entrada) o 'OUT' (salida).
type EntityWarehouseProductTransit struct {
	StoreProductId       string  `db:"store_product_id"`
	WarehouseIDReference string  `db:"warehouse_id_reference"`
	Direction            string  `db:"direction"`
	InStock              float32 `db:"in_stock"`
	AvgCost              float32 `db:"avg_cost"`
}
