package ports

import "sofia-backend/domain/models"

type PortRequest interface {
	CreateRequest(request *models.ModelRequest) (*models.ModelRequest, error)
}
