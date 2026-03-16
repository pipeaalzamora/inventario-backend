package services

// DEPRECATED: ProductCompanyService ha sido reemplazado por ProductPerStoreService
// Este servicio se mantiene comentado para referencia durante la migración.
// Ver: product_per_store.service.go
/*
import (
	"context"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
)

type ProductCompanyService struct {
	PowerChecker
	repo ports.PortProductCompany
}

func NewProductCompanyService(repo ports.PortProductCompany) *ProductCompanyService {
	return &ProductCompanyService{
		repo: repo,
	}
}

func (s *ProductCompanyService) GetProductsByCompany(ctx context.Context, companyID string) ([]models.ModelProductCompany, error) {
	if ok := s.EveryPower(ctx, PowerPrefixCompany+companyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta empresa")
	}
	return s.repo.GetProductsByCompany(companyID)
}

func (s *ProductCompanyService) GetProductCompanyByID(ctx context.Context, productCompanyID string) (*models.ModelProductCompany, error) {
	productCompany, err := s.repo.GetProductCompanyByID(productCompanyID)
	if err != nil {
		return nil, err
	}
	if productCompany == nil {
		return nil, nil
	}

	// Validate company ownership
	if ok := s.EveryPower(ctx, PowerPrefixCompany+productCompany.CompanyID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a este producto")
	}

	return productCompany, nil
}

func (s *ProductCompanyService) GetProductsResquestRestriction(ctx context.Context, storeId string) ([]models.ModelProductRequestRestriction, int, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeId); !ok {
		return nil, 0, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	products, err := s.repo.GetAllProductRequestRestrictionByStoreId(storeId)
	if err != nil {
		return nil, 0, err
	}
	return products, 0, nil
}

func (s *ProductCompanyService) GetOnlyRestrictedProducts(ctx context.Context, storeId string) ([]models.ModelProductRequestRestriction, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeId); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.repo.GetProductsRequestRetrictedByStoreId(storeId)
}
*/
