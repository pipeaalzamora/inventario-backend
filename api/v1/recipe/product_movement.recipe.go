package recipe

import "time"

type RecipeNewMovement struct {
	StoreId     string                  `json:"storeId"`
	CompanyId   string                  `json:"companyId"`
	MovedFrom   *string                 `json:"movedFrom"`
	MovedTo     *string                 `json:"movedTo"`
	MovedAt     time.Time               `json:"movedAt"`
	WasteKind   WasteKind               `json:"outputKind"`
	Products    []RecipeMovementProduct `json:"products"`
	Observation string                  `json:"observation"`
	MetaData    *map[string]any         `json:"metaData,omitempty"`
}
type WasteKind string

const (
	noneKind     WasteKind = "NONE"
	wasteKind    WasteKind = "WASTE"
	expiredKind  WasteKind = "EXPIRED"
	adjustedKind WasteKind = "ADJUSTED"
	otherKind    WasteKind = "OTHER"
)

type RecipeMovementProduct struct {
	StoreProductID string  `json:"storeProductId" binding:"required,uuid"`
	Quantity       float32 `json:"quantity" binding:"required,gt=0"`
	//UnitCost       float32 `json:"unitCost"`
	//TotalCost      float32 `json:"totalCost"`
	//ProductState   string  `json:"productState"`
	Reason string `json:"reason"`
}

type RecipeProductMovementByWarehouseIDs struct {
	WarehouseIDs []string `json:"warehouseIds"`
}

// RecipeTransferMovement representa la entrada para transferir productos entre bodegas
type RecipeTransferMovement struct {
	StoreId           string                    `json:"storeId" binding:"required,uuid" errMsg:"El ID de la tienda es obligatorio"`
	FromWarehouseId   string                    `json:"fromWarehouseId" binding:"required,uuid" errMsg:"El ID de la bodega origen es obligatorio"`
	ToWarehouseId     string                    `json:"toWarehouseId" binding:"required,uuid" errMsg:"El ID de la bodega destino es obligatorio"`
	Products          []RecipeTransferProduct   `json:"products" binding:"required,min=1,dive" errMsg:"Debe incluir al menos un producto"`
	Observation       string                    `json:"observation"`
}

// RecipeTransferProduct representa un producto a transferir con su cantidad y costo
type RecipeTransferProduct struct {
	StoreProductID string  `json:"storeProductId" binding:"required,uuid" errMsg:"El ID del producto es obligatorio"`
	Quantity       float32 `json:"quantity" binding:"required,gt=0" errMsg:"La cantidad debe ser mayor a 0"`
	UnitCost       float32 `json:"unitCost" binding:"required,gte=0" errMsg:"El costo unitario debe ser mayor o igual a 0"`
}
