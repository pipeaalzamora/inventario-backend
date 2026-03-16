package services

import (
	"sofia-backend/domain/ports"
)

type InventoryRequestService struct {
	PowerChecker
	inventoryRequestRepo ports.PortInventoryRequest
}

func NewInventoryRequestService(inventoryRequestRepo ports.PortInventoryRequest) *InventoryRequestService {
	return &InventoryRequestService{
		inventoryRequestRepo: inventoryRequestRepo,
	}
}

/*
func (s *InventoryRequestService) GetAllInventoryRequests(ctx context.Context, storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelInventoryRequest, int, error) {
	if ok := s.EveryPower(ctx, PowerPrefixStore+storeID); !ok {
		return nil, 0, types.ThrowPower("No tienes permiso para acceder a esta tienda")
	}
	return s.inventoryRequestRepo.GetInventoryRequestsByStore(storeID, page, size, filter)
}

func (s *InventoryRequestService) GetInventoryRequestByID(ctx context.Context, requestID string) (*models.ModelInventoryRequest, error) {
	request, err := s.inventoryRequestRepo.GetInventoryRequestByID(requestID)
	if err != nil {
		return nil, err
	}
	if request == nil {
		return nil, types.ThrowMsg("Solicitud no encontrada")
	}

	if ok := s.EveryPower(ctx, PowerPrefixCompany+request.CompanyID, PowerPrefixStore+request.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta solicitud")
	}

	return request, nil
}

func (s *InventoryRequestService) CreateInventoryRequest(ctx context.Context, request *recipe.RecipeInventoryRequest, status entities.RequestStatus) (*models.ModelInventoryRequest, error) {
	if ok := s.EveryPower(ctx, PowerRequestCreate); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear solicitudes de inventario")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+request.CompanyID, PowerPrefixStore+request.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}

	if request == nil {
		return nil, types.ThrowData("La solicitud no puede ser nula")
	}
	if len(request.Items) == 0 {
		return nil, types.ThrowData("Se requiere al menos un item en la solicitud")
	}

	modelRequest := &models.ModelInventoryRequest{
		CompanyID:   request.CompanyID,
		StoreID:     request.StoreID,
		WarehouseID: request.WarehouseID,
		Status:      entities.RequestStatus(status),
		RequestType: entities.RequestType(request.RequestType),
		RequesterID: request.RequesterID,
		Items:       make([]models.ModelInventoryRequestItem, len(request.Items)),
	}

	for i, item := range request.Items {
		modelRequest.Items[i] = models.ModelInventoryRequestItem{
			Quantity:       item.Quantity,
			StoreProductID: item.ItemID,
		}
	}

	return s.inventoryRequestRepo.CreateInventoryRequest(modelRequest, request.Observation)
}

func (s *InventoryRequestService) UpdateInventoryRequest(ctx context.Context, id string, request *recipe.RecipeInventoryRequest, status entities.RequestStatus) (*models.ModelInventoryRequest, error) {
	if ok := s.EveryPower(ctx, PowerRequestUpdate); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar solicitudes de inventario")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+request.CompanyID, PowerPrefixStore+request.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}

	user, ok := s.GetUserFromContext(ctx)
	if !ok || user == nil {
		return nil, types.ThrowData("Usuario no encontrado en el contexto")
	}

	if request == nil {
		return nil, types.ThrowData("La solicitud no puede ser nula")
	}
	if len(request.Items) == 0 {
		return nil, types.ThrowData("Se requiere al menos un item en la solicitud")
	}

	modelRequest := &models.ModelInventoryRequest{
		CompanyID:   request.CompanyID,
		StoreID:     request.StoreID,
		WarehouseID: request.WarehouseID,
		Status:      status,
		RequestType: entities.RequestType(request.RequestType),
		RequesterID: request.RequesterID,
		Items:       make([]models.ModelInventoryRequestItem, len(request.Items)),
	}

	for i, item := range request.Items {
		modelRequest.Items[i] = models.ModelInventoryRequestItem{
			Quantity:       item.Quantity,
			StoreProductID: item.ItemID,
		}
	}

	return s.inventoryRequestRepo.UpdateInventoryRequest(id, modelRequest, request.Observation, user.ID)
}

func (s *InventoryRequestService) ChangeStatus(ctx context.Context, id string, status entities.RequestStatus, observation *string) (*models.ModelInventoryRequest, error) {
	user, ok := s.GetUserFromContext(ctx)
	if !ok || user == nil {
		return nil, types.ThrowData("Usuario no encontrado en el contexto")
	}
	return s.inventoryRequestRepo.ChangeInventoryRequestStatus(id, status, observation, user.ID)
}

func (s *InventoryRequestService) ApproveAndUpdateRequest(ctx context.Context, id string, request *recipe.RecipeInventoryRequest) (*models.ModelInventoryRequest, error) {
	if ok := s.EveryPower(ctx, PowerRequestResolve, fmt.Sprintf("%s%s", PowerPrefixStore, request.StoreID)); !ok {
		return nil, shared.PowerError{
			Message: fmt.Sprintf(
				"User has not the required powers: %v",
				[]string{PowerPurchaseCreate, PowerRequestCreate}),
		}
	}

	user, ok := s.GetUserFromContext(ctx)
	if !ok || user == nil {
		return nil, types.ThrowData("Usuario no encontrado en el contexto")
	}

	if len(request.Items) == 0 {
		return nil, types.ThrowData("Se requiere al menos un item en la solicitud")
	}

	modelRequest := &models.ModelInventoryRequest{
		StoreID:     request.StoreID,
		WarehouseID: request.WarehouseID,
		Status:      "approved", // Assuming "approved" is the status for approved requests
		RequestType: entities.RequestType(request.RequestType),
		RequesterID: request.RequesterID,
		Items:       make([]models.ModelInventoryRequestItem, len(request.Items)),
	}

	for i, item := range request.Items {
		modelRequest.Items[i] = models.ModelInventoryRequestItem{
			Quantity:       item.Quantity,
			StoreProductID: item.ItemID,
		}
	}

	return s.inventoryRequestRepo.ApproveAndUpdateRequest(id, modelRequest, request.Observation, user.ID)
}
*/
