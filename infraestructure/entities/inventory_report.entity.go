package entities

type EntityProductGeneralStock struct {
	ProductID     string  `json:"product_id" db:"product_id"`
	ProductName   string  `json:"product_name" db:"product_name"`
	SKU           string  `json:"sku" db:"sku"`
	Quantity      float32 `json:"quantity" db:"quantity"`
	MinStock      float32 `json:"min_stock" db:"min_stock"`
	MaxStock      float32 `json:"max_stock" db:"max_stock"`
	WarehouseID   string  `json:"warehouse_id" db:"warehouse_id"`
	WarehouseName string  `json:"warehouse_name" db:"warehouse_name"`
}

type EntityProductDetailStock struct {
	WarehouseID   string  `json:"warehouseId" db:"warehouse_id"`
	WarehouseName string  `json:"warehouseName" db:"warehouse_name"`
	ProductID     string  `json:"productId" db:"product_company_id"`
	ProductName   string  `json:"productName" db:"product_name"`
	SKU           string  `json:"sku" db:"sku"`
	ProductImage  *string `json:"productImage" db:"product_image"`
	CurrentStock  float32 `json:"currentStock" db:"current_stock"`
	MinStock      float32 `json:"minStock" db:"minimal_stock"`
	MaxStock      float32 `json:"maxStock" db:"maximal_stock"`
	AvgCost       float32 `json:"avgCost" db:"avg_cost"`
}
