package facades

import (
	"sofia-backend/config"
	"sofia-backend/domain/external"
	"sofia-backend/domain/services"
)

type PurchaseFacade struct {
	appService       *services.ServiceContainer
	externalservices *external.ServiceContainer
	isDebug          bool
}

func NewPurchaseFacade(appService *services.ServiceContainer, externalService *external.ServiceContainer, config *config.Config) *PurchaseFacade {
	return &PurchaseFacade{appService: appService, externalservices: externalService, isDebug: config.Debug}
}

/*
func (f *PurchaseFacade) GetAllPurchase(ctx context.Context, storeID string, page int, size int, filter *map[string]interface{}) (shared.PaginationResponse[models.ModelPurchase], error) {
	purchases, total, err := f.appService.PurchaseService.GetAllPurchase(ctx, storeID, page, size, filter)
	if err != nil {
		return shared.PaginationResponse[models.ModelPurchase]{}, err
	}
	return shared.NewPagination(purchases, total, page, size), nil
}

func (f *PurchaseFacade) GetPurchasesByInventoryRequestID(ctx context.Context, inventoryRequestID string) ([]models.ModelPurchase, error) {
	return f.appService.PurchaseService.GetPurchasesByInventoryRequestID(inventoryRequestID)
}

func (f *PurchaseFacade) GetPurchaseByID(ctx context.Context, purchaseID string) (*dto.DTOPurchase, error) {
	purchase, err := f.appService.PurchaseService.GetPurchaseByID(ctx, purchaseID, false)
	if err != nil {
		return nil, err
	}
	return f.toDto(purchase), nil
}

func (f *PurchaseFacade) CreatePurchaseOrder(ctx context.Context, purchase *recipe.RecipePurchase) (*dto.DTOPurchase, error) {
	// create a inventory request approved for purchase
	userId, ok := f.appService.AuthService.GetUserIDFromContext(ctx)
	if !ok {
		return nil, types.ThrowData("Error al obtener el ID del usuario del contexto")
	}

	// Get companyId from store if not provided
	if purchase.CompanyID == "" {
		store, err := f.appService.StoreService.GetStoreByID(ctx, purchase.StoreID)
		if err != nil {
			return nil, types.ThrowData("Error al obtener la tienda para extraer la compañía")
		}
		purchase.CompanyID = store.CompanyID
	}

	// Validate Products
	storeProductIDs := make([]string, len(purchase.Items))
	for i, item := range purchase.Items {
		storeProductIDs[i] = item.StoreProductID
	}

	supplierProducts, err := f.appService.SupplierProductService.GetSupplierProductsByStoreIDAndSupplierIDWithProductCompanyIDs(
		purchase.StoreID,
		purchase.SupplierID,
		storeProductIDs,
	)
	if err != nil {
		return nil, err
	}

	for _, item := range purchase.Items {
		sp, exists := supplierProducts[item.StoreProductID]
		if !exists {
			return nil, types.ThrowMsg("El StoreProductID " + item.StoreProductID + " no existe para el SupplierID " + purchase.SupplierID)
		}
		if item.UnitPrice != sp.SupplierProductPrice {
			return nil, types.ThrowMsg("El UnitPrice para el StoreProductID " + item.StoreProductID + " no es igual al precio del producto del proveedor de " + fmt.Sprintf("%f", sp.SupplierProductPrice))
		}
	}

	created, err := f.appService.PurchaseService.CreatePurchaseOrderWithInventoryRequest(ctx, purchase, userId)
	if err != nil {
		fmt.Printf("Failed to create purchase order for supplier %s: %v\n", purchase.SupplierID, err)
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
	err = f.appService.SupplierOCService.CreateSupplierOC(created.ID, token, &exp)
	if err != nil {
		// si falla la creacion de el token para el supplier se retorna el error?
		fmt.Println("Error creating supplier token:", err)
	}

	// traer email de suplier
	supplierEmail, err := f.appService.SupplierService.GetSupplierContactsEmailByID(purchase.SupplierID)
	if err != nil {
		fmt.Println("Approve Request Error getting supplier email:", err)
	}

	// enviar el correo con la url junto al token
	err = f.externalservices.EmailService.SendSupplierViewEmail(supplierEmail, token, exp)
	if err != nil {
		fmt.Println("Error sending supplier email:", err)
	}

	// Send Notification to users with purchase power

	return f.GetPurchaseByID(ctx, created.ID)
}

func (f *PurchaseFacade) RetryWithOtherSupplier(ctx context.Context, purchaseID string) (*dto.DTOPurchase, error) {
	purchase, err := f.appService.PurchaseService.GetPurchaseByID(ctx, purchaseID, false)
	if err != nil {
		return nil, err
	}

	// Revisar qué productos están rechazados
	var rejectedItems []models.ModelPurchaseItem

	// duplicated slice to avoid modifying the original purchase items
	var itemsToUpdate []models.ModelPurchaseItem

	for _, item := range purchase.Items {
		if item.Status == entities.ItemPurchaseStatusRejected {
			rejectedItems = append(rejectedItems, item)
			itemsToUpdate = append(itemsToUpdate, item)
		}
	}

	if len(rejectedItems) == 0 {
		return nil, types.ThrowMsg("No hay ítems rechazados para reintentar")
	}

	// Para cada producto rechazado, sacamos el primer supplier y lo guardamos en un map[supplierId][]items
	supplierItemsMap := make(map[string][]models.ModelPurchaseItem)

	// También preparamos un array con los items que vamos a actualizar de estado a reintentado o no_supplier
	for index, item := range rejectedItems {

		if len(item.SupplierOptions) > 0 {
			// tomamos el primer supplier disponible
			supplierID := item.SupplierOptions[0].SupplierID
			// lo añadimos al map de supplier a items
			supplierItemsMap[supplierID] = append(supplierItemsMap[supplierID], item)

			// actualizamos el item rechazado a reintentado
			itemsToUpdate[index].Status = entities.ItemPurchaseStatusRetried
		} else {
			// si no hay más suppliers, actualizamos el item a no_supplier
			itemsToUpdate[index].Status = entities.ItemPurchaseStatusNoSupplier
		}
	}

	// Actualizamos los items del purchase original
	err = f.appService.PurchaseService.UpdatePurchaseItemsStatus(purchase.ID, itemsToUpdate)
	if err != nil {
		return nil, err
	}

	// Luego creamos las órdenes de compra para cada supplier
	for supplierID, items := range supplierItemsMap {
		purchaseItems := make([]recipe.RecipePurchaseItem, len(items))
		for i, item := range items {
			newPrice := item.SupplierOptions[0].Price
			purchaseItems[i] = recipe.RecipePurchaseItem{
				StoreProductID:  item.StoreProductID,
				Quantity:        item.Quantity,
				PurchaseUnit:    item.PurchaseUnit,
				UnitPrice:       newPrice,
				SupplierOptions: item.SupplierOptions[1:], // Remove the first supplier to avoid using it again
			}
		}

		created, err := f.appService.PurchaseService.CreatePurchaseOrder(ctx, &recipe.RecipePurchase{
			Description:        fmt.Sprintf("Purchase order for inventory request %s", purchase.ID),
			SupplierID:         supplierID,
			StoreID:            purchase.StoreID,
			CompanyID:          purchase.CompanyID,
			InventoryRequestID: purchase.InventoryRequestID,
			Items:              purchaseItems,
		})

		if err != nil {
			fmt.Printf("Failed to create purchase order for supplier %s: %v\n", supplierID, err)
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
		err = f.appService.SupplierOCService.CreateSupplierOC(created.ID, token, &exp)
		if err != nil {
			// si falla la creacion de el token para el supplier se retorna el error?
			fmt.Println("Error creating supplier token:", err)
		}

		// traer email de suplier
		supplierEmail, err := f.appService.SupplierService.GetSupplierContactsEmailByID(supplierID)
		if err != nil {
			fmt.Println("Approve Request Error getting supplier email:", err)
		}

		// enviar el correo con la url junto al token
		err = f.externalservices.EmailService.SendSupplierViewEmail(supplierEmail, token, exp)
		if err != nil {
			fmt.Println("Error sending supplier email:", err)
		}

		err = f.appService.PurchaseService.AddSonOCToPurchase(purchase.ID, created.ID)
		if err != nil {
			return nil, err
		}

	}

	return f.GetPurchaseByID(ctx, purchaseID)
}

func (f *PurchaseFacade) CancelPurchase(ctx context.Context, purchaseID string, observation string) (*dto.DTOPurchase, error) {
	purchase, err := f.appService.PurchaseService.GetPurchaseByID(ctx, purchaseID, false)
	if err != nil {
		return nil, err
	}

	if purchase.Status != entities.PurchaseStatusPending {
		return nil, types.ThrowMsg("Solo se pueden cancelar las compras pendientes")
	}

	err = f.appService.PurchaseService.CancelPurchase(ctx, purchaseID, observation)
	if err != nil {
		return nil, err
	}

	return f.GetPurchaseByID(ctx, purchaseID)
}

func (f *PurchaseFacade) ApprovePurchase(ctx context.Context, purchaseID string) (*dto.DTOPurchase, error) {
	purchase, err := f.appService.PurchaseService.GetPurchaseByID(ctx, purchaseID, false)
	if err != nil {
		return nil, err
	}

	if purchase.Status != entities.PurchaseStatusPending {
		return nil, types.ThrowMsg("Solo se pueden aprobar las compras pendientes")
	}

	err = f.appService.PurchaseService.ApprovePurchase(ctx, purchaseID)
	if err != nil {
		return nil, err
	}

	// Crear movimientos de ingreso automáticos hacia bodega de transición
	// Obtener purchase actualizada con items aprobados
	updatedPurchase, err := f.appService.PurchaseService.GetPurchaseByID(ctx, purchaseID, true)
	if err != nil {
		return nil, err
	}

	if err := f.createAndBookIncomingMovements(ctx, updatedPurchase); err != nil {
		return nil, types.ThrowMsg("error al crear los movimientos de ingreso: " + err.Error())
	}

	return f.GetPurchaseByID(ctx, purchaseID)
}

// createAndBookIncomingMovements crea movimientos INPUT automáticos al aprobar OC
// y actualiza stock en bodega de transición con dirección "IN"
func (f *PurchaseFacade) createAndBookIncomingMovements(ctx context.Context, purchase *models.ModelPurchase) error {
	// 1. Obtener bodega de transición de la tienda
	transitionWarehouse, err := f.appService.WarehouseService.GetTransitionWarehouseByStoreID(ctx, purchase.StoreID)
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
		product, err := f.appService.ProductPerStoreService.GetProductPerStoreByID(ctx, item.StoreProductID)
		if err != nil {
			return types.ThrowMsg("no se pudo obtener el producto: " + err.Error())
		}

		var inventoryUnit *string
		if product.UnitInventory.ID > 0 {
			inventoryUnit = &product.UnitInventory.Abbreviation
		}

		// Obtener stock actual en la bodega de transición (antes del movimiento)
		var stockBefore float32 = 0
		warehouseStock, err := f.appService.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
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
			MovedAt:           time.Now(),
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
	createdMovements, err := f.appService.ProductMovementService.CreateNewMovements(ctx, movements, true)
	if err != nil {
		return types.ThrowMsg("error al crear los movimientos: " + err.Error())
	}

	// 4. Actualizar stock en warehouse_per_product para cada producto
	for _, movement := range createdMovements {
		// Verificar si ya existe registro en warehouse_per_product
		existingStock, err := f.appService.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
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
			_, err := f.appService.WarehousePerProductService.CreateNewWPP(ctx, newStock)
			if err != nil {
				return types.ThrowMsg("error al crear stock en bodega de transición: " + err.Error())
			}
		} else {
			// Actualizar stock existente sumando la cantidad
			oldStock := existingStock.InStock
			existingStock.InStock += movement.Quantity
			existingStock.Direction = &directionIn
			existingStock.WarehouseIdReference = &warehouseRef // Bodega destino de la tienda

			// Recalcular costo promedio
			if existingStock.InStock > 0 {
				totalCost := oldStock * existingStock.CostAvg
				totalCost += movement.TotalCost
				existingStock.CostAvg = totalCost / existingStock.InStock
			}

			_, err := f.appService.WarehousePerProductService.UpdateWPP(ctx, existingStock)
			if err != nil {
				return types.ThrowMsg("error al actualizar stock en bodega de transición: " + err.Error())
			}
		}
	}

	return nil
}

func (f *PurchaseFacade) toDto(purchase *models.ModelPurchase) *dto.DTOPurchase {
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
