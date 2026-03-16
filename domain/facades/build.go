package facades

import (
	"sofia-backend/config"
	"sofia-backend/domain/external"
	"sofia-backend/domain/services"
)

type FacadeContainer struct {
	UserFacade                 *UserFacade
	NotificationFacade         *NotificationFacade
	AuthFacade                 *AuthFacade
	ProfileFacade              *ProfileFacade
	ProductFacade              *ProductFacade
	StoreFacade                *StoreFacade
	CompanyFacade              *CompanyFacade
	WarehouseFacade            *WarehouseFacade
	InventoryFacade            *InventoryFacade
	InventoryCountFacade       *InventoryCountFacade
	InventoryRequestFacade     *InventoryRequestFacade
	PurchaseFacade             *PurchaseFacade
	SupplierFacade             *SupplierFacade
	SupplierOCFacade           *SupplierOCFacade
	ProductMovementFacade      *ProductMovementFacade
	DeliveryPurchaseNoteFacade *DeliveryPurchaseNoteFacade
	MeasurementFacade          *MeasurementFacade
	StoreProductFacade         *StoreProductFacade
	RequestFacade              *RequestFacade
}

func Build(
	config *config.Config,
	services *services.ServiceContainer,
	external *external.ServiceContainer,
) *FacadeContainer {
	return &FacadeContainer{
		UserFacade: NewUserFacade(
			services,
			external,
			config,
		),
		RequestFacade:      NewRequestFacade(services),
		NotificationFacade: NewNotificationFacade(services),
		AuthFacade: NewAuthFacade(
			services,
			external,
			config,
		),
		ProfileFacade:        NewProfileFacade(services),
		ProductFacade:        NewProductFacade(services),
		StoreFacade:          NewStoreFacade(services),
		CompanyFacade:        NewCompanyFacade(services),
		WarehouseFacade:      NewWarehouseFacade(services),
		InventoryFacade:      NewInventoryFacade(services),
		InventoryCountFacade: NewInventoryCountFacade(services),
		//InventoryRequestFacade: NewInventoryRequestFacade(services, external, config),
		//PurchaseFacade:         NewPurchaseFacade(services, external, config),
		SupplierFacade: NewSupplierFacade(services),
		//SupplierOCFacade:       NewSupplierOCFacade(services),
		ProductMovementFacade: NewProductMovementFacade(services),
		DeliveryPurchaseNoteFacade: NewDeliveryPurchaseNoteFacade(
			services,
			external,
			config,
		),
		MeasurementFacade:  NewMeasurementFacade(services),
		StoreProductFacade: NewStoreProductFacade(services),
	}
}
