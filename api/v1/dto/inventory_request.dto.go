package dto

import (
	"time"
)

type DTOInventoryRequest struct {
	ID             string                      `json:"id"`
	DisplayID      string                      `json:"displayId"`
	StoreID        string                      `json:"storeId"`
	StoreName      string                      `json:"storeName,omitempty"`
	WarehouseID    string                      `json:"warehouseId"`
	WarehouseName  string                      `json:"warehouseName,omitempty"`
	Status         string                      `json:"status"`
	RequestType    string                      `json:"requestType"`
	RequesterID    string                      `json:"requesterId"`
	RequesterName  string                      `json:"requesterName,omitempty"`
	CreatedAt      time.Time                   `json:"createdAt"`
	UpdatedAt      time.Time                   `json:"updatedAt"`
	Items          []DTOInventoryRequestItem   `json:"items,omitempty"`
	RequestHistory []DTOInventoryHistoryStatus `json:"requestHistory,omitempty"`
	Conflicts      []DTORequestConflictWrapper `json:"conflicts,omitempty"`
}

type DTOInventoryRequestItem struct {
	ItemID       string  `json:"itemId"`
	Quantity     float32 `json:"quantity"`
	PurchaseUnit string  `json:"purchaseUnit"`
}

type DTOInventoryHistoryStatus struct {
	NewStatus     string    `json:"newStatus"`
	Observation   *string   `json:"observation"`
	ChangedAt     time.Time `json:"changedAt"`
	ChangedByName string    `json:"changedByName"`
}

type DTORequestConflictWrapper struct {
	ItemID string      `json:"itemId"`
	Type   string      `json:"type"`
	Detail interface{} `json:"detail"`
}
