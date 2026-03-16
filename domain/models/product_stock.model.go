package models

type ModelProductGeneralStock struct {
	ProductID     string  `json:"product_id"`
	ProductName   string  `json:"product_name"`
	SKU           string  `json:"sku"`
	Quantity      float32 `json:"quantity"`
	InventoryUnit string  `json:"inventoryUnit"`
	MinStock      float32 `json:"min_stock"`
	MaxStock      float32 `json:"max_stock"`
	WarehouseID   string  `json:"warehouse_id"`
	WarehouseName string  `json:"warehouse_name"`
}

type ModelProductDetailStock struct {
	ProductID      string                `json:"productId"`
	ProductName    string                `json:"productName"`
	SKU            string                `json:"sku"`
	ProductImage   *string               `json:"productImage"`
	CurrentStock   float32               `json:"currentStock"`
	MinStock       float32               `json:"minStock"`
	MaxStock       float32               `json:"maxStock"`
	WarehouseStock []ModelWareHouseStock `json:"warehouseStock"`
}

type ModelWareHouseStock struct {
	WarehouseID   string  `json:"warehouseId"`
	WarehouseName string  `json:"warehouseName"`
	AvgCost       float32 `json:"avgCost"`
	CurrentStock  float32 `json:"currentStock"`
	MinStock      float32 `json:"minStock"`
	MaxStock      float32 `json:"maxStock"`
}

////////////////////////////////////////////////////////////////////

// ModelWarehouseProductInventory agrupa productos de inventario por bodega
type ModelWarehouseProductInventory struct {
	WarehouseID   string                   `json:"warehouseId"`
	WarehouseName string                   `json:"warehouseName"`
	Products      []ModelProductStoreStock `json:"products"`
}

type ModelProductStoreStock struct {
	ModelProductPerStore
	Totals ModelProductStoreStockTotals `json:"totals"`
}

type ModelProductStoreStockTotals struct {
	CurrentStock float32 `json:"currentStock"`
	StockIn      float32 `json:"stockIn"`
	StockOut     float32 `json:"stockOut"`
	AreMinAlert  bool    `json:"areMinAlert"`
	AvgCost      float32 `json:"avgCost"`
	TotalCost    float32 `json:"totalCost"`
	AreMaxAlert  bool    `json:"areMaxAlert"`
}

type ModelWarehouseProductStock struct {
	StoreProductId string  `json:"storeProductId"`
	WarehouseID    string  `json:"warehouseId"`
	WarehouseName  string  `json:"warehouseName"`
	CurrentStock   float32 `json:"currentStock"`
	AvgCost        float32 `json:"avgCost"`
}

// ModelWarehouseProductTransit representa stock en tránsito en bodegas de traspaso
type ModelWarehouseProductTransit struct {
	StoreProductId       string  `json:"storeProductId"`
	WarehouseIDReference string  `json:"warehouseIdReference"`
	Direction            string  `json:"direction"` // IN o OUT
	InStock              float32 `json:"inStock"`
	AvgCost              float32 `json:"avgCost"`
}
