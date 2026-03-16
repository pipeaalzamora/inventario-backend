package services

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/shared"
	"sofia-backend/types"
	"strings"
	"time"
)

type InventoryCountService struct {
	PowerChecker
	repo          ports.PortInventoryCount
	cacheService  ports.PortCache
	bucketService ports.PortBucket
}

func NewInventoryCountService(repo ports.PortInventoryCount, cacheService ports.PortCache, bucketService ports.PortBucket) *InventoryCountService {
	return &InventoryCountService{
		repo:          repo,
		cacheService:  cacheService,
		bucketService: bucketService,
	}
}

func (s *InventoryCountService) GetAll(ctx context.Context) ([]models.ModelInventoryCount, error) {
	powers, ok := s.GetPowersFromContext(ctx)
	if !ok {
		return nil, nil
	}

	filteredCompanyID := map[string]bool{}
	filteredStoreID := map[string]bool{}
	for _, p := range powers {
		if strings.HasPrefix(p, PowerPrefixCompany) {
			companyId := p[len(PowerPrefixCompany):]
			filteredCompanyID[companyId] = true
		}
		if strings.HasPrefix(p, PowerPrefixStore) {
			storeId := p[len(PowerPrefixStore):]
			filteredStoreID[storeId] = true
		}
	}

	allCounts, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	result := []models.ModelInventoryCount{}
	for _, count := range allCounts {
		if filteredCompanyID[count.CompanyID] && filteredStoreID[count.StoreID] {
			result = append(result, count)
		}
	}

	return result, nil
}

func (s *InventoryCountService) GetAllByUserId(ctx context.Context, userId string) ([]models.ModelInventoryCount, error) {
	powers, ok := s.GetPowersFromContext(ctx)
	if !ok {
		return make([]models.ModelInventoryCount, 0), nil
	}

	filteredCompanyID := map[string]bool{}
	filteredStoreID := map[string]bool{}
	for _, p := range powers {
		if strings.HasPrefix(p, PowerPrefixCompany) {
			companyId := p[len(PowerPrefixCompany):]
			filteredCompanyID[companyId] = true
		}
		if strings.HasPrefix(p, PowerPrefixStore) {
			storeId := p[len(PowerPrefixStore):]
			filteredStoreID[storeId] = true
		}
	}

	allCounts, err := s.repo.GetAllByUserId(userId)
	if err != nil {
		return make([]models.ModelInventoryCount, 0), err
	}

	result := make([]models.ModelInventoryCount, 0)
	for _, count := range allCounts {
		if filteredCompanyID[count.CompanyID] && filteredStoreID[count.StoreID] {
			result = append(result, count)
		}
	}

	return result, nil
}

func (s *InventoryCountService) GetByID(ctx context.Context, id string) (*models.ModelInventoryCount, error) {
	count, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if count == nil {
		return nil, nil
	}

	if ok := s.EveryPower(ctx, PowerPrefixCompany+count.CompanyID, PowerPrefixStore+count.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso para acceder a este conteo de inventario")
	}

	return count, nil
}

// GetByIDWithoutPermissionCheck obtiene un conteo sin validar permisos de propiedad
// Útil para validaciones internas cuando se valida por asignación de usuario
func (s *InventoryCountService) GetByIDWithoutPermissionCheck(id string) (*models.ModelInventoryCount, error) {
	return s.repo.GetByID(id)
}

// GetIncidenceByProduct obtiene la incidencia existente para un producto en un conteo
func (s *InventoryCountService) GetIncidenceByProduct(countId string, productId string) (*models.ModelInventoryCountItem, error) {
	return s.repo.GetIncidenceByProduct(countId, productId)
}

// DeleteIncidenceImage elimina la imagen de una incidencia existente
func (s *InventoryCountService) DeleteIncidenceImage(ctx context.Context, countId string, productId string, observation *string) error {
	// Obtener la incidencia existente para obtener la URL de la imagen
	existingIncidence, err := s.repo.GetIncidenceByProduct(countId, productId)
	if err != nil {
		return err
	}

	// Validar que existe la incidencia y tiene imagen
	if existingIncidence == nil {
		return types.ThrowMsg("No hay imagen para eliminar")
	}

	if existingIncidence.IncidenceImageURL == nil || *existingIncidence.IncidenceImageURL == "" {
		return types.ThrowMsg("No hay imagen para eliminar")
	}

	// Eliminar el archivo del storage
	imageUrl := *existingIncidence.IncidenceImageURL
	err = s.bucketService.DeleteFile(ctx, imageUrl)
	if err != nil {
		return types.ThrowData("Error al eliminar la imagen del storage")
	}

	// Actualizar la base de datos (poner incidence_image_url en NULL)
	return s.repo.DeleteIncidenceImage(countId, productId, observation)
}
func (s *InventoryCountService) GetCompletedByID(itemsMetadata []models.ModelInventoryCountMetadata) ([]models.ModelInventoryCountItem, error) {
	return s.repo.GetCompletedByID(itemsMetadata)
}

func (s *InventoryCountService) GetItemsByInventoryCountID(id string) ([]models.ModelInventoryCountItem, error) {
	return s.repo.GetItemsByInventoryCountID(id)
}

func (s *InventoryCountService) Create(ctx context.Context, model *models.ModelInventoryCount) (*models.ModelInventoryCount, error) {
	if ok := s.EveryPower(ctx, PowerInventoryCountCreate); !ok {
		return nil, types.ThrowPower("No tienes permiso para crear conteos de inventario")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+model.CompanyID, PowerPrefixStore+model.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}
	return s.repo.Create(model)
}

func (s *InventoryCountService) Update(ctx context.Context, model *models.ModelInventoryCount) (*models.ModelInventoryCount, error) {
	if ok := s.EveryPower(ctx, PowerInventoryCountUpdate); !ok {
		return nil, types.ThrowPower("No tienes permiso para actualizar conteos de inventario")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+model.CompanyID, PowerPrefixStore+model.StoreID); !ok {
		return nil, types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}
	return s.repo.Update(model)
}

func (s *InventoryCountService) Delete(ctx context.Context, id string) error {
	// First get the entity to validate ownership
	count, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	if count == nil {
		return types.ThrowMsg("No se encontro el conteo de inventario")
	}
	if ok := s.EveryPower(ctx, PowerPrefixCompany+count.CompanyID, PowerPrefixStore+count.StoreID); !ok {
		return types.ThrowPower("No tienes permiso de propiedad sobre esta compañía o tienda")
	}
	return s.repo.Delete(id)
}

func (s *InventoryCountService) Commit(model *models.ModelInventoryCount) (*models.ModelInventoryCount, error) {
	return s.repo.Commit(model)
}

func (s *InventoryCountService) ChangeState(id string, newState string) (*models.ModelInventoryCount, error) {
	return s.repo.ChangeState(id, newState)
}

func (s *InventoryCountService) SetNewAssigned(id string, userId *string) (*models.ModelInventoryCount, error) {
	return s.repo.ChangeAssigned(id, userId)
}

func (s *InventoryCountService) SetNewDate(id string, newDate time.Time) (*models.ModelInventoryCount, error) {
	return s.repo.ChangeDate(id, newDate)
}

func (s *InventoryCountService) Reject(model *models.ModelInventoryCount) (*models.ModelInventoryCount, error) {
	return s.repo.Reject(model)
}

func (s *InventoryCountService) SaveDraft(id string, draft *models.ModelInventoryCount) error {
	key := fmt.Sprintf("DRAFT:INVENTORY_COUNT:%s", id)
	return s.cacheService.AddKeyValue(key, draft, 0)
}

func (s *InventoryCountService) GetDraft(id string) (*models.ModelInventoryCount, error) {
	key := fmt.Sprintf("DRAFT:INVENTORY_COUNT:%s", id)
	var draft models.ModelInventoryCount
	err := s.cacheService.GetWithKey(key, &draft)
	if err != nil {
		return nil, err
	}
	return &draft, nil
}

func (s *InventoryCountService) DeleteDraft(id string) error {
	key := fmt.Sprintf("DRAFT:INVENTORY_COUNT:%s", id)
	return s.cacheService.DeleteByKey(key)
}

func (s *InventoryCountService) SaveIncidence(ctx context.Context, countId string, productId string, base64Image *string, mimeType *string, observation *string) error {
	// Validar que el conteo existe
	count, err := s.repo.GetByID(countId)
	if err != nil {
		return err
	}
	if count == nil {
		return types.ThrowMsg("Conteo de inventario no encontrado")
	}

	// NO validar permisos de propiedad aquí - se valida en el facade que el usuario esté asignado

	// Validar que el producto existe en el conteo
	productFound := false
	for _, item := range count.CountItems {
		if item.ProductID == productId {
			productFound = true
			break
		}
	}
	if !productFound {
		return types.ThrowMsg("El producto no está en este conteo de inventario")
	}

	var imageUrl *string
	var finalObservation *string

	// Si se envía imagen, procesarla
	if base64Image != nil && *base64Image != "" {
		if mimeType == nil || *mimeType == "" {
			return types.ThrowMsg("mimeType es requerido cuando se envía imagen")
		}

		// Validar mimeType
		if !shared.ValidateMimeType(*mimeType) {
			return types.ThrowMsg("Tipo de imagen no soportado. Solo se permiten: image/jpeg, image/png, image/webp")
		}

		// Limpiar y convertir base64 a bytes
		cleanedBase64 := shared.CleanBase64String(*base64Image)
		imageData, err := base64.StdEncoding.DecodeString(cleanedBase64)
		if err != nil {
			return types.ThrowMsg("Formato base64 inválido")
		}

		// Validar tamaño
		if int64(len(cleanedBase64)) > shared.MaxBase64Size {
			return types.ThrowMsg("La imagen excede el tamaño máximo permitido (2.7MB)")
		}

		// Generar nombre de archivo
		ext := shared.GetExtensionFromMimeType(*mimeType)
		fileName := fmt.Sprintf("incidencias/incidencia_%s.%s", shared.GenerateUUID(), ext)

		// Crear bytes.Reader que implementa io.Reader
		imageReader := bytes.NewReader(imageData)

		// Crear un tipo que implemente multipart.File
		file := &bytesReaderFile{
			reader: imageReader,
		}

		// Subir imagen a bucket
		uploadedUrl, err := s.bucketService.UploadFile(ctx, file, fileName)
		if err != nil {
			return types.ThrowData("Error al subir la imagen")
		}

		imageUrl = &uploadedUrl
	}

	// Si se envía observación, usarla
	if observation != nil {
		finalObservation = observation
	}

	// Guardar en repositorio (solo actualizar campos no nulos)
	return s.repo.SaveIncidence(countId, productId, imageUrl, finalObservation)
}

// bytesReaderFile implementa multipart.File para poder usar con BucketService
type bytesReaderFile struct {
	reader *bytes.Reader
}

func (b *bytesReaderFile) Read(p []byte) (n int, err error) {
	return b.reader.Read(p)
}

func (b *bytesReaderFile) ReadAt(p []byte, off int64) (n int, err error) {
	return b.reader.ReadAt(p, off)
}

func (b *bytesReaderFile) Seek(offset int64, whence int) (int64, error) {
	return b.reader.Seek(offset, whence)
}

func (b *bytesReaderFile) Close() error {
	return nil
}
