package entities

type EntityCurrency struct {
	ID            int     `db:"id"`
	IsoCode       string  `db:"iso_code"`
	NumericCode   int     `db:"numeric_code"`
	Name          string  `db:"currency_name"`
	Symbol        string  `db:"currency_symbol"`
	DecimalPlaces float32 `db:"decimal_places"`
	Rate          float32 `db:"rate"`
	Available     bool    `db:"available"`
}
