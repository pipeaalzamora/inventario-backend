package services

import (
	"sofia-backend/config"
	"sofia-backend/infraestructure/data"
	externalservices "sofia-backend/infraestructure/external-services"
)

type ServiceContainer struct {
	UserService         *UserService
	NotificationService *NotificationService
	ProfileService      *ProfileService
	AuthService         *AuthService
	CurrencyService     *ServiceCurrency
	ProductService      *ProductService
	StoreService        *StoreService
	CompanyService      *CompanyService
	WarehouseService    *WarehouseService
	//InventoryRequestService     *InventoryRequestService
	InventoryCountService *InventoryCountService
	SupplierService       *SupplierService
	//PurchaseService             *PurchaseService
	ProductPerStoreService *ProductPerStoreService // Nueva: Reemplaza ProductCompanyService
	SupplierProductService *SupplierProductService
	//SupplierOCService           *SupplierOCService
	DeliveryPurchaseNoteService *DeliveryPurchaseNoteService
	WarehousePerProductService  *WarehousePerProductService
	ProductMovementService      *ProductMovementService
	MeasurementService          *MeasurementService
	InventoryService            *InventoryService
	RequestService              *RequestService
	PriceHistoryService         *PriceHistoryService
	// DEPRECATED: ProductCompanyService - usar ProductPerStoreService
}

func Build(
	config *config.Config,
	data *data.DataContainer,
	externalServices *externalservices.ExternalServicesContainer,
) *ServiceContainer {
	_inventoryService := NewInventoryService(data.InventoryRepo)
	_requestService := NewRequestService(data.RequestRepo, data.StoreRepo, data.CompanyRepo)
	return &ServiceContainer{
		UserService:         NewUserService(data.UserAccountRepo, data.ProfileAccountRepo),
		NotificationService: NewNotificationService(data.NotificationRepo, externalServices.SSEservice),
		ProfileService:      NewProfileService(data.ProfileAccountRepo, data.PowerAccountRepo, externalServices.CacheService),
		AuthService:         NewAuthService(externalServices.CacheService, config),
		CurrencyService:     NewCurrencyService(data.CurrencyRepo),
		ProductService:      NewProductService(data.ProductRepo, data.ProductCodeRepo, externalServices.BucketService, data.CategoryRepo),
		StoreService:        NewStoreService(data.StoreRepo, data.ProfileAccountRepo, data.UserAccountRepo, externalServices.CacheService, data.PowerAccountRepo),

		CompanyService:   NewCompanyService(data.CompanyRepo, data.ProfileAccountRepo, externalServices.BucketService, externalServices.CacheService, data.UserAccountRepo),
		WarehouseService: NewWarehouseService(data.WarehouseRepo, data.StoreRepo),
		//InventoryRequestService:     NewInventoryRequestService(data.InventoryRequestRepo),
		InventoryCountService: NewInventoryCountService(data.InventoryCountRepo, externalServices.CacheService, externalServices.BucketService),
		SupplierService:       NewSupplierService(data.SupplierRepo),
		//PurchaseService:             NewPurchaseService(data.PurchaseRepo),
		ProductPerStoreService: NewProductPerStoreService(data.ProductPerStoreRepo, data.StoreRepo, data.PriceHistoryRepo), // Nueva: Reemplaza ProductCompanyService
		SupplierProductService: NewSupplierProductService(data.SupplierProductRepo),
		//SupplierOCService:           NewSupplierOCService(data.SupplierOCRepo),
		DeliveryPurchaseNoteService: NewDeliveryPurchaseNoteService(data.DeliveryPurchaseNoteRepo, externalServices.BucketService),
		WarehousePerProductService:  NewWarehousePerProductService(data.WarehouseProductRepo),
		ProductMovementService:      NewProductMovementService(data.WasteRepo),
		MeasurementService:          NewMeasurementService(data.MeasurementRepo),
		InventoryService:            _inventoryService,
		RequestService:              _requestService,
		PriceHistoryService:         NewPriceHistoryService(data.PriceHistoryRepo),
	}
}
