package facades

import (
	"sofia-backend/domain/services"
)

type SupplierOCFacade struct {
	serviceContainer *services.ServiceContainer
}

func NewSupplierOCFacade(service *services.ServiceContainer) *SupplierOCFacade {
	return &SupplierOCFacade{
		serviceContainer: service,
	}
}

/*
func (s *SupplierOCFacade) GetSupplierOC(ctx context.Context, token string) (*dto.DTOPurchase, error) {
	supplierOC, err := s.serviceContainer.SupplierOCService.GetSupplierOC(token)
	if err != nil {
		return nil, err
	}

	// Using context.Background() as this is a public supplier endpoint (no user auth)
	purchase, err := s.serviceContainer.PurchaseService.GetPurchaseByID(ctx, supplierOC.PurchaseID, true)
	if err != nil {
		return nil, err
	}

	if purchase == nil {
		return nil, types.ThrowMsg("orden de compra no encontrada")
	}

	if purchase.Status != entities.PurchaseStatusPending {
		return nil, types.ThrowMsg("la orden de compra no está en estado pendiente")
	}

	return s.toDtoPurchase(purchase), nil
}

func (s *SupplierOCFacade) UpdateSupplierOC(ctx context.Context, token string, recipeSuppOC recipe.SupplierOCRecipe) error {
	supplierOC, err := s.serviceContainer.SupplierOCService.GetSupplierOC(token)
	if err != nil {
		return err
	}

	purchase, err := s.serviceContainer.PurchaseService.GetPurchaseByID(ctx, supplierOC.PurchaseID, true)
	if err != nil {
		return err
	}

	if purchase == nil || len(recipeSuppOC.Items) == 0 {
		return types.ThrowMsg("los ítems no pueden estar vacíos")
	}

	if purchase.ID != recipeSuppOC.PurchaseId {
		return types.ThrowMsg("el token no coincide con la orden de compra")
	}

	approveds := make(map[string]recipe.SupplierOCItem, 0)
	rejecteds := make(map[string]recipe.SupplierOCItem, 0)

	for _, item := range recipeSuppOC.Items {
		if item.ItemId == "" {
			return types.ThrowMsg("el ID del ítem no puede estar vacío")
		}

		if item.Accepted {
			approveds[item.ItemId] = item
		} else {
			rejecteds[item.ItemId] = item
		}
	}

	for index, item := range purchase.Items {
		if _, ok := approveds[item.ID]; ok {
			purchase.Items[index].Status = entities.ItemPurchaseStatusApproved
		} else if _, ok := rejecteds[item.ID]; ok {
			purchase.Items[index].Status = entities.ItemPurchaseStatusRejected
		} else {
			return types.ThrowMsg("El ID del ítem " + item.ID + " no se encontró en la solicitud")
		}
	}

	err = s.serviceContainer.PurchaseService.UpdatePurchaseItemsStatus(recipeSuppOC.PurchaseId, purchase.Items)
	if err != nil {
		return err
	}

	var state entities.PurchaseStatus = entities.PurchaseStatusPending
	if len(approveds) == len(recipeSuppOC.Items) {
		state = entities.PurchaseStatusOnDelivery
	} else if len(rejecteds) == len(recipeSuppOC.Items) {
		state = entities.PurchaseStatusCancelled
	} else if len(approveds) > 0 && len(rejecteds) > 0 {
		state = entities.PurchaseStatusSunk
	}

	err = s.serviceContainer.PurchaseService.UpdatePurchaseState(recipeSuppOC.PurchaseId, state, recipeSuppOC.Observation)
	if err != nil {
		return err
	}

	if state == entities.PurchaseStatusSunk {
		// crear nueva orden de compra con los items aprovados
		purchase, err := s.serviceContainer.PurchaseService.GetPurchaseByID(ctx, recipeSuppOC.PurchaseId, true)
		if err != nil {
			return err
		}

		if purchase == nil {
			return types.ThrowMsg("no se pudo crear la nueva orden de compra")
		}

		var newItemsRecipe []recipe.RecipePurchaseItem
		for _, item := range purchase.Items {
			if _, ok := approveds[item.ID]; ok {
				newItemsRecipe = append(newItemsRecipe, recipe.RecipePurchaseItem{
					StoreProductID: item.StoreProductID,
					Quantity:       item.Quantity,
					PurchaseUnit:   item.PurchaseUnit,
					UnitPrice:      item.UnitPrice,
				})
			}
		}

		newRecipe := &recipe.RecipePurchase{
			SupplierID:         purchase.SupplierID,
			WarehouseID:        purchase.WarehouseID,
			StoreID:            purchase.StoreID,
			InventoryRequestID: purchase.InventoryRequestID,
			Items:              newItemsRecipe,
		}

		newPurchase, err := s.serviceContainer.PurchaseService.CreatePurchaseOrderApprovedWithoutAuth(ctx, newRecipe)
		if err != nil {
			return err
		}

		err = s.serviceContainer.PurchaseService.AddSonOCToPurchase(recipeSuppOC.PurchaseId, newPurchase.ID)
		if err != nil {
			return err
		}

	}

	hashedToken := shared.CreateSaltyHash(token)
	err = s.serviceContainer.SupplierOCService.UpdateSupplierOC(hashedToken, recipeSuppOC)
	if err != nil {
		return err
	}

	// Crear movimientos de ingreso si el estado es ON_DELIVERY
	if state == entities.PurchaseStatusOnDelivery {
		purchase, err := s.serviceContainer.PurchaseService.GetPurchaseByID(ctx, recipeSuppOC.PurchaseId, true)
		if err != nil {
			return err
		}
		if err := s.createAndBookIncomingMovements(ctx, purchase); err != nil {
			// Log error but don't fail the whole operation
			// TODO: Implement retry mechanism or flag for pending movements
			return types.ThrowMsg("error al crear los movimientos de ingreso: " + err.Error())
		}
	}

	// Si creamos child purchase (sunk), también crear movimientos para ese purchase
	if state == entities.PurchaseStatusSunk {
		childPurchase, err := s.serviceContainer.PurchaseService.GetPurchaseByID(ctx, recipeSuppOC.PurchaseId, true)
		if err != nil {
			return err
		}
		if childPurchase != nil && len(childPurchase.ChildrenPurchase) > 0 {
			for _, child := range childPurchase.ChildrenPurchase {
				childPurchaseData, err := s.serviceContainer.PurchaseService.GetPurchaseByID(ctx, child.PurchaseChildID, true)
				if err != nil {
					continue
				}
				if err := s.createAndBookIncomingMovements(ctx, childPurchaseData); err != nil {
					// Log error but continue
					continue
				}
			}
		}
	}

	return nil
}

// createAndBookIncomingMovements crea movimientos INPUT automáticos al aprobar OC
// y actualiza stock en bodega de transición con dirección "IN"
func (s *SupplierOCFacade) createAndBookIncomingMovements(ctx context.Context, purchase *models.ModelPurchase) error {
	// 1. Obtener bodega de transición de la tienda
	transitionWarehouse, err := s.serviceContainer.WarehouseService.GetTransitionWarehouseByStoreID(ctx, purchase.StoreID)
	if err != nil {
		return types.ThrowMsg("no se pudo obtener la bodega de transición: " + err.Error())
	}

	// 2. Construir movimientos INPUT para cada producto aprobado
	movements := make([]models.ModelProductMovement, 0, len(purchase.Items))
	for _, item := range purchase.Items {
		if item.Status != entities.ItemPurchaseStatusApproved {
			continue
		}

		// Obtener información del producto para la unidad de inventario
		product, err := s.serviceContainer.ProductPerStoreService.GetProductPerStoreByID(ctx, item.StoreProductID)
		if err != nil {
			return types.ThrowMsg("no se pudo obtener el producto: " + err.Error())
		}

		var inventoryUnit *string
		if product.UnitInventory.ID > 0 {
			inventoryUnit = &product.UnitInventory.Abbreviation
		}

		// Obtener stock actual en la bodega de transición (antes del movimiento)
		var stockBefore float32 = 0
		warehouseStock, err := s.serviceContainer.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
			ctx, item.StoreProductID, transitionWarehouse.ID,
		)
		if err == nil && warehouseStock != nil {
			stockBefore = warehouseStock.InStock
		}

		// Calcular stock después del movimiento
		stockAfter := stockBefore + item.Quantity

		docType := "OC"
		movedBy := "SYSTEM"
		movements = append(movements, models.ModelProductMovement{
			StoreProductID:    item.StoreProductID,
			Observation:       "Ingreso OC: " + purchase.DisplayID,
			Quantity:          item.Quantity,
			InventoryUnit:     inventoryUnit,
			UnitCost:          item.UnitPrice,
			TotalCost:         item.Subtotal,
			MovedFrom:         nil, // Ingreso desde proveedor (no warehouse)
			MovedTo:           &transitionWarehouse.ID,
			MovedAt:           time.Now().UTC(),
			MovedBy:           movedBy,
			MovementDocType:   &docType,
			DocumentReference: &purchase.DisplayID,
			MovementType:      "NEWINPUT",
			PurchaseID:        &purchase.ID,
			StockBefore:       &stockBefore,
			StockAfter:        &stockAfter,
		})
	}

	if len(movements) == 0 {
		return nil // No hay productos aprobados
	}

	// 3. Crear movimientos batch
	createdMovements, err := s.serviceContainer.ProductMovementService.CreateNewMovements(ctx, movements, true)
	if err != nil {
		return types.ThrowMsg("error al crear los movimientos: " + err.Error())
	}

	// 4. Actualizar stock en warehouse_per_product para cada producto
	for _, movement := range createdMovements {
		// Verificar si ya existe registro en warehouse_per_product
		existingStock, err := s.serviceContainer.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
			ctx, movement.StoreProductID, transitionWarehouse.ID,
		)

		directionIn := "IN"
		// warehouse_id_reference apunta a la bodega destino del purchase
		warehouseRef := purchase.WarehouseID

		if err != nil || existingStock == nil {
			// Crear nuevo registro con dirección "IN"
			newStock := &models.ModelProductWarehouse{
				StoreProductId:       movement.StoreProductID,
				WarehouseId:          transitionWarehouse.ID,
				WarehouseIdReference: &warehouseRef, // Bodega destino de la tienda
				Direction:            &directionIn,
				InStock:              movement.Quantity,
				CostAvg:              movement.UnitCost,
			}
			_, err := s.serviceContainer.WarehousePerProductService.CreateNewWPP(ctx, newStock)
			if err != nil {
				return types.ThrowMsg("error al crear stock en bodega de transición: " + err.Error())
			}
		} else {
			// Actualizar stock existente sumando la cantidad
			existingStock.InStock += movement.Quantity
			existingStock.Direction = &directionIn
			existingStock.WarehouseIdReference = &warehouseRef // Bodega destino de la tienda
			// Recalcular costo promedio
			totalCost := (existingStock.InStock - movement.Quantity) * existingStock.CostAvg
			totalCost += movement.TotalCost
			existingStock.CostAvg = totalCost / existingStock.InStock

			_, err := s.serviceContainer.WarehousePerProductService.UpdateWPP(ctx, existingStock)
			if err != nil {
				return types.ThrowMsg("error al actualizar stock en bodega de transición: " + err.Error())
			}
		}
	}

	return nil
}

func (s *SupplierOCFacade) toDtoPurchase(purchase *models.ModelPurchase) *dto.DTOPurchase {
	children := make([]dto.DTOPurchaseHierarchy, len(purchase.ChildrenPurchase))
	for i, child := range purchase.ChildrenPurchase {
		children[i] = dto.DTOPurchaseHierarchy{
			PurchaseChildID:        child.PurchaseChildID,
			PurchaseChildDisplayID: child.PurchaseChildDisplayID,
		}
	}

	items := make([]dto.DTOPurchaseItem, len(purchase.Items))
	for i, item := range purchase.Items {
		items[i] = dto.DTOPurchaseItem{
			ID:             item.ID,
			PurchaseID:     item.PurchaseID,
			StoreProductID: item.StoreProductID,
			ProductName:    item.ProductName,
			Quantity:       item.Quantity,
			PurchaseUnit:   item.PurchaseUnit,
			UnitPrice:      item.UnitPrice,
			Subtotal:       item.Subtotal,
			Status:         item.Status.ToString(),
		}
	}

	history := make([]dto.DTOPurchaseHistory, len(purchase.PurchaseHistory))
	for i, h := range purchase.PurchaseHistory {
		history[i] = dto.DTOPurchaseHistory{
			ID:          h.ID,
			PurchaseID:  h.PurchaseID,
			NewStatus:   h.NewStatus,
			Observation: h.Observation,
			ChangedAt:   h.ChangedAt,
		}
	}

	purchase.Items = nil
	purchase.PurchaseHistory = nil

	dtoPurchase := &dto.DTOPurchase{
		ID:                            purchase.ID,
		DisplayID:                     purchase.DisplayID,
		SupplierID:                    purchase.SupplierID,
		SupplierName:                  purchase.SupplierName,
		SupplierPhone:                 purchase.SupplierPhone,
		WarehouseID:                   purchase.WarehouseID,
		WarehouseName:                 purchase.WarehouseName,
		WarehouseAddress:              purchase.WarehouseAddress,
		WarehousePhone:                purchase.WarehousePhone,
		StoreID:                       purchase.StoreID,
		StoreName:                     purchase.StoreName,
		InventoryRequestID:            purchase.InventoryRequestID,
		Status:                        purchase.Status.ToString(),
		CreatedAt:                     purchase.CreatedAt,
		UpdatedAt:                     purchase.UpdatedAt,
		ParentPurchaseID:              purchase.ParentPurchaseID,
		ParentPurchaseDisplayID:       purchase.ParentPurchaseDisplayID,
		DeliveryPurchaseNoteID:        purchase.DeliveryPurchaseNoteID,
		DeliveryPurchaseNoteDisplayID: purchase.DeliveryPurchaseNoteDisplayID,
		ChildrenPurchase:              children,
		Items:                         items,
		PurchaseHistory:               history,
	}

	return dtoPurchase
}
*/
