package services

import (
	"sofia-backend/domain/ports"
)

type PurchaseService struct {
	PowerChecker
	repo ports.PortPurchase
}

func NewPurchaseService(repo ports.PortPurchase) *PurchaseService {
	return &PurchaseService{
		repo: repo,
	}
}

/*
func (s *PurchaseService) CreatePurchaseOrder(ctx context.Context, purchase *recipe.RecipePurchase) (*models.ModelPurchase, error) {
	if ok := s.EveryPower(ctx, PowerPurchaseCreate); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear ordenes de compra")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+purchase.CompanyID, PowerPrefixStore+purchase.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}
	modelItems := make([]models.ModelPurchaseItem, len(purchase.Items))
	for i, item := range purchase.Items {
		modelItems[i] = models.ModelPurchaseItem{
			StoreProductID:  item.StoreProductID,
			Quantity:        item.Quantity,
			UnitPrice:       item.UnitPrice,
			SupplierOptions: item.SupplierOptions,
		}
	}
	modelPurchase := &models.ModelPurchase{
		SupplierID:         purchase.SupplierID,
		CompanyID:          purchase.CompanyID,
		StoreID:            purchase.StoreID,
		InventoryRequestID: purchase.InventoryRequestID,
		Items:              modelItems,
	}
	return s.repo.CreatePurchaseOrder(modelPurchase)
}

func (s *PurchaseService) CreatePurchaseOrderApprovedWithoutAuth(ctx context.Context, purchase *recipe.RecipePurchase) (*models.ModelPurchase, error) {
	// if ok := s.EveryPower(ctx, PowerPurchaseCreate); !ok {
	// 	return nil, types.ThrowPower("No tienes permiso para crear ordenes de compra")
	// }
	// if ok := s.EveryPower(ctx, PowerPrefixCompany+purchase.CompanyID, PowerPrefixStore+purchase.StoreID); !ok {
	// 	return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	// }
	modelItems := make([]models.ModelPurchaseItem, len(purchase.Items))
	for i, item := range purchase.Items {
		modelItems[i] = models.ModelPurchaseItem{
			StoreProductID:  item.StoreProductID,
			Quantity:        item.Quantity,
			UnitPrice:       item.UnitPrice,
			SupplierOptions: item.SupplierOptions,
		}
	}
	modelPurchase := &models.ModelPurchase{
		SupplierID:         purchase.SupplierID,
		CompanyID:          purchase.CompanyID,
		StoreID:            purchase.StoreID,
		InventoryRequestID: purchase.InventoryRequestID,
		Items:              modelItems,
	}
	return s.repo.CreatePurchaseOrderApproved(modelPurchase)
}

func (s *PurchaseService) UpdatePurchaseItemsStatus(purchaseID string, items []models.ModelPurchaseItem) error {

	if len(items) == 0 {
		return types.ThrowData("Los items no pueden estar vacíos")
	}

	return s.repo.UpdatePurchaseItemsStatus(purchaseID, items)
}

func (s *PurchaseService) UpdatePurchaseState(purchaseID string, state entities.PurchaseStatus, obsevation string) error {
	return s.repo.UpdatePurchaseState(purchaseID, state, obsevation)
}

func (s *PurchaseService) AddDeliveryNoteIdAndSetArrivedStatus(purchaseID string, deliveryNoteID string) error {
	return s.repo.AddDeliveryNoteIdAndSetArrivedStatus(purchaseID, deliveryNoteID)
}

func (s *PurchaseService) AddSonOCToPurchase(parentID string, childID string) error {
	return s.repo.AddSonOCToPurchase(parentID, childID)
}

func (s *PurchaseService) GetAllPurchase(ctx context.Context, storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelPurchase, int, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeID); !ok {
		return nil, 0, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.repo.GetAllPurchase(storeID, page, size, filter)
}

func (s *PurchaseService) GetPurchaseByID(ctx context.Context, purchaseID string, baypassPower bool) (*models.ModelPurchase, error) {
	purchase, err := s.repo.GetPurchaseByID(purchaseID)
	if err != nil {
		return nil, err
	}
	if purchase == nil {
		return nil, types.ThrowMsg("Orden de compra no encontrada")
	}

	if !baypassPower {
		if ok := s.EveryPower(ctx, PowerPrefixCompany+purchase.CompanyID, PowerPrefixStore+purchase.StoreID); !ok {
			return nil, types.ThrowPower("No tienes permiso para acceder a esta orden de compra")
		}
	}

	return purchase, nil
}

func (s *PurchaseService) GetPurchasesByInventoryRequestID(inventoryRequestID string) ([]models.ModelPurchase, error) {
	return s.repo.GetPurchasesByInventoryRequestID(inventoryRequestID)
}

func (s *PurchaseService) CreatePurchaseOrderWithInventoryRequest(
	ctx context.Context,
	purchase *recipe.RecipePurchase,
	userId string,
) (*models.ModelPurchase, error) {
	if ok := s.EveryPower(ctx, PowerPurchaseCreate, PowerRequestCreate); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear ordenes de compra o solicitudes")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+purchase.CompanyID, PowerPrefixStore+purchase.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}

	modelPurchase := &models.ModelPurchase{
		SupplierID:         purchase.SupplierID,
		CompanyID:          purchase.CompanyID,
		StoreID:            purchase.StoreID,
		InventoryRequestID: purchase.InventoryRequestID,
		Items:              make([]models.ModelPurchaseItem, len(purchase.Items)),
	}

	for i, item := range purchase.Items {
		modelPurchase.Items[i] = models.ModelPurchaseItem{
			StoreProductID: item.StoreProductID,
			Quantity:       item.Quantity,
			PurchaseUnit:   item.PurchaseUnit,
			UnitPrice:      item.UnitPrice,
		}
	}

	modelRequest := &models.ModelInventoryRequest{
		CompanyID:   purchase.CompanyID,
		StoreID:     purchase.StoreID,
		WarehouseID: purchase.WarehouseID,
		Status:      entities.RequestStatusApproved,
		RequestType: entities.RequestTypePurchase,
		RequesterID: userId,
		Items:       make([]models.ModelInventoryRequestItem, len(purchase.Items)),
	}

	for i, item := range purchase.Items {
		modelRequest.Items[i] = models.ModelInventoryRequestItem{
			Quantity:       item.Quantity,
			StoreProductID: item.StoreProductID,
		}
	}

	return s.repo.CreatePurchaseOrderWithInventoryRequest(modelPurchase, modelRequest)
}

func (s *PurchaseService) CancelPurchase(ctx context.Context, purchaseID string, observation string) error {
	if ok := s.EveryPower(ctx, PowerPurchaseUpdate, PowerRequestUpdate); !ok {
		return shared.PowerError{
			Message: fmt.Sprintf(
				"User has not the required powers: %v",
				[]string{PowerPurchaseUpdate, PowerRequestUpdate}),
		}
	}

	return s.repo.CancelPurchase(purchaseID, observation)
}

func (s *PurchaseService) ApprovePurchase(ctx context.Context, purchaseID string) error {
	if ok := s.EveryPower(ctx, PowerPurchaseApprove, PowerPurchaseUpdate); !ok {
		return shared.PowerError{
			Message: fmt.Sprintf(
				"User has not the required power: %v",
				PowerPurchaseApprove),
		}
	}

	return s.repo.ApprovePurchase(purchaseID)
}
*/
