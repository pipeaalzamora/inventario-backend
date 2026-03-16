package facades

import (
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"

	"github.com/gin-gonic/gin"
)

type MeasurementFacade struct {
	serviceBox *services.ServiceContainer
}

func NewMeasurementFacade(serviceBox *services.ServiceContainer) *MeasurementFacade {
	return &MeasurementFacade{
		serviceBox: serviceBox,
	}
}

func (f *MeasurementFacade) GetMeasurements() ([]models.ModelMeasurementUnique, error) {
	return f.serviceBox.MeasurementService.GetMeasurements()
}

func (f *MeasurementFacade) CreateMeasurement(gctx *gin.Context, measurementRecipe recipe.MeasurementRecipe) (*models.ModelMeasurementUnique, error) {
	return f.serviceBox.MeasurementService.CreateMeasurement(gctx, measurementRecipe)
}

func (f *MeasurementFacade) GetRelatedUnits(gctx *gin.Context, unitID int) ([]models.ModelMeasurementConversionUnit, error) {
	return f.serviceBox.MeasurementService.GetRelatedUnits(gctx, unitID)
}
