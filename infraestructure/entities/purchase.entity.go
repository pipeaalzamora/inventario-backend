package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type EntityPurchase struct {
	ID                    string         `db:"id"`
	DisplayID             string         `db:"display_id"`
	SupplierID            string         `db:"supplier_id"`
	SupplierName          string         `db:"supplier_name"` // Solo aparece cuando vienen del repo
	SupplierPhone         string         `db:"supplier_phone"`
	CompanyID             string         `db:"company_id"`
	StoreID               string         `db:"store_id"`
	StoreName             string         `db:"store_name"` // Solo aparece cuando vienen del repo
	WarehouseID           string         `db:"warehouse_id"`
	WarehouseName         string         `db:"warehouse_name"`
	WarehouseAddress      string         `db:"warehouse_address"`
	WarehousePhone        string         `db:"warehouse_phone"`
	InventoryRequestID    string         `db:"inventory_request_id"`
	Status                PurchaseStatus `db:"status"`
	DeliveryNoteID        *string        `db:"delivery_purchase_note_id"`
	DeliveryNoteDisplayID *string        `db:"delivery_purchase_note_display_id"`
	ParentPurchaseID      *string        `db:"parent_purchase_id"`
	ParentDisplayID       *string        `db:"parent_display_id"`
	CreatedAt             time.Time      `db:"created_at"`
	UpdatedAt             time.Time      `db:"updated_at"`
}

type EntityPurchaseItem struct {
	ID              string             `db:"id"`
	PurchaseID      string             `db:"purchase_id"`
	StoreProductID  string             `db:"store_product_id"` // Referencia a product_per_store
	ProductName     string             `db:"product_name"`
	ProductID       string             `db:"product_id"`
	SupplierOptions SupplierOptions    `db:"available_suppliers"`
	Quantity        float32            `db:"quantity"`
	PurchaseUnit    string             `db:"purchase_unit"`
	UnitPrice       float32            `db:"unit_price"`
	Subtotal        float32            `db:"subtotal"`
	Observation     string             `db:"observation"`
	Status          ItemPurchaseStatus `db:"status"`
}

type EntityPurchaseHierarchy struct {
	ID               string `db:"id"`
	ParentPurchaseID string `db:"parent_purchase_id"`
	ParentDisplayID  string `db:"parent_display_id"`
	ChildPurchaseID  string `db:"child_purchase_id"`
	ChildDisplayID   string `db:"child_display_id"`
}

type EntityPurchaseHistory struct {
	ID          string    `db:"id"`
	PurchaseID  string    `db:"purchase_id"`
	NewStatus   string    `db:"new_status"`
	Observation string    `db:"observation"` // TODO: Remove
	ChangedAt   time.Time `db:"changed_at"`
}
type EntitySupplierOption struct {
	SupplierID string  `json:"supplier_id" db:"supplier_id"`
	Price      float32 `json:"supplier_price" db:"supplier_price"`
}

type SupplierOptions []EntitySupplierOption

func (s SupplierOptions) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *SupplierOptions) Scan(src interface{}) error {
	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return fmt.Errorf("cannot scan %T into SupplierOptions", src)
	}
	return json.Unmarshal(data, s)
}

type PurchaseStatus string
type ItemPurchaseStatus string

const (
	ItemPurchaseStatusPending    ItemPurchaseStatus = "pending"
	ItemPurchaseStatusApproved   ItemPurchaseStatus = "approved"
	ItemPurchaseStatusRejected   ItemPurchaseStatus = "rejected"
	ItemPurchaseStatusDelivered  ItemPurchaseStatus = "delivered"
	ItemPurchaseStatusRetried    ItemPurchaseStatus = "retried"
	ItemPurchaseStatusNoSupplier ItemPurchaseStatus = "no_supplier"
)

const (
	PurchaseStatusPending    PurchaseStatus = "pending"
	PurchaseStatusRejected   PurchaseStatus = "rejected"
	PurchaseStatusCompleted  PurchaseStatus = "completed"
	PurchaseStatusSunk       PurchaseStatus = "sunk"
	PurchaseStatusOnDelivery PurchaseStatus = "on_delivery"
	PurchaseStatusCancelled  PurchaseStatus = "cancelled"
	PurchaseStatusEdited     PurchaseStatus = "edited"
	PurchaseStatusArrived    PurchaseStatus = "arrived"
)

func (s ItemPurchaseStatus) ToString() string {
	return string(s)
}
func (s PurchaseStatus) ToString() string {
	return string(s)
}
