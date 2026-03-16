package ports

import (
	"sofia-backend/domain/models"
)

type CountryPort interface {
	GetCountries() ([]models.CountryModel, int, error)
	GetCountryByID(id int) (*models.CountryModel, error)
	CreateCountry(country *models.CountryModel) (*models.CountryModel, error)
	UpdateCountry(id int, country *models.CountryModel) (*models.CountryModel, error)
	DeleteCountry(id int) error
	GetCountryByIsoCode(isoCode string) (*models.CountryModel, error)
}
