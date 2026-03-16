package ports

// PortPriceHistory define las operaciones para gestionar el historial de precios
type PortPriceHistory interface {
	// GetPreviousPrice obtiene el precio anterior de un producto (plantilla)
	GetPreviousPrice(productID string) (*float32, error)
}
