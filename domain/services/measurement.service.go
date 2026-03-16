package services

import (
	"context"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
)

type MeasurementService struct {
	PowerChecker

	MeasurementRepo ports.PortMeasurement
}

func NewMeasurementService(measurementRepo ports.PortMeasurement) *MeasurementService {
	return &MeasurementService{
		MeasurementRepo: measurementRepo,
	}
}

func (ms *MeasurementService) GetMeasurements() ([]models.ModelMeasurementUnique, error) {
	return ms.MeasurementRepo.GetMeasurements()
}

func (ms *MeasurementService) CreateMeasurement(ctx context.Context, measurementRecipe recipe.MeasurementRecipe) (*models.ModelMeasurementUnique, error) {
	if measurementRecipe.ConversionFactor <= 0 {
		return nil, types.ThrowRecipe("El factor de conversión debe ser mayor a 0", "conversionFactor")
	}

	exists, err := ms.MeasurementRepo.ExistsMeasurementById(measurementRecipe.BaseUnitId)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, types.ThrowRecipe("La unidad base no existe", "baseUnitId")
	}

	return ms.MeasurementRepo.CreateMeasurement(
		measurementRecipe.Name,
		measurementRecipe.Abbreviation,
		measurementRecipe.Description,
		measurementRecipe.BaseUnitId,
		measurementRecipe.ConversionFactor,
	)
}

func (ms *MeasurementService) GetRelatedUnits(ctx context.Context, unitID int) ([]models.ModelMeasurementConversionUnit, error) {
	exists, err := ms.MeasurementRepo.ExistsMeasurementById(unitID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, types.ThrowRecipe("La unidad de medida no existe", "unitId")
	}

	return ms.MeasurementRepo.GetRelatedUnits(unitID)
}
