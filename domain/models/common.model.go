package models

type ModelCurrency struct {
	ID            int     `json:"id"`
	IsoCode       string  `json:"isoCode"`
	NumericCode   int     `json:"numericCode"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	DecimalPlaces float32 `json:"decimalPlaces"`
	Rate          float32 `json:"rate"`
	Available     bool    `json:"available"`
}
