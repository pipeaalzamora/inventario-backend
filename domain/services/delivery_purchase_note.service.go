package services

import (
	"context"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"
)

type DeliveryPurchaseNoteService struct {
	PowerChecker
	repo     ports.PortDeliveryPurchaseNote
	external ports.PortBucket
}

func NewDeliveryPurchaseNoteService(repo ports.PortDeliveryPurchaseNote, external ports.PortBucket) *DeliveryPurchaseNoteService {
	return &DeliveryPurchaseNoteService{
		repo:     repo,
		external: external,
	}
}

func (s *DeliveryPurchaseNoteService) CreateDeliveryPurchaseNote(ctx context.Context, note *recipe.RecipeDeliveryPurchaseNote) (*models.ModelDeliveryPurchaseNote, error) {
	if ok := s.EveryPower(ctx, PowerDeliveryPurchaseNoteCreate); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear notas de entrega")
	}

	if ok := s.EveryPower(ctx, PowerPrefixCompany+note.CompanyID, PowerPrefixStore+note.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}

	userId, ok := s.GetUserIDFromContext(ctx)
	if !ok {
		return nil, types.ThrowPower("No autorizado")
	}

	modelDeliveryNotePurchaseNote := &models.ModelDeliveryPurchaseNote{
		SupplierID:  note.SupplierID,
		CompanyID:   note.CompanyID,
		StoreID:     note.StoreID,
		WarehouseID: note.WarehouseID,
		Comment:     note.Comment,
		PurchaseID:  note.PurchaseID,
		Total:       note.Total,
		UserID:      userId,
		Status:      entities.DeliveryPurchaseNoteStatus(note.Status),
		Items:       make([]models.ModelDeliveryPurchaseNoteItem, 0),
	}

	for _, item := range note.Items {
		modelDeliveryNotePurchaseNote.Items = append(modelDeliveryNotePurchaseNote.Items, models.ModelDeliveryPurchaseNoteItem{
			StoreProductID: item.StoreProductID,
			Quantity:       item.Quantity,
			PurchaseUnit:   item.PurchaseUnit,
			UnitPrice:      item.UnitPrice,
			Subtotal:       item.Subtotal,
			TaxTotal:       item.TaxTotal,
			Difference:     item.Difference,
			Status:         entities.DeliveryPurchaseNoteItemStatus(item.Status),
		})
	}

	return s.repo.CreateDeliveryPurchaseNote(modelDeliveryNotePurchaseNote)
}

func (s *DeliveryPurchaseNoteService) UpdateDeliveryPurchaseNote(ctx context.Context, id string, note *recipe.RecipeDeliveryPurchaseNote) (*models.ModelDeliveryPurchaseNote, error) {
	if ok := s.EveryPower(ctx, PowerDeliveryPurchaseNoteUpdate); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar notas de entrega")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+note.CompanyID, PowerPrefixStore+note.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}

	modelDeliveryNotePurchaseNote := &models.ModelDeliveryPurchaseNote{
		SupplierID:  note.SupplierID,
		CompanyID:   note.CompanyID,
		StoreID:     note.StoreID,
		WarehouseID: note.WarehouseID,
		Comment:     note.Comment,
		PurchaseID:  note.PurchaseID,
		Total:       note.Total,
		Status:      entities.DeliveryPurchaseNoteStatus(note.Status),
		Items:       make([]models.ModelDeliveryPurchaseNoteItem, len(note.Items)),
	}

	for i, item := range note.Items {
		modelDeliveryNotePurchaseNote.Items[i] = models.ModelDeliveryPurchaseNoteItem{
			StoreProductID: item.StoreProductID,
			Quantity:       item.Quantity,
			PurchaseUnit:   item.PurchaseUnit,
			UnitPrice:      item.UnitPrice,
			Subtotal:       item.Subtotal,
			TaxTotal:       item.TaxTotal,
			Difference:     item.Difference,
			Status:         entities.DeliveryPurchaseNoteItemStatus(item.Status),
		}
	}

	return s.repo.UpdateDeliveryPurchaseNote(id, modelDeliveryNotePurchaseNote)
}

func (s *DeliveryPurchaseNoteService) GetDeliveryPurchaseNoteByID(ctx context.Context, id string) (*models.ModelDeliveryPurchaseNote, error) {
	note, err := s.repo.GetDetailDeliveryPurchaseNote(id)
	if err != nil {
		return nil, err
	}
	if note == nil {
		return nil, types.ThrowMsg("Nota de entrega no encontrada")
	}

	if ok := s.EveryPower(ctx, PowerPrefixCompany+note.CompanyID, PowerPrefixStore+note.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a esta nota de entrega")
	}

	return note, nil
}

func (s *DeliveryPurchaseNoteService) GetAllDeliveryPurchaseNotes(ctx context.Context, storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelDeliveryPurchaseNote, int, error) {
	return s.repo.GetAllDeliveryPurchaseNotes(storeID, page, size, filter)
}

func (s *DeliveryPurchaseNoteService) UpdateDeliveryPurchaseNoteStatus(ctx context.Context, id string, status entities.DeliveryPurchaseNoteStatus) error {
	note, err := s.repo.GetDetailDeliveryPurchaseNote(id)
	if err != nil {
		return err
	}
	if note == nil {
		return types.ThrowMsg("Nota de entrega no encontrada")
	}

	if ok := s.EveryPower(ctx, PowerPrefixCompany+note.CompanyID, PowerPrefixStore+note.StoreID); !ok {
		return types.ThrowPower("No tienes permiso de propiedad sobre esta nota de entrega")
	}

	return s.repo.UpdateDeliveryPurchaseNoteStatus(id, status)
}

func (s *DeliveryPurchaseNoteService) CompleteDeliveryPurchaseNote(ctx context.Context, id string, status entities.DeliveryPurchaseNoteStatus, invoiceFolio string, invoiceGuide string) error {
	note, err := s.repo.GetDetailDeliveryPurchaseNote(id)
	if err != nil {
		return err
	}
	if note == nil {
		return types.ThrowMsg("Nota de entrega no encontrada")
	}

	if ok := s.EveryPower(ctx, PowerPrefixCompany+note.CompanyID, PowerPrefixStore+note.StoreID); !ok {
		return types.ThrowPower("No tienes permiso de propiedad sobre esta nota de entrega")
	}

	return s.repo.CompleteDeliveryPurchaseNote(id, status, invoiceFolio, invoiceGuide)
}

func (s *DeliveryPurchaseNoteService) AddFileToDeliveryPurchaseNote(ctx context.Context, purchaseDeliveryNoteID string, file *recipe.UploadForm) error {
	note, err := s.repo.GetDetailDeliveryPurchaseNote(purchaseDeliveryNoteID)
	if err != nil {
		return err
	}
	if note == nil {
		return types.ThrowMsg("Nota de entrega no encontrada")
	}

	if ok := s.EveryPower(ctx, PowerPrefixCompany+note.CompanyID, PowerPrefixStore+note.StoreID); !ok {
		return types.ThrowPower("No tienes permiso de propiedad sobre esta nota de entrega")
	}

	// Open the file to get its content
	f, err := file.File.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	// Generate a unique name for the file
	url, err := s.external.UploadFile(ctx, f, file.File.Filename)
	if err != nil {
		return err
	}

	modelFile := &models.ModelFile{
		// Get the file type from the file header's MIME type
		FileType: file.File.Header.Get("Content-Type"),
		// Store the URL returned by the bucket service
		FileURL: url,
	}

	return s.repo.AddFileToDeliveryPurchaseNote(purchaseDeliveryNoteID, modelFile)
}

func (s *DeliveryPurchaseNoteService) RemoveFileFromDeliveryPurchaseNote(ctx context.Context, noteID string, fileID string) error {
	note, err := s.repo.GetDetailDeliveryPurchaseNote(noteID)
	if err != nil {
		return err
	}
	if note == nil {
		return types.ThrowMsg("Nota de entrega no encontrada")
	}

	if ok := s.EveryPower(ctx, PowerPrefixCompany+note.CompanyID, PowerPrefixStore+note.StoreID); !ok {
		return types.ThrowPower("No tienes permiso de propiedad sobre esta nota de entrega")
	}

	file, err := s.repo.GetFileByID(fileID)
	if err != nil {
		return err
	}

	err = s.external.DeleteFile(ctx, file.FileURL)
	if err != nil {
		return err
	}

	return s.repo.RemoveFileFromDeliveryPurchaseNote(fileID)
}
