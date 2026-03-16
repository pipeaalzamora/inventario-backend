package models

import (
	"time"
)

type ModelCompany struct {
	ID          string          `json:"id"`
	CountryID   int             `json:"countryId"`
	FiscalData  ModelFiscalData `json:"fiscalData"`
	CompanyName string          `json:"companyName"`
	Description *string         `json:"description"`
	ImageLogo   *string         `json:"imageLogo"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
}

type CompanySupplierModel struct {
	SupplierID    string `json:"id"`
	SupplierName  string `json:"supplierName"`
	Description   string `json:"description"`
	Available     bool   `json:"available"`
	CountryID     int    `json:"countryId"`
	IDFiscal      string `json:"idFiscal"`
	RawFiscalID   string `json:"rawFiscalId"`
	FiscalName    string `json:"fiscalName"`
	FiscalAddress string `json:"fiscalAddress"`
	Email         string `json:"email"`
}
