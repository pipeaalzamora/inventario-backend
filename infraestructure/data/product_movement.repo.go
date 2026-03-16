package data

import (
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProdutMovementRepo struct {
	db *sqlx.DB
}

func NewProductMovementRepo(db *sqlx.DB) ports.PortProductMovement {
	return &ProdutMovementRepo{
		db: db,
	}
}

func (r *ProdutMovementRepo) GetAllProductMovements() ([]models.ModelProductMovement, error) {
	query := `
		SELECT
			id,
			store_product_id,
			observation,
			quantity,
			unit_cost,
			total_cost,
			moved_from,
			moved_to,
			moved_at,
			moved_by,
			movement_doc_type,
			movement_doc_reference,
			movement_type,
			purchase_id,
			inventory_unit,
			stock_before,
			stock_after
		FROM product_movement
	`

	var result []entities.EntityProductMovement
	if err := r.db.Select(&result, query); err != nil {
		return nil, types.ThrowData("Error al obtener los movimientos de productos")
	}
	return r.toModelList(result), nil
}

func (r *ProdutMovementRepo) GetByProductMovementId(movementId string) (*models.ModelProductMovement, error) {
	query := `
		SELECT
			id,
			store_product_id,
			observation,
			quantity,
			unit_cost,
			total_cost,
			moved_from,
			moved_to,
			moved_at,
			moved_by,
			movement_doc_type,
			movement_doc_reference,
			movement_type,
			purchase_id,
			inventory_unit,
			stock_before,
			stock_after
		FROM product_movement
		WHERE id = $1
	`
	var result entities.EntityProductMovement
	if err := r.db.Get(&result, query, movementId); err != nil {
		return nil, types.ThrowData("Error al obtener el movimiento de producto")
	}
	model := r.toModel(result)
	return &model, nil
}

// GetAllProductMovementsByStoreProductID obtiene movimientos por ID de producto por tienda.
// Reemplaza GetAllProductMovementsByProductCompanyID.
func (r *ProdutMovementRepo) GetAllProductMovementsByStoreProductID(storeProductId string) ([]models.ModelProductMovement, error) {
	query := `
		SELECT
			id,
			store_product_id,
			observation,
			quantity,
			unit_cost,
			total_cost,
			moved_from,
			moved_to,
			moved_at,
			moved_by,
			movement_doc_type,
			movement_doc_reference,
			movement_type,
			purchase_id,
			inventory_unit,
			stock_before,
			stock_after
		FROM product_movement
		WHERE store_product_id = $1
	`
	var result []entities.EntityProductMovement
	if err := r.db.Select(&result, query, storeProductId); err != nil {
		return nil, types.ThrowData("Error al obtener los movimientos de productos " + err.Error())
	}
	return r.toModelList(result), nil
}

// GetAllProductMovementsByStoreID obtiene movimientos de productos por tienda.
func (r *ProdutMovementRepo) GetAllProductMovementsByStoreID(storeId string) ([]models.ModelProductMovement, error) {
	query := `
		SELECT
			pm.id,
			pm.store_product_id,
			pm.observation,
			pm.quantity,
			pm.unit_cost,
			pm.total_cost,
			pm.moved_from,
			pm.moved_to,
			pm.moved_at,
			pm.moved_by,
			pm.movement_doc_type,
			pm.movement_doc_reference,
			pm.movement_type,
			pm.purchase_id,
			pm.inventory_unit,
			pm.stock_before,
			pm.stock_after
		FROM product_movement pm
		JOIN product_per_store pps ON pm.store_product_id = pps.id
		WHERE pps.store_id = $1
	`
	var result []entities.EntityProductMovement
	if err := r.db.Select(&result, query, storeId); err != nil {
		return nil, types.ThrowData("Error al obtener los movimientos de productos")
	}
	return r.toModelList(result), nil
}

func (r *ProdutMovementRepo) GetAllProductMovementsByWarehouseIDs(warehouseIDs []string) ([]models.ModelProductMovement, error) {
	query := `
		SELECT
			pm.id,
			pm.store_product_id,
			pm.observation,
			pm.quantity,
			pm.unit_cost,
			pm.total_cost,
			pm.moved_from,
			pm.moved_to,
			pm.moved_at,
			pm.moved_by,
			pm.movement_doc_type,
			pm.movement_doc_reference,
			pm.movement_type,
			pm.purchase_id,
			pm.inventory_unit,
			pm.stock_before,
			pm.stock_after
		FROM product_movement pm
		JOIN warehouse_per_product wpp ON wpp.store_product_id = pm.store_product_id
		WHERE wpp.warehouse_id = ANY($1)
	`
	var result []entities.EntityProductMovement
	if err := r.db.Select(&result, query, pq.Array(warehouseIDs)); err != nil {
		return nil, types.ThrowData("Error al obtener los movimientos de productos")
	}
	return r.toModelList(result), nil
}

// GetAllProductMovementsByDateRange obtiene movimientos donde la bodega es origen o destino.
func (r *ProdutMovementRepo) GetAllProductMovementsByDateRange(warehouseID string) ([]models.ModelProductMovement, error) {
	query := `
        SELECT
            id,
            store_product_id,
            observation,
            quantity,
            unit_cost,
            total_cost,
            moved_from,
            moved_to,
            moved_at,
            moved_by,
            movement_doc_type,
            movement_doc_reference,
            movement_type,
            purchase_id,
            inventory_unit,
            stock_before,
            stock_after
        FROM product_movement
        WHERE (moved_from = $1 OR moved_to = $1)
        ORDER BY moved_at DESC
    `
	var result []entities.EntityProductMovement
	if err := r.db.Select(&result, query, warehouseID); err != nil {
		return nil, types.ThrowData("Error al obtener los movimientos de productos de la bodega: " + err.Error())
	}
	return r.toModelList(result), nil
}

func (r *ProdutMovementRepo) CreateNewSingleMovement(movement models.ModelProductMovement) (*models.ModelProductMovement, error) {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	insertQuery := `
		INSERT INTO product_movement
		(
			store_product_id,
			observation,
			quantity,
			unit_cost,
			total_cost,
			moved_from,
			moved_to,
			moved_at,
			moved_by,
			movement_type,
			movement_doc_type,
			movement_doc_reference,
			purchase_id,
			inventory_unit,
			stock_before,
			stock_after
		) 
		VALUES
		(
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
		)
		RETURNING *
	`

	// Query para reducir stock en la bodega de origen (para salidas)
	reduceStockQuery := `
	UPDATE warehouse_per_product
	SET in_stock = in_stock - $1
	WHERE store_product_id = $2 AND warehouse_id = $3
	`

	// Query para aumentar stock en la bodega de destino (para entradas)
	increaseStockQuery := `
	UPDATE warehouse_per_product
	SET in_stock = in_stock + $1
	WHERE store_product_id = $2 AND warehouse_id = $3
	`

	var newMovement entities.EntityProductMovement
	err := tx.QueryRowx(insertQuery,
		movement.StoreProductID,
		movement.Observation,
		movement.Quantity,
		movement.UnitCost,
		movement.TotalCost,
		movement.MovedFrom,
		movement.MovedTo,
		movement.MovedAt,
		movement.MovedBy,
		movement.MovementType,
		movement.MovementDocType,
		movement.DocumentReference,
		movement.PurchaseID,
		movement.InventoryUnit,
		movement.StockBefore,
		movement.StockAfter,
	).StructScan(&newMovement)

	if err != nil {
		return nil, types.ThrowData("Error al crear el movimiento de producto: " + err.Error())
	}

	// Reducir stock de la bodega de origen si existe (salida/retiro)
	if movement.MovedFrom != nil {
		_, err := tx.Exec(reduceStockQuery, movement.Quantity, movement.StoreProductID, *movement.MovedFrom)
		if err != nil {
			return nil, types.ThrowData("Error al reducir el stock de la bodega de origen: " + err.Error())
		}
	}

	// Aumentar stock en la bodega de destino si existe (entrada)
	if movement.MovedTo != nil {
		_, err := tx.Exec(increaseStockQuery, movement.Quantity, movement.StoreProductID, *movement.MovedTo)
		if err != nil {
			return nil, types.ThrowData("Error al aumentar el stock de la bodega de destino: " + err.Error())
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción: " + err.Error())
	}

	model := r.toModel(newMovement)
	return &model, nil
}

func (r *ProdutMovementRepo) CreateNewMultiplesMovements(movements []models.ModelProductMovement) ([]models.ModelProductMovement, error) {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	insertQuery := `
	INSERT INTO product_movement (
		store_product_id,
		observation,
		quantity,
		unit_cost,
		total_cost,
		moved_from,
		moved_to,
		moved_at,
		moved_by,
		movement_doc_type,
		movement_doc_reference,
		movement_type,
		purchase_id,
		inventory_unit,
		stock_before,
		stock_after
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
	) RETURNING *
	`

	// Query para reducir stock en la bodega de origen (para salidas)
	reduceStockQuery := `
	UPDATE warehouse_per_product
	SET in_stock = in_stock - $1
	WHERE store_product_id = $2 AND warehouse_id = $3
	`

	// Query para aumentar stock en la bodega de destino (para entradas)
	increaseStockQuery := `
	UPDATE warehouse_per_product
	SET in_stock = in_stock + $1
	WHERE store_product_id = $2 AND warehouse_id = $3
	`

	var newMovements []entities.EntityProductMovement
	for _, movement := range movements {
		var newMovement entities.EntityProductMovement
		err := tx.QueryRowx(insertQuery,
			movement.StoreProductID,
			movement.Observation,
			movement.Quantity,
			movement.UnitCost,
			movement.TotalCost,
			movement.MovedFrom,
			movement.MovedTo,
			movement.MovedAt,
			movement.MovedBy,
			movement.MovementDocType,
			movement.DocumentReference,
			movement.MovementType,
			movement.PurchaseID,
			movement.InventoryUnit,
			movement.StockBefore,
			movement.StockAfter,
		).StructScan(&newMovement)
		if err != nil {
			return nil, types.ThrowData("Error al crear el movimiento de producto: " + err.Error())
		}

		// Reducir stock de la bodega de origen si existe (salida/retiro)
		if movement.MovedFrom != nil {
			_, err := tx.Exec(reduceStockQuery, movement.Quantity, movement.StoreProductID, *movement.MovedFrom)
			if err != nil {
				return nil, types.ThrowData("Error al reducir el stock de la bodega de origen: " + err.Error())
			}
		}

		// Aumentar stock en la bodega de destino si existe (entrada)
		if movement.MovedTo != nil {
			_, err := tx.Exec(increaseStockQuery, movement.Quantity, movement.StoreProductID, *movement.MovedTo)
			if err != nil {
				return nil, types.ThrowData("Error al aumentar el stock de la bodega de destino: " + err.Error())
			}
		}

		newMovements = append(newMovements, newMovement)
	}

	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción: " + err.Error())
	}

	return r.toModelList(newMovements), nil
}

func (r *ProdutMovementRepo) MakeProductWithdrawal(movement models.ModelProductMovement) (*models.ModelProductMovement, error) {
	// Implementation pending
	return nil, nil
}

// CreateTransferMovements crea movimientos de transferencia entre bodegas
// Actualiza el stock en ambas bodegas y el costo promedio en la bodega destino
// newAvgCosts: mapa con los nuevos costos promedio por producto (clave: storeProductId|warehouseId)
func (r *ProdutMovementRepo) CreateTransferMovements(movements []models.ModelProductMovement, newAvgCosts map[string]float32) ([]models.ModelProductMovement, error) {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	insertQuery := `
		INSERT INTO product_movement (
			store_product_id,
			observation,
			quantity,
			unit_cost,
			total_cost,
			moved_from,
			moved_to,
			moved_at,
			moved_by,
			movement_type,
			movement_doc_type,
			movement_doc_reference,
			purchase_id,
			inventory_unit,
			stock_before,
			stock_after
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
		) RETURNING id, store_product_id, observation, quantity, unit_cost, total_cost,
		            moved_from, moved_to, moved_at, moved_by, movement_type,
		            movement_doc_type, movement_doc_reference, purchase_id,
		            inventory_unit, stock_before, stock_after
	`

	// Validar stock disponible en bodega origen dentro de la transacción (con lock)
	checkStockQuery := `
		SELECT in_stock
		FROM warehouse_per_product
		WHERE store_product_id = $1 AND warehouse_id = $2
		FOR UPDATE
	`

	reduceStockQuery := `
		UPDATE warehouse_per_product
		SET in_stock = in_stock - $1
		WHERE store_product_id = $2 AND warehouse_id = $3
	`

	// UPSERT nativo para bodega destino: crea si no existe, actualiza si existe
	upsertStockAndCostQuery := `
		INSERT INTO warehouse_per_product (store_product_id, warehouse_id, in_stock, cost_avg)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (store_product_id, warehouse_id)
		DO UPDATE SET 
			in_stock = warehouse_per_product.in_stock + EXCLUDED.in_stock,
			cost_avg = EXCLUDED.cost_avg
	`

	var newMovements []entities.EntityProductMovement
	for _, movement := range movements {
		// Validar stock disponible en bodega origen (dentro de la transacción con lock)
		if movement.MovedFrom != nil {
			var currentStock float32
			err := tx.QueryRow(checkStockQuery, movement.StoreProductID, *movement.MovedFrom).Scan(&currentStock)
			if err != nil {
				return nil, types.ThrowData("Error al verificar stock en bodega origen: " + err.Error())
			}
			if currentStock < movement.Quantity {
				return nil, types.ThrowData(fmt.Sprintf(
					"Stock insuficiente en bodega origen. Disponible: %.2f, Solicitado: %.2f",
					currentStock, movement.Quantity,
				))
			}
		}

		// Validar que el costo promedio exista en el mapa
		var newAvgCost float32 = 0
		if movement.MovedTo != nil {
			key := movement.StoreProductID + "|" + *movement.MovedTo
			cost, exists := newAvgCosts[key]
			if !exists {
				return nil, types.ThrowData(fmt.Sprintf(
					"No se encontró el costo promedio calculado para el producto %s en bodega %s",
					movement.StoreProductID, *movement.MovedTo,
				))
			}
			newAvgCost = cost
		}

		var newMovement entities.EntityProductMovement
		err := tx.QueryRowx(insertQuery,
			movement.StoreProductID,
			movement.Observation,
			movement.Quantity,
			movement.UnitCost,
			movement.TotalCost,
			movement.MovedFrom,
			movement.MovedTo,
			movement.MovedAt,
			movement.MovedBy,
			movement.MovementType,
			movement.MovementDocType,
			movement.DocumentReference,
			movement.PurchaseID,
			movement.InventoryUnit,
			movement.StockBefore,
			movement.StockAfter,
		).StructScan(&newMovement)
		if err != nil {
			return nil, types.ThrowData("Error al crear el movimiento de transferencia: " + err.Error())
		}

		// Reducir stock en bodega origen
		if movement.MovedFrom != nil {
			_, err := tx.Exec(reduceStockQuery, movement.Quantity, movement.StoreProductID, *movement.MovedFrom)
			if err != nil {
				return nil, types.ThrowData("Error al reducir el stock de la bodega origen: " + err.Error())
			}
		}

		// UPSERT nativo: Aumentar stock y actualizar costo promedio en bodega destino
		if movement.MovedTo != nil {
			_, err := tx.Exec(upsertStockAndCostQuery, movement.StoreProductID, *movement.MovedTo, movement.Quantity, newAvgCost)
			if err != nil {
				return nil, types.ThrowData("Error al actualizar el stock y costo promedio de la bodega destino: " + err.Error())
			}
		}

		newMovements = append(newMovements, newMovement)
	}

	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción: " + err.Error())
	}

	return r.toModelList(newMovements), nil
}

func (r *ProdutMovementRepo) toModel(entity entities.EntityProductMovement) models.ModelProductMovement {
	return models.ModelProductMovement{
		ID:                entity.ID,
		StoreProductID:    entity.StoreProductID,
		Observation:       entity.Observation,
		Quantity:          entity.Quantity,
		InventoryUnit:     entity.InventoryUnit,
		UnitCost:          entity.UnitCost,
		TotalCost:         entity.TotalCost,
		MovedFrom:         entity.MovedFrom,
		MovedTo:           entity.MovedTo,
		MovedAt:           entity.MovedAt,
		MovementType:      entity.MovementType,
		MovementDocType:   entity.MovementDocType,
		DocumentReference: entity.DocumentReference,
		MovedBy:           entity.MovedBy,
		PurchaseID:        entity.PurchaseID,
		StockBefore:       entity.StockBefore,
		StockAfter:        entity.StockAfter,
	}
}

func (r *ProdutMovementRepo) toModelList(entities []entities.EntityProductMovement) []models.ModelProductMovement {
	var movements []models.ModelProductMovement

	if len(entities) == 0 {
		return make([]models.ModelProductMovement, 0)
	}

	for _, entity := range entities {
		movements = append(movements, r.toModel(entity))
	}
	return movements
}
