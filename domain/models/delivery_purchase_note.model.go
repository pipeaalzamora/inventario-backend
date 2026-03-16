package models

import (
	"sofia-backend/infraestructure/entities"
	"time"
)

type ModelDeliveryPurchaseNote struct {
	ID                string                              `json:"id"`
	DisplayID         string                              `json:"display_id"`
	SupplierID        string                              `json:"supplier_id"`
	SupplierName      string                              `json:"supplier_name,omitempty"`
	CompanyID         string                              `json:"company_id"`
	StoreID           string                              `json:"store_id"`
	StoreName         string                              `json:"store_name,omitempty"`
	FolioInvoice      string                              `json:"folio_invoice"`
	FolioGuide        string                              `json:"folio_guide"`
	WarehouseID       string                              `json:"warehouse_id"`
	WarehouseName     string                              `json:"warehouse_name,omitempty"`
	DueDate           time.Time                           `json:"due_date"`
	Status            entities.DeliveryPurchaseNoteStatus `json:"status"`
	Comment           string                              `json:"comment"`
	UserID            string                              `json:"user_id"`
	UserName          string                              `json:"user_name"`
	Total             float32                             `json:"total"`
	Items             []ModelDeliveryPurchaseNoteItem     `json:"items,omitempty"`
	Files             []ModelFile                         `json:"files,omitempty"`
	PurchaseID        string                              `json:"purchase,omitempty"`
	PurchaseDisplayID string                              `json:"purchase_display_id,omitempty"`
	CreatedAt         time.Time                           `json:"created_at"`
	UpdatedAt         time.Time                           `json:"updated_at"`
}

type ModelDeliveryPurchaseNoteItem struct {
	ID                     string                                  `json:"id"`
	DeliveryPurchaseNoteID string                                  `json:"delivery_purchase_note_id"`
	StoreProductID         string                                  `json:"store_product_id"` // Referencia a product_per_store
	Status                 entities.DeliveryPurchaseNoteItemStatus `json:"status"`
	ProductName            string                                  `json:"product_name"`
	Quantity               float32                                 `json:"quantity"`
	PurchaseUnit           string                                  `json:"purchase_unit"`
	Difference             float32                                 `json:"difference"`
	UnitPrice              float32                                 `json:"unit_price"`
	Subtotal               float32                                 `json:"subtotal"`
	TaxTotal               float32                                 `json:"tax_total"`
}

type ModelFile struct {
	ID       string `json:"id" db:"id"`
	FileType string `json:"file_type" db:"file_type"`
	FileURL  string `json:"file_url" db:"file_url"`
}
