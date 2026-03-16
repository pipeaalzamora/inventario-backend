package facades

import (
	"context"
	"fmt"
	"sofia-backend/api/v1/dto"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/shared"
	"sofia-backend/types"
	"time"

	"github.com/gin-gonic/gin"
)

type InventoryCountFacade struct {
	box *services.ServiceContainer
}

func NewInventoryCountFacade(box *services.ServiceContainer) *InventoryCountFacade {
	return &InventoryCountFacade{box: box}
}

func (f *InventoryCountFacade) GetAll(ctx *gin.Context) (shared.PaginationResponse[models.ModelInventoryCount], error) {
	emptyResponse := shared.PaginationResponse[models.ModelInventoryCount]{}

	response, err := f.box.InventoryCountService.GetAll(ctx.Request.Context())
	if err != nil {
		return emptyResponse, err
	}

	return shared.NewPagination(response, len(response), 1, len(response)), nil
}

func (f *InventoryCountFacade) GetAllByUserId(ctx *gin.Context, userId string) ([]models.ModelInventoryCount, error) {
	emptyResponse := make([]models.ModelInventoryCount, 0)

	_, err := f.box.UserService.GetUserByID(ctx, userId)
	if err != nil {
		return emptyResponse, err
	}

	response, err := f.box.InventoryCountService.GetAllByUserId(ctx, userId)
	if err != nil {
		return emptyResponse, err
	}

	for i, ic := range response {
		draft, err := f.box.InventoryCountService.GetDraft(ic.ID)
		if err == nil && draft != nil {
			response[i].Metadata = draft.Metadata
		}
	}

	return response, nil
	//return shared.NewPagination(f.toGeneralDTOList(response), len(response), 1, len(response)), nil
}

func (f *InventoryCountFacade) GetById(ctx *gin.Context, id string) (*dto.DtoInventoryCountDetail, error) {

	inventoryCount, err := f.box.InventoryCountService.GetDraft(id)

	if err == nil && inventoryCount != nil {
		return f.toDTO(inventoryCount), nil
	}

	inventoryCount, err = f.box.InventoryCountService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	/* if inventoryCount.Status == "completed" {
		completedItems, err := f.getCompleatedItems(inventoryCount.Metadata)
		if err != nil {
			return nil, err
		}
		inventoryCount.CountItems = completedItems
	} */

	return f.toDTO(inventoryCount), nil

}

func (f *InventoryCountFacade) Create(ctx *gin.Context, companyId string, recipeNew recipe.RecipeCreateInventoryCount) (*dto.DtoInventoryCountDetail, error) {
	//TODO: validar permisos

	//todo: deberia haber un endpoint para traer la tienda por id (para verificar si existe?)

	_, err := f.box.CompanyService.GetCompanyByID(ctx, companyId)
	if err != nil {
		return nil, err
	}

	storeWarehouses, err := f.box.WarehouseService.GetWarehousesByStoreId(ctx, recipeNew.StoreID)
	if err != nil {
		return nil, err
	}

	var warehouse *models.ModelWarehouse
	for _, w := range storeWarehouses {
		if w.ID == recipeNew.WarehouseID {
			warehouse = &w
		}
	}

	if warehouse == nil {
		return nil, types.ThrowMsg("La bodega no pertenece a la tienda seleccionada")
	}

	userId := ctx.Keys[shared.UserIdKey()].(string)

	status := "pending"
	if recipeNew.AssignedTo == nil {
		status = "un_assigned"
	}

	//truncar now a utc con dia
	now := time.Now().UTC()

	if recipeNew.DueDate == nil || recipeNew.DueDate.Before(now) {
		recipeNew.DueDate = &now
	}

	// normalizar tiempo para que sea solo yyyy-mm-dd HH:00:00
	dateTruncated := recipeNew.DueDate.UTC()
	recipeNew.DueDate = &dateTruncated

	if len(recipeNew.Observations) == 0 {
		recipeNew.Observations = "Sin Observaciones"
	}

	if len(recipeNew.ProductsID) == 0 {
		return nil, types.ThrowMsg("Se debe proporcionar al menos un ID de producto")
	}

	products := make([]models.ModelInventoryCountItem, 0)
	metadata := make([]models.ModelInventoryCountMetadata, 0)

	for _, ID := range recipeNew.ProductsID {
		if product, err := f.box.ProductPerStoreService.GetProductPerStoreByID(ctx, ID); err != nil {
			return nil, types.ThrowMsg("Producto con ID: '" + ID + "' no fue encontrado")
		} else {

			products = append(products, models.ModelInventoryCountItem{
				ProductID:   product.ID,
				ScheduledAt: *recipeNew.DueDate,
			})

			//se inicializa la lista de unidades con la unidad base
			unitCounts := []models.ModelInventoryUnitsCount{
				{
					UnitId:  product.UnitInventory.ID,
					UnitAbv: product.UnitInventory.Abbreviation,
					Count:   0,
					Factor:  1,
				},
			}

			//se cargan las demas unidades si existen
			for _, unit := range product.UnitMatrix {
				metadataUnit := models.ModelInventoryUnitsCount{
					UnitId:  unit.ID,
					UnitAbv: unit.Abbreviation,
					Count:   0,
					Factor:  unit.Factor,
				}

				unitCounts = append(unitCounts, metadataUnit)
			}

			icMetadata := models.ModelInventoryCountMetadata{
				ProductID:  product.ID,
				UnitsCount: unitCounts,
			}

			metadata = append(metadata, icMetadata)
		}
	}

	createModel := &models.ModelInventoryCount{
		StoreID:     recipeNew.StoreID,
		CompanyID:   companyId,
		WarehouseID: recipeNew.WarehouseID,
		CreatedBy:   userId,
		AssignedTo:  recipeNew.AssignedTo,
		Status:      status,
		ScheduledAt: *recipeNew.DueDate,
		CountItems:  products,
		Metadata:    metadata,
	}

	inventoryCountCreated, err := f.box.InventoryCountService.Create(ctx, createModel)
	if err != nil {
		return nil, err
	}

	return f.toDTO(inventoryCountCreated), nil
}

func (f *InventoryCountFacade) Update(ctx context.Context, id string, recipeEdit recipe.RecipeCreateInventoryCount) (*dto.DtoInventoryCountDetail, error) {
	//TODO: validar permisos

	inventoryCount, err := f.box.InventoryCountService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if inventoryCount.Status != "un_assigned" && inventoryCount.Status != "pending" {
		return nil, types.ThrowMsg("no puedes editar productos en una lista de conteo en progeso o completada")
	}

	status := "pending"
	if recipeEdit.AssignedTo == nil {
		status = "un_assigned"
	}

	products := make([]models.ModelInventoryCountItem, 0)
	metadata := make([]models.ModelInventoryCountMetadata, 0)
	for _, ID := range recipeEdit.ProductsID {
		if product, err := f.box.ProductPerStoreService.GetProductPerStoreByID(ctx, ID); err != nil {
			return nil, types.ThrowMsg("Producto con ID: '" + ID + "' no fue encontrado")
		} else {
			products = append(products, models.ModelInventoryCountItem{
				ProductID:   product.ID,
				ScheduledAt: *recipeEdit.DueDate,
			})

			//se inicializa la lista de unidades con la unidad base
			unitCounts := []models.ModelInventoryUnitsCount{
				{
					UnitId:  product.UnitInventory.ID,
					UnitAbv: product.UnitInventory.Abbreviation,
					Count:   0,
				},
			}

			//se cargan las demas unidades
			for _, unit := range product.UnitMatrix {
				metadataUnit := models.ModelInventoryUnitsCount{
					UnitId:  unit.ID,
					UnitAbv: unit.Abbreviation,
					Count:   0,
				}

				unitCounts = append(unitCounts, metadataUnit)
			}

			icMetadata := models.ModelInventoryCountMetadata{
				ProductID:  product.ID,
				UnitsCount: unitCounts,
			}

			metadata = append(metadata, icMetadata)
		}
	}

	truncatedDate := recipeEdit.DueDate.UTC()
	recipeEdit.DueDate = &truncatedDate

	inventoryCount.ScheduledAt = *recipeEdit.DueDate
	inventoryCount.AssignedTo = recipeEdit.AssignedTo
	inventoryCount.CountItems = products
	inventoryCount.Metadata = metadata
	inventoryCount.Status = status

	inventoryCount, err = f.box.InventoryCountService.Update(ctx, inventoryCount)
	if err != nil {
		return nil, err
	}

	return f.toDTO(inventoryCount), nil

}

func (f *InventoryCountFacade) NewAssigned(ctx context.Context, id string, recipeAssigned recipe.RecipeChangeInventoryCountAssigned) (*dto.DtoInventoryCountDetail, error) {
	//TODO: validar permisos

	var inventoryCount *models.ModelInventoryCount

	inventoryCount, err := f.box.InventoryCountService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	currentStatus := inventoryCount.Status

	if currentStatus == "completed" || currentStatus == "cancelled" {
		return nil, types.ThrowMsg("No se puede cambiar el asignado para conteos de inventario completados o cancelados.")
	}

	//______________________________________//

	if inventoryCount.AssignedTo == nil && recipeAssigned.NewId == nil {
		return nil, types.ThrowMsg("Se debe proporcionar un nuevo asignado.")
	}

	if recipeAssigned.NewId != nil && inventoryCount.AssignedTo != nil && *recipeAssigned.NewId == *inventoryCount.AssignedTo {
		return nil, types.ThrowMsg("Se debe proporcionar un asignado diferente.")
	}

	if recipeAssigned.NewId != nil {
		_, err := f.box.UserService.GetUserByID(ctx, *recipeAssigned.NewId)
		if err != nil {
			return nil, err
		}

		currentStatus = "pending"
	} else {
		currentStatus = "un_assigned"
	}

	inventoryCount, err = f.box.InventoryCountService.SetNewAssigned(id, recipeAssigned.NewId)
	if err != nil {
		return nil, err
	}

	if currentStatus != inventoryCount.Status {
		inventoryCount, err = f.box.InventoryCountService.ChangeState(id, currentStatus)
		if err != nil {
			return nil, err
		}
	}

	return f.toDTO(inventoryCount), nil

}

func (f *InventoryCountFacade) NewDate(ctx context.Context, id string, recipeDate recipe.RecipeChangeInventoryCountDate) (*models.ModelInventoryCount, error) {
	//TODO: validar permisos

	inventoryCount, err := f.box.InventoryCountService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if inventoryCount.Status == "completed" || inventoryCount.Status == "cancelled" {
		return nil, types.ThrowMsg("No se puede cambiar la fecha para conteos de inventario completados o cancelados.")
	}

	//______________________________________//

	now := time.Now().UTC().Truncate(time.Minute)
	recipeDate.NewDate = recipeDate.NewDate.UTC().Truncate(time.Minute)

	if recipeDate.NewDate.Before(now) {
		return nil, types.ThrowMsg("La nueva fecha debe ser en el futuro o el día de hoy.")
	}

	if inventoryCount.ScheduledAt.Equal(recipeDate.NewDate) {
		return nil, types.ThrowMsg("Se debe proporcionar una fecha diferente.")
	}

	return f.box.InventoryCountService.SetNewDate(id, recipeDate.NewDate)

}

func (f *InventoryCountFacade) StartInventoryCount(ctx *gin.Context, id string) error {
	//TODO: validar permisos
	inventoryCount, err := f.box.InventoryCountService.GetByID(ctx, id)
	if err != nil {
		return err
	}

	userId := ctx.Keys[shared.UserIdKey()].(string)

	if inventoryCount.AssignedTo == nil || *inventoryCount.AssignedTo != userId {
		return types.ThrowMsg("Usuario no asignado a esta lista")
	}

	if inventoryCount.Status != "pending" && inventoryCount.Status != "rejected" {
		return types.ThrowMsg("No se puede iniciar el conteo de inventario")
	}

	now := time.Now().UTC()
	scheduledTruncated := inventoryCount.ScheduledAt.UTC().Truncate(time.Hour * 24)

	if scheduledTruncated.After(now) {
		return types.ThrowMsg("No se puede iniciar el conteo de inventario antes de la fecha programada.")
	}

	inventoryCount, err = f.box.InventoryCountService.ChangeState(id, "in_progress")
	if err != nil {
		return err
	}

	return nil
}

func (f *InventoryCountFacade) Commit(ctx *gin.Context, id string, recipeCommit recipe.RecipeInventoryCount) error {
	//TODO: validar permisos

	var inventoryCount *models.ModelInventoryCount

	inventoryCount, err := f.box.InventoryCountService.GetDraft(id)
	if err != nil || inventoryCount == nil {

		inventoryCount, err = f.box.InventoryCountService.GetByID(ctx.Request.Context(), id)
		if err != nil {
			return err
		}
	}

	userId := ctx.Keys[shared.UserIdKey()].(string)

	if inventoryCount.AssignedTo == nil || *inventoryCount.AssignedTo != userId {
		return types.ThrowMsg("Usuario no asignado a este conteo de inventario.")
	}

	if inventoryCount.Status != "in_progress" {
		return types.ThrowMsg("No se puede confirmar el conteo de inventario.")
	}

	for _, recipeItem := range recipeCommit.Items {

		if !recipeItem.Completed {
			return types.ThrowMsg("Todos los artículos deben estar completos para confirmar el conteo de inventario.")
		}

		f.calculateTotalAndUpdateModel(inventoryCount, recipeItem)

	}

	inventoryCount.Status = "completed"

	inventoryCount, err = f.box.InventoryCountService.Commit(inventoryCount)
	if err != nil {
		return err
	}

	err = f.box.InventoryCountService.DeleteDraft(id)
	if err != nil {
		key := fmt.Sprintf("DRAFT:INVENTORY_COUNT:%s", id)
		fmt.Println("Hubo un problema tratando de borrar el borrador con llave: ", key)
	}

	return nil
}

func (f *InventoryCountFacade) Cancel(ctx context.Context, id string) error {
	inventoryCount, err := f.box.InventoryCountService.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if inventoryCount.Status != "un_assigned" && inventoryCount.Status != "pending" {
		return types.ThrowMsg("ocurrió un error al cancelar el conteo de inventario")
	}

	err = f.box.InventoryCountService.Delete(ctx, inventoryCount.ID)
	if err != nil {
		return err
	}

	/* inventoryCount, err = f.box.InventoryCountService.ChangeState(id, "rejected")

	if err != nil {
		return err
	} */

	return nil
}

func (f *InventoryCountFacade) SaveDraft(ctx context.Context, id string, recipeDraft recipe.RecipeInventoryCount) error {
	inventoryCount, err := f.box.InventoryCountService.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if inventoryCount == nil {
		return types.ThrowMsg("Conteo de inventario no encontrado.")
	}

	if inventoryCount.Status != "in_progress" {
		return types.ThrowMsg("No se puede guardar el borrador del conteo de inventario.")
	}

	for _, recipeItem := range recipeDraft.Items {

		f.calculateTotalAndUpdateModel(inventoryCount, recipeItem)

	}

	return f.box.InventoryCountService.SaveDraft(id, inventoryCount)
}

func (f *InventoryCountFacade) RejectInventoryCount(ctx context.Context, id string) (*dto.DtoInventoryCountDetail, error) {
	inventoryCount, err := f.box.InventoryCountService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if inventoryCount.Status != "completed" {
		return nil, types.ThrowMsg("No se puede rechazar el conteo de inventario.")
	}

	/* completedItems, err := f.getCompleatedItems(inventoryCount.Metadata)
	if err != nil {
		return nil, err
	} */

	for i, item := range inventoryCount.Metadata {
		item.Completed = false
		for i := range item.UnitsCount {
			item.UnitsCount[i].Count = 0
		}

		inventoryCount.Metadata[i] = item
	}

	// inventoryCount.CountItems = completedItems

	inventoryCount, err = f.box.InventoryCountService.Reject(inventoryCount)
	if err != nil {
		return nil, err
	}

	return f.toDTO(inventoryCount), nil
}

/* func (f *InventoryCountFacade) CreateProductMovement(id string, pcID string) error {
	inventoryCount, err := f.box.InventoryCountService.GetByID(id)
	if err != nil {
		return err
	}

	pc, err := f.box.ProductCompanyService.GetProductCompanyByID(pcID)
	if err != nil {
		return err
	}

	measurementUnits := f.box.ProductCompanyService.GetBaseUnitOfMeasurement()
	if len(measurementUnits) == 0 {
		return shared.DomainError{
			Message: "Error handling measurement units",
		}
	}

	var pcFind *models.ModelInventoryCountMetadata
	for _, item := range inventoryCount.Metadata {
		if item.ProductID == pc.ID {
			pcFind = &item
			break
		}
	}

	if pcFind == nil || len(pcFind.UnitsCount) == 0 {
		return shared.DomainError{
			Message: "No inventory count metadata found for the specified product",
		}
	}

	for _, uc := range pcFind.UnitsCount {
		for _, mu := range measurementUnits {
			if uc.Unit != mu.Abbreviation {
				continue
			}

		}
	}

	return nil

} */

//______________________________________//

// func (f *InventoryCountFacade) getCompleatedItems(metadata []models.ModelInventoryCountMetadata) ([]models.ModelInventoryCountItem, error) {
// 	return f.box.InventoryCountService.GetCompletedByID(metadata)
// }

func (f *InventoryCountFacade) toDTO(model *models.ModelInventoryCount) *dto.DtoInventoryCountDetail {
	return &dto.DtoInventoryCountDetail{
		ID:              model.ID,
		DisplayID:       model.DisplayID,
		StoreID:         model.StoreID,
		StoreName:       model.StoreName,
		CompanyID:       model.CompanyID,
		WarehouseID:     model.WarehouseID,
		WarehouseName:   model.WarehouseName,
		CreatedBy:       model.CreatedBy,
		CreatedByName:   model.CreatedByName,
		AssignedTo:      model.AssignedTo,
		AssignedToName:  model.AssignedToName,
		Status:          model.Status,
		ScheduledAt:     model.ScheduledAt,
		CompletedAt:     model.CompletedAt,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		MovementTrackId: model.MovementTrackId,
		CountItems:      f.toMetadataDTO(model.Metadata, model.CountItems),
	}
}

func (f *InventoryCountFacade) toMetadataDTO(metadata []models.ModelInventoryCountMetadata, countItem []models.ModelInventoryCountItem) []dto.DtoInventoryCountItem {
	dtos := make([]dto.DtoInventoryCountItem, 0)

	for _, md := range metadata {
		dtoItem := dto.DtoInventoryCountItem{
			ProductID:    md.ProductID,
			UnitsCount:   make([]dto.DtoUnitsCount, 0),
			Completed:    md.Completed,
			ProductName:  "",
			ProductSKU:   "",
			ProductImage: nil,
			Total:        md.Total,
		}

		for _, ci := range countItem {
			if ci.ProductID == md.ProductID {
				dtoItem.ProductName = ci.ProductName
				dtoItem.ProductSKU = ci.ProductSKU
				dtoItem.ProductImage = ci.ProductImage
				dtoItem.IncidenceImageURL = ci.IncidenceImageURL
				dtoItem.IncidenceObservation = ci.IncidenceObservation
				break
			}
		}

		for _, uc := range md.UnitsCount {
			dtoItem.UnitsCount = append(dtoItem.UnitsCount, dto.DtoUnitsCount{
				UnitId:  uc.UnitId,
				UnitAbv: uc.UnitAbv,
				Count:   uc.Count,
				Factor:  uc.Factor,
			})
		}

		dtos = append(dtos, dtoItem)
	}
	return dtos
}

func (f *InventoryCountFacade) toGeneralDTO(model *models.ModelInventoryCount) *dto.DtoInventoryCountGeneral {
	completedItems := 0
	for _, item := range model.Metadata {
		if item.Completed {
			completedItems++
		}
	}

	return &dto.DtoInventoryCountGeneral{
		ID:              model.ID,
		DisplayID:       model.DisplayID,
		StoreID:         model.StoreID,
		StoreName:       model.StoreName,
		WarehouseID:     model.WarehouseID,
		WarehouseName:   model.WarehouseName,
		CompanyID:       model.CompanyID,
		CreatedBy:       model.CreatedBy,
		CreatedByName:   model.CreatedByName,
		AssignedTo:      model.AssignedTo,
		AssignedToName:  model.AssignedToName,
		Status:          model.Status,
		ScheduledAt:     model.ScheduledAt,
		CompletedAt:     model.CompletedAt,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		MovementTrackId: model.MovementTrackId,
		TotalItems:      len(model.Metadata),
		CompletedItems:  completedItems,
	}
}

func (f *InventoryCountFacade) toGeneralDTOList(models []models.ModelInventoryCount) []dto.DtoInventoryCountGeneral {
	dtos := make([]dto.DtoInventoryCountGeneral, 0, len(models))
	for _, model := range models {
		dtos = append(dtos, *f.toGeneralDTO(&model))
	}
	return dtos
}

func (*InventoryCountFacade) calculateTotalAndUpdateModel(inventoryCount *models.ModelInventoryCount, recipeItem recipe.RecipeCountItem) error {
	for i, md := range inventoryCount.Metadata {
		if md.ProductID == recipeItem.ProductID {
			total := float32(0)

			unitsMap := md.GetUnitsMap()

			if len(unitsMap) == 0 {
				return types.ThrowMsg("Error al manejar las unidades en los metadatos del conteo de inventario.")
			}

			unitCounts := make([]models.ModelInventoryUnitsCount, 0)
			for _, ruc := range recipeItem.UnitsCount {

				total += ruc.Count * unitsMap[ruc.UnitId].Factor

				unitCounts = append(unitCounts, models.ModelInventoryUnitsCount{
					UnitId:  ruc.UnitId,
					UnitAbv: ruc.UnitAbv,
					Count:   ruc.Count,
					Factor:  unitsMap[ruc.UnitId].Factor,
				})
			}

			inventoryCount.Metadata[i].Total = total
			inventoryCount.Metadata[i].UnitsCount = unitCounts
			inventoryCount.Metadata[i].Completed = recipeItem.Completed
		}
	}
	return nil
}

func (f *InventoryCountFacade) SaveIncidence(ctx *gin.Context, countId string, items recipe.RecipeInventoryCountIncidence) error {
	// Validar que el array no esté vacío
	if len(items.Items) == 0 {
		return types.ThrowMsg("El array de incidencias no puede estar vacío")
	}

	// Validar que el conteo existe (sin validar permisos de propiedad)
	inventoryCount, err := f.box.InventoryCountService.GetByIDWithoutPermissionCheck(countId)
	if err != nil {
		return err
	}
	if inventoryCount == nil {
		return types.ThrowMsg("Conteo de inventario no encontrado")
	}

	// Validar que el usuario está asignado al conteo
	userId := ctx.Keys[shared.UserIdKey()].(string)
	if inventoryCount.AssignedTo == nil || *inventoryCount.AssignedTo != userId {
		return types.ThrowMsg("Usuario no asignado a este conteo de inventario")
	}

	// Validar que el conteo está en progreso
	if inventoryCount.Status != "in_progress" {
		return types.ThrowMsg("Solo se pueden guardar incidencias en conteos en progreso")
	}

	// Procesar cada item
	for _, item := range items.Items {
		// Validar que el producto existe en el conteo
		productFound := false
		for _, countItem := range inventoryCount.CountItems {
			if countItem.ProductID == item.IDProducto {
				productFound = true
				break
			}
		}
		if !productFound {
			return types.ThrowMsg("El producto con ID '" + item.IDProducto + "' no está en este conteo de inventario")
		}

		// Verificar si ya existe una incidencia previa
		existingIncidence, err := f.box.InventoryCountService.GetIncidenceByProduct(countId, item.IDProducto)
		if err != nil {
			return err
		}

		hasExistingIncidence := existingIncidence != nil && existingIncidence.IncidenceImageURL != nil && *existingIncidence.IncidenceImageURL != ""

		// Observación (opcional, puede ser vacía o nil)
		// Si el campo está presente en el JSON (no nil), actualizarlo (incluso si está vacío)
		// Si es nil, no actualizarlo (mantener el existente)
		var observation *string
		if item.Incidencias.Observaciones != nil {
			observation = item.Incidencias.Observaciones
		} else {
			observation = nil
		}

		// Caso 1: imagen es null explícito → eliminar imagen
		if item.Incidencias.Imagen.IsExplicitNull {
			// Validar que existe incidencia previa con imagen
			if !hasExistingIncidence {
				return types.ThrowMsg("No hay imagen para eliminar")
			}

			// Eliminar imagen
			err = f.box.InventoryCountService.DeleteIncidenceImage(
				ctx.Request.Context(),
				countId,
				item.IDProducto,
				observation,
			)
			if err != nil {
				return err
			}
			continue // Continuar con el siguiente item
		}

		// Caso 2: imagen tiene valor (string base64) → subir nueva imagen
		if item.Incidencias.Imagen.Value != nil && *item.Incidencias.Imagen.Value != "" {
			// Validar mimeType cuando se envía imagen
			if item.Incidencias.MimeType == nil || *item.Incidencias.MimeType == "" {
				return types.ThrowMsg("mimeType es requerido cuando se envía imagen")
			}

			if !shared.ValidateMimeType(*item.Incidencias.MimeType) {
				return types.ThrowMsg("Tipo de imagen no soportado. Solo se permiten: image/jpeg, image/png, image/webp")
			}

			cleanedBase64 := shared.CleanBase64String(*item.Incidencias.Imagen.Value)
			base64Image := &cleanedBase64
			mimeType := item.Incidencias.MimeType

			// Guardar incidencia con nueva imagen
			err = f.box.InventoryCountService.SaveIncidence(
				ctx.Request.Context(),
				countId,
				item.IDProducto,
				base64Image,
				mimeType,
				observation,
			)
			if err != nil {
				return err
			}
			continue // Continuar con el siguiente item
		}

		// Caso 3: imagen no está presente (campo omitido) → mantener imagen existente
		if !hasExistingIncidence {
			return types.ThrowMsg("Debe incluir una imagen al crear una nueva incidencia")
		}

		// Si solo se actualiza la observación (imagen no presente), usar SaveIncidence sin imagen
		if observation != nil {
			err = f.box.InventoryCountService.SaveIncidence(
				ctx.Request.Context(),
				countId,
				item.IDProducto,
				nil, // No actualizar imagen
				nil, // No mimeType
				observation,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
