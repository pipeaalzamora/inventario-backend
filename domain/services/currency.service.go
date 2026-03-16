package services

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
)

type ServiceCurrency struct {
	PowerChecker
	repo ports.PortCurrency
}

func NewCurrencyService(repo ports.PortCurrency) *ServiceCurrency {
	return &ServiceCurrency{
		repo: repo,
	}
}

func (s *ServiceCurrency) GetById(id int) (*models.ModelCurrency, error) {
	return s.repo.GetById(id)
}

func (s *ServiceCurrency) GetByIsoCode(isoCode string) (*models.ModelCurrency, error) {
	return s.repo.GetByIsoCode(isoCode)
}

func (s *ServiceCurrency) GetByNumericCode(numericCode int) (*models.ModelCurrency, error) {
	return s.repo.GetByNumericCode(numericCode)
}

func (s *ServiceCurrency) GetAll() ([]models.ModelCurrency, error) {
	return s.repo.GetAll()
}

func (s *ServiceCurrency) Create(currency *models.ModelCurrency) (*models.ModelCurrency, error) {
	return s.repo.Create(currency)
}

func (s *ServiceCurrency) Update(currency *models.ModelCurrency) (*models.ModelCurrency, error) {
	return s.repo.Update(currency)
}

func (s *ServiceCurrency) Enable(currency *models.ModelCurrency) (*models.ModelCurrency, error) {
	return s.repo.EnableDisable(currency, true)
}

func (s *ServiceCurrency) Disable(currency *models.ModelCurrency) (*models.ModelCurrency, error) {
	return s.repo.EnableDisable(currency, false)
}
