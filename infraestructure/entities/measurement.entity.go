package entities

type EntityMeasurementUnique struct {
	ID           int    `db:"id"`
	Name         string `db:"unit_name"`
	Abbreviation string `db:"abbreviation"`
	Description  string `db:"description"`
	IsBasic      bool   `db:"basic_unit"`
}

type EntityMeasurementUniqueWithFactor struct {
	EntityMeasurementUnique
	ConversionFactor float64 `db:"conversion_factor"`
}
