package facades

import (
	"context"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
)

type RequestFacade struct {
	appService *services.ServiceContainer
}

func NewRequestFacade(appService *services.ServiceContainer) *RequestFacade {
	return &RequestFacade{
		appService: appService,
	}
}

func (f *RequestFacade) CreateRequest(ctx context.Context, request *recipe.RecipeNewRequest) (*models.ModelRequest, error) {
	return f.appService.RequestService.CreateRequest(ctx, request)
}
