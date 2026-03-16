package ports

import "sofia-backend/domain/models"

type PortCategory interface {
	GetById(id string) (*models.ModelProductCategory, error)
	GetByName(name string) (*models.ModelProductCategory, error)

	GetAll() ([]models.ModelProductCategory, error)

	Create(category *models.ModelProductCategory) (*models.ModelProductCategory, error)
	Update(category *models.ModelProductCategory) (*models.ModelProductCategory, error)

	EnableDisable(category *models.ModelProductCategory, value bool) (*models.ModelProductCategory, error)

	GetAllByProductId(productId string) ([]models.ModelProductCategory, error)
}
