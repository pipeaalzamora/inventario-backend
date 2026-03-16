package dto

import "time"

type DTODeliveryPurchaseNote struct {
	ID                string                        `json:"id"`
	DisplayID         string                        `json:"displayId"`
	SupplierID        string                        `json:"supplierId"`
	SupplierName      string                        `json:"supplierName"`
	PurchaseID        string                        `json:"purchaseId"`
	FolioInvoice      string                        `json:"folioInvoice"`
	FolioGuide        string                        `json:"folioGuide"`
	PurchaseDisplayID string                        `json:"purchaseDisplayId"`
	StoreID           string                        `json:"storeId"`
	Status            string                        `json:"status"`
	StoreName         string                        `json:"storeName"`
	WarehouseID       string                        `json:"warehouseId"`
	WarehouseName     string                        `json:"warehouseName"`
	Comment           string                        `json:"comment"`
	Total             float32                       `json:"total"`
	UserID            string                        `json:"userId"`
	UserName          string                        `json:"userName"`
	CreatedAt         time.Time                     `json:"createdAt"`
	UpdatedAt         time.Time                     `json:"updatedAt"`
	Items             []DTODeliveryPurchaseNoteItem `json:"items"`
	Files             []DTOFile                     `json:"files"`
}

type DTODeliveryPurchaseNoteItem struct {
	ID                     string  `json:"id"`
	DeliveryPurchaseNoteID string  `json:"deliveryPurchaseNoteId"`
	StoreProductID         string  `json:"storeProductId"`
	ProductName            string  `json:"productName"`
	Quantity               float32 `json:"quantity"`
	PurchaseUnit           string  `json:"purchaseUnit"`
	Difference             float32 `json:"difference"`
	Status                 string  `json:"status"`
	UnitPrice              float32 `json:"unitPrice"`
	Subtotal               float32 `json:"subtotal"`
	TaxTotal               float32 `json:"taxTotal"`
}

type DTOFile struct {
	ID       string `json:"id"`
	FileType string `json:"fileType"`
	FileURL  string `json:"fileUrl"`
}
