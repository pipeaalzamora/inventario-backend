package entities

import (
	"time"
)

type EntityCompany struct {
	ID           string    `json:"id" db:"id"`
	CountryID    int       `json:"country_id" db:"country_id"`
	FiscalDataID string    `json:"fiscal_data_id" db:"fiscal_data_id"`
	CompanyName  string    `json:"company_name" db:"company_name"`
	Description  *string   `json:"description" db:"description"`
	ImageLogo    *string   `json:"image_logo" db:"image_logo"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
