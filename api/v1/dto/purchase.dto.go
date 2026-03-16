package dto

import "time"

type DTOPurchase struct {
	ID                            string                 `json:"id"`
	DisplayID                     string                 `json:"displayId"`
	SupplierID                    string                 `json:"supplierId"`
	SupplierName                  string                 `json:"supplierName"`
	SupplierPhone                 string                 `json:"supplierPhone"`
	StoreID                       string                 `json:"storeId"`
	StoreName                     string                 `json:"storeName"`
	WarehouseID                   string                 `json:"warehouseId"`
	WarehouseName                 string                 `json:"warehouseName"`
	WarehouseAddress              string                 `json:"warehouseAddress"`
	WarehousePhone                string                 `json:"warehousePhone"`
	InventoryRequestID            string                 `json:"inventoryRequestId"`
	Status                        string                 `json:"status"`
	ParentPurchaseID              *string                `json:"parentPurchaseId"`
	ParentPurchaseDisplayID       *string                `json:"parentDisplayId"`
	DeliveryPurchaseNoteID        *string                `json:"deliveryPurchaseNoteId"`
	DeliveryPurchaseNoteDisplayID *string                `json:"deliveryPurchaseNoteDisplayId"`
	ChildrenPurchase              []DTOPurchaseHierarchy `json:"childrenPurchase"`
	CreatedAt                     time.Time              `json:"createdAt"`
	UpdatedAt                     time.Time              `json:"updatedAt"`
	Items                         []DTOPurchaseItem      `json:"items"`
	PurchaseHistory               []DTOPurchaseHistory   `json:"purchaseHistory"`
}

type DTOPurchaseHierarchy struct {
	PurchaseChildID        string `json:"purchaseChildId"`
	PurchaseChildDisplayID string `json:"purchaseChildDisplayId"`
}

type DTOPurchaseItem struct {
	ID             string  `json:"id"`
	PurchaseID     string  `json:"purchaseId"`
	StoreProductID string  `json:"storeProductId"`
	ProductName    string  `json:"productName"`
	Quantity       float32 `json:"quantity"`
	PurchaseUnit   string  `json:"purchaseUnit"`
	UnitPrice      float32 `json:"unitPrice"`
	Subtotal       float32 `json:"subtotal"`
	Status         string  `json:"status"`
}

type DTOPurchaseHistory struct {
	ID          string    `json:"id"`
	PurchaseID  string    `json:"purchaseId"`
	NewStatus   string    `json:"newStatus"`
	Observation string    `json:"observation,omitempty"`
	ChangedAt   time.Time `json:"changedAt"`
}
