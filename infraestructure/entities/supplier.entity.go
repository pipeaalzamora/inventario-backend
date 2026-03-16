package entities

import "time"

type EntitySupplier struct {
	ID           string    `json:"id" db:"id"`
	FiscalDataID string    `json:"fiscal_data_id" db:"fiscal_data_id"`
	CountryID    int       `json:"country_id" db:"country_id"`
	SupplierName string    `json:"supplier_name" db:"supplier_name"`
	Description  string    `json:"description" db:"description"`
	Available    bool      `json:"available" db:"available"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

/*
	type EntitySupplierStore struct {
		ID         string `json:"id" db:"id"`
		SupplierID string `json:"supplier_id" db:"supplier_id"`
		StoreID    string `json:"store_id" db:"store_id"`
		Available  bool   `json:"available" db:"available"`
	}
*/

type EntitySupplierContact struct {
	ID          string    `json:"id" db:"id"`
	SupplierID  string    `json:"supplier_id" db:"supplier_id"`
	ContactName string    `json:"contact_name" db:"contact_name"`
	Description string    `json:"description" db:"description"`
	Email       string    `json:"email" db:"email"`
	Phone       string    `json:"phone" db:"phone"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type EntitySupplierStoreProduct struct {
	ID             string    `json:"id" db:"id"`
	StoreProductID string    `json:"store_product_id" db:"store_product_id"`
	SupplierID     string    `json:"supplier_id" db:"supplier_id"`
	SupplierName   string    `json:"supplier_name" db:"supplier_name"`
	ProductID      string    `json:"product_id" db:"product_id"`
	ProductName    string    `json:"product_name" db:"product_name"`
	Priority       int       `json:"priority" db:"priority"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	// Fields from supplier_product via LEFT JOIN
	UnitPrice    *float64 `db:"unit_price"`
	PurchaseUnit *string  `db:"purchase_unit"`
}
