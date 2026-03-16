package ports

import (
	"sofia-backend/domain/models"
	"sofia-backend/infraestructure/entities"
)

type PortPurchase interface {
	CreatePurchaseOrder(purchase *models.ModelPurchase) (*models.ModelPurchase, error)
	CreatePurchaseOrderApproved(purchase *models.ModelPurchase) (*models.ModelPurchase, error)
	GetAllPurchase(storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelPurchase, int, error)
	GetPurchaseByID(purchaseID string) (*models.ModelPurchase, error)
	GetPurchasesByInventoryRequestID(inventoryRequestID string) ([]models.ModelPurchase, error)
	CancelPurchase(purchaseID string, observation string) error
	CreatePurchaseOrderWithInventoryRequest(purchase *models.ModelPurchase, request *models.ModelInventoryRequest) (*models.ModelPurchase, error)
	AddDeliveryNoteIdAndSetArrivedStatus(purchaseID string, deliveryNoteID string) error
	AddSonOCToPurchase(purchaseID string, sonDisplayID string) error
	UpdatePurchaseState(purchaseID string, state entities.PurchaseStatus, observation string) error
	UpdatePurchaseItemsStatus(purchaseID string, items []models.ModelPurchaseItem) error
	GetPurchasesByDeliveryPurchaseNote(id string) ([]models.ModelPurchase, error)
	ApprovePurchase(purchaseID string) error
}
