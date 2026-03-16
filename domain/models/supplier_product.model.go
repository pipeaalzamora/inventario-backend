package models

import "time"

type ModelSupplierProduct struct {
	ID           string                   `json:"id"`
	SupplierID   string                   `json:"supplierId"`
	ProductID    string                   `json:"productId"`
	Name         string                   `json:"name"`
	Description  *string                  `json:"description"`
	SKU          string                   `json:"sku"`
	Price        float32                  `json:"price"`
	PurchaseUnit ModelProductPurchaseUnit `json:"purchaseUnit"`
	Available    bool                     `json:"available"`
	CreatedAt    time.Time                `json:"createdAt"`
	UpdatedAt    time.Time                `json:"updatedAt"`
}

type ModelProductPurchaseUnit struct {
	UnitID   int    `json:"unitId"`
	UnitAbv  string `json:"unitAbv"`
	UnitName string `json:"unitName"`
}

type ModelSupplierProductLegacy struct {
	ProductID                  string  `json:"productId"`
	ProductName                string  `json:"productName"`
	ProductDescription         *string `json:"productDescription"`
	ProductImage               *string `json:"productImage"`
	CompanyID                  string  `json:"companyId"`
	StoreID                    string  `json:"storeId"`
	SupplierID                 string  `json:"supplierId"`
	ProductCompanyID           string  `json:"productCompanyId"`
	ProductCompanySKU          string  `json:"productCompanySku"`
	SupplierProductID          string  `json:"supplierProductId"`
	SupplierProductPrice       float32 `json:"supplierProductPrice"`
	SupplierProductMinQuantity float32 `json:"supplierProductMinQuantity"`
	UnitPurchase               string  `json:"unitPurchase"`
}

type SupplierOption struct {
	SupplierID string  `json:"supplierId"`
	Price      float32 `json:"price"`
}

type ModelMapProductCompanyKeySupplierProduct map[string]ModelSupplierProductLegacy
