package ports

import "sofia-backend/domain/models"

type PortProduct interface {
	GetById(id string) (*models.ModelProduct, error)

	GetAll() ([]models.ModelProduct, error)
	GetAllFull() ([]models.ModelProduct, error)

	GetByCategory(categoryID int) ([]models.ModelProduct, error)

	Create(product *models.ModelProduct) (*models.ModelProduct, error)
	Update(product *models.ModelProduct, oldProduct *models.ModelProduct) (*models.ModelProduct, error)
	Delete(id string) error

	CheckCodeExists(productID *string, codes map[int]string) (map[int]string, error)
}
