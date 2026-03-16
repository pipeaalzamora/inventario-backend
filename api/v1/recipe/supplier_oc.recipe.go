package recipe

type SupplierOCRecipe struct {
	PurchaseId  string           `json:"purchaseId" binding:"required"`
	Observation string           `json:"observation"`
	Items       []SupplierOCItem `json:"items" binding:"required"`
}

type SupplierOCItem struct {
	ItemId   string `json:"itemId" binding:"required"`
	Accepted bool   `json:"accepted" binding:"required"`
}
