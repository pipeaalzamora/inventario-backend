package recipe

type RecipeInventoryRequest struct {
	CompanyID   string                       `json:"companyId" binding:"required"`
	StoreID     string                       `json:"storeId" binding:"required"`
	WarehouseID string                       `json:"warehouseId" binding:"required"`
	RequestType string                       `json:"requestType" binding:"required"`
	RequesterID string                       `json:"requesterId" binding:"required"`
	Items       []RecipeInventoryRequestItem `json:"items" binding:"required"`
	Observation *string                      `json:"observation"`
}

type RecipeInventoryRequestItem struct {
	ItemID   string  `json:"itemId" binding:"required"`
	Quantity float32 `json:"quantity" binding:"required"`
}

type RecipeInventoryRequestStatus struct {
	Observation *string `json:"observation"`
}
