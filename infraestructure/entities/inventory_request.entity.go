package entities

import (
	"time"
)

type EntityInventoryRequest struct {
	ID            string        `db:"id"`
	DisplayID     string        `db:"display_id"`
	CompanyID     string        `db:"company_id"`
	StoreID       string        `db:"store_id"`
	StoreName     string        `db:"store_name"` // Solo aparece cuando vienen del repo
	WarehouseID   string        `db:"warehouse_id"`
	WarehouseName string        `db:"warehouse_name"` // Solo aparece cuando vienen del repo
	Status        RequestStatus `db:"status"`
	RequestType   RequestType   `db:"request_type"`
	RequesterID   string        `db:"requester_id"`
	RequesterName string        `db:"requester_name"` // Solo aparece cuando vienen del repo
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at"`
}
type EntityInventoryRequestItem struct {
	ID                 string  `db:"id"`
	InventoryRequestID string  `db:"inventory_request_id"`
	StoreProductID     string  `db:"store_product_id"` // Referencia a product_per_store
	ProductID          string  `db:"product_id"`
	Quantity           float32 `db:"quantity"`
	PurchaseUnit       string  `db:"purchase_unit"`
}

type EntityInventoryHistoryStatus struct {
	ID                 string        `db:"id"`
	InventoryRequestID string        `db:"inventory_request_id"`
	NewStatus          RequestStatus `db:"new_status"`
	Observation        *string       `db:"observation"`
	ChangedAt          time.Time     `db:"changed_at"`
	ChangedBy          string        `db:"changed_by"`
	ChangedByName      string        `db:"changed_by_name"`
}

type FilterInventoryRequest struct {
	Status *RequestStatus `json:"status"`
}

// ENUMS de Postgres en Go
type RequestStatus string
type RequestType string

const (
	// Valores de request_status
	RequestStatusPending    RequestStatus = "pending"
	RequestStatusApproved   RequestStatus = "approved"
	RequestStatusRejected   RequestStatus = "rejected"
	RequestStatusConflicted RequestStatus = "conflicted"
	RequestStatusCompleted  RequestStatus = "completed"
	RequestStatusCanceled   RequestStatus = "cancelled"

	// Valores de request_type
	RequestTypeSupplier RequestType = "supplier_request"
	RequestTypePurchase RequestType = "purchase_request"
	RequestTypeInternal RequestType = "internal_request"
)

func (s RequestStatus) ToString() string {
	return string(s)
}

func (s RequestType) ToString() string {
	return string(s)
}
