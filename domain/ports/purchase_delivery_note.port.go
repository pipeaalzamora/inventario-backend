package ports

import (
	"sofia-backend/domain/models"
	"sofia-backend/infraestructure/entities"
)

type PortDeliveryPurchaseNote interface {
	CreateDeliveryPurchaseNote(deliveryPurchaseNote *models.ModelDeliveryPurchaseNote) (*models.ModelDeliveryPurchaseNote, error)
	UpdateDeliveryPurchaseNote(id string, deliveryPurchaseNote *models.ModelDeliveryPurchaseNote) (*models.ModelDeliveryPurchaseNote, error)
	GetAllDeliveryPurchaseNotes(storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelDeliveryPurchaseNote, int, error)
	GetDetailDeliveryPurchaseNote(id string) (*models.ModelDeliveryPurchaseNote, error)
	GetDetailDeliveryPurchaseNoteByOC(ocID string) ([]models.ModelDeliveryPurchaseNote, error)
	AddFileToDeliveryPurchaseNote(deliveryPurchaseNoteID string, file *models.ModelFile) error
	RemoveFileFromDeliveryPurchaseNote(fileID string) error
	GetFileByID(fileID string) (*models.ModelFile, error)
	UpdateDeliveryPurchaseNoteStatus(id string, status entities.DeliveryPurchaseNoteStatus) error
	CompleteDeliveryPurchaseNote(id string, status entities.DeliveryPurchaseNoteStatus, invoiceFolio string, invoiceGuide string) error
}
