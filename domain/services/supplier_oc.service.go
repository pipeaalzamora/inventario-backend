package services

import (
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/shared"
	"sofia-backend/types"
	"time"
)

type SupplierOCService struct {
	repo ports.PortSupplierOC
}

func NewSupplierOCService(repo ports.PortSupplierOC) *SupplierOCService {
	return &SupplierOCService{
		repo: repo,
	}
}

func (s *SupplierOCService) GetSupplierOC(token string) (*models.ModelSupplierToken, error) {
	// hashear token
	hashedToken := shared.CreateSaltyHash(token)

	// verificar si el token está persistido

	// buscar OC
	supplierOC, err := s.repo.GetSupplierOC(hashedToken)
	if err != nil {
		return nil, err
	}

	if supplierOC.Used {
		return nil, types.ThrowMsg("El token ya fue utilizado")
	}

	if supplierOC.Exp.Before(time.Now()) {
		return nil, types.ThrowMsg("El token ha expirado")
	}

	return supplierOC, nil
}

func (s *SupplierOCService) CreateSupplierOC(inventoryRequestID string, token string, exp *time.Time) error {

	// hashear token
	tokenHash := shared.CreateSaltyHash(token)

	// crear modelo
	model := &models.ModelSupplierToken{
		PurchaseID: inventoryRequestID,
		TokenHash:  tokenHash,
		Exp:        exp,
		Used:       false,
	}

	_, err := s.repo.CreateSupplierOC(model)
	if err != nil {
		return err
	}

	return nil
}

func (s *SupplierOCService) UpdateSupplierOC(hashedToken string, recipe recipe.SupplierOCRecipe) error {
	supplierOC, err := s.repo.GetSupplierOC(hashedToken)
	if err != nil {
		return err
	}

	if supplierOC.Used {
		return types.ThrowMsg("el token ya fue utilizado")
	}

	if supplierOC.Exp.Before(time.Now()) {
		return types.ThrowMsg("el token ha expirado")
	}

	updatedSupplierOC, err := s.repo.UpdateSupplierOC(supplierOC)
	if err != nil {
		return err
	}

	if !updatedSupplierOC.Used {
		return types.ThrowMsg("no se pudo actualizar el token de proveedor")
	}

	// aqui hay que actualizar la orden de compra con los aprobados y rechazados

	return nil
}
