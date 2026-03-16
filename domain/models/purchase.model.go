package models

import (
	"sofia-backend/infraestructure/entities"
	"time"
)

type ModelPurchase struct {
	ID                            string                   `json:"id"`
	DisplayID                     string                   `json:"displayId"`
	SupplierID                    string                   `json:"supplierId"`
	SupplierName                  string                   `json:"supplierName,omitempty"`
	SupplierPhone                 string                   `json:"supplierPhone,omitempty"`
	CompanyID                     string                   `json:"companyId"`
	StoreID                       string                   `json:"storeId"`
	StoreName                     string                   `json:"storeName,omitempty"`
	WarehouseID                   string                   `json:"warehouseId"`
	WarehouseName                 string                   `json:"warehouseName,omitempty"`
	WarehouseAddress              string                   `json:"warehouseAddress,omitempty"`
	WarehousePhone                string                   `json:"warehousePhone,omitempty"`
	InventoryRequestID            string                   `json:"inventoryRequestId"`
	Status                        entities.PurchaseStatus  `json:"status"`
	CreatedAt                     time.Time                `json:"createdAt"`
	UpdatedAt                     time.Time                `json:"updatedAt"`
	ParentPurchaseID              *string                  `json:"parentPurchaseId"`
	ParentPurchaseDisplayID       *string                  `json:"parentDisplayId,omitempty"`
	DeliveryPurchaseNoteID        *string                  `json:"deliveryPurchaseNoteId"`
	DeliveryPurchaseNoteDisplayID *string                  `json:"deliveryPurchaseNoteDisplayId,omitempty"`
	ChildrenPurchase              []ModelPurchaseHierarchy `json:"childrenPurchase,omitempty"`
	Items                         []ModelPurchaseItem      `json:"items,omitempty"`
	PurchaseHistory               []ModelPurchaseHistory   `json:"purchaseHistory,omitempty"`
}

type ModelPurchaseHierarchy struct {
	PurchaseChildID        string `json:"purchaseChildId"`
	PurchaseChildDisplayID string `json:"purchaseChildDisplayId"`
}

type ModelPurchaseItem struct {
	ID              string                      `json:"id"`
	PurchaseID      string                      `json:"purchaseId"`
	StoreProductID  string                      `json:"storeProductId"` // Referencia a product_per_store
	ProductName     string                      `json:"productName"`
	ProductID       string                      `json:"productId,omitempty"`
	SupplierOptions []SupplierOption            `json:"supplierOptions,omitempty"`
	Quantity        float32                     `json:"quantity"`
	PurchaseUnit    string                      `json:"purchaseUnit"`
	UnitPrice       float32                     `json:"unitPrice"`
	Subtotal        float32                     `json:"subtotal"`
	Status          entities.ItemPurchaseStatus `json:"status"`
}

type ModelPurchaseHistory struct {
	ID          string    `json:"id"`
	PurchaseID  string    `json:"purchaseId"`
	NewStatus   string    `json:"newStatus"`
	ChangedAt   time.Time `json:"changedAt"`
	Observation string    `json:"observation,omitempty"`
}
