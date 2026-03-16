package ports

import (
	"sofia-backend/domain/models"
)

type PortFiscalData interface {
	CreateFiscalData(fiscalData *models.ModelFiscalData) (*models.ModelFiscalData, error)
	UpdateFiscalData(id string, fiscalData *models.ModelFiscalData) (*models.ModelFiscalData, error)
	DeleteFiscalData(id string) error
	GetFiscalDataByCompanyID(companyID string) ([]models.ModelFiscalData, error)
	UpdateFiscalDataWithEconomicActivities(fiscalDataID string, input *models.ModelFiscalData, economicActivityIDs []string) (*models.ModelFiscalData, error)

	GetEconomicActivities() ([]models.EconomicActivityModel, error)
	GetEconomicActivitiesByFiscalDataID(fiscalDataID string) ([]models.EconomicActivityModel, error)
	CreateEconomicActivity(economicActivity *models.EconomicActivityModel) (*models.EconomicActivityModel, error)
	UpdateEconomicActivity(id string, economicActivity *models.EconomicActivityModel) (*models.EconomicActivityModel, error)
	DeleteEconomicActivity(id string) error

	GetEconomicActivityClasses() ([]models.EconomicActivityClassModel, error)
	CreateEconomicActivityClass(economicActivityClass *models.EconomicActivityClassModel) (*models.EconomicActivityClassModel, error)
	UpdateEconomicActivityClass(id string, economicActivityClass *models.EconomicActivityClassModel) (*models.EconomicActivityClassModel, error)
	DeleteEconomicActivityClass(id string) error
}
