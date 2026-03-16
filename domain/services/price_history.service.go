package services

import (
	"sofia-backend/domain/ports"
)

// PriceHistoryService maneja la consulta del historial de precios
// El registro lo hace el trigger de PostgreSQL automáticamente
type PriceHistoryService struct {
	priceHistoryRepo ports.PortPriceHistory
}

// NewPriceHistoryService crea una nueva instancia del servicio
func NewPriceHistoryService(priceHistoryRepo ports.PortPriceHistory) *PriceHistoryService {
	return &PriceHistoryService{
		priceHistoryRepo: priceHistoryRepo,
	}
}

// GetPreviousPrice obtiene el precio anterior de un producto (plantilla)
func (s *PriceHistoryService) GetPreviousPrice(productID string) (*float32, error) {
	return s.priceHistoryRepo.GetPreviousPrice(productID)
}
