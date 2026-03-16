package models

// ModelProductWarehouse representa el stock de un producto en una bodega.
// Ahora referencia store_product_id (producto por tienda) en lugar de product_company_id.
type ModelProductWarehouse struct {
	ID                   string  `json:"id" db:"id"`
	StoreProductId       string  `json:"store_product_id"` // Referencia a product_per_store
	WarehouseId          string  `json:"warehouse_id"`
	WarehouseIdReference *string `json:"warehouse_id_reference"` // Bodega de referencia (origen/destino)
	Direction            *string `json:"direction"`              // Dirección del movimiento: "IN" o "OUT"
	InStock              float32 `json:"in_stock"`
	CostAvg              float32 `json:"cost_avg"` // Costo promedio del producto
	InTransit            float32 `json:"in_transit"`
	Ordered              float32 `json:"ordered"`
}
