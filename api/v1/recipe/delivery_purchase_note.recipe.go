package recipe

import "mime/multipart"

type RecipeDeliveryPurchaseNote struct {
	ID          string                           `json:"id"`
	SupplierID  string                           `json:"supplierId" validate:"required"`
	CompanyID   string                           `json:"companyId" validate:"required"`
	StoreID     string                           `json:"storeId" validate:"required"`
	WarehouseID string                           `json:"warehouseId" validate:"required"`
	PurchaseID  string                           `json:"purchaseId" validate:"required"`
	DueDate     string                           `json:"dueDate" validate:"required,datetime=2006-01-02"`
	Comment     string                           `json:"comment"`
	Status      string                           `json:"status"`
	Total       float32                          `json:"total"`
	Items       []RecipeDeliveryPurchaseNoteItem `json:"items" validate:"dive"`
}

type RecipeDeliveryPurchaseNoteItem struct {
	StoreProductID string  `json:"storeProductId" validate:"required"`
	Quantity       float32 `json:"quantity" validate:"required"`
	PurchaseUnit   string  `json:"purchaseUnit" validate:"required"`
	UnitPrice      float32 `json:"unitPrice"`
	Subtotal       float32 `json:"subtotal"`
	TaxTotal       float32 `json:"taxTotal"`
	Total          float32 `json:"total"`
	Difference     float32 `json:"difference"`
	Status         string  `json:"status"`
}

type UploadForm struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
