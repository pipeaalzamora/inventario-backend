package dto

import "time"

type SimpleCompanyDTO struct {
	ID          string    `json:"id"`
	CountryID   int       `json:"countryId"`
	IDFiscal    string    `json:"idFiscal"`
	CompanyName string    `json:"companyName"`
	Description *string   `json:"description"`
	ImageLogo   *string   `json:"imageLogo"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type DetailedCompanyDTO struct {
	ID          string               `json:"id"`
	CountryID   int                  `json:"countryId"`
	CompanyName string               `json:"companyName"`
	Description *string              `json:"description"`
	ImageLogo   *string              `json:"imageLogo"`
	FiscalData  CompanyFiscalDataDto `json:"fiscalData"`
	CreatedAt   time.Time            `json:"createdAt"`
	UpdatedAt   time.Time            `json:"updatedAt"`
}

type CompanyFiscalDataDto struct {
	ID            string `json:"id"`
	IDFiscal      string `json:"idFiscal"`
	FiscalName    string `json:"fiscalName"`
	FiscalAddress string `json:"fiscalAddress"`
	FiscalState   string `json:"fiscalState"`
	FiscalCity    string `json:"fiscalCity"`
	Email         string `json:"email"`
}

type CompanySupplierDTO struct {
	ID            string `json:"id"`
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
