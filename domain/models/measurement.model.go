package models

type ModelMeasurementUnique struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	Description  string `json:"description"`
	IsBasic      bool   `json:"isBasic"`
}
