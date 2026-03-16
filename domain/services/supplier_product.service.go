package services

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
)

type SupplierProductService struct {
	PowerChecker
	SupplierProductRepo ports.PortSupplierProduct
}

func NewSupplierProductService(supplierProductRepo ports.PortSupplierProduct) *SupplierProductService {
	return &SupplierProductService{
		SupplierProductRepo: supplierProductRepo,
	}
}

func (s *SupplierProductService) GetSupplierProductsByStoreIDAndSupplierIDWithProductCompanyIDs(
	storeID string,
	supplierID string,
	productCompanyIDs []string,
) (models.ModelMapProductCompanyKeySupplierProduct, error) {
	supplierProducts, err := s.SupplierProductRepo.GetSupplierProductsByStoreIDAndSupplierIDWithProductCompanyIDs(storeID, supplierID, productCompanyIDs)
	if err != nil {
		return nil, err
	}

	result := make(models.ModelMapProductCompanyKeySupplierProduct)
	for _, sp := range supplierProducts {
		result[sp.ProductCompanyID] = sp
	}

	return result, nil
}

func (s *SupplierProductService) GetSupplierProductsByStoreIDAndSupplierID(storeID string, supplierID string) ([]models.ModelSupplierProductLegacy, error) {
	return s.SupplierProductRepo.GetSupplierProductsByStoreIDAndSupplierID(storeID, supplierID)
}
