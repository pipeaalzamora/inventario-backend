package entities

import "time"

type EntitySupplierProduct struct {
	ID               string    `db:"id"`
	SupplierID       string    `db:"supplier_id"`
	ProductID        string    `db:"product_id"`
	Name             string    `db:"product_name"`
	Description      *string   `db:"description"`
	SKU              string    `db:"sku"`
	Price            float32   `db:"unit_price"`
	PurchaseUnitID   int       `db:"purchase_unit_id"`
	PurchaseUnit     string    `db:"purchase_unit"`
	PurchaseUnitName string    `db:"purchase_unit_name"`
	Available        bool      `db:"available"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

type EntitySupplierProductLegacy struct {
	ProductID                  string  `db:"product_id"`
	ProductName                string  `db:"product_name"`
	ProductDescription         *string `db:"product_description"`
	ProductImage               *string `db:"product_image"`
	CompanyID                  string  `db:"company_id"`
	StoreID                    string  `db:"store_id"`
	SupplierID                 string  `db:"supplier_id"`
	ProductCompanyID           string  `db:"product_company_id"`
	ProductCompanySKU          string  `db:"product_company_sku"`
	SupplierProductID          string  `db:"supplier_product_id"`
	SupplierProductPrice       float32 `db:"supplier_product_price"`
	SupplierProductMinQuantity float32 `db:"supplier_product_min_quantity"`
	UnitPurchase               string  `db:"unit_purchase"`
}
