package services

import (
	"context"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/shared"
	"sofia-backend/types"

	"github.com/gin-gonic/gin"
)

type ProductMovementService struct {
	PowerChecker
	repo ports.PortProductMovement
}

func NewProductMovementService(repo ports.PortProductMovement) *ProductMovementService {
	return &ProductMovementService{
		repo: repo,
	}
}

func (s *ProductMovementService) GetAllProductMovements() ([]models.ModelProductMovement, error) {
	return s.repo.GetAllProductMovements()
}

func (s *ProductMovementService) GetProductMovementByID(movementId string) (*models.ModelProductMovement, error) {
	return s.repo.GetByProductMovementId(movementId)
}

func (s *ProductMovementService) GetAllProductMovementsByStoreProductID(ctx context.Context, storeId string, storeProductId string) ([]models.ModelProductMovement, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeId); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.repo.GetAllProductMovementsByStoreProductID(storeProductId)
}

func (s *ProductMovementService) GetAllProductMovementsByStoreId(ctx context.Context, storeId string) ([]models.ModelProductMovement, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeId); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.repo.GetAllProductMovementsByStoreID(storeId)
}

func (s *ProductMovementService) GetAllProductMovementsByWarehouseIDs(warehouseIDs []string) ([]models.ModelProductMovement, error) {
	return s.repo.GetAllProductMovementsByWarehouseIDs(warehouseIDs)
}

func (s *ProductMovementService) GetAllProductMovementsByDateRange(warehouseID string) ([]models.ModelProductMovement, error) {
	return s.repo.GetAllProductMovementsByDateRange(warehouseID)
}

func (s *ProductMovementService) CreateNewMovement(ctx context.Context, model models.ModelProductMovement) (*models.ModelProductMovement, error) {
	if ok := s.EveryPower(ctx, PowerProductMovementCreate); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear movimientos de productos")
	}
	userId := ctx.(*gin.Context).Keys[shared.UserIdKey()].(string)
	model.MovedBy = userId

	return s.repo.CreateNewSingleMovement(model)
}

func (s *ProductMovementService) CreateNewMovements(ctx context.Context, models []models.ModelProductMovement, bypass bool) ([]models.ModelProductMovement, error) {

	if !bypass {
		if ok := s.EveryPower(ctx, PowerProductMovementCreate); !ok {
			return nil, types.ThrowPower("No tienes permiso para crear movimientos de productos")
		}
		userId := ctx.(*gin.Context).Keys[shared.UserIdKey()].(string)
		for i := range models {
			models[i].MovedBy = userId
		}
	}

	return s.repo.CreateNewMultiplesMovements(models)
}

// CreateTransferMovements crea movimientos de transferencia entre bodegas
// Requiere permiso de crear movimientos de productos
func (s *ProductMovementService) CreateTransferMovements(ctx context.Context, movements []models.ModelProductMovement, newAvgCosts map[string]float32) ([]models.ModelProductMovement, error) {
	if ok := s.EveryPower(ctx, PowerProductMovementCreate); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear movimientos de productos")
	}

	userId := ctx.(*gin.Context).Keys[shared.UserIdKey()].(string)
	for i := range movements {
		movements[i].MovedBy = userId
	}

	return s.repo.CreateTransferMovements(movements, newAvgCosts)
}
