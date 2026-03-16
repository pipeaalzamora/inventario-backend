package data

import (
	"sofia-backend/domain/ports"

	"github.com/jmoiron/sqlx"
)

type DataContainer struct {
	UserAccountRepo          ports.PortUser
	NotificationRepo         ports.PortNotification
	ProfileAccountRepo       ports.PortProfile
	PowerAccountRepo         ports.PortPower
	CurrencyRepo             ports.PortCurrency
	CategoryRepo             ports.PortCategory
	ProductCodeRepo          ports.PortProductCode
	ProductRepo              ports.PortProduct
	SupplierRepo             ports.PortSupplier
	StoreRepo                ports.PortStore
	CompanyRepo              ports.PortCompany
	WarehouseRepo            ports.PortWarehouse
	InventoryRequestRepo     ports.PortInventoryRequest
	InventoryCountRepo       ports.PortInventoryCount
	InventoryRepo            ports.PortInventory
	PurchaseRepo             ports.PortPurchase
	ProductPerStoreRepo      ports.PortProductPerStore // Nueva: Reemplaza ProductCompanyRepo
	SupplierProductRepo      ports.PortSupplierProduct
	SupplierOCRepo           ports.PortSupplierOC
	DeliveryPurchaseNoteRepo ports.PortDeliveryPurchaseNote
	WarehouseProductRepo     ports.PortWarehouseProduct
	WasteRepo                ports.PortProductMovement
	MeasurementRepo          ports.PortMeasurement
	RequestRepo              ports.PortRequest
	PriceHistoryRepo         ports.PortPriceHistory		
	// DEPRECATED: ProductCompanyRepo - usar ProductPerStoreRepo
}

func Build(
	// mongoDB *mongo.Database,
	postgresDB *sqlx.DB,
) *DataContainer {
	powerRepo := NewPowerAccountRepo(postgresDB)
	_categoryRepo := NewCategoryRepo(postgresDB)
	_productCodeRepo := NewProductCodeRepo(postgresDB)
	_wareHouseRepo := NewWarehouseRepo(postgresDB)
	_productRepo := NewProductRepo(postgresDB, _categoryRepo, _productCodeRepo)
	_measurementRepo := NewMeasurementRepo(postgresDB)
	_requestRepo := NewRequestRepo(postgresDB)
	return &DataContainer{
		UserAccountRepo:          NewUserAccountRepo(postgresDB),
		NotificationRepo:         NewNotificationRepo(postgresDB),
		ProfileAccountRepo:       NewProfileAccountRepo(postgresDB),
		PowerAccountRepo:         powerRepo,
		CurrencyRepo:             NewCurrencyRepo(postgresDB),
		CategoryRepo:             _categoryRepo,
		ProductCodeRepo:          _productCodeRepo,
		ProductRepo:              _productRepo,
		SupplierRepo:             NewSupplierRepo(postgresDB),
		StoreRepo:                NewStoreRepo(postgresDB, powerRepo, _wareHouseRepo),
		CompanyRepo:              NewCompanyRepo(postgresDB, powerRepo),
		WarehouseRepo:            NewWarehouseRepo(postgresDB),
		InventoryRequestRepo:     NewInventoryRequestRepo(postgresDB),
		InventoryCountRepo:       NewInventoryCountRepo(postgresDB),
		InventoryRepo:            NewInventoryRepo(postgresDB),
		PurchaseRepo:             NewPurchaseRepo(postgresDB),
		ProductPerStoreRepo:      NewProductPerStoreRepo(postgresDB, _productRepo, _measurementRepo), // Nueva: Reemplaza ProductCompanyRepo
		SupplierProductRepo:      NewSupplierProductRepo(postgresDB),
		SupplierOCRepo:           NewSupplierOCRepo(postgresDB),
		DeliveryPurchaseNoteRepo: NewDeliveryPurchaseNoteRepo(postgresDB),
		WarehouseProductRepo:     NewWarehouseProductRepo(postgresDB),
		WasteRepo:                NewProductMovementRepo(postgresDB),
		MeasurementRepo:          _measurementRepo,
		RequestRepo:              _requestRepo,
		PriceHistoryRepo: 		  NewPriceHistoryRepo(postgresDB),
	}
}
