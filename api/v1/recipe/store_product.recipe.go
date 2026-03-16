package recipe

// RecipeCreateStoreProduct define la entrada para crear un producto de tienda.
type RecipeCreateStoreProduct struct {
	CompanyID       string                       `json:"companyId" binding:"required,uuid" errMsg:"El ID de la compañía es obligatorio"`
	StoreID         string                       `json:"storeId" binding:"required,uuid" errMsg:"El ID de la tienda es obligatorio"`
	ProductID       string                       `json:"productId" binding:"required,uuid" errMsg:"El ID de la plantilla es obligatorio"`
	ProductName     string                       `json:"productName" binding:"required" errMsg:"El nombre del producto es obligatorio"`
	TagID           int                          `json:"tagId"`
	SKU             string                       `json:"sku" binding:"required" errMsg:"El SKU es obligatorio"`
	IsSellable      bool                         `json:"isSellable"`
	UseRecipe       bool                         `json:"useRecipe"`
	UnitInventoryID int                          `json:"unitInventoryId" binding:"required" errMsg:"La unidad de inventario es obligatoria"`
	UnitMatrix      []RecipeUnitConversion       `json:"unitMatrix" binding:"required" errMsg:"La matriz de unidades es obligatoria y debe tener al menos una unidad"`
	Description     string                       `json:"description"`
	EstimatedCost   float32                      `json:"estimatedCost"`
	Quantities      RecipeStoreProductQuantities `json:"quantities"`
	Suppliers       []RecipeStoreProductSupplier `json:"suppliers"`
}

// RecipeUpdateStoreProduct define la entrada para actualizar un producto de tienda.
// Todos los campos son opcionales (partial update).
type RecipeUpdateStoreProduct struct {
	ProductName     *string                       `json:"productName"`
	TagID           *int                          `json:"tagId"`
	SKU             *string                       `json:"sku"`
	IsSellable      *bool                         `json:"isSellable"`
	UseRecipe       *bool                         `json:"useRecipe"`
	UnitInventoryID *int                          `json:"unitInventoryId"`
	UnitMatrix      []RecipeUnitConversion        `json:"unitMatrix"`
	Description     *string                       `json:"description"`
	EstimatedCost   *float32                      `json:"estimatedCost"`
	Quantities      *RecipeStoreProductQuantities `json:"quantities"`
	Suppliers       []RecipeStoreProductSupplier  `json:"suppliers"`
}

// RecipeStoreProductPathParams define los parámetros de ruta para operaciones con producto de tienda.
type RecipeStoreProductPathParams struct {
	ID string `uri:"id" binding:"required,uuid" errMsg:"El ID del producto de tienda es obligatorio"`
}

// RecipeUnitConversion define una unidad de conversión para la matriz de unidades.
type RecipeUnitConversion struct {
	UnitID int `json:"unitId" binding:"required"`
	// Name         string  `json:"name" binding:"required"`
	// Abbreviation string  `json:"abbreviation" binding:"required"`
	Factor float32 `json:"factor" binding:"required,gt=0"`
}

// RecipeStoreProductCost define los costos del producto de tienda.
type RecipeStoreProductCost struct {
	CostAvg       float32 `json:"costAvg"`
	CostEstimated float32 `json:"costEstimated"`
}

// RecipeStoreProductQuantities define las cantidades del producto de tienda.
type RecipeStoreProductQuantities struct {
	MinimalStock float32  `json:"minimalStock"`
	MaximalStock float32  `json:"maximalStock"`
	MaxQuantity  *float32 `json:"maxQuantity"`
}

// RecipeStoreProductSupplier define un proveedor asignado al producto de tienda con su prioridad.
type RecipeStoreProductSupplier struct {
	SupplierID string `json:"supplierId" binding:"required"`
	Priority   int    `json:"priority" binding:"required"`
}
