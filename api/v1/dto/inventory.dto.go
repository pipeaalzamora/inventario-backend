package dto

type WarehouseProductsDTO struct {
	ProductID    string  `json:"productId"`
	ProductName  string  `json:"productName"`
	ProductImage *string `json:"productImage"`
	ProductSku   string  `json:"productSku"`

	WarehouseID   string `json:"warehouseId"`
	WarehouseName string `json:"warehouseName"`

	InStock    float32 `json:"inStock"`
	BaseUnit   string  `json:"baseUnit"`
	BaseUnitId int     `json:"baseUnitId"`
}
