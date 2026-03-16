package models

import (
	"sofia-backend/infraestructure/entities"
	"time"
)

type ModelInventoryRequest struct {
	ID             string                        `json:"id"`
	DisplayID      string                        `json:"displayId"`
	CompanyID      string                        `json:"companyId"`
	StoreID        string                        `json:"storeId"`
	StoreName      string                        `json:"storeName"` // Solo aparece cuando vienen del repo
	WarehouseID    string                        `json:"warehouseId"`
	WarehouseName  string                        `json:"warehouseName"` // Solo aparece cuando vienen del repo
	Status         entities.RequestStatus        `json:"status"`
	RequestType    entities.RequestType          `json:"requestType"`
	RequesterID    string                        `json:"requesterId"`
	RequesterName  string                        `json:"requesterName"` // Solo aparece cuando vienen del repo
	CreatedAt      time.Time                     `json:"createdAt"`
	UpdatedAt      time.Time                     `json:"updatedAt"`
	Items          []ModelInventoryRequestItem   `json:"items,omitempty"`          // Solo aparece cuando vienen del repo
	RequestHistory []ModelInventoryHistoryStatus `json:"requestHistory,omitempty"` // Solo aparece cuando vienen del repo
}

type ModelInventoryRequestItem struct {
	ID                 string  `json:"id"`
	InventoryRequestID string  `json:"inventoryRequestId"`
	StoreProductID     string  `json:"storeProductId"` // Referencia a product_per_store
	ProductID          string  `json:"productId"`
	Quantity           float32 `json:"quantity"`
	PurchaseUnit       string  `json:"purchaseUnit"`
}

type ModelInventoryHistoryStatus struct {
	ID                 string    `json:"id"`
	InventoryRequestID string    `json:"inventoryRequestId"`
	NewStatus          string    `json:"newStatus"`
	Observation        *string   `json:"observation"`
	ChangedAt          time.Time `json:"changedAt"`
	ChangedByName      string    `json:"changedByName"`
}

type ModelMaxQuantityConflict struct {
	MaxQuantity  float32 `json:"maxQuantity"`
	CurrQuantity float32 `json:"currQuantity"`
}

func (c *ModelMaxQuantityConflict) GetConflictType() string {
	return "max_quantity_conflict"
}

type ModelModRequestConflict struct {
	Mod int `json:"mod"`
}

func (c *ModelModRequestConflict) GetConflictType() string {
	return "mod_request_conflict"
}
