package facades

import (
	"context"
	"sofia-backend/api/v1/dto"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
)

type InventoryFacade struct {
	serviceBox *services.ServiceContainer
}

func NewInventoryFacade(serviceBox *services.ServiceContainer) *InventoryFacade {
	return &InventoryFacade{
		serviceBox: serviceBox,
	}
}

func (f *InventoryFacade) GetAllInventoryByStoreId(ctx context.Context, companyId string, storeId string, warehouseIds []string) ([]models.ModelWarehouseProductInventory, error) {
	// 1. Obtener todos los productos de la tienda
	productsStore, err := f.serviceBox.ProductPerStoreService.GetProductsByStore(ctx, storeId)
	if err != nil {
		return nil, err
	}

	// 2. Obtener bodegas de la tienda
	allWarehouses, err := f.serviceBox.WarehouseService.GetWarehousesByStoreId(ctx, storeId)
	if err != nil {
		return nil, err
	}

	// 3. Filtrar bodegas si se especificaron IDs
	var targetWarehouses []models.ModelWarehouse
	if len(warehouseIds) > 0 {
		warehouseIDSet := make(map[string]bool)
		for _, id := range warehouseIds {
			warehouseIDSet[id] = true
		}
		for _, w := range allWarehouses {
			if warehouseIDSet[w.ID] {
				targetWarehouses = append(targetWarehouses, w)
			}
		}
	} else {
		targetWarehouses = allWarehouses
	}

	// 4. Obtener stock por bodega (bodegas normales)
	warehouseStock, err := f.serviceBox.InventoryService.GetProductStoreInventoryByStoreID(ctx, storeId, warehouseIds)
	if err != nil {
		return nil, err
	}

	// 5. Obtener stock en tránsito (bodegas de traspaso)
	transitStock, err := f.serviceBox.InventoryService.GetProductTransitByReferences(ctx, storeId, warehouseIds)
	if err != nil {
		return nil, err
	}

	// 6. Crear mapa de bodegas con sus productos
	type warehouseData struct {
		warehouseName string
		productMap    map[string]models.ModelProductStoreStock
	}
	warehouseMap := make(map[string]*warehouseData)

	// 7. Inicializar estructura de bodegas (vacías)
	for _, warehouse := range targetWarehouses {
		warehouseMap[warehouse.ID] = &warehouseData{
			warehouseName: warehouse.WarehouseName,
			productMap:    make(map[string]models.ModelProductStoreStock),
		}
	}

	// 8. Crear mapa de productos para lookup rápido
	productStoreMap := make(map[string]models.ModelProductPerStore)
	for _, product := range productsStore {
		productStoreMap[product.ID] = product
	}

	// 9. Agregar solo productos que tienen registro en warehouse_per_product (desde warehouseStock)
	for _, ws := range warehouseStock {
		// Solo procesar si la bodega existe en nuestro mapa
		if wData, ok := warehouseMap[ws.WarehouseID]; ok {
			// Buscar el producto en el mapa de productos de la tienda
			if product, exists := productStoreMap[ws.StoreProductId]; exists {
				// Agregar el producto a esta bodega con su stock real
				wData.productMap[ws.StoreProductId] = models.ModelProductStoreStock{
					ModelProductPerStore: product,
					Totals: models.ModelProductStoreStockTotals{
						CurrentStock: ws.CurrentStock,
						StockIn:      0,
						StockOut:     0,
						AvgCost:      ws.AvgCost,
						TotalCost:    ws.CurrentStock * ws.AvgCost,
						AreMinAlert:  ws.CurrentStock < product.Quantities.MinimalStock,
						AreMaxAlert:  ws.CurrentStock > product.Quantities.MaximalStock,
					},
				}
			}
		}
	}

	// 10. Procesar stock en tránsito y agregarlo a los productos correspondientes
	for _, ts := range transitStock {
		warehouseID := ts.WarehouseIDReference

		// Solo procesar si la bodega de referencia existe en nuestro mapa
		if _, ok := warehouseMap[warehouseID]; !ok {
			continue
		}

		// Solo procesar si el producto existe en esa bodega
		if prod, ok := warehouseMap[warehouseID].productMap[ts.StoreProductId]; ok {
			if ts.Direction == "IN" {
				prod.Totals.StockIn += ts.InStock
			} else if ts.Direction == "OUT" {
				prod.Totals.StockOut += ts.InStock
			}
			warehouseMap[warehouseID].productMap[ts.StoreProductId] = prod
		}
	}

	// 11. Calcular costo promedio ponderado para cada producto en cada bodega
	for warehouseID, wData := range warehouseMap {
		for productID, prod := range wData.productMap {
			if prod.Totals.CurrentStock > 0 {
				prod.Totals.AvgCost = prod.Totals.TotalCost / prod.Totals.CurrentStock
			}
			warehouseMap[warehouseID].productMap[productID] = prod
		}
	}

	// 12. Convertir mapa a lista de ModelWarehouseProductInventory
	var result []models.ModelWarehouseProductInventory
	for warehouseID, wData := range warehouseMap {
		var products []models.ModelProductStoreStock
		for _, prod := range wData.productMap {
			products = append(products, prod)
		}

		result = append(result, models.ModelWarehouseProductInventory{
			WarehouseID:   warehouseID,
			WarehouseName: wData.warehouseName,
			Products:      products,
		})
	}

	return result, nil
}

func (f *InventoryFacade) GetSingleProductStockByWarehouse(ctx context.Context, companyId string, storeId string, warehouseId string, storeProductId string) (*models.ModelProductStoreStock, error) {
	// 1. Obtener el producto de la tienda
	productStore, err := f.serviceBox.ProductPerStoreService.GetProductPerStoreByID(ctx, storeProductId)
	if err != nil {
		return nil, err
	}

	// 2. Obtener stock del producto en la bodega específica
	warehouseStock, err := f.serviceBox.InventoryService.GetSingleProductStock(ctx, storeId, warehouseId, storeProductId)
	if err != nil {
		return nil, err
	}

	// 3. Obtener stock en tránsito del producto
	transitStock, err := f.serviceBox.InventoryService.GetSingleProductTransit(ctx, storeId, warehouseId, storeProductId)
	if err != nil {
		return nil, err
	}

	// 4. Construir el modelo de respuesta
	result := models.ModelProductStoreStock{
		ModelProductPerStore: *productStore,
		Totals: models.ModelProductStoreStockTotals{
			CurrentStock: warehouseStock.CurrentStock,
			StockIn:      0,
			StockOut:     0,
			AvgCost:      warehouseStock.AvgCost,
			TotalCost:    warehouseStock.CurrentStock * warehouseStock.AvgCost,
			AreMinAlert:  warehouseStock.CurrentStock < productStore.Quantities.MinimalStock,
			AreMaxAlert:  warehouseStock.CurrentStock > productStore.Quantities.MaximalStock,
		},
	}

	// 5. Procesar stock en tránsito
	for _, ts := range transitStock {
		switch ts.Direction {
		case "IN":
			result.Totals.StockIn += ts.InStock
		case "OUT":
			result.Totals.StockOut += ts.InStock
		}
	}

	return &result, nil
}

/*
func (f *InventoryFacade) GetAllProductsByCompanyId(ctx context.Context, companyId string) (shared.PaginationResponse[models.ModelProductGeneralStock], error) {
	return shared.PaginationResponse[models.ModelProductGeneralStock]{}, nil
}

func (f *InventoryFacade) GetDetailedProduct(ctx context.Context, companyId string, productId string) (*models.ModelProductDetailStock, error) {
	return &models.ModelProductDetailStock{}, nil
}
*/

/*
func (f *InventoryFacade) GetInventoryByWarehouseId(ctx context.Context, warehouseId string) (shared.PaginationResponse[dto.WarehouseProductsDTO], error) {
	emptyResponse := shared.PaginationResponse[dto.WarehouseProductsDTO]{}

	wProducts, err := f.serviceBox.WarehousePerProductService.GetAllByWarehouse(ctx, warehouseId)
	if err != nil {
		return emptyResponse, err
	}

	var dtoItems []dto.WarehouseProductsDTO

	product, err := f.serviceBox.ProductService.GetAllProducts()
	warehouses, err := f.serviceBox.WarehouseService.GetAll(ctx)
	for _, wp := range wProducts {
		pc, err := f.serviceBox.ProductCompanyService.GetProductCompanyByID(ctx, wp.ProductId)
		if err != nil {
			return emptyResponse, err
		}

		var prod *models.ModelProduct
		for _, p := range product {
			if pc.ProductID == p.ID {
				prod = &p
				break
			}
		}

		var ware *models.ModelWarehouse
		for _, w := range warehouses {
			if wp.WarehouseId == w.ID {
				ware = &w
				break
			}
		}

		dtoItem := f.toDTO(wp, prod, pc, ware)
		dtoItems = append(dtoItems, dtoItem)
	}

	return shared.NewPagination(dtoItems, 1, 1, len(dtoItems)), nil
}
*/

func (f *InventoryFacade) toDTO(model models.ModelProductWarehouse, p *models.ModelProduct, pps *models.ModelProductPerStore, w *models.ModelWarehouse) dto.WarehouseProductsDTO {
	return dto.WarehouseProductsDTO{
		ProductID:    pps.ID,
		ProductName:  p.Name,
		ProductImage: p.Image,
		ProductSku:   p.SKU,

		WarehouseID:   w.ID,
		WarehouseName: w.WarehouseName,

		InStock:    model.InStock,
		BaseUnit:   pps.UnitInventory.Abbreviation,
		BaseUnitId: pps.UnitInventory.ID,
	}
}
