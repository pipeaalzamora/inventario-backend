package services

import (
	"context"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
	"strings"
)

type SupplierService struct {
	PowerChecker
	SupplierRepo ports.PortSupplier
}

func NewSupplierService(supplierRepo ports.PortSupplier) *SupplierService {
	return &SupplierService{
		SupplierRepo: supplierRepo,
	}
}

// //////////////// SUPPLIER CRUD ///////////////////
func (s *SupplierService) GetSuppliersByStoreProductId(storeID string, productID []string) (models.ModelProductSuppliersMap, error) {
	// Create a dictionary for each product ID with a {price, supplierID} structure
	productSupplier, err := s.SupplierRepo.GetSuppliersByStoreProductId(storeID, productID)
	if err != nil {
		return nil, err
	}

	suppliers := make(models.ModelProductSuppliersMap)
	for _, p := range productSupplier {
		suppliers[p.ProductID] = append(suppliers[p.ProductID], models.SupplierOption{
			SupplierID: p.SupplierID,
			Price:      p.Price,
		})
	}

	return suppliers, nil
}

func (s *SupplierService) GetSupplierProductsByProductID(productID string) ([]models.ModelSupplierProduct, error) {
	return s.SupplierRepo.GetSupplierProductsByProductID(productID)
}

func (s *SupplierService) GetSuppliers() ([]models.ModelSupplier, error) {
	return s.SupplierRepo.GetAllSuppliers()
}

func (s *SupplierService) GetSupplierByID(ctx context.Context, supplierID string) (*models.ModelSupplier, error) {
	return s.SupplierRepo.GetSupplierByID(supplierID)
}

func (s *SupplierService) GetSupplierByFiscalIDAndCountry(fiscalID string, countryID int) (*models.ModelSupplier, error) {
	return s.SupplierRepo.GetSupplierByFiscalIDAndCountry(fiscalID, countryID)
}

func (s *SupplierService) CreateSupplier(ctx context.Context, supplier recipe.RecipeCreateSupplier) (*models.ModelSupplier, error) {
	if ok := s.SomePower(ctx, PowerSupplierCreate); !ok {
		return nil, types.ThrowPower("No tienes los permisos para crear proveedor")
	}

	rawFiscalID := strings.TrimPrefix(supplier.IDFiscal, "CL-")
	rawFiscalID = strings.ReplaceAll(rawFiscalID, ".", "")

	contacts := make([]models.ModelSupplierContact, len(supplier.Contacts))
	for i, contact := range supplier.Contacts {
		contacts[i] = models.ModelSupplierContact{
			Name:        contact.Name,
			Description: contact.Description,
			Email:       contact.Email,
			Phone:       contact.Phone,
		}
	}

	model := &models.ModelSupplier{
		CountryID:    1,
		SupplierName: supplier.Name,
		Description:  *supplier.Description,
		Available:    true,
		FiscalData: models.ModelFiscalData{
			IDFiscal:      supplier.IDFiscal,
			RawFiscalID:   rawFiscalID,
			FiscalName:    supplier.FiscalName,
			FiscalAddress: supplier.FiscalAddress,
			FiscalState:   supplier.FiscalState,
			FiscalCity:    supplier.FiscalCity,
			Email:         supplier.Email,
		},
		Contacts: contacts,
	}

	return s.SupplierRepo.CreateSupplier(model)
}

func (s *SupplierService) UpdateSupplier(ctx context.Context, supplier *models.ModelSupplier, ogSupplier *models.ModelSupplier) (*models.ModelSupplier, error) {
	if ok := s.SomePower(ctx, PowerSupplierUpdate); !ok {
		return nil, types.ThrowPower("No tienes los permisos para actualizar proveedor")
	}

	return s.SupplierRepo.UpdateSupplier(supplier, ogSupplier)
}

func (s *SupplierService) EnableDisableSupplier(id string, available bool) error {
	return s.SupplierRepo.EnableDisableSupplier(id, available)
}

func (s *SupplierService) GetSupplierByFiscalNameAndCountry(fiscalName string, countryID int) (*models.ModelSupplier, error) {
	return s.SupplierRepo.GetSupplierByFiscalIDAndCountry(fiscalName, countryID)
}

func (s *SupplierService) GetSupplierContactsEmailByID(supplierId string) (string, error) {
	supplier, err := s.SupplierRepo.GetSupplierByID(supplierId)
	if err != nil {
		return "", err
	}

	if supplier == nil {
		return "", types.ThrowData("Proveedor no encontrado")
	}

	email, err := s.SupplierRepo.GetSupplierEmail(supplier.FiscalData.ID)
	if err != nil {
		return "", err
	}
	return email, nil
}

// //////////////// SUPPLIER PRODUCTS CRUD ///////////////////
func (s *SupplierService) GetSupplierProducts(supplierID string) ([]models.ModelSupplierProduct, error) {
	return s.SupplierRepo.GetSupplierProducts(supplierID)
}

func (s *SupplierService) GetSupplierProductById(supplierID, productID string) (*models.ModelSupplierProduct, error) {
	return s.SupplierRepo.GetSupplierProductById(supplierID, productID)
}

func (s *SupplierService) GetSupplierProductBySku(supplierID, sku string) (*models.ModelSupplierProduct, error) {
	return s.SupplierRepo.GetSupplierProductBySku(supplierID, sku)
}

func (s *SupplierService) AddProductToSupplier(supplierID string, product *models.ModelSupplierProduct) (*models.ModelSupplierProduct, error) {
	return s.SupplierRepo.AddProductToSupplier(supplierID, product)
}

func (s *SupplierService) UpdateSupplierProductsPrice(ctx context.Context, supplierID string, products []models.ModelSupplierProduct) ([]models.ModelSupplierProduct, error) {
	if ok := s.SomePower(ctx, PowerSupplierUpdate); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar productos de proveedor")
	}

	return s.SupplierRepo.UpdateSupplierProductsPrices(supplierID, products)
}

func (s *SupplierService) UpdateSupplierProduct(ctx context.Context, supplierID string, product *models.ModelSupplierProduct) (*models.ModelSupplierProduct, error) {
	if ok := s.SomePower(ctx, PowerSupplierUpdate); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar productos de proveedor")
	}
	return s.SupplierRepo.UpdateSupplierProduct(supplierID, product)
}

func (s *SupplierService) DeleteSupplierProduct(ctx context.Context, supplierID, productID string) (*models.ModelSupplierProduct, error) {
	if ok := s.EveryPower(ctx, PowerSupplierDelete); !ok {
		return nil, types.ThrowPower("No tienes permiso para eliminar productos de proveedor")
	}
	return s.SupplierRepo.DeleteSupplierProduct(supplierID, productID)
}
