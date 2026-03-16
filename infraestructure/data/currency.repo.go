package data

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"

	"github.com/jmoiron/sqlx"
)

type currencyRepo struct {
	db *sqlx.DB
}

func NewCurrencyRepo(db *sqlx.DB) ports.PortCurrency {
	return &currencyRepo{
		db: db,
	}
}

// GetById implements ports.PortCurrency.
func (c *currencyRepo) GetById(id int) (*models.ModelCurrency, error) {
	var currencyEntity entities.EntityCurrency

	query := "SELECT * FROM currencies WHERE id = $1"
	if err := c.db.Get(&currencyEntity, query, id); err != nil {
		return nil, err
	}
	return c.toModel(&currencyEntity), nil
}

// GetByIsoCode implements ports.PortCurrency.
func (c *currencyRepo) GetByIsoCode(isoCode string) (*models.ModelCurrency, error) {
	var currencyEntity entities.EntityCurrency

	query := "SELECT * FROM currencies WHERE iso_code = $1"
	if err := c.db.Get(&currencyEntity, query, isoCode); err != nil {
		return nil, err
	}
	return c.toModel(&currencyEntity), nil
}

// GetByNumericCode implements ports.PortCurrency.
func (c *currencyRepo) GetByNumericCode(numericCode int) (*models.ModelCurrency, error) {
	var currencyEntity entities.EntityCurrency

	query := "SELECT * FROM currencies WHERE numeric_code = $1"
	if err := c.db.Get(&currencyEntity, query, numericCode); err != nil {
		return nil, err
	}
	return c.toModel(&currencyEntity), nil
}

// GetAll implements ports.PortCurrency.
func (c *currencyRepo) GetAll() ([]models.ModelCurrency, error) {
	var currencyEntities []entities.EntityCurrency

	query := "SELECT * FROM currencies"
	if err := c.db.Select(&currencyEntities, query); err != nil {
		return nil, err
	}
	return c.toModelMap(currencyEntities), nil
}

// Create implements ports.PortCurrency.
func (c *currencyRepo) Create(currency *models.ModelCurrency) (*models.ModelCurrency, error) {
	query := `
		INSERT INTO currency (iso_code, numeric_code, currency_name, currency_symbol, decimal_places, rate, available)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`
	var currencyCreated entities.EntityCurrency
	if err := c.db.QueryRow(
		query,
		currency.IsoCode,
		currency.NumericCode,
		currency.Name,
		currency.Symbol,
		currency.DecimalPlaces,
		currency.Rate,
		currency.Available,
	).Scan(
		&currencyCreated.ID,
		&currencyCreated.IsoCode,
		&currencyCreated.NumericCode,
		&currencyCreated.Name,
		&currencyCreated.Symbol,
		&currencyCreated.DecimalPlaces,
		&currencyCreated.Rate,
		&currencyCreated.Available,
	); err != nil {
		return nil, err
	}
	return c.toModel(&currencyCreated), nil
}

// Update implements ports.PortCurrency.
func (c *currencyRepo) Update(currency *models.ModelCurrency) (*models.ModelCurrency, error) {
	query := `
		UPDATE currency
		SET 
			iso_code = $1, 
			numeric_code = $2, 
			currency_name = $3, 
			currency_symbol = $4, 
			decimal_places = $5, 
			rate = $6, 
		WHERE id = $7
	`
	if _, err := c.db.Exec(
		query,
		currency.IsoCode,
		currency.NumericCode,
		currency.Name,
		currency.Symbol,
		currency.DecimalPlaces,
		currency.Rate,
		currency.ID,
	); err != nil {
		return nil, err
	}
	return currency, nil
}

// Delete implements ports.PortCurrency.
func (c *currencyRepo) EnableDisable(currency *models.ModelCurrency, value bool) (*models.ModelCurrency, error) {
	query := "UPDATE currency SET available = $1 WHERE id = $2"
	if _, err := c.db.Exec(query, value, currency.ID); err != nil {
		return nil, err
	}

	currency.Available = value
	return currency, nil
}

func (c *currencyRepo) toModel(entity *entities.EntityCurrency) *models.ModelCurrency {
	return &models.ModelCurrency{
		ID:            entity.ID,
		IsoCode:       entity.IsoCode,
		NumericCode:   entity.NumericCode,
		Name:          entity.Name,
		Symbol:        entity.Symbol,
		DecimalPlaces: entity.DecimalPlaces,
		Rate:          entity.Rate,
		Available:     entity.Available,
	}
}

func (c *currencyRepo) toModelMap(entities []entities.EntityCurrency) []models.ModelCurrency {
	models := make([]models.ModelCurrency, len(entities))
	for i, e := range entities {
		models[i] = *c.toModel(&e)
	}
	return models
}
