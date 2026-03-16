package recipe

import "sofia-backend/domain/models"

type RecipePurchase struct {
	Description        string `json:"description" validate:"required"`
	SupplierID         string `json:"supplierId" validate:"required"`
	CompanyID          string `json:"companyId" validate:"required"`
	StoreID            string `json:"storeId" validate:"required"`
	WarehouseID        string `json:"warehouseId" validate:"required"`
	InventoryRequestID string `json:"inventoryRequestId"`
	// Status      string               `json:"status" validate:"required"`
	Items []RecipePurchaseItem `json:"items" validate:"required,dive,required"`
}

// RecipeCreatePurchase represents the structure for creating a new purchase from endpoint

type RecipePurchaseItem struct {
	StoreProductID  string                  `json:"itemId" validate:"required"`
	Quantity        float32                 `json:"quantity" validate:"required"`
	PurchaseUnit    string                  `json:"purchaseUnit"`
	UnitPrice       float32                 `json:"unitPrice" validate:"required"`
	SupplierOptions []models.SupplierOption `json:"supplierOptions,omitempty"`
	// Status    string  `json:"status" validate:"required"`
}

type RecipeCancelPurchase struct {
	Observation string `json:"observation" validate:"required"`
}
