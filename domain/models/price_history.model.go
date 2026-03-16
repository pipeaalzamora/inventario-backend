package models

import "time"

// ModelPriceHistory representa el historial de cambios de precio de un producto (plantilla)
type ModelPriceHistory struct {
	ID            string    `json:"id"`
	ProductID     string    `json:"productId"`
	PreviousPrice float32   `json:"previousPrice"`
	NewPrice      float32   `json:"newPrice"`
	CreatedAt     time.Time `json:"createdAt"`
}
