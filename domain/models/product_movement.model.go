package models

import "time"

type ModelProductMovement struct {
	ID             string  `json:"id"`
	StoreProductID string  `json:"storeProductId"` // Referencia a product_per_store
	Observation    string  `json:"observation"`
	Quantity       float32 `json:"quantity"`
	InventoryUnit  *string `json:"inventoryUnit"`
	UnitCost       float32 `json:"unitCost"`
	TotalCost      float32 `json:"totalCost"`
	//ProductState   string    `json:"productState"`
	MovedFrom         *string   `json:"movedFrom"`
	MovedTo           *string   `json:"movedTo"`
	MovedAt           time.Time `json:"movedAt"`
	MovementType      string    `json:"movementType"`
	MovementDocType   *string   `json:"movementDocType"`
	DocumentReference *string   `json:"documentReference"`
	MovedBy           string    `json:"movedBy"`
	PurchaseID        *string   `json:"purchaseId"`  // Referencia a purchase (NULL para movimientos manuales)
	StockBefore       *float32  `json:"stockBefore"` // Stock antes del movimiento en la bodega afectada
	StockAfter        *float32  `json:"stockAfter"`  // Stock después del movimiento en la bodega afectada
}
