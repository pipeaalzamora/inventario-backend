package entities

import "time"

type EntityDeliveryPurchaseNote struct {
	ID                string                     `json:"id" db:"id"`
	DisplayID         string                     `json:"display_id" db:"display_id"`
	SupplierID        string                     `json:"supplier_id" db:"supplier_id"`
	FolioInvoice      string                     `json:"folio_invoice" db:"folio_invoice"`
	FolioGuide        string                     `json:"folio_guide" db:"folio_guide"`
	SupplierName      string                     `json:"supplier_name" db:"supplier_name"`
	CompanyID         string                     `json:"company_id" db:"company_id"`
	StoreID           string                     `json:"store_id" db:"store_id"`
	StoreName         string                     `json:"store_name" db:"store_name"`
	WarehouseID       string                     `json:"warehouse_id" db:"warehouse_id"`
	WarehouseName     string                     `json:"warehouse_name" db:"warehouse_name"`
	PurchaseID        string                     `json:"purchase_id" db:"purchase_id"`
	PurchaseDisplayID string                     `json:"purchase_display_id" db:"purchase_display_id"`
	Status            DeliveryPurchaseNoteStatus `json:"status" db:"note_status"`
	DueDate           time.Time                  `json:"due_date" db:"due_date"`
	Comment           string                     `json:"comment" db:"comment"`
	UserID            string                     `json:"user_id" db:"user_id"`
	UserName          string                     `json:"user_name" db:"user_name"`
	Total             float32                    `json:"total" db:"total"`
	CreatedAt         time.Time                  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time                  `json:"updated_at" db:"updated_at"`
}

type EntityDeliveryPurchaseNoteItem struct {
	ID                     string                         `json:"id" db:"id"`
	DeliveryPurchaseNoteID string                         `json:"delivery_purchase_note_id" db:"delivery_purchase_note_id"`
	StoreProductID         string                         `json:"store_product_id" db:"store_product_id"` // Referencia a product_per_store
	ProductName            string                         `json:"product_name" db:"product_name"`
	Quantity               float32                        `json:"quantity" db:"quantity"`
	PurchaseUnit           string                         `json:"purchase_unit" db:"purchase_unit"`
	Status                 DeliveryPurchaseNoteItemStatus `json:"status" db:"item_status"`
	Difference             float32                        `json:"difference" db:"difference"`
	UnitPrice              float32                        `json:"unit_price" db:"unit_price"`
	Subtotal               float32                        `json:"subtotal" db:"subtotal"`
	TaxTotal               float32                        `json:"tax_total" db:"tax_total"`
}

type DeliveryPurchaseNoteStatus string

const (
	DeliveryPurchaseNoteStatusPending   DeliveryPurchaseNoteStatus = "pending"
	DeliveryPurchaseNoteStatusCompleted DeliveryPurchaseNoteStatus = "completed"
	DeliveryPurchaseNoteStatusDisputed  DeliveryPurchaseNoteStatus = "disputed"
	DeliveryPurchaseNoteStatusCancelled DeliveryPurchaseNoteStatus = "cancelled"
)

func (d DeliveryPurchaseNoteStatus) String() string {
	return string(d)
}

type DeliveryPurchaseNoteItemStatus string

const (
	DeliveryPurchaseNoteItemStatusAccepted   DeliveryPurchaseNoteItemStatus = "accepted"
	DeliveryPurchaseNoteItemStatusSubstock   DeliveryPurchaseNoteItemStatus = "substock"
	DeliveryPurchaseNoteItemStatusRejected   DeliveryPurchaseNoteItemStatus = "rejected"
	DeliveryPurchaseNoteItemStatusSuprastock DeliveryPurchaseNoteItemStatus = "suprastock"
)

func (d DeliveryPurchaseNoteItemStatus) String() string {
	return string(d)
}

type EntityFile struct {
	ID        string    `json:"id" db:"id"`
	FileType  string    `json:"file_type" db:"file_type"`
	FileURL   string    `json:"file_url" db:"file_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
