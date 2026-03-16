package dto

import "time"

type DtoSupplierSimple struct {
	ID           string `json:"id"`
	IDFiscal     string `json:"idFiscal"`
	RawFiscalID  string `json:"rawFiscalId"`
	CountryID    int    `json:"countryId"`
	SupplierName string `json:"supplierName"`
	Description  string `json:"description"`
	Available    bool   `json:"available"`
}

type DtoSupplier struct {
	ID          string                `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Available   bool                  `json:"available"`
	CountryID   int                   `json:"countryId"`
	FiscalData  DtoSupplierFiscalData `json:"fiscalData"`
	Contacts    []DtoSupplierContact  `json:"contacts"`
	Products    []DtoSupplierProduct  `json:"products"`
}

type DtoSupplierFiscalData struct {
	ID            string `json:"id"`
	IDFiscal      string `json:"idFiscal"`
	RawFiscalID   string `json:"rawFiscalId"`
	FiscalName    string `json:"fiscalName"`
	FiscalAddress string `json:"fiscalAddress"`
	FiscalState   string `json:"fiscalState"`
	FiscalCity    string `json:"fiscalCity"`
	Email         string `json:"email"`
}

type DtoSupplierContact struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
}

type DtoSupplierProduct struct {
	ID          string    `json:"id"`
	SupplierID  string    `json:"supplierId"`
	ProductID   string    `json:"productId"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	SKU         string    `json:"sku"`
	Price       float32   `json:"price"`
	Unit        int       `json:"unit"`
	UnitAbv     string    `json:"unitAbv"`
	UnitName    string    `json:"unitName"`
	Available   bool      `json:"available"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
