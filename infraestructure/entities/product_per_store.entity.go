package entities

// EntityProductPerStore representa la entidad de base de datos para product_per_store.
// Reemplaza a EntityProductCompany, ahora con granularidad a nivel de tienda.
type EntityProductPerStore struct {
	ID          string  `db:"id"`
	StoreID     string  `db:"store_id"`
	ProductID   string  `db:"product_id"`
	TagID       *int    `db:"tag_id"`
	SKU         string  `db:"sku"`
	ProductName string  `db:"product_name"`
	Image       *string `db:"image"` // Imagen del producto base
	// DEPRECATED: ItemPurchase    bool             `db:"item_purchase"`
	ItemSale bool `db:"item_sale"`
	// DEPRECATED: ItemInventory   bool             `db:"item_inventory"`
	// DEPRECATED: IsFrozen        bool             `db:"is_frozen"`
	UseRecipe       bool `db:"use_recipe"`
	UnitInventoryID int  `db:"unit_inventory_id"`
	//UnitInventory   EntityMeasurementUnique `db:"unit_inventory"`
	//UnitMatrix      []EntityMeasurementUnique `db:"unit_matrix"`
	// DEPRECATED: UnitPurchaseID  int              `db:"unit_purchase_id"`
	// DEPRECATED: UnitPurchase    string           `db:"unit_purchase"`
	// DEPRECATED: CostLast        float32          `db:"cost_last"`
	Description   string  `db:"description"`
	CostEstimated float32 `db:"cost_estimated"`
	// DEPRECATED: CostEstimated float32 `db:"cost_estimated"` // Ya no está en product_per_store
	CostAvg       float32 `db:"cost_avg"`
	MinimalStock  float32 `db:"minimal_stock"`
	MaximalStock  float32 `db:"maximal_stock"`
	// DEPRECATED: MinimalOrder    float32          `db:"minimal_order"`
	MaxQuantity *float32 `db:"max_quantity"` // Campo de restricción de solicitud
	CreatedAt   string   `db:"created_at"`
	UpdatedAt   string   `db:"updated_at"`
}

// EntityProductRequestRestriction representa la vista de restricciones de solicitud.
// Usado para consultas simplificadas en el flujo de inventory_request.
type EntityProductRequestRestriction struct {
	ID          string   `db:"id"`
	Name        string   `db:"product_name"`
	Description *string  `db:"description"`
	Image       *string  `db:"image"`
	MaxQuantity *float32 `db:"max_quantity"`
	Unit        string   `db:"unit_inventory"`
}
