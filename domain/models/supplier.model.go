package models

import "time"

type ModelSupplier struct {
	ID           string                 `json:"id"`
	FiscalData   ModelFiscalData        `json:"fiscalData"`
	CountryID    int                    `json:"countryId"`
	SupplierName string                 `json:"supplierName"`
	Description  string                 `json:"description"`
	Available    bool                   `json:"available"`
	Contacts     []ModelSupplierContact `json:"contacts"`
	Products     []ModelSupplierProduct `json:"products"`
	CreatedAt    time.Time              `json:"createdAt"`
	UpdatedAt    time.Time              `json:"updatedAt"`
}

type ModelSupplierContact struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ModelSupplierStoreProduct struct {
	ID                string  `json:"id"`
	SupplierProductID string  `json:"supplierProductId"`
	SupplierID        string  `json:"supplierId"`
	SupplierName      string  `json:"supplierName"`
	ProductID         string  `json:"productId"`
	ProductName       string  `json:"productName"`
	CompanyID         string  `json:"companyId"`
	StoreID           string  `json:"storeId"`
	Priority          int     `json:"priority"`
	Preferred         bool    `json:"preferred"`
	ServiceZone       string  `json:"serviceZone"`
	Price             float32 `json:"price"`
	MinOrderQuantity  float32 `json:"minOrderQuantity"`
	PurchaseUnit      string  `json:"purchaseUnit"`
	PaymentTerms      string  `json:"paymentTerms"`
	LeadTimeDays      int     `json:"leadTimeDays"`
	Available         bool    `json:"available"`
	// SaleUnit		  string  `json:"saleUnit"`
}

// Create a dictionary for each product ID with a {priority, supplierID} structure

type ModelProductSuppliersMap map[string][]SupplierOption
