package ports

import "sofia-backend/domain/models"

type PortCurrency interface {
	GetById(id int) (*models.ModelCurrency, error)
	GetByIsoCode(isoCode string) (*models.ModelCurrency, error)
	GetByNumericCode(numericCode int) (*models.ModelCurrency, error)

	GetAll() ([]models.ModelCurrency, error)

	Create(currency *models.ModelCurrency) (*models.ModelCurrency, error)
	Update(currency *models.ModelCurrency) (*models.ModelCurrency, error)

	EnableDisable(currency *models.ModelCurrency, value bool) (*models.ModelCurrency, error)
}
