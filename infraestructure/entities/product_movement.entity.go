package entities

import "time"

type EntityProductMovement struct {
	ID             string  `db:"id"`
	StoreProductID string  `db:"store_product_id"` // Referencia a product_per_store
	Observation    string  `db:"observation"`
	Quantity       float32 `db:"quantity"`
	UnitCost       float32 `db:"unit_cost"`
	TotalCost      float32 `db:"total_cost"`
	//ProductState   string    `db:"product_state"`
	MovedFrom         *string   `db:"moved_from"`
	MovedTo           *string   `db:"moved_to"`
	MovedAt           time.Time `db:"moved_at"`
	MovedBy           string    `db:"moved_by"`
	MovementType      string    `db:"movement_type"`
	MovementDocType   *string   `db:"movement_doc_type"`      // Tipo de documento asociado al movimiento (NULL para movimientos manuales)
	DocumentReference *string   `db:"movement_doc_reference"` // Referencia al documento asociado (NULL para movimientos manuales)
	PurchaseID        *string   `db:"purchase_id"`            // Referencia a purchase (NULL para movimientos manuales)
	InventoryUnit     *string   `db:"inventory_unit"`         // Unidad de inventario del producto al momento del movimiento
	StockBefore       *float32  `db:"stock_before"`           // Stock antes del movimiento en la bodega afectada
	StockAfter        *float32  `db:"stock_after"`            // Stock después del movimiento en la bodega afectada
}
