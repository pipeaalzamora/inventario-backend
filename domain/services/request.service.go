package services

import (
	"context"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/types"
)

type RequestService struct {
	PowerChecker
	requestRepo ports.PortRequest
	storeRepo   ports.PortStore
	companyRepo ports.PortCompany
}

func NewRequestService(
	requestRepo ports.PortRequest,
	storeRepo ports.PortStore,
	companyRepo ports.PortCompany,
) *RequestService {
	return &RequestService{
		requestRepo: requestRepo,
		storeRepo:   storeRepo,
		companyRepo: companyRepo,
	}
}

func (s *RequestService) CreateRequest(ctx context.Context, request *recipe.RecipeNewRequest) (*models.ModelRequest, error) {

	switch request.RequestType {
	case models.RequestKindStore.ToString():
		if ok := s.EveryPower(ctx, PowerRequestCreateForStore); !ok {
			return nil, types.ThrowPower("No tienes permiso para crear solicitudes para tiendas")
		}
	case models.RequestKindCompany.ToString():
		if ok := s.EveryPower(ctx, PowerRequestCreateForCompany); !ok {
			return nil, types.ThrowPower("No tienes permiso para crear solicitudes para tiendas de otra empresa")
		}
	case models.RequestKindPurchase.ToString():
		if ok := s.EveryPower(ctx, PowerRequestCreateForSupplier); !ok {
			return nil, types.ThrowPower("No tienes permiso para crear solicitudes para proveedores")
		}
	default:
		return nil, types.ThrowRecipe("Tipo de solicitud no válido", "requestType")
	}

	//verificar si la empresa y la tienda son de la misma empresa
	_store, err := s.storeRepo.GetStoreByCompanyID(request.CompanyID)
	if err != nil {
		return nil, err
	}
	if _store.CompanyID != request.CompanyID {
		return nil, types.ThrowPower("La empresa y la tienda no son de la misma empresa")
	}

	_company, err := s.companyRepo.GetCompanyByID(request.CompanyID)
	if err != nil {
		return nil, err
	}
	if _company == nil {
		return nil, types.ThrowData("Empresa no encontrada")
	}

	_userContext, ok := s.GetUserFromContext(ctx)
	if !ok {
		return nil, types.ThrowData("No se pudo obtener el usuario del contexto")
	}

	_status := models.ModelRequestStatus{Id: 1, Name: "Creada"}
	_newRequest := &models.ModelRequest{
		CompanyId:   request.CompanyID,
		StoreId:     _store.ID,
		WarehouseId: nil,
		Status:      _status,
		RequestKind: models.ModelRequestKind{
			Id:   models.RequestKind(request.RequestType).ToInt(),
			Name: request.RequestType,
		},
		CreatedBy: models.ModelRequestUser{
			Id:        _userContext.ID,
			UserName:  _userContext.UserName,
			UserEmail: _userContext.UserEmail,
		},
		Items: make([]models.ModelRequestItem, 0),
	}
	return s.requestRepo.CreateRequest(_newRequest)
}
