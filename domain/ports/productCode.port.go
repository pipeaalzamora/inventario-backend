package ports

import "sofia-backend/domain/models"

type PortProductCode interface {
	GetAll() ([]models.ModelProductCodeKind, error)
	GetAllByProductId(productId string) ([]models.ModelProductCode, error)
}
