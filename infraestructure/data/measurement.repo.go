package data

import (
	"math"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

type measurementRepo struct {
	db *sqlx.DB
}

func NewMeasurementRepo(db *sqlx.DB) ports.PortMeasurement {
	return &measurementRepo{
		db: db,
	}
}

// GetById implements ports.PortMeasurement.
func (m *measurementRepo) GetById(id int) (*models.ModelMeasurementUnique, error) {
	var entity entities.EntityMeasurementUnique
	query := `SELECT id, unit_name, abbreviation, description, basic_unit FROM measurement_unit WHERE id = $1`
	err := m.db.Get(&entity, query, id)
	if err != nil {
		return nil, types.ThrowData("ocurrió un error al traer la unidad de medida")
	}

	model := &models.ModelMeasurementUnique{
		ID:           entity.ID,
		Name:         entity.Name,
		Abbreviation: entity.Abbreviation,
		Description:  entity.Description,
		IsBasic:      entity.IsBasic,
	}

	return model, nil
}

func (m *measurementRepo) GetMeasurements() ([]models.ModelMeasurementUnique, error) {
	var entitiesList []entities.EntityMeasurementUnique
	query := `SELECT id, unit_name, abbreviation, description, basic_unit FROM measurement_unit`
	err := m.db.Select(&entitiesList, query)
	if err != nil {
		return nil, types.ThrowData("ocurrió un error al traer la lista de unidades")
	}

	measurements := make([]models.ModelMeasurementUnique, len(entitiesList))
	for i, entity := range entitiesList {
		measurements[i] = models.ModelMeasurementUnique{
			ID:           entity.ID,
			Name:         entity.Name,
			Abbreviation: entity.Abbreviation,
			Description:  entity.Description,
			IsBasic:      entity.IsBasic,
		}
	}

	return measurements, nil
}

func (m *measurementRepo) ExistsMeasurementById(id int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM measurement_unit WHERE id = $1)`
	err := m.db.Get(&exists, query, id)
	if err != nil {
		return false, types.ThrowData("error al verificar la unidad de medida")
	}
	return exists, nil
}

func (m *measurementRepo) CreateMeasurement(name, abbreviation, description string, baseUnitId int, conversionFactor float32) (*models.ModelMeasurementUnique, error) {
	tx, err := m.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("error al iniciar la transacción")
	}
	defer tx.Rollback()

	var existsName bool
	queryCheckName := `SELECT EXISTS(SELECT 1 FROM measurement_unit WHERE unit_name = $1)`
	err = tx.Get(&existsName, queryCheckName, name)
	if err != nil {
		return nil, types.ThrowData("error al verificar el nombre")
	}
	if existsName {
		return nil, types.ThrowRecipe("El nombre de la unidad ya existe", "name")
	}

	var existsAbbr bool
	queryCheckAbbr := `SELECT EXISTS(SELECT 1 FROM measurement_unit WHERE abbreviation = $1)`
	err = tx.Get(&existsAbbr, queryCheckAbbr, abbreviation)
	if err != nil {
		return nil, types.ThrowData("error al verificar la abreviación")
	}
	if existsAbbr {
		return nil, types.ThrowRecipe("La abreviación ya existe", "abbreviation")
	}

	queryUnit := `
		INSERT INTO measurement_unit (unit_name, abbreviation, description, basic_unit)
		VALUES ($1, $2, $3, false)
		RETURNING *
	`

	var id int
	var unitName, abbr, desc string
	var basicUnit bool
	err = tx.QueryRow(queryUnit, name, abbreviation, description).Scan(&id, &abbr, &unitName, &desc, &basicUnit)
	if err != nil {
		return nil, types.ThrowData("error al crear la unidad de medida")
	}

	queryConversion := `
		INSERT INTO unit_conversion (from_unit_id, to_unit_id, conversion_factor)
		VALUES ($1, $2, $3)
	`

	_, err = tx.Exec(queryConversion, id, baseUnitId, conversionFactor)
	if err != nil {
		return nil, types.ThrowData("error al crear la conversión de unidad")
	}

	if err = tx.Commit(); err != nil {
		return nil, types.ThrowData("error al confirmar la transacción")
	}

	model := &models.ModelMeasurementUnique{
		ID:           id,
		Name:         unitName,
		Abbreviation: abbr,
		Description:  desc,
		IsBasic:      basicUnit,
	}

	return model, nil
}

// GetRelatedUnits implements ports.PortMeasurement.
// It collects related measurement units via unit_conversion edges, normalizing
// conversion factors relative to the provided unitID. If a related unit is a
// basic unit, it expands one additional level and multiplies factors accordingly.
// The starting unit itself is excluded from results.
func (m *measurementRepo) GetRelatedUnits(unitID int) ([]models.ModelMeasurementConversionUnit, error) {
	// Query to get all conversions related to a unit (both directions)
	// Outgoing: from_unit_id = $1 -> factor as is
	// Incoming: to_unit_id = $1 -> invert factor (1/factor)
	query := `
		SELECT 
			mu.id, 
			mu.unit_name, 
			mu.abbreviation, 
			mu.description, 
			mu.basic_unit,
			CASE 
				WHEN uc.from_unit_id = $1 THEN uc.conversion_factor
				ELSE 1.0 / uc.conversion_factor
			END AS conversion_factor
		FROM unit_conversion uc
		JOIN measurement_unit mu ON 
			CASE 
				WHEN uc.from_unit_id = $1 THEN uc.to_unit_id = mu.id
				ELSE uc.from_unit_id = mu.id
			END
		WHERE uc.from_unit_id = $1 OR uc.to_unit_id = $1
	`

	visited := make(map[int]bool)
	visited[unitID] = true // exclude start unit from results

	results := make(map[int]models.ModelMeasurementConversionUnit)

	// Step 1: Get direct conversions from the starting unit
	var directUnits []entities.EntityMeasurementUniqueWithFactor
	if err := m.db.Select(&directUnits, query, unitID); err != nil {
		return nil, types.ThrowData("ocurrió un error al traer unidades relacionadas")
	}

	// Collect basic units that need expansion
	var basicUnitsToExpand []struct {
		id     int
		factor float32
	}

	for _, unit := range directUnits {
		if visited[unit.ID] {
			continue
		}
		visited[unit.ID] = true

		factor := roundFactor(unit.ConversionFactor)

		results[unit.ID] = models.ModelMeasurementConversionUnit{
			ID:           unit.ID,
			Name:         unit.Name,
			Abbreviation: unit.Abbreviation,
			Description:  unit.Description,
			Factor:       factor,
		}

		// If it's a basic unit, mark it for expansion
		if unit.IsBasic {
			basicUnitsToExpand = append(basicUnitsToExpand, struct {
				id     int
				factor float32
			}{id: unit.ID, factor: factor})
		}
	}

	// Step 2: Expand one level from each basic unit found
	for _, basicUnit := range basicUnitsToExpand {
		var secondLevelUnits []entities.EntityMeasurementUniqueWithFactor
		if err := m.db.Select(&secondLevelUnits, query, basicUnit.id); err != nil {
			return nil, types.ThrowData("ocurrió un error al traer unidades relacionadas (segundo nivel)")
		}

		for _, unit := range secondLevelUnits {
			if visited[unit.ID] {
				continue
			}
			visited[unit.ID] = true

			// Multiply: factor from start to basic unit * factor from basic unit to this unit
			accumulatedFactor := roundFactor(float64(basicUnit.factor) * unit.ConversionFactor)

			results[unit.ID] = models.ModelMeasurementConversionUnit{
				ID:           unit.ID,
				Name:         unit.Name,
				Abbreviation: unit.Abbreviation,
				Description:  unit.Description,
				Factor:       accumulatedFactor,
			}
		}
	}

	// Build slice from map
	out := make([]models.ModelMeasurementConversionUnit, 0, len(results))
	for _, v := range results {
		out = append(out, v)
	}
	return out, nil
}

// roundFactor rounds a conversion factor to avoid floating-point precision issues
func roundFactor(val float64) float32 {
	if val == 0 {
		return 0
	}
	r := math.Round(val*1e9) / 1e9
	if math.Abs(r-math.Round(r)) < 1e-6 {
		r = math.Round(r)
	}
	return float32(r)
}
