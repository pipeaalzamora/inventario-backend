package facades

import (
	"sofia-backend/config"
	"sofia-backend/domain/external"
	"sofia-backend/domain/services"
)

type DeliveryPurchaseNoteFacade struct {
	appServices      *services.ServiceContainer
	externalservices *external.ServiceContainer
	isDebug          bool
}

func NewDeliveryPurchaseNoteFacade(appServices *services.ServiceContainer, externalServices *external.ServiceContainer, config *config.Config) *DeliveryPurchaseNoteFacade {
	return &DeliveryPurchaseNoteFacade{
		appServices:      appServices,
		externalservices: externalServices,
		isDebug:          config.Debug,
	}
}

/*
func (f *DeliveryPurchaseNoteFacade) CreateDeliveryPurchaseNote(ctx context.Context, note *recipe.RecipeDeliveryPurchaseNote) (*dto.DTODeliveryPurchaseNote, error) {
	purchase, err := f.validateAndProcessDeliveryNote(ctx, note)
	if err != nil {
		return nil, err
	}

	err = f.processDeliveryNoteItems(note, purchase, false)
	if err != nil {
		return nil, err
	}

	model, err := f.appServices.DeliveryPurchaseNoteService.CreateDeliveryPurchaseNote(ctx, note)
	if err != nil {
		return nil, err
	}

	err = f.appServices.PurchaseService.AddDeliveryNoteIdAndSetArrivedStatus(model.PurchaseID, model.ID)
	if err != nil {
		return nil, err
	}

	dto, err := f.GetDeliveryPurchaseNoteByID(ctx, model.ID)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (f *DeliveryPurchaseNoteFacade) UpdateDeliveryPurchaseNote(ctx context.Context, id string, note *recipe.RecipeDeliveryPurchaseNote) (*dto.DTODeliveryPurchaseNote, error) {
	purchase, err := f.validateAndProcessDeliveryNote(ctx, note)
	if err != nil {
		return nil, err
	}

	err = f.processDeliveryNoteItems(note, purchase, true)
	if err != nil {
		return nil, err
	}

	model, err := f.appServices.DeliveryPurchaseNoteService.UpdateDeliveryPurchaseNote(ctx, id, note)
	if err != nil {
		return nil, err
	}

	dto, err := f.GetDeliveryPurchaseNoteByID(ctx, model.ID)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (f *DeliveryPurchaseNoteFacade) ConfirmDeliveryPurchaseNote(ctx context.Context, deliveryPurchaseNoteID string, invoiceFolio string, invoiceGuide string) (*dto.DTODeliveryPurchaseNote, error) {
	model, err := f.appServices.DeliveryPurchaseNoteService.GetDeliveryPurchaseNoteByID(ctx, deliveryPurchaseNoteID)
	if err != nil {
		return nil, err
	}
	if model == nil {
		return nil, types.ThrowMsg("Recepción no encontrada.")
	}

	if model.Status != entities.DeliveryPurchaseNoteStatusPending {
		return nil, types.ThrowMsg("Solo las recepciones en estado pendiente pueden ser confirmadas.")
	}

	if len(model.Files) == 0 {
		return nil, types.ThrowMsg("La recepción no tiene archivos adjuntos.")
	}

	// update delivery note status to completed
	err = f.appServices.DeliveryPurchaseNoteService.CompleteDeliveryPurchaseNote(
		ctx,
		deliveryPurchaseNoteID,
		entities.DeliveryPurchaseNoteStatusCompleted,
		invoiceFolio,
		invoiceGuide,
	)
	if err != nil {
		return nil, err
	}

	// Crear movimientos TRANSFER desde bodega transición a bodega destino
	err = f.transferProductsFromTransitionWarehouse(ctx, model)
	if err != nil {
		return nil, types.ThrowMsg("error al transferir productos desde bodega de transición: " + err.Error())
	}

	err = f.appServices.PurchaseService.UpdatePurchaseState(model.PurchaseID, entities.PurchaseStatusCompleted, "Purchase completed by delivery note "+model.DisplayID)
	if err != nil {
		return nil, err
	}

	modelPurchase, err := f.appServices.PurchaseService.GetPurchaseByID(ctx, model.PurchaseID, false)
	if err != nil {
		return nil, err
	}

	// Check if inventory request needs to be completed
	modelPurchases, err := f.appServices.PurchaseService.GetPurchasesByInventoryRequestID(modelPurchase.InventoryRequestID)
	if err != nil {
		return nil, err
	}

	// check if purchase has children
	var purchaseParents map[string]bool = make(map[string]bool)
	for _, p := range modelPurchases {
		if p.ParentPurchaseID != nil {
			purchaseParents[*(p.ParentPurchaseID)] = true
		}
	}

	// check if all purchases are completed
	mustComplete := true

purchaseLoop:
	for _, p := range modelPurchases {
		if purchaseParents[p.ID] {
			// check if exists products rejected
			for _, item := range p.Items {
				if item.Status == entities.ItemPurchaseStatusRejected {
					fmt.Println("Found rejected item in purchase with children, cannot complete inventory request")
					// if exists, inventory request cannot be completed
					mustComplete = false
					break purchaseLoop // This will exit the outer loop immediately
				}
			}

			// if purchase has children
			continue
		}

		// if one purchase is not completed or cancelled, inventory request cannot be completed
		if p.Status == entities.PurchaseStatusCompleted || p.Status == entities.PurchaseStatusCancelled {
			continue
		} else {
			fmt.Println("Found non-completed purchase with children, cannot complete inventory request")
			mustComplete = false
			break purchaseLoop
		}
	}

	// if all purchases are completed or cancell, complete inventory request
	if mustComplete {
		_, err = f.appServices.InventoryRequestService.ChangeStatus(ctx, modelPurchase.InventoryRequestID, entities.RequestStatusCompleted, nil)
		if err != nil {
			return nil, err
		}
	}

	return f.GetDeliveryPurchaseNoteByID(ctx, deliveryPurchaseNoteID)
}

func (f *DeliveryPurchaseNoteFacade) GeneratePurchaseCorrectionNote(ctx context.Context, deliveryNoteId string) (*dto.DTODeliveryPurchaseNote, error) {
	modelDeliveryNote, err := f.appServices.DeliveryPurchaseNoteService.GetDeliveryPurchaseNoteByID(ctx, deliveryNoteId)
	if err != nil {
		return nil, err
	}

	if modelDeliveryNote == nil {
		return nil, types.ThrowMsg("Recepción no encontrada.")
	}

	if modelDeliveryNote.Status != entities.DeliveryPurchaseNoteStatusDisputed {
		return nil, types.ThrowMsg("Solo las recepciones en estado en disputa pueden generar notas de corrección.")
	}

	purchase, err := f.appServices.PurchaseService.GetPurchaseByID(ctx, modelDeliveryNote.PurchaseID, false)
	if err != nil {
		return nil, err
	}

	itemsCorrected := make(map[string]float32)
	for _, item := range modelDeliveryNote.Items {
		switch item.Status {
		case entities.DeliveryPurchaseNoteItemStatusAccepted:
			itemsCorrected[item.StoreProductID] = item.Quantity
		case entities.DeliveryPurchaseNoteItemStatusSubstock:
			itemsCorrected[item.StoreProductID] = item.Quantity
		case entities.DeliveryPurchaseNoteItemStatusRejected:
			continue
		case entities.DeliveryPurchaseNoteItemStatusSuprastock:
			itemsCorrected[item.StoreProductID] = item.Quantity - item.Difference
		}
	}

	// Create Purchase Correction with the items that were rejected or suprastock
	purchaseItems := make([]recipe.RecipePurchaseItem, len(modelDeliveryNote.Items))
	for i, item := range modelDeliveryNote.Items {
		purchaseItems[i] = recipe.RecipePurchaseItem{
			StoreProductID: item.StoreProductID,
			Quantity:       itemsCorrected[item.StoreProductID],
			UnitPrice:      item.UnitPrice,
		}
	}

	created, err := f.appServices.PurchaseService.CreatePurchaseOrder(ctx, &recipe.RecipePurchase{
		Description:        fmt.Sprintf("Purchase order for inventory request %s", purchase.ID),
		SupplierID:         purchase.SupplierID,
		StoreID:            purchase.StoreID,
		InventoryRequestID: purchase.InventoryRequestID,
		Items:              purchaseItems,
	})
	if err != nil {
		return nil, err
	}

	// aqui debe ir la creacion del token para las OC
	var token string
	if f.isDebug {
		token = fmt.Sprintf("token-%s", created.DisplayID)
	} else {
		token, err = shared.CreateRandomURLString(64)
		if err != nil {
			fmt.Println("Error creating token:", err)
		}
	}

	// expira en 7 dias
	exp := time.Now().Add(time.Hour * 24 * 7).UTC() // una semana en UTC

	// crear el supplier token
	err = f.appServices.SupplierOCService.CreateSupplierOC(created.ID, token, &exp)
	if err != nil {
		// si falla la creacion de el token para el supplier se retorna el error?
		fmt.Println("Error creating supplier token:", err)
	}

	// traer email de suplier
	supplierEmail, err := f.appServices.SupplierService.GetSupplierContactsEmailByID(purchase.SupplierID)
	if err != nil {
		fmt.Println("Approve Request Error getting supplier email:", err)
	}

	// enviar el correo con la url junto al token
	err = f.externalservices.EmailService.SendSupplierViewEmail(supplierEmail, token, exp)
	if err != nil {
		fmt.Println("Error sending supplier email:", err)
	}

	err = f.appServices.PurchaseService.AddSonOCToPurchase(purchase.ID, created.ID)
	if err != nil {
		return nil, err
	}

	// update delivery note status to corrected
	err = f.appServices.DeliveryPurchaseNoteService.UpdateDeliveryPurchaseNoteStatus(
		ctx,
		deliveryNoteId,
		entities.DeliveryPurchaseNoteStatusCancelled,
	)

	if err != nil {
		return nil, err
	}

	// update status of original purchase to corrected
	err = f.appServices.PurchaseService.UpdatePurchaseState(
		purchase.ID,
		entities.PurchaseStatusEdited,
		"Purchase corrected by delivery note "+modelDeliveryNote.DisplayID,
	)
	if err != nil {
		return nil, err
	}

	return f.GetDeliveryPurchaseNoteByID(ctx, deliveryNoteId)
}

func (f *DeliveryPurchaseNoteFacade) GetAllDeliveryPurchaseNotes(ctx context.Context, storeID string, page int, size int, filter *map[string]interface{}) (shared.PaginationResponse[models.ModelDeliveryPurchaseNote], error) {
	data, total, err := f.appServices.DeliveryPurchaseNoteService.GetAllDeliveryPurchaseNotes(ctx, storeID, page, size, filter)
	if err != nil {
		return shared.PaginationResponse[models.ModelDeliveryPurchaseNote]{}, err
	}

	return shared.NewPagination(data, total, page, size), nil
}

func (f *DeliveryPurchaseNoteFacade) GetDeliveryPurchaseNoteByID(ctx context.Context, id string) (*dto.DTODeliveryPurchaseNote, error) {
	model, err := f.appServices.DeliveryPurchaseNoteService.GetDeliveryPurchaseNoteByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return f.toDTO(model), nil
}

func (f *DeliveryPurchaseNoteFacade) UploadFileToDeliveryPurchaseNote(ctx context.Context, deliveryPurchaseNoteID string, form *recipe.UploadForm) (*dto.DTODeliveryPurchaseNote, error) {
	model, err := f.appServices.DeliveryPurchaseNoteService.GetDeliveryPurchaseNoteByID(ctx, deliveryPurchaseNoteID)
	if err != nil {
		return nil, err
	}

	// check if delivery note exists
	if model == nil {
		return nil, types.ThrowMsg("Recepción no encontrada.")
	}

	err = f.appServices.DeliveryPurchaseNoteService.AddFileToDeliveryPurchaseNote(ctx, deliveryPurchaseNoteID, form)
	if err != nil {
		return nil, err
	}

	return f.GetDeliveryPurchaseNoteByID(ctx, deliveryPurchaseNoteID)
}

func (f *DeliveryPurchaseNoteFacade) RemoveFileFromDeliveryPurchaseNote(ctx context.Context, deliveryPurchaseNoteID string, fileID string) (*dto.DTODeliveryPurchaseNote, error) {
	modelNote, err := f.appServices.DeliveryPurchaseNoteService.GetDeliveryPurchaseNoteByID(ctx, deliveryPurchaseNoteID)
	if err != nil {
		return nil, err
	}

	// check if file belongs to the delivery note
	var file *models.ModelFile
	for _, f := range modelNote.Files {
		if f.ID == fileID {
			file = &f
			break
		}
	}

	if file == nil {
		return nil, types.ThrowMsg("El archivo no pertenece a la recepción especificada.")
	}

	err = f.appServices.DeliveryPurchaseNoteService.RemoveFileFromDeliveryPurchaseNote(ctx, deliveryPurchaseNoteID, fileID)
	if err != nil {
		return nil, err
	}

	return f.GetDeliveryPurchaseNoteByID(ctx, deliveryPurchaseNoteID)
}

func (f *DeliveryPurchaseNoteFacade) toDTO(model *models.ModelDeliveryPurchaseNote) *dto.DTODeliveryPurchaseNote {
	items := make([]dto.DTODeliveryPurchaseNoteItem, len(model.Items))
	for i, item := range model.Items {
		items[i] = dto.DTODeliveryPurchaseNoteItem{
			ID:                     item.ID,
			DeliveryPurchaseNoteID: item.DeliveryPurchaseNoteID,
			ProductName:            item.ProductName,
			StoreProductID:         item.StoreProductID,
			Quantity:               item.Quantity,
			PurchaseUnit:           item.PurchaseUnit,
			Status:                 item.Status.String(),
			UnitPrice:              item.UnitPrice,
			Subtotal:               item.Subtotal,
			TaxTotal:               item.TaxTotal,
			Difference:             item.Difference,
		}
	}

	file := make([]dto.DTOFile, len(model.Files))
	for i, f := range model.Files {
		file[i] = dto.DTOFile{
			ID:       f.ID,
			FileType: f.FileType,
			FileURL:  f.FileURL,
		}
	}

	return &dto.DTODeliveryPurchaseNote{
		ID:                model.ID,
		DisplayID:         model.DisplayID,
		SupplierID:        model.SupplierID,
		SupplierName:      model.SupplierName,
		PurchaseID:        model.PurchaseID,
		FolioInvoice:      model.FolioInvoice,
		FolioGuide:        model.FolioGuide,
		PurchaseDisplayID: model.PurchaseDisplayID,
		StoreID:           model.StoreID,
		StoreName:         model.StoreName,
		WarehouseID:       model.WarehouseID,
		WarehouseName:     model.WarehouseName,
		Status:            model.Status.String(),
		Comment:           model.Comment,
		Total:             model.Total,
		UserID:            model.UserID,
		UserName:          model.UserName,
		Items:             items,
		Files:             file,
		CreatedAt:         model.CreatedAt,
		UpdatedAt:         model.UpdatedAt,
	}
}

// Extract common validation and processing logic
func (f *DeliveryPurchaseNoteFacade) validateAndProcessDeliveryNote(ctx context.Context, note *recipe.RecipeDeliveryPurchaseNote) (*models.ModelPurchase, error) {
	purchase, err := f.appServices.PurchaseService.GetPurchaseByID(ctx, note.PurchaseID, false)
	if err != nil {
		return nil, err
	}
	if purchase.StoreID != note.StoreID {
		return nil, types.ThrowMsg("La compra no pertenece a la tienda especificada.")
	}

	if purchase.SupplierID != note.SupplierID {
		return nil, types.ThrowMsg("La compra no pertenece al proveedor especificado.")
	}

	// Validar que la compra no tenga ya una nota de entrega asociada
	if purchase.DeliveryPurchaseNoteID != nil && *purchase.DeliveryPurchaseNoteID != "" {
		return nil, types.ThrowMsg("Esta orden de compra ya tiene una nota de entrega asociada.")
	}

	// Validar que el estado de la compra sea 'en camino'
	if purchase.Status != entities.PurchaseStatusOnDelivery {
		return nil, types.ThrowMsg("Solo se pueden crear recepciones para órdenes en estado 'en camino'.")
	}

	return purchase, nil
}

func (f *DeliveryPurchaseNoteFacade) processDeliveryNoteItems(note *recipe.RecipeDeliveryPurchaseNote, purchase *models.ModelPurchase, useOriginalSubtotal bool) error {
	// Check if note has the same items as the purchase
	mapPurchaseItem := make(map[string]models.ModelPurchaseItem)
	for _, item := range purchase.Items {
		mapPurchaseItem[item.StoreProductID] = item
	}

	disputed := 0
	var total float32 = 0.0

	for i, item := range note.Items {

		purchaseItem, ok := mapPurchaseItem[item.StoreProductID]
		if !ok {
			return types.ThrowMsg("El artículo no pertenece a la compra especificada.")
		}

		subtotal := purchaseItem.UnitPrice * item.Quantity
		taxTotal := float32(math.Round(float64(subtotal) * 0.19)) // Hardcoded tax 19%
		diff := item.Quantity - purchaseItem.Quantity

		var status entities.DeliveryPurchaseNoteItemStatus

		switch {
		case item.Quantity == 0:
			status = entities.DeliveryPurchaseNoteItemStatusRejected
			disputed++
		case diff < 0:
			status = entities.DeliveryPurchaseNoteItemStatusSubstock
			disputed++
		case diff > 0:
			status = entities.DeliveryPurchaseNoteItemStatusSuprastock
		case diff == 0:
			status = entities.DeliveryPurchaseNoteItemStatusAccepted
		}

		// Use original subtotal for updates, calculated subtotal for creates
		itemSubtotal := subtotal
		if useOriginalSubtotal {
			itemSubtotal = purchaseItem.Subtotal
		}

		note.Items[i] = recipe.RecipeDeliveryPurchaseNoteItem{
			StoreProductID: item.StoreProductID,
			Quantity:       item.Quantity,
			UnitPrice:      purchaseItem.UnitPrice,
			PurchaseUnit:   purchaseItem.PurchaseUnit,
			Subtotal:       itemSubtotal,
			TaxTotal:       taxTotal,
			Total:          itemSubtotal + taxTotal,
			Difference:     diff,
			Status:         status.String(),
		}

		total += subtotal
	}

	if disputed > 0 {
		note.Status = entities.DeliveryPurchaseNoteStatusDisputed.String()
	} else {
		note.Status = entities.DeliveryPurchaseNoteStatusPending.String()
	}

	note.Total = total
	return nil
}

// transferProductsFromTransitionWarehouse transfiere productos desde bodega de transición
// a la bodega destino especificada en la delivery note, creando movimientos TRANSFER
func (f *DeliveryPurchaseNoteFacade) transferProductsFromTransitionWarehouse(ctx context.Context, deliveryNote *models.ModelDeliveryPurchaseNote) error {
	// 1. Validar que existe purchase asociada con movimientos previos
	purchase, err := f.appServices.PurchaseService.GetPurchaseByID(ctx, deliveryNote.PurchaseID, true)
	if err != nil {
		return types.ThrowMsg("no se pudo obtener la orden de compra: " + err.Error())
	}

	// 2. Obtener bodega de transición
	transitionWarehouse, err := f.appServices.WarehouseService.GetTransitionWarehouseByStoreID(ctx, deliveryNote.StoreID)
	if err != nil {
		return types.ThrowMsg("no se pudo obtener la bodega de transición: " + err.Error())
	}

	// 3. Crear movimientos TRANSFER solo para items ACCEPTED
	movements := make([]models.ModelProductMovement, 0)
	for _, item := range deliveryNote.Items {
		if item.Status != entities.DeliveryPurchaseNoteItemStatusAccepted {
			continue // Skip rejected, substock, suprastock items
		}

		// Buscar el purchase item correspondiente para verificar que fue aprobado
		var purchaseItem *models.ModelPurchaseItem
		for _, pItem := range purchase.Items {
			if pItem.StoreProductID == item.StoreProductID {
				purchaseItem = &pItem
				break
			}
		}

		if purchaseItem == nil || purchaseItem.Status != entities.ItemPurchaseStatusApproved {
			continue // Skip if not found or not approved
		}

		// Obtener información del producto para la unidad de inventario
		product, err := f.appServices.ProductPerStoreService.GetProductPerStoreByID(ctx, item.StoreProductID)
		if err != nil {
			return types.ThrowMsg("no se pudo obtener el producto: " + err.Error())
		}

		var inventoryUnit *string
		if product.UnitInventory.ID > 0 {
			inventoryUnit = &product.UnitInventory.Abbreviation
		}

		// Obtener stock actual en la bodega de transición (origen, antes del movimiento)
		var stockBefore float32 = 0
		transitionStock, err := f.appServices.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
			ctx, item.StoreProductID, transitionWarehouse.ID,
		)
		if err == nil && transitionStock != nil {
			stockBefore = transitionStock.InStock
		}

		// Calcular stock después del movimiento (en bodega origen, resta)
		stockAfter := stockBefore - item.Quantity

		docType := "DN"
		movements = append(movements, models.ModelProductMovement{
			StoreProductID:    item.StoreProductID,
			Observation:       "Transferencia de " + transitionWarehouse.WarehouseName + " a bodega destino - DN: " + deliveryNote.DisplayID,
			Quantity:          item.Quantity,
			InventoryUnit:     inventoryUnit,
			UnitCost:          item.UnitPrice,
			TotalCost:         item.Subtotal,
			MovedFrom:         &transitionWarehouse.ID,
			MovedTo:           &deliveryNote.WarehouseID,
			MovedAt:           time.Now(),
			MovementType:      "TRANSFER",
			MovementDocType:   &docType,
			DocumentReference: &deliveryNote.DisplayID,
			PurchaseID:        &purchase.ID,
			StockBefore:       &stockBefore,
			StockAfter:        &stockAfter,
		})
	}

	if len(movements) == 0 {
		return nil // No hay productos para transferir
	}

	// 4. Crear movimientos batch
	createdMovements, err := f.appServices.ProductMovementService.CreateNewMovements(ctx, movements, false)
	if err != nil {
		return types.ThrowMsg("error al crear los movimientos de transferencia: " + err.Error())
	}

	// 5. Actualizar stock: restar de bodega transición y sumar a bodega destino
	directionOut := "OUT"
	directionIn := "IN"

	for _, movement := range createdMovements {
		// 5.1 Restar stock de bodega de transición
		transitionStock, err := f.appServices.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
			ctx, movement.StoreProductID, transitionWarehouse.ID,
		)
		if err != nil || transitionStock == nil {
			return types.ThrowMsg("no se encontró stock en bodega de transición para producto " + movement.StoreProductID)
		}

		if transitionStock.InStock < movement.Quantity {
			return types.ThrowMsg("stock insuficiente en bodega de transición para producto " + movement.StoreProductID)
		}

		transitionStock.InStock -= movement.Quantity
		transitionStock.Direction = &directionOut
		transitionStock.WarehouseIdReference = &deliveryNote.WarehouseID

		_, err = f.appServices.WarehousePerProductService.UpdateWPP(ctx, transitionStock)
		if err != nil {
			return types.ThrowMsg("error al actualizar stock en bodega de transición: " + err.Error())
		}

		// 5.2 Sumar stock a bodega destino
		destStock, err := f.appServices.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
			ctx, movement.StoreProductID, deliveryNote.WarehouseID,
		)

		if err != nil || destStock == nil {
			// Crear nuevo registro en bodega destino
			newStock := &models.ModelProductWarehouse{
				StoreProductId:       movement.StoreProductID,
				WarehouseId:          deliveryNote.WarehouseID,
				WarehouseIdReference: &transitionWarehouse.ID,
				Direction:            &directionIn,
				InStock:              movement.Quantity,
				CostAvg:              movement.UnitCost,
			}
			_, err := f.appServices.WarehousePerProductService.CreateNewWPP(ctx, newStock)
			if err != nil {
				return types.ThrowMsg("error al crear stock en bodega destino: " + err.Error())
			}
		} else {
			// Actualizar stock existente sumando la cantidad
			oldStock := destStock.InStock
			destStock.InStock += movement.Quantity
			destStock.Direction = &directionIn
			destStock.WarehouseIdReference = &transitionWarehouse.ID

			// Recalcular costo promedio
			totalCost := oldStock * destStock.CostAvg
			totalCost += movement.TotalCost
			destStock.CostAvg = totalCost / destStock.InStock

			_, err := f.appServices.WarehousePerProductService.UpdateWPP(ctx, destStock)
			if err != nil {
				return types.ThrowMsg("error al actualizar stock en bodega destino: " + err.Error())
			}
		}
	}

	return nil
}
*/
