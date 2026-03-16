package ports

import "sofia-backend/domain/models"

type PortMeasurement interface {
	GetById(id int) (*models.ModelMeasurementUnique, error)
	GetMeasurements() ([]models.ModelMeasurementUnique, error)
	CreateMeasurement(name, abbreviation, description string, baseUnitId int, conversionFactor float32) (*models.ModelMeasurementUnique, error)
	ExistsMeasurementById(id int) (bool, error)
	// GetRelatedUnits returns related measurement units via unit_conversion,
	// with factors normalized relative to the provided unitID.
	// If a related unit is a basic unit, it expands one additional level
	// and multiplies factors accordingly.
	GetRelatedUnits(unitID int) ([]models.ModelMeasurementConversionUnit, error)
}
