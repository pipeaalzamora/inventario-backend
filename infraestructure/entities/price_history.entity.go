package entities

import "time"

// EntityPriceHistory representa la entidad de historial de precios en la base de datos
type EntityPriceHistory struct {
	ID            string    `db:"id"`
	ProductID     string    `db:"product_id"`
	PreviousPrice float32   `db:"previous_price"`
	NewPrice      float32   `db:"new_price"`
	CreatedAt     time.Time `db:"created_at"`
}
