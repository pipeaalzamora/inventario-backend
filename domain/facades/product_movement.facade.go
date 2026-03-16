package facades

import (
	"context"
	"fmt"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/shared"
	"sofia-backend/types"
	"time"
)

type ProductMovementFacade struct {
	box *services.ServiceContainer
}

func NewProductMovementFacade(box *services.ServiceContainer) *ProductMovementFacade {
	return &ProductMovementFacade{
		box: box,
	}
}

func (f *ProductMovementFacade) GetAllProductMovements() (shared.PaginationResponse[models.ModelProductMovement], error) {
	movementList := make([]models.ModelProductMovement, 0)

	movementList, err := f.box.ProductMovementService.GetAllProductMovements()
	if err != nil {
		return shared.PaginationResponse[models.ModelProductMovement]{}, err
	}

	return shared.NewPagination(movementList, len(movementList), 1, len(movementList)), nil
}

func (f *ProductMovementFacade) GetProductMovementByID(movementId string) (*models.ModelProductMovement, error) {
	movement, err := f.box.ProductMovementService.GetProductMovementByID(movementId)
	if err != nil {
		return nil, err
	}

	return movement, nil
}

func (f *ProductMovementFacade) GetAllProductMovementsByStoreProductID(ctx context.Context, storeId string, storeProductId string) (shared.PaginationResponse[models.ModelProductMovement], error) {

	emptyResponse := shared.PaginationResponse[models.ModelProductMovement]{}

	// Validar que la tienda existe y el usuario tiene permisos
	_, err := f.box.StoreService.GetStoreByID(ctx, storeId)
	if err != nil {
		return emptyResponse, err
	}

	// Obtener el producto y verificar que pertenece a la tienda especificada
	product, err := f.box.ProductPerStoreService.GetProductPerStoreByID(ctx, storeProductId)
	if err != nil {
		return emptyResponse, err
	}

	// Verificar que el producto pertenece a la tienda especificada
	if product.StoreID != storeId {
		return emptyResponse, types.ThrowData(fmt.Sprintf(
			"El producto %s no pertenece a la tienda especificada (store_id: %s)",
			storeProductId,
			storeId,
		))
	}

	// Obtener movimientos del producto
	movements, err := f.box.ProductMovementService.GetAllProductMovementsByStoreProductID(ctx, storeId, storeProductId)
	if err != nil {
		return emptyResponse, err
	}

	return shared.NewPagination(movements, len(movements), 1, len(movements)), nil
}

func (f *ProductMovementFacade) GetAllProductMovementsByStoreID(ctx context.Context, storeId string) (shared.PaginationResponse[models.ModelProductMovement], error) {
	emptyResponse := shared.PaginationResponse[models.ModelProductMovement]{}
	_, err := f.box.StoreService.GetStoreByID(ctx, storeId)
	if err != nil {
		return emptyResponse, err
	}

	movements, err := f.box.ProductMovementService.GetAllProductMovementsByStoreId(ctx, storeId)
	if err != nil {
		return emptyResponse, err
	}

	return shared.NewPagination(movements, len(movements), 1, len(movements)), nil
}

func (f *ProductMovementFacade) GetProductMovementsByWharehouseIDs(warehouseIDs []string) (shared.PaginationResponse[models.ModelProductMovement], error) {
	movements := make([]models.ModelProductMovement, 0)

	if len(warehouseIDs) == 0 {
		return shared.NewPagination(movements, 0, 1, 0), types.ThrowMsg("Se debe proporcionar al menos un ID de bodega.")
	}

	movements, err := f.box.ProductMovementService.GetAllProductMovementsByWarehouseIDs(warehouseIDs)
	if err != nil {
		return shared.PaginationResponse[models.ModelProductMovement]{}, err
	}
	return shared.NewPagination(movements, len(movements), 1, len(movements)), nil
}

func (f *ProductMovementFacade) GetAllProductMovementsByDateRange(warehouseID string) (shared.PaginationResponse[models.ModelProductMovement], error) {
	movements, err := f.box.ProductMovementService.GetAllProductMovementsByDateRange(warehouseID)
	if err != nil {
		return shared.PaginationResponse[models.ModelProductMovement]{}, err
	}
	return shared.NewPagination(movements, len(movements), 1, len(movements)), nil
}

func (f *ProductMovementFacade) CreateNewInputMovement(ctx context.Context, inputRecipe recipe.RecipeNewMovement) ([]models.ModelProductMovement, error) {
	_, err := f.box.StoreService.GetStoreByID(ctx, inputRecipe.StoreId)
	if err != nil {
		return nil, err
	}

	if inputRecipe.MovedTo == nil {
		return nil, types.ThrowMsg("Debe proporcionarse el ID de la bodega destino para los movimientos de entrada.")
	}

	moveToWarehouse, err := f.box.WarehouseService.GetWarehouseByID(ctx, *inputRecipe.MovedTo)
	if err != nil {
		return nil, err
	}

	if err := f.validateRecipeProducts(ctx, inputRecipe); err != nil {
		return nil, err
	}

	_movements := make([]models.ModelProductMovement, 0)
	for _, item := range inputRecipe.Products {
		// Obtener información del producto para la unidad de inventario
		product, err := f.box.ProductPerStoreService.GetProductPerStoreByID(ctx, item.StoreProductID)
		if err != nil {
			return nil, err
		}

		var inventoryUnit *string
		if product.UnitInventory.ID > 0 {
			inventoryUnit = &product.UnitInventory.Abbreviation
		}

		// Obtener stock actual en la bodega destino (antes del movimiento)
		var stockBefore float32 = 0
		warehouseStock, err := f.box.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
			ctx, item.StoreProductID, moveToWarehouse.ID,
		)
		if err == nil && warehouseStock != nil {
			stockBefore = warehouseStock.InStock
		}

		// Calcular stock después del movimiento
		stockAfter := stockBefore + item.Quantity

		// Obtener costo promedio para calcular unit_cost y total_cost
		var unitCost float32 = 0
		if warehouseStock != nil {
			unitCost = warehouseStock.CostAvg
		}
		totalCost := unitCost * item.Quantity

		newMovement := models.ModelProductMovement{
			StoreProductID: item.StoreProductID,
			Observation:    inputRecipe.Observation,
			Quantity:       item.Quantity,
			InventoryUnit:  inventoryUnit,
			UnitCost:       unitCost,
			TotalCost:      totalCost,
			MovedFrom:      nil,
			MovedTo:        &moveToWarehouse.ID,
			MovementType:   "NEWINPUT",
			MovedAt:        time.Now(),
			StockBefore:    &stockBefore,
			StockAfter:     &stockAfter,
		}
		_movements = append(_movements, newMovement)
	}

	return f.box.ProductMovementService.CreateNewMovements(ctx, _movements, false)
}

func (f *ProductMovementFacade) CreateNewOutputMovement(ctx context.Context, outputRecipe recipe.RecipeNewMovement) ([]models.ModelProductMovement, error) {
	/*
		_, err := f.box.StoreService.GetStoreByID(ctx, outputRecipe.StoreId)
		if err != nil {
			return nil, err
		}
	*/

	if outputRecipe.MovedFrom == nil {
		return nil, types.ThrowMsg("Debe proporcionarse el ID de la bodega de origen para los movimientos de salida.")
	}

	if !f.validateWasteKind(outputRecipe.WasteKind) {
		return nil, types.ThrowMsg("Tipo de salida proporcionada no válida.")
	}

	moveFromWarehouse, err := f.box.WarehouseService.GetWarehouseByID(ctx, *outputRecipe.MovedFrom)
	if err != nil {
		return nil, err
	}

	if err := f.validateRecipeProducts(ctx, outputRecipe); err != nil {
		return nil, err
	}

	_movements := make([]models.ModelProductMovement, 0)
	for _, item := range outputRecipe.Products {
		// Obtener información del producto para la unidad de inventario
		product, err := f.box.ProductPerStoreService.GetProductPerStoreByID(ctx, item.StoreProductID)
		if err != nil {
			return nil, err
		}

		var inventoryUnit *string
		if product.UnitInventory.ID > 0 {
			inventoryUnit = &product.UnitInventory.Abbreviation
		}

		// Obtener stock actual en la bodega origen (antes del movimiento)
		var stockBefore float32 = 0
		var unitCost float32 = 0
		warehouseStock, err := f.box.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
			ctx, item.StoreProductID, moveFromWarehouse.ID,
		)
		if err == nil && warehouseStock != nil {
			stockBefore = warehouseStock.InStock
			unitCost = warehouseStock.CostAvg
		}

		// Calcular stock después del movimiento (salida resta)
		stockAfter := stockBefore - item.Quantity
		totalCost := unitCost * item.Quantity

		newMovement := models.ModelProductMovement{
			StoreProductID: item.StoreProductID,
			Observation:    outputRecipe.Observation,
			Quantity:       item.Quantity,
			InventoryUnit:  inventoryUnit,
			UnitCost:       unitCost,
			TotalCost:      totalCost,
			MovedFrom:      &moveFromWarehouse.ID,
			MovedTo:        nil,
			MovementType:   string(outputRecipe.WasteKind),
			MovedAt:        time.Now(),
			StockBefore:    &stockBefore,
			StockAfter:     &stockAfter,
		}
		_movements = append(_movements, newMovement)
	}

	return f.box.ProductMovementService.CreateNewMovements(ctx, _movements, false)
}

func (f *ProductMovementFacade) validateRecipeProducts(ctx context.Context, movementRecipe recipe.RecipeNewMovement) error {
	for _, item := range movementRecipe.Products {
		product, err := f.box.ProductPerStoreService.GetProductPerStoreByID(ctx, item.StoreProductID)
		if err != nil {
			return err
		}

		// Verificar que el producto pertenezca a la tienda especificada
		if product.StoreID != movementRecipe.StoreId {
			return types.ThrowData(fmt.Sprintf(
				"El producto %s no pertenece a la tienda especificada (store_id: %s)",
				item.StoreProductID,
				movementRecipe.StoreId,
			))
		}
	}
	return nil
}

func (f *ProductMovementFacade) validateWasteKind(kind recipe.WasteKind) bool {
	switch kind {
	case "NONE",
		"WASTE",
		"EXPIRED",
		"ADJUSTED",
		"OTHER":
		return true
	default:
		return false
	}
}

// CreateTransferMovement transfiere productos entre bodegas de la misma tienda
// Actualiza el costo promedio en la bodega destino usando la fórmula:
// avgCost = ((cantidadActual * costoPromedioActual) + (cantidadIngreso * costoPorUnidad)) / (cantidadActual + cantidadIngreso)
func (f *ProductMovementFacade) CreateTransferMovement(ctx context.Context, transferRecipe recipe.RecipeTransferMovement) ([]models.ModelProductMovement, error) {
	// 1. Validar que la tienda existe y el usuario tiene permisos
	_, err := f.box.StoreService.GetStoreByID(ctx, transferRecipe.StoreId)
	if err != nil {
		return nil, err
	}

	// 2. Validar que las bodegas origen y destino son diferentes
	if transferRecipe.FromWarehouseId == transferRecipe.ToWarehouseId {
		return nil, types.ThrowRecipe("La bodega origen y destino deben ser diferentes", "toWarehouseId")
	}

	// 3. Validar que las bodegas existen y pertenecen a la tienda
	fromWarehouse, err := f.box.WarehouseService.GetWarehouseByID(ctx, transferRecipe.FromWarehouseId)
	if err != nil {
		return nil, types.ThrowRecipe("La bodega origen no existe", "fromWarehouseId")
	}
	if fromWarehouse.StoreId != transferRecipe.StoreId {
		return nil, types.ThrowRecipe("La bodega origen no pertenece a la tienda especificada", "fromWarehouseId")
	}

	toWarehouse, err := f.box.WarehouseService.GetWarehouseByID(ctx, transferRecipe.ToWarehouseId)
	if err != nil {
		return nil, types.ThrowRecipe("La bodega destino no existe", "toWarehouseId")
	}
	if toWarehouse.StoreId != transferRecipe.StoreId {
		return nil, types.ThrowRecipe("La bodega destino no pertenece a la tienda especificada", "toWarehouseId")
	}

	// 4. Validar productos y stock disponible
	for _, item := range transferRecipe.Products {
		product, err := f.box.ProductPerStoreService.GetProductPerStoreByID(ctx, item.StoreProductID)
		if err != nil {
			return nil, types.ThrowRecipe("Producto no encontrado: "+item.StoreProductID, "products")
		}
		if product.StoreID != transferRecipe.StoreId {
			return nil, types.ThrowRecipe("El producto "+item.StoreProductID+" no pertenece a la tienda", "products")
		}

		// Verificar stock disponible en bodega origen
		stockOrigen, err := f.box.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
			ctx, item.StoreProductID, transferRecipe.FromWarehouseId,
		)
		if err != nil {
			return nil, types.ThrowRecipe("El producto no existe en la bodega origen", "products")
		}
		if stockOrigen.InStock < item.Quantity {
			return nil, types.ThrowRecipe(
				fmt.Sprintf("Stock insuficiente para el producto. Disponible: %.2f, Solicitado: %.2f", stockOrigen.InStock, item.Quantity),
				"products",
			)
		}
	}

	// 5. Crear los movimientos de transferencia
	_movements := make([]models.ModelProductMovement, 0)
	newAvgCosts := make(map[string]float32)

	for _, item := range transferRecipe.Products {
		product, _ := f.box.ProductPerStoreService.GetProductPerStoreByID(ctx, item.StoreProductID)

		var inventoryUnit *string
		if product.UnitInventory.ID > 0 {
			inventoryUnit = &product.UnitInventory.Abbreviation
		}

		// Obtener stock en bodega origen (antes del movimiento)
		stockOrigen, _ := f.box.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
			ctx, item.StoreProductID, transferRecipe.FromWarehouseId,
		)
		stockBeforeFrom := stockOrigen.InStock

		// Obtener stock en bodega destino (antes del movimiento)
		stockDestino, err := f.box.WarehousePerProductService.GetWarehousePerProductByStoreProductAndWarehouse(
			ctx, item.StoreProductID, transferRecipe.ToWarehouseId,
		)
		
		var stockBeforeTo float32 = 0
		var currentAvgCost float32 = 0
		if err == nil && stockDestino != nil {
			stockBeforeTo = stockDestino.InStock
			currentAvgCost = stockDestino.CostAvg
		}

		stockAfterTo := stockBeforeTo + item.Quantity

		// Calcular nuevo costo promedio ponderado para la bodega destino
		// Formula: avgCost = ((cantidadActual * costoPromedioActual) + (cantidadIngreso * costoPorUnidad)) / (cantidadActual + cantidadIngreso)
		var newAvgCost float32
		if stockAfterTo > 0 {
			newAvgCost = ((stockBeforeTo * currentAvgCost) + (item.Quantity * item.UnitCost)) / stockAfterTo
		} else {
			newAvgCost = item.UnitCost
		}

		// Guardar el nuevo costo promedio en el mapa usando una clave única
		key := item.StoreProductID + "|" + transferRecipe.ToWarehouseId
		newAvgCosts[key] = newAvgCost

		totalCost := item.Quantity * item.UnitCost

		newMovement := models.ModelProductMovement{
			StoreProductID: item.StoreProductID,
			Observation:    transferRecipe.Observation,
			Quantity:       item.Quantity,
			InventoryUnit:  inventoryUnit,
			UnitCost:       item.UnitCost,
			TotalCost:      totalCost,
			MovedFrom:      &transferRecipe.FromWarehouseId,
			MovedTo:        &transferRecipe.ToWarehouseId,
			MovementType:   "TRANSFER",
			MovedAt:        time.Now(),
			StockBefore:    &stockBeforeFrom,
			StockAfter:     &stockAfterTo,
		}

		_movements = append(_movements, newMovement)
	}

	// 6. Ejecutar los movimientos con actualización de costo promedio
	return f.box.ProductMovementService.CreateTransferMovements(ctx, _movements, newAvgCosts)
}
