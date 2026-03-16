package models

// DEPRECATED: ModelProductCompany ha sido reemplazado por ModelProductPerStore
// Esta estructura se mantiene comentada para referencia durante la migración.
// Ver: product_per_store.model.go
/*
type ModelProductCompany struct {
	ID            string                           `json:"id"`
	CompanyID     string                           `json:"companyId"`
	ProductID     string                           `json:"productId"`
	TagID         int                              `json:"tagId"`
	SKU           string                           `json:"sku"`
	ProductName   string                           `json:"productName"`
	ItemPurchase  bool                             `json:"itemPurchase"`
	ItemSale      bool                             `json:"itemSale"`
	ItemInventory bool                             `json:"itemInventory"`
	IsFrozen      bool                             `json:"isFrozen"`
	UseRecipe     bool                             `json:"useRecipe"`
	UnitPurchase  ModelMeasurementUnit             `json:"unitPurchase"`
	UnitInventory ModelMeasurementUnit             `json:"unitInventory"`
	UnitMatrix    []ModelMeasurementConversionUnit `json:"unitMatrix"`
	CostLast      float64                          `json:"costLast"`
	Description   string                           `json:"description"`
	CostEstimated float64                          `json:"costEstimated"`
	CostAvg       float64                          `json:"costAvg"`
	MinimalStock  float64                          `json:"minimalStock"`
	MaximalStock  float64                          `json:"maximalStock"`
	MinimalOrder  float64                          `json:"minimalOrder"`
	CreatedAt     string                           `json:"createdAt"`
	UpdatedAt     string                           `json:"updatedAt"`
}
*/

// ModelMeasurementUnit representa una unidad de medida
type ModelMeasurementUnit struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
	IsBase       bool   `json:"isBase"`
}

// ModelMeasurementConversionUnit representa una unidad de conversión con factor
type ModelMeasurementConversionUnit struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Abbreviation string  `json:"abbreviation"`
	Description  string  `json:"description"`
	Factor       float32 `json:"factor"`
}
