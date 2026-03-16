package facades

import (
	"sofia-backend/config"
	"sofia-backend/domain/external"
	"sofia-backend/domain/services"
)

type InventoryRequestFacade struct {
	appServices      *services.ServiceContainer
	externalservices *external.ServiceContainer
	isDebug          bool
}

func NewInventoryRequestFacade(appServices *services.ServiceContainer, externalServices *external.ServiceContainer, config *config.Config) *InventoryRequestFacade {
	return &InventoryRequestFacade{
		appServices:      appServices,
		externalservices: externalServices,
		isDebug:          config.Debug,
	}
}

/*
func (f *InventoryRequestFacade) GetAllInventoryRequests(ctx context.Context, storeID string, page int, size int, filter *map[string]interface{}) (shared.PaginationResponse[models.ModelInventoryRequest], error) {
	data, total, err := f.appServices.InventoryRequestService.GetAllInventoryRequests(ctx, storeID, page, size, filter)
	if err != nil {
		return shared.PaginationResponse[models.ModelInventoryRequest]{}, err
	}

	return shared.NewPagination(data, total, page, size), nil
}

func (f *InventoryRequestFacade) GetInventoryRequestByID(ctx context.Context, requestID string) (*dto.DTOInventoryRequest, error) {
	request, err := f.appServices.InventoryRequestService.GetInventoryRequestByID(ctx, requestID)
	if err != nil {
		return nil, err
	}

	products, err := f.appServices.ProductPerStoreService.GetOnlyRestrictedProducts(ctx, request.StoreID)
	if err != nil {
		return nil, err
	}

	if request == nil {
		return nil, types.ThrowData("Solicitud de inventario no encontrada")
	}

	user, err := f.appServices.UserService.GetUserByID(ctx, request.RequesterID)
	if err != nil {
		return nil, err
	}

	// Create a map for quick lookup of product IDs
	maxQuantityProdudctMap := f.createMap(products)

	dtoRequest := &dto.DTOInventoryRequest{
		ID:             request.ID,
		DisplayID:      request.DisplayID,
		StoreID:        request.StoreID,
		StoreName:      request.StoreName,
		WarehouseID:    request.WarehouseID,
		WarehouseName:  request.WarehouseName,
		Status:         request.Status.ToString(),
		RequestType:    request.RequestType.ToString(),
		RequesterID:    request.RequesterID,
		RequesterName:  user.UserName, // Assuming UserName is the name of the requester
		CreatedAt:      request.CreatedAt,
		UpdatedAt:      request.UpdatedAt,
		Items:          make([]dto.DTOInventoryRequestItem, len(request.Items)),
		RequestHistory: make([]dto.DTOInventoryHistoryStatus, len(request.RequestHistory)),
		Conflicts:      []dto.DTORequestConflictWrapper{},
	}

	for i, item := range request.Items {
		dtoRequest.Items[i] = dto.DTOInventoryRequestItem{
			ItemID:       item.StoreProductID,
			Quantity:     item.Quantity,
			PurchaseUnit: item.PurchaseUnit,
		}
		if product, exists := maxQuantityProdudctMap[item.StoreProductID]; exists {
			// El tipo de conflicto es "MaxQuantity" si la cantidad supera el máximo permitido
			if item.Quantity > *product.MaxQuantity {
				modelConflict := models.ModelMaxQuantityConflict{
					MaxQuantity:  *product.MaxQuantity,
					CurrQuantity: item.Quantity,
				}
				dtoRequest.Conflicts = append(dtoRequest.Conflicts, dto.DTORequestConflictWrapper{
					ItemID: item.StoreProductID,
					Type:   modelConflict.GetConflictType(),
					Detail: modelConflict,
				})
			}
		}
	}

	for i, status := range request.RequestHistory {
		dtoRequest.RequestHistory[i] = dto.DTOInventoryHistoryStatus{
			NewStatus:     status.NewStatus,
			ChangedAt:     status.ChangedAt,
			Observation:   status.Observation,
			ChangedByName: status.ChangedByName,
		}
	}

	return dtoRequest, err
}

func (f *InventoryRequestFacade) CreateInventoryRequest(ctx context.Context, rec *recipe.RecipeInventoryRequest) (*dto.DTOInventoryRequest, error) {
	// Get restricted products for the store
	productRestricted, err := f.appServices.ProductPerStoreService.GetOnlyRestrictedProducts(ctx, rec.StoreID)
	if err != nil {
		return nil, err
	}
	// Validar las cantidades de los productos en la solicitud
	// Si no hay conflictos, se aprueba la solicitud
	status := f.validateProducts(productRestricted, rec.Items, entities.RequestStatusApproved)

	request, err := f.appServices.InventoryRequestService.CreateInventoryRequest(ctx, rec, status)
	if err != nil {
		return nil, err
	}

	user, err := f.appServices.UserService.GetUserByID(ctx, rec.RequesterID)
	if err != nil {
		return nil, err
	}

	usersIds, err := f.getResolverUsers(ctx, request.StoreID)
	if err != nil {
		return nil, err
	}

	userId, ok := f.appServices.AuthService.GetUserIDFromContext(ctx)
	if !ok {
		return nil, types.ThrowData("Error al obtener el ID del usuario del contexto")
	}

	// Notify all users with purchase:resolve and store:<store_id> powers except the one who made the change
	if request.Status == entities.RequestStatusConflicted {
		f.notify(userId, usersIds, request.ID, "created")
	}

	if request.Status == entities.RequestStatusApproved {
		// Notify requester that their request has been approved
		f.approveRequest(ctx, request.ID, userId)
	}

	dtoRequest := &dto.DTOInventoryRequest{
		ID:            request.ID,
		DisplayID:     request.DisplayID,
		StoreID:       request.StoreID,
		StoreName:     request.StoreName,
		WarehouseID:   request.WarehouseID,
		WarehouseName: request.WarehouseName,
		Status:        request.Status.ToString(),
		RequestType:   request.RequestType.ToString(),
		RequesterID:   request.RequesterID,
		RequesterName: user.UserName, // Assuming UserName is the name of the requester
		CreatedAt:     request.CreatedAt,
		UpdatedAt:     request.UpdatedAt,
	}

	return dtoRequest, nil
}

func (f *InventoryRequestFacade) UpdateInventoryRequest(ctx context.Context, id string, recipe *recipe.RecipeInventoryRequest) (*dto.DTOInventoryRequest, error) {
	requestModel, err := f.appServices.InventoryRequestService.GetInventoryRequestByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if requestModel == nil {
		return nil, types.ThrowData("Solicitud de inventario no encontrada")
	}

	// Flujo Normal, actualizar la solicitud para que aprobada si no tiene conflictos
	// Get restricted products for the store
	productRestricted, err := f.appServices.ProductPerStoreService.GetOnlyRestrictedProducts(ctx, recipe.StoreID)
	if err != nil {
		return nil, err
	}

	// Validar las cantidades de los productos en la solicitud
	userId, ok := f.appServices.ProfileService.GetUserIDFromContext(ctx)
	if !ok {
		return nil, types.ThrowData("Error al obtener el ID del usuario del contexto")
	}

	powers, err := f.appServices.ProfileService.GetPowersByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	canSetApproved := false
	for _, power := range powers {
		if power.PowerName == "request:resolve" {
			// If the user has the purchase:resolve power, they can set the status to approved directly
			canSetApproved = true
			break
		}
	}

	var status entities.RequestStatus
	if canSetApproved {
		status = requestModel.Status
	} else {
		status = f.validateProducts(productRestricted, recipe.Items, requestModel.Status)
	}

	// Update the request
	request, err := f.appServices.InventoryRequestService.UpdateInventoryRequest(ctx, id, recipe, status)
	if err != nil {
		return nil, err
	}

	// If the status is changing to conflicted, notify the resolvers
	usersIds, err := f.getResolverUsers(ctx, requestModel.StoreID)
	if err != nil {
		return nil, err
	}

	// Always notify the requester
	// check if the requester is already in the list
	var found bool
	for _, uid := range usersIds {
		if uid == requestModel.RequesterID {
			found = true
			break
		}
	}
	if !found {
		usersIds = append(usersIds, requestModel.RequesterID)
	}
	// Notify all users with purchase:resolve and store:<store_id> powers except the one who made the change
	f.notify(userId, usersIds, requestModel.ID, "updated")

	dtoRequest := &dto.DTOInventoryRequest{
		ID:      request.ID,
		StoreID: request.StoreID,
		// StoreName:     request.StoreName,
		WarehouseID: request.WarehouseID,
		// WarehouseName: request.WarehouseName,
		Status:      request.Status.ToString(),
		RequestType: request.RequestType.ToString(),
		RequesterID: request.RequesterID,
		// RequesterName: request.RequesterName, // Assuming UserName is the name of the requester
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}

	return dtoRequest, nil
}

func (f *InventoryRequestFacade) GetProductWithConflictByRequestId(ctx context.Context, storeId string, requestId string) ([]models.ModelProductRequestRestriction, error) {
	products, err := f.appServices.ProductPerStoreService.GetOnlyRestrictedProducts(ctx, storeId)
	if err != nil {
		return nil, err
	}

	request, err := f.appServices.InventoryRequestService.GetInventoryRequestByID(ctx, requestId)
	if err != nil {
		return nil, err
	}

	// Create a map for quick lookup of product IDs
	maxQuantityProdudctMap := f.createMap(products)

	// Filtrar los productos restringidos basados en la solicupdateditud
	var restrictedProducts []models.ModelProductRequestRestriction
	for _, item := range request.Items {
		if product, isRestricted := maxQuantityProdudctMap[item.StoreProductID]; isRestricted {
			if item.Quantity > *product.MaxQuantity {
				restrictedProducts = append(restrictedProducts, product)
			}
		}
	}

	return restrictedProducts, nil
}

func (f *InventoryRequestFacade) ApprovedAndUpdate(ctx context.Context, id string, recipe *recipe.RecipeInventoryRequest) (*dto.DTOInventoryRequest, error) {
	requestModel, err := f.appServices.InventoryRequestService.GetInventoryRequestByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if requestModel == nil {
		return nil, types.ThrowData("Solicitud de inventario no encontrada")
	}

	request, err := f.appServices.InventoryRequestService.ApproveAndUpdateRequest(ctx, id, recipe)
	if err != nil {
		return nil, err
	}

	userId, ok := f.appServices.AuthService.GetUserIDFromContext(ctx)
	if !ok {
		return nil, types.ThrowData("Error al obtener el ID del usuario del contexto")
	}

	// Notify requester that their request has been approved
	f.approveRequest(ctx, request.ID, userId)

	user, err := f.appServices.UserService.GetUserByID(ctx, recipe.RequesterID)
	if err != nil {
		return nil, err
	}

	dtoRequest := &dto.DTOInventoryRequest{
		ID:            request.ID,
		DisplayID:     request.DisplayID,
		StoreID:       request.StoreID,
		WarehouseID:   request.WarehouseID,
		Status:        request.Status.ToString(),
		RequestType:   request.RequestType.ToString(),
		RequesterID:   request.RequesterID,
		RequesterName: user.UserName, // Assuming UserName is the name of the requester
		CreatedAt:     request.CreatedAt,
		UpdatedAt:     request.UpdatedAt,
	}

	return dtoRequest, nil
}

func (f *InventoryRequestFacade) ChangeInventoryRequestStatus(ctx context.Context, id string, status string, recipe *recipe.RecipeInventoryRequestStatus) (*dto.DTOInventoryRequest, error) {
	request, err := f.appServices.InventoryRequestService.GetInventoryRequestByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if request == nil {
		return nil, types.ThrowData("Solicitud de inventario no encontrada")
	}

	if request.Status == entities.RequestStatus(status) {
		return nil, types.ThrowData("La solicitud ya está en el estado solicitado")
	}

	// Check if the status is valid
	if request.Status == entities.RequestStatusApproved {
		return nil, types.ThrowData("La solicitud ya está aprobada y no puede ser modificada")
	}

	updatedRequest, err := f.appServices.InventoryRequestService.ChangeStatus(
		ctx,
		id,
		entities.RequestStatus(status),
		recipe.Observation,
	)
	if err != nil {
		return nil, err
	}

	usersIds, err := f.getResolverUsers(ctx, updatedRequest.StoreID)
	if err != nil {
		return nil, err
	}

	userId, ok := f.appServices.AuthService.GetUserIDFromContext(ctx)
	if !ok {
		return nil, types.ThrowData("Error al obtener el ID del usuario del contexto")
	}

	// Always notify the requester that their request status has changed
	// check if the requester is already in the list
	var found bool
	for _, uid := range usersIds {
		if uid == updatedRequest.RequesterID {
			found = true
			break
		}
	}
	if !found {
		usersIds = append(usersIds, updatedRequest.RequesterID)
	}
	// Notify all users with purchase:resolve and store:<store_id> powers except the one who made the change
	f.notify(userId, usersIds, updatedRequest.ID, "status_changed")

	user, err := f.appServices.UserService.GetUserByID(ctx, updatedRequest.RequesterID)
	if err != nil {
		return nil, err
	}

	dtoRequest := &dto.DTOInventoryRequest{
		ID:            updatedRequest.ID,
		DisplayID:     updatedRequest.DisplayID,
		StoreID:       updatedRequest.StoreID,
		WarehouseID:   updatedRequest.WarehouseID,
		Status:        updatedRequest.Status.ToString(),
		RequestType:   updatedRequest.RequestType.ToString(),
		RequesterID:   updatedRequest.RequesterID,
		RequesterName: user.UserName, // Assuming UserName is the name of the requester
		CreatedAt:     updatedRequest.CreatedAt,
		UpdatedAt:     updatedRequest.UpdatedAt,
	}

	return dtoRequest, nil
}

func (f *InventoryRequestFacade) validateProducts(restrictedProducts []models.ModelProductRequestRestriction, items []recipe.RecipeInventoryRequestItem, oldStatus entities.RequestStatus) entities.RequestStatus {
	// Create a map for quick lookup of restricted product IDs
	restrictedProductMap := f.createMap(restrictedProducts)

	// Validar las cantidades de los productos en la solicitud
	status := oldStatus
	for _, item := range items {
		if product, isRestricted := restrictedProductMap[item.ItemID]; isRestricted {
			if item.Quantity > *product.MaxQuantity {
				status = entities.RequestStatusConflicted
				break
			}
		}
	}
	return status
}

func (f *InventoryRequestFacade) createMap(restrictedProducts []models.ModelProductRequestRestriction) map[string]models.ModelProductRequestRestriction {
	// Create a map for quick lookup of restricted product IDs
	restrictedProductMap := make(map[string]models.ModelProductRequestRestriction)
	for _, product := range restrictedProducts {
		if product.MaxQuantity == nil {
			continue
		}
		restrictedProductMap[product.ID] = product
	}
	return restrictedProductMap
}

func (f *InventoryRequestFacade) getResolverUsers(ctx context.Context, storeId string) ([]string, error) {
	// TODO: fix this to use the constant
	profilesStore, err := f.appServices.ProfileService.GetProfilesByPowerID(ctx, fmt.Sprintf("store:%s", storeId))
	if err != nil {
		fmt.Printf("Failed to get profiles by power ID: %v\n", err)
		return nil, err
	}

	// TODO: fix this to use the constant
	profilesPurchase, err := f.appServices.ProfileService.GetProfilesByPowerID(ctx, "request:resolve")
	if err != nil {
		fmt.Printf("Failed to get profiles by power ID: %v\n", err)
		return nil, err
	}

	profilePurchaseIDs := make([]string, len(profilesPurchase))
	for i, profile := range profilesPurchase {
		profilePurchaseIDs[i] = profile.ID
	}
	userPurchase, err := f.appServices.UserService.GetUsersByProfileIDs(ctx, profilePurchaseIDs)
	if err != nil {
		fmt.Printf("Failed to get users by profile IDs: %v\n", err)
		return nil, err
	}

	profileStoreIDs := make([]string, len(profilesStore))
	for i, profile := range profilesStore {
		profileStoreIDs[i] = profile.ID
	}
	userStore, err := f.appServices.UserService.GetUsersByProfileIDs(ctx, profileStoreIDs)
	if err != nil {
		fmt.Printf("Failed to get users by profile IDs: %v\n", err)
		return nil, err
	}

	purchaseMap := make(map[string]string, len(userPurchase))
	for _, u := range userPurchase {
		purchaseMap[u.ID] = u.ID
	}
	// filter store users that also exist in purchase users
	var userIDsToNotify []string
	for _, u := range userStore {
		if _, ok := purchaseMap[u.ID]; ok {
			userIDsToNotify = append(userIDsToNotify, u.ID)
		}
	}

	return userIDsToNotify, nil
}

func (f *InventoryRequestFacade) notify(from string, userIDsToNotify []string, requestID string, event string) {
	// send notification to users in both
	var notificationsMessage []recipe.SendNotificationRecipe
	for _, userID := range userIDsToNotify {
		if userID == from {
			continue
		}
		notificationsMessage = append(notificationsMessage, recipe.SendNotificationRecipe{
			From:             from,
			To:               userID,
			NotificationType: "inventory_request",
			Payload: map[string]interface{}{
				"message":   "There is an update on the inventory request",
				"event":     event,
				"timestamp": time.Now().Format(time.RFC3339),
				"requestId": requestID,
			},
		})
	}

	if len(notificationsMessage) == 0 {
		return
	}

	if _, err := f.appServices.NotificationService.SendMultipleMessages(notificationsMessage); err != nil {
		fmt.Printf("Failed to send notifications: %v\n", err)
		return
	}

}

func (f *InventoryRequestFacade) approveRequest(ctx context.Context, requestId string, userId string) error {
	request, err := f.appServices.InventoryRequestService.GetInventoryRequestByID(ctx, requestId)
	if err != nil {
		return err
	}
	itemIds := make([]string, len(request.Items))
	for i, item := range request.Items {
		itemIds[i] = item.ProductID
	}

	productSupplierMap, err := f.appServices.SupplierService.GetSuppliersByStoreProductId(request.StoreID, itemIds)
	if err != nil {
		return err
	}

	// group items by supplier taking the first supplier found for each product
	supplierItemsMap := make(map[string][]models.ModelInventoryRequestItem)
	for _, item := range request.Items {
		if suppliers, ok := productSupplierMap[item.ProductID]; ok && len(suppliers) > 0 {
			supplierID := suppliers[0].SupplierID
			supplierItemsMap[supplierID] = append(supplierItemsMap[supplierID], item)
		} else {
			fmt.Printf("No supplier found for product %s\n", item.ProductID)
		}
	}

	// Create purchase orders for each supplier
	for supplierID, items := range supplierItemsMap {
		purchaseItems := make([]recipe.RecipePurchaseItem, len(items))
		for i, item := range items {
			supplierItem := productSupplierMap[item.ProductID][0]
			purchaseItems[i] = recipe.RecipePurchaseItem{
				StoreProductID:  item.StoreProductID,
				Quantity:        item.Quantity,
				UnitPrice:       supplierItem.Price,
				SupplierOptions: productSupplierMap[item.ProductID][1:], // Tomo solamente los suppliers alternativos
			}
		}

		purchase, err := f.appServices.PurchaseService.CreatePurchaseOrder(ctx, &recipe.RecipePurchase{
			Description:        fmt.Sprintf("Purchase order for inventory request %s", request.ID),
			SupplierID:         supplierID,
			CompanyID:          request.CompanyID,
			StoreID:            request.StoreID,
			InventoryRequestID: request.ID,
			Items:              purchaseItems,
		})
		if err != nil {
			// Continue creating orders for other suppliers even if one fails
			fmt.Printf("Failed to create purchase order for supplier %s: %v\n", supplierID, err)
			// Continue creating orders for other suppliers even if one fails
			continue
		}

		////////// Supplier token creation //////////

		// aqui debe ir la creacion del token para las OC
		var token string
		if f.isDebug {
			token = fmt.Sprintf("token-%s", purchase.DisplayID)
		} else {
			token, err = shared.CreateRandomURLString(64)
			if err != nil {
				fmt.Println("Error creating token:", err)
			}
		}

		// expira en 7 dias
		exp := time.Now().Add(time.Hour * 24 * 7).UTC() // una semana en UTC

		// crear el supplier token
		err = f.appServices.SupplierOCService.CreateSupplierOC(purchase.ID, token, &exp)
		if err != nil {
			// si falla la creacion de el token para el supplier se retorna el error?
			fmt.Println("Error creating supplier token:", err)
		}

		// traer email de suplier
		supplierEmail, err := f.appServices.SupplierService.GetSupplierContactsEmailByID(supplierID)
		if err != nil {
			fmt.Println("Approve Request Error getting supplier email:", err)
		}

		// enviar el correo con la url junto al token
		err = f.externalservices.EmailService.SendSupplierViewEmail(supplierEmail, token, exp)
		if err != nil {
			fmt.Println("Error sending supplier email:", err)
		}

		////////// Supplier token creation //////////

	}

	f.notify(userId, []string{request.RequesterID}, request.ID, "approved")

	return nil
}
*/
