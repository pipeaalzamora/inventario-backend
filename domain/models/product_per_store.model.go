package models

// ModelProductPerStore representa la configuración de un producto para una tienda específica.
// Reemplaza a ModelProductCompany y unifica la información de request_restriction.
// Cada tienda maneja sus propios costos y configuraciones de producto.
type ModelProductPerStore struct {
	ID              string                           `json:"id"`
	StoreID         string                           `json:"storeId"`
	ProductTemplate ModelProduct                     `json:"productTemplate"`
	TagID           int                              `json:"tagId"` // el TAG es para etiqueda de si e smateria prima, receta, procesado, etc.
	ProductName     string                           `json:"productName"`
	Image           *string                          `json:"image"` // Imagen del producto base
	ItemSale        bool                             `json:"itemSale"`
	UseRecipe       bool                             `json:"useRecipe"`
	UnitInventory   ModelMeasurementUnique           `json:"unitInventory"`
	UnitMatrix      []ModelMeasurementConversionUnit `json:"unitMatrix"`
	Description     string                           `json:"description"`
	Costs           ModelStoreProductCost            `json:"costs"`
	Quantities      ModelStoreProductQuantities      `json:"quantities"`
	Suppliers       []ModelSupplierStoreProduct      `json:"suppliers"` // Lista de proveedores asignados
	CreatedAt       string                           `json:"createdAt"`
	UpdatedAt       string                           `json:"updatedAt"`
	// ItemInventory bool                             `json:"itemInventory"`
	// ItemPurchase  bool                             `json:"itemPurchase"`
	// IsFrozen      bool                             `json:"isFrozen"`
	// UnitPurchase  ModelMeasurementUnit             `json:"unitPurchase"`
	// CostLast      float32                          `json:"costLast"`
	// MinimalOrder  float32  `json:"minimalOrder"`
}

// ModelProductRequestRestriction es un modelo simplificado para consultas de restricciones de solicitud.
// Usado en el flujo de inventory_request para validar cantidades máximas.
type ModelProductRequestRestriction struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Image       *string  `json:"image"`
	MaxQuantity *float32 `json:"maxQuantity"`
	Unit        string   `json:"unit"`
}

type ModelStoreProductCost struct {
	//CostAvg       float32 `json:"costAvg"`
	// DEPRECATED: CostEstimated ahora está en ModelProduct (plantilla)
	// CostEstimated float32 `json:"costEstimated"`
}

type ModelStoreProductQuantities struct {
	MinimalStock float32  `json:"minimalStock"`
	MaximalStock float32  `json:"maximalStock"`
	MaxQuantity  *float32 `json:"maxQuantity"`
}
