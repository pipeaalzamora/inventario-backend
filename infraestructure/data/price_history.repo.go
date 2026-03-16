package data

import (
	"sofia-backend/domain/ports"

	"github.com/jmoiron/sqlx"
)

// PriceHistoryRepo implementa PortPriceHistory
type PriceHistoryRepo struct {
	db *sqlx.DB
}

// NewPriceHistoryRepo crea una nueva instancia del repositorio
func NewPriceHistoryRepo(db *sqlx.DB) ports.PortPriceHistory {
	return &PriceHistoryRepo{db: db}
}

// GetPreviousPrice obtiene el precio anterior de un producto desde el historial
func (r *PriceHistoryRepo) GetPreviousPrice(productID string) (*float32, error) {
	var previousPrice float32
	query := `
		SELECT previous_price FROM price_history 
		WHERE product_id = $1 
		ORDER BY created_at DESC 
		LIMIT 1
	`
	err := r.db.Get(&previousPrice, query, productID)
	if err != nil {
		return nil, nil // No hay historial, retornar nil
	}
	return &previousPrice, nil
}
