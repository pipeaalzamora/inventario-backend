package data

import (
	"context"
	"database/sql"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type PurchaseRepo struct {
	db *sqlx.DB
}

func NewPurchaseRepo(db *sqlx.DB) ports.PortPurchase {
	return &PurchaseRepo{
		db: db,
	}
}

func (r *PurchaseRepo) CreatePurchaseOrder(purchase *models.ModelPurchase) (*models.ModelPurchase, error) {
	items := make([]entities.EntityPurchaseItem, len(purchase.Items))
	for i, item := range purchase.Items {
		suppliers := make([]entities.EntitySupplierOption, len(item.SupplierOptions))
		for j, sup := range item.SupplierOptions {
			suppliers[j] = entities.EntitySupplierOption{
				SupplierID: sup.SupplierID,
				Price:      sup.Price,
			}
		}
		items[i] = entities.EntityPurchaseItem{
			StoreProductID:  item.StoreProductID,
			Quantity:        item.Quantity,
			UnitPrice:       item.UnitPrice,
			Subtotal:        item.Subtotal,
			Status:          entities.ItemPurchaseStatusPending,
			SupplierOptions: suppliers,
		}
	}
	entity := &entities.EntityPurchase{
		SupplierID:         purchase.SupplierID,
		StoreID:            purchase.StoreID,
		InventoryRequestID: purchase.InventoryRequestID,
		Status:             entities.PurchaseStatusPending,
	}

	// insertar todo en transacción
	tx := r.db.MustBegin()
	defer tx.Rollback()

	err := tx.QueryRowx(
		`INSERT INTO purchase (supplier_id, store_id, inventory_request_id, status) 
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at, display_id`,
		entity.SupplierID, entity.StoreID, entity.InventoryRequestID, entity.Status,
	).Scan(&entity.ID, &entity.CreatedAt, &entity.UpdatedAt, &entity.DisplayID)
	if err != nil {
		return nil, types.ThrowData("Error al insertar la orden de compra")
	}

	for i := range items {
		items[i].PurchaseID = entity.ID
		err = tx.QueryRowx(
			`INSERT INTO purchase_item (
            purchase_id,
            store_product_id,
            quantity,
            unit_price,
            status,
            available_suppliers
        ) 
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
			items[i].PurchaseID,
			items[i].StoreProductID,
			items[i].Quantity,
			items[i].UnitPrice,
			items[i].Status,
			items[i].SupplierOptions, // 👈 aquí el json []byte
		).Scan(&items[i].ID)

		if err != nil {
			return nil, types.ThrowData("Error al insertar el ítem de la orden de compra")
		}
	}
	// insertar el historial
	_, err = tx.Exec(
		`INSERT INTO purchase_history (purchase_id, new_status, observation) 
		VALUES ($1, $2, $3)`,
		entity.ID, entity.Status, "Purchase order created",
	)
	if err != nil {
		return nil, types.ThrowData("Error al insertar el historial de la orden de compra")
	}

	err = tx.Commit()
	if err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toModel(entity), nil
}

func (r *PurchaseRepo) AddDeliveryNoteIdAndSetArrivedStatus(purchaseID string, deliveryNoteID string) error {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	// Update the purchase with the delivery note ID and set the status to "arrived"
	_, err := tx.Exec(
		`UPDATE purchase
		SET delivery_purchase_note_id = $1, status = $2
		WHERE id = $3`,
		deliveryNoteID, entities.PurchaseStatusArrived, purchaseID,
	)

	if err != nil {
		return types.ThrowData("Error al actualizar la orden de compra")
	}

	// insertar el historial
	_, err = tx.Exec(
		`INSERT INTO purchase_history (purchase_id, new_status, observation) 
		VALUES ($1, $2, $3)`,
		purchaseID, entities.PurchaseStatusArrived, "Delivery note added",
	)

	if err != nil {
		tx.Rollback()
		return types.ThrowData("Error al insertar el historial de la orden de compra")
	}

	err = tx.Commit()
	if err != nil {
		return types.ThrowData("Error al confirmar la transacción")
	}

	return nil
}

func (r *PurchaseRepo) CreatePurchaseOrderApproved(purchase *models.ModelPurchase) (*models.ModelPurchase, error) {
	items := make([]entities.EntityPurchaseItem, len(purchase.Items))
	for i, item := range purchase.Items {
		suppliers := make([]entities.EntitySupplierOption, len(item.SupplierOptions))
		for j, sup := range item.SupplierOptions {
			suppliers[j] = entities.EntitySupplierOption{
				SupplierID: sup.SupplierID,
				Price:      sup.Price,
			}
		}
		items[i] = entities.EntityPurchaseItem{
			StoreProductID:  item.StoreProductID,
			Quantity:        item.Quantity,
			UnitPrice:       item.UnitPrice,
			Subtotal:        item.Subtotal,
			Status:          entities.ItemPurchaseStatusApproved,
			SupplierOptions: suppliers,
		}
	}
	entity := &entities.EntityPurchase{
		SupplierID:         purchase.SupplierID,
		StoreID:            purchase.StoreID,
		InventoryRequestID: purchase.InventoryRequestID,
		Status:             entities.PurchaseStatusOnDelivery,
	}

	// insertar todo en transacción
	tx := r.db.MustBegin()
	err := tx.QueryRowx(
		`INSERT INTO purchase (supplier_id, store_id, inventory_request_id, status) 
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at, display_id`,
		entity.SupplierID, entity.StoreID, entity.InventoryRequestID, entity.Status,
	).Scan(&entity.ID, &entity.CreatedAt, &entity.UpdatedAt, &entity.DisplayID)
	if err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al insertar la orden de compra")
	}

	for i := range items {
		items[i].PurchaseID = entity.ID
		err = tx.QueryRowx(
			`INSERT INTO purchase_item (
            purchase_id,
            store_product_id,
            quantity,
            unit_price,
            status,
            available_suppliers
        ) 
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
			items[i].PurchaseID,
			items[i].StoreProductID,
			items[i].Quantity,
			items[i].UnitPrice,
			items[i].Status,
			items[i].SupplierOptions, // 👈 aquí el json []byte
		).Scan(&items[i].ID)

		if err != nil {
			tx.Rollback()
			return nil, types.ThrowData("Error al insertar el ítem de la orden de compra")
		}
	}
	// insertar el historial
	_, err = tx.Exec(
		`INSERT INTO purchase_history (purchase_id, new_status, observation) 
		VALUES ($1, $2, $3)`,
		entity.ID, entity.Status, "Purchase order created",
	)
	if err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al insertar el historial de la orden de compra")
	}

	err = tx.Commit()
	if err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toModel(entity), nil
}

func (r *PurchaseRepo) UpdatePurchaseItemsStatus(purchaseID string, items []models.ModelPurchaseItem) error {
	tx := r.db.MustBegin()

	query := `
		UPDATE purchase_item
		SET status = $1
		WHERE id = $2 
	`

	for i := range items {
		_, err := tx.Exec(
			query,
			items[i].Status,
			items[i].ID,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err := tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *PurchaseRepo) UpdatePurchaseState(purchaseID string, newStatus entities.PurchaseStatus, observation string) error {
	ctx := context.Background()

	tx, err := r.db.BeginTxx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return types.ThrowData("Error al iniciar la transacción")
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err = tx.ExecContext(ctx, `
        UPDATE purchase
           SET status = $1::purchase_status
         WHERE id = $2::uuid
    `, newStatus.ToString(), purchaseID); err != nil {
		return types.ThrowData("Error al actualizar el estado de la orden de compra")
	}

	if _, err = tx.ExecContext(ctx, `
        INSERT INTO purchase_history (purchase_id, new_status, observation)
        VALUES ($1::uuid, $2::purchase_status, $3)
    `, purchaseID, newStatus.ToString(), observation); err != nil {
		return types.ThrowData("Error al insertar el historial de la orden de compra")
	}

	if err = tx.Commit(); err != nil {
		return types.ThrowData("Error al confirmar la transacción")
	}
	return nil
}

func (r *PurchaseRepo) AddSonOCToPurchase(parentID string, childID string) error {
	_, err := r.db.Exec(`
		INSERT INTO purchase_hierarchy (parent_purchase_id, child_purchase_id)
		VALUES ($1, $2)
	`, parentID, childID)
	if err != nil {
		return types.ThrowData("Error al actualizar la orden de compra")
	}
	return nil
}

func (r *PurchaseRepo) GetPurchaseByID(id string) (*models.ModelPurchase, error) {
	var entity entities.EntityPurchase
	// Join supplier, store, warehouse to get names
	err := r.db.Get(&entity, `SELECT
			p.*,
			s.company_id,
			s.store_name,
			sup.supplier_name,
			ir.warehouse_id,
			w.warehouse_name,
			w.warehouse_address,
			w.warehouse_phone,
			dpn.id as delivery_purchase_note_id,
			dpn.display_id as delivery_purchase_note_display_id
		FROM purchase p
		JOIN store s ON p.store_id = s.id
		JOIN supplier sup ON p.supplier_id = sup.id
		JOIN inventory_request ir ON p.inventory_request_id = ir.id
		LEFT JOIN delivery_purchase_note dpn ON p.delivery_purchase_note_id = dpn.id
		JOIN warehouse w ON ir.warehouse_id = w.id
		WHERE p.id=$1`, id)
	if err != nil {
		return nil, types.ThrowData("Error al obtener la orden de compra por ID" + err.Error())
	}

	var items []entities.EntityPurchaseItem
	// Join product_per_store to get product_name
	err = r.db.Select(&items, `SELECT
			pi.*,
			pps.product_name,
			pps.product_id,
			mu.abbreviation as purchase_unit
		FROM purchase_item pi
		JOIN product_per_store pps ON pi.store_product_id = pps.id
		JOIN measurement_unit mu ON mu.id = pps.unit_inventory_id
		WHERE pi.purchase_id=$1`, id)
	if err != nil {
		return nil, types.ThrowData("Error al obtener los ítems de la orden de compra")
	}

	var history []entities.EntityPurchaseHistory
	err = r.db.Select(&history, "SELECT * FROM purchase_history WHERE purchase_id=$1 ORDER BY changed_at DESC", id)
	if err != nil {
		return nil, types.ThrowData("Error al obtener el historial de la orden de compra")
	}

	// Search for parent purchase
	var parent entities.EntityPurchaseHierarchy
	err = r.db.Get(&parent, `
		SELECT parent_purchase_id, p.display_id as parent_display_id
		FROM purchase_hierarchy
		JOIN purchase p ON p.id = parent_purchase_id
		WHERE child_purchase_id=$1
	`, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, types.ThrowData("Error al obtener la orden de compra padre")
	}

	if err != nil {
		if err == sql.ErrNoRows {
			// no parent → dejar nil
			entity.ParentPurchaseID = nil
			entity.ParentDisplayID = nil
		} else {
			return nil, types.ThrowData("Error al obtener la orden de compra padre")
		}
	} else {
		// hay parent → asignar punteros
		entity.ParentPurchaseID = &parent.ParentPurchaseID
		entity.ParentDisplayID = &parent.ParentDisplayID
	}

	// Search for child purchases
	var children []entities.EntityPurchaseHierarchy
	err = r.db.Select(&children, `
		SELECT child_purchase_id, p.display_id as child_display_id
		FROM purchase_hierarchy
		JOIN purchase p ON p.id = child_purchase_id
		WHERE parent_purchase_id=$1
	`, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, types.ThrowData("Error al obtener las órdenes de compra hijas")
	}

	childrenModels := r.toModelHierarchyList(children)
	itemsModels := r.toModelItemList(items)
	itemsHistory := r.toModelHistoryList(history)

	model := r.toModel(&entity)

	model.Items = itemsModels
	model.PurchaseHistory = itemsHistory
	model.ChildrenPurchase = childrenModels

	return model, nil
}

func (r *PurchaseRepo) GetAllPurchase(storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelPurchase, int, error) {
	var entities []entities.EntityPurchase
	var total int

	query := `SELECT
			p.*,
			s.store_name,
			ir.warehouse_id,
			w.warehouse_name,
			w.warehouse_address,
			w.warehouse_phone,
			sup.supplier_name,
			ph.parent_purchase_id,
			pp.display_id as parent_display_id,
			dpn.display_id as delivery_purchase_note_display_id
		FROM purchase p
		JOIN store s ON p.store_id = s.id
		LEFT JOIN purchase_hierarchy ph ON p.id = ph.child_purchase_id
		LEFT JOIN purchase pp ON ph.parent_purchase_id = pp.id
		LEFT JOIN delivery_purchase_note dpn ON p.delivery_purchase_note_id = dpn.id
		JOIN inventory_request ir ON p.inventory_request_id = ir.id
		JOIN warehouse w ON ir.warehouse_id = w.id
		JOIN supplier sup ON p.supplier_id = sup.id
		WHERE p.store_id=$1`
	countQuery := "SELECT COUNT(*) FROM purchase WHERE store_id=$1"
	args := []interface{}{storeID}
	argIndex := 2

	// if filter != nil {
	// 		for key, value := range *filter {
	// 			query += " AND " + key + "=$" + string(rune(argIndex))
	// 			countQuery += " AND " + key + "=$" + string(rune(argIndex))
	// 			args = append(args, value)
	// 			argIndex++
	// 		}
	// }
	// }

	err := r.db.Get(&total, countQuery, args...)
	if err != nil {
		return nil, 0, types.ThrowData("Error al contar las órdenes de compra")
	}

	query += " ORDER BY p.created_at DESC LIMIT $" + strconv.Itoa(argIndex) +
		" OFFSET $" + strconv.Itoa(argIndex+1)
	args = append(args, size, (page-1)*size)

	err = r.db.Select(&entities, query, args...)
	if err != nil {
		return nil, 0, types.ThrowData("Error al obtener las órdenes de compra" + err.Error())
	}

	models := make([]models.ModelPurchase, len(entities))
	for i, entity := range entities {
		model := *r.toModel(&entity)
		models[i] = model
	}

	return models, total, nil
}

func (r *PurchaseRepo) GetPurchasesByInventoryRequestID(inventoryRequestID string) ([]models.ModelPurchase, error) {
	var entitiesPurchase []entities.EntityPurchase

	query := `SELECT
			p.*,
			s.store_name,
			ir.warehouse_id,
			w.warehouse_name,
			w.warehouse_address,
			w.warehouse_phone,
			sup.supplier_name,
			ph.parent_purchase_id,
			pp.display_id as parent_display_id,
			dpn.display_id as delivery_purchase_note_display_id
		FROM purchase p
		JOIN store s ON p.store_id = s.id
		LEFT JOIN purchase_hierarchy ph ON p.id = ph.child_purchase_id
		LEFT JOIN purchase pp ON ph.parent_purchase_id = pp.id
		LEFT JOIN delivery_purchase_note dpn ON p.delivery_purchase_note_id = dpn.id
		JOIN inventory_request ir ON p.inventory_request_id = ir.id
		JOIN warehouse w ON ir.warehouse_id = w.id
		JOIN supplier sup ON p.supplier_id = sup.id
		WHERE p.inventory_request_id=$1
	`

	args := []interface{}{inventoryRequestID}
	err := r.db.Select(&entitiesPurchase, query, args...)
	if err != nil {
		return nil, types.ThrowData("Error al obtener las órdenes de compra por ID de solicitud de inventario")
	}

	models := make([]models.ModelPurchase, len(entitiesPurchase))
	for i, entity := range entitiesPurchase {
		model := *r.toModel(&entity)

		// Obtener historial de la compra
		var history []entities.EntityPurchaseHistory
		err = r.db.Select(&history, "SELECT * FROM purchase_history WHERE purchase_id=$1 ORDER BY changed_at DESC", entity.ID)
		if err != nil {
			return nil, types.ThrowData("Error al obtener el historial de la orden de compra")
		}
		model.PurchaseHistory = r.toModelHistoryList(history)

		models[i] = model
	}

	return models, nil
}

func (r *PurchaseRepo) CreatePurchaseOrderWithInventoryRequest(purchase *models.ModelPurchase, request *models.ModelInventoryRequest) (*models.ModelPurchase, error) {
	items := make([]entities.EntityPurchaseItem, len(purchase.Items))
	for i, item := range purchase.Items {
		items[i] = entities.EntityPurchaseItem{
			StoreProductID: item.StoreProductID,
			Quantity:       item.Quantity,
			UnitPrice:      item.UnitPrice,
			Subtotal:       item.Subtotal,
			Status:         entities.ItemPurchaseStatusPending,
		}
	}
	entity := &entities.EntityPurchase{
		SupplierID: purchase.SupplierID,
		StoreID:    purchase.StoreID,
		Status:     entities.PurchaseStatusPending,
	}

	// insertar todo en transacción
	tx := r.db.MustBegin()

	// Insertar el inventory request
	entityRequest := &entities.EntityInventoryRequest{
		StoreID:     request.StoreID,
		WarehouseID: request.WarehouseID,
		Status:      entities.RequestStatus(request.Status),
		RequestType: entities.RequestType(request.RequestType),
		RequesterID: request.RequesterID,
	}

	var created entities.EntityInventoryRequest

	// Insert cabecera con QueryRow + Scan
	err := tx.QueryRow(
		`INSERT INTO inventory_request 
			(store_id, warehouse_id, status, request_type, requester_id, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		 RETURNING id, store_id, warehouse_id, status, request_type, requester_id, created_at, updated_at, display_id`,
		entityRequest.StoreID,
		entityRequest.WarehouseID,
		entityRequest.Status,
		entityRequest.RequestType,
		entityRequest.RequesterID,
	).Scan(
		&created.ID,
		&created.StoreID,
		&created.WarehouseID,
		&created.Status,
		&created.RequestType,
		&created.RequesterID,
		&created.CreatedAt,
		&created.UpdatedAt,
		&created.DisplayID,
	)
	if err != nil {
		return nil, types.ThrowData("Error al insertar la solicitud")
	}

	// Insert ítems
	for _, item := range request.Items {
		_, err := tx.Exec(`
			INSERT INTO inventory_request_item (inventory_request_id, store_product_id, quantity)
			VALUES ($1, $2, $3)`,
			created.ID,
			item.StoreProductID,
			item.Quantity,
		)
		if err != nil {
			return nil, types.ThrowData("Error al insertar el ítem de la solicitud")
		}
	}

	// Insertar estado inicial en historial
	_, err = tx.Exec(`
			INSERT INTO inventory_request_history (inventory_request_id, new_status, changed_at, changed_by)
			VALUES ($1, $2, NOW(), $3)
		`, created.ID, created.Status, created.RequesterID)
	if err != nil {
		return nil, types.ThrowData("Error al insertar el estado inicial")
	}

	var createdPurchase entities.EntityPurchase
	// Insert Purchase
	err = tx.QueryRowx(
		`INSERT INTO purchase (
			supplier_id,
			store_id,
			inventory_request_id,
			status
		) 
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			store_id,
			status,
			created_at,
			updated_at,
			display_id`,
		entity.SupplierID, entity.StoreID, created.ID, entity.Status,
	).Scan(
		&createdPurchase.ID,
		&createdPurchase.StoreID,
		&createdPurchase.Status,
		&createdPurchase.CreatedAt,
		&createdPurchase.UpdatedAt,
		&createdPurchase.DisplayID,
	)
	if err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al crear la orden de compra")
	}

	for i := range items {
		items[i].PurchaseID = createdPurchase.ID
		err := tx.QueryRowx(
			`INSERT INTO purchase_item (purchase_id, store_product_id, quantity, unit_price, status) 
			VALUES ($1, $2, $3, $4, $5) RETURNING id`,
			items[i].PurchaseID, items[i].StoreProductID, items[i].Quantity, items[i].UnitPrice, items[i].Status,
		).Scan(&items[i].ID)
		if err != nil {
			tx.Rollback()
			return nil, types.ThrowData("Error al insertar el ítem de la orden de compra")
		}
	}

	// insertar el historial
	_, err = tx.Exec(
		`INSERT INTO purchase_history (purchase_id, new_status, observation) 
		VALUES ($1, $2, $3)`,
		createdPurchase.ID, createdPurchase.Status, "Purchase order created",
	)
	if err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al insertar el historial de la orden de compra")
	}

	err = tx.Commit()
	if err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toModel(&createdPurchase), nil
}

func (r *PurchaseRepo) GetPurchasesByDeliveryPurchaseNote(deliveryNoteId string) ([]models.ModelPurchase, error) {
	var entities []entities.EntityPurchase
	query := `SELECT
			p.*,
			s.store_name,
			ir.warehouse_id,
			w.warehouse_name,
			w.warehouse_address,
			w.warehouse_phone,
			sup.supplier_name,
			sup.supplier_phone,
			ph.parent_purchase_id,
			pp.display_id as parent_display_id,
			dpn.display_id as delivery_purchase_note_display_id
		FROM purchase p
		JOIN store s ON p.store_id = s.id
		JOIN inventory_request ir ON p.inventory_request_id = ir.id
		JOIN warehouse w ON ir.warehouse_id = w.id
		LEFT JOIN purchase_hierarchy ph ON p.id = ph.child_purchase_id
		LEFT JOIN purchase pp ON ph.parent_purchase_id = pp.id
		LEFT JOIN delivery_purchase_note dpn ON p.delivery_purchase_note_id = dpn.id
		JOIN supplier sup ON p.supplier_id = sup.id
		JOIN delivery_note dn ON dn.purchase_id = p.id
		WHERE dn.id=$1`

	err := r.db.Select(&entities, query, deliveryNoteId)
	if err != nil {
		return nil, types.ThrowData("Error al obtener las órdenes de compra por ID de nota de entrega")
	}

	models := make([]models.ModelPurchase, len(entities))
	for i, entity := range entities {
		models[i] = *r.toModel(&entity)
	}

	return models, nil
}

func (r *PurchaseRepo) CancelPurchase(purchaseID string, observation string) error {
	// Inside a transaction
	tx := r.db.MustBegin()

	// Update purchase status
	res, err := tx.Exec(
		`UPDATE purchase SET status=$1, updated_at=NOW() WHERE id=$2`,
		entities.PurchaseStatusCancelled.ToString(), purchaseID,
	)
	if err != nil {
		tx.Rollback()
		return types.ThrowData("Error al actualizar el estado de la orden de compra")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return types.ThrowData("Error al obtener las filas afectadas")
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return types.ThrowData("No se encontró ninguna orden de compra con el ID proporcionado")
	}

	// Insert into purchase_history
	_, err = tx.Exec(
		`INSERT INTO purchase_history (purchase_id, new_status, observation) 
		VALUES ($1, $2, $3)`,
		purchaseID, entities.PurchaseStatusCancelled.ToString(), observation,
	)
	if err != nil {
		tx.Rollback()
		return types.ThrowData("Error al insertar el historial de la orden de compra")
	}

	err = tx.Commit()
	if err != nil {
		return types.ThrowData("Error al confirmar la transacción")
	}

	return nil
}

func (r *PurchaseRepo) ApprovePurchase(purchaseID string) error {
	// Inside a transaction
	tx := r.db.MustBegin()

	// Update purchase status
	res, err := tx.Exec(
		`UPDATE purchase SET status=$1, updated_at=NOW() WHERE id=$2`,
		entities.PurchaseStatusOnDelivery, purchaseID,
	)

	if err != nil {
		tx.Rollback()
		return types.ThrowData("Error al actualizar el estado de la orden de compra")
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return types.ThrowData("Error al obtener las filas afectadas")
	}

	if rowsAffected == 0 {
		tx.Rollback()
		return types.ThrowData("No se encontró ninguna orden de compra con el ID proporcionado")
	}

	// Insert into purchase_history
	_, err = tx.Exec(
		`INSERT INTO purchase_history (purchase_id, new_status, observation) 
		VALUES ($1, $2, $3)`,
		purchaseID, entities.PurchaseStatusOnDelivery, "Purchase approved",
	)
	if err != nil {
		tx.Rollback()
		return types.ThrowData("Error al insertar el historial de la orden de compra")
	}

	_, err = tx.Exec(
		`UPDATE purchase_item SET status=$1 WHERE purchase_id=$2 AND status=$3`,
		entities.ItemPurchaseStatusApproved, purchaseID, entities.ItemPurchaseStatusPending,
	)
	if err != nil {
		tx.Rollback()
		return types.ThrowData("Error al actualizar el estado de los ítems de la orden de compra")
	}

	err = tx.Commit()
	if err != nil {
		return types.ThrowData("Error al confirmar la transacción")
	}

	return nil
}

func (r *PurchaseRepo) toModel(entity *entities.EntityPurchase) *models.ModelPurchase {
	if entity == nil {
		return nil
	}

	model := &models.ModelPurchase{
		ID:                            entity.ID,
		DisplayID:                     entity.DisplayID,
		SupplierID:                    entity.SupplierID,
		SupplierName:                  entity.SupplierName,
		SupplierPhone:                 entity.SupplierPhone,
		StoreID:                       entity.StoreID,
		CompanyID:                     entity.CompanyID,
		StoreName:                     entity.StoreName,
		InventoryRequestID:            entity.InventoryRequestID,
		WarehouseID:                   entity.WarehouseID,
		WarehouseName:                 entity.WarehouseName,
		WarehouseAddress:              entity.WarehouseAddress,
		WarehousePhone:                entity.WarehousePhone,
		ParentPurchaseID:              entity.ParentPurchaseID,
		ParentPurchaseDisplayID:       entity.ParentDisplayID,
		DeliveryPurchaseNoteDisplayID: entity.DeliveryNoteDisplayID,
		DeliveryPurchaseNoteID:        entity.DeliveryNoteID,
		Status:                        entity.Status,
		CreatedAt:                     entity.CreatedAt,
		UpdatedAt:                     entity.UpdatedAt,
	}

	if entity.ParentPurchaseID != nil && entity.ParentDisplayID != nil {
		model.ParentPurchaseDisplayID = entity.ParentDisplayID
		model.ParentPurchaseID = entity.ParentPurchaseID
	}

	return model
}

func (r *PurchaseRepo) toModelItem(entity *entities.EntityPurchaseItem) *models.ModelPurchaseItem {
	if entity == nil {
		return nil
	}

	model := &models.ModelPurchaseItem{
		ID:             entity.ID,
		PurchaseID:     entity.PurchaseID,
		StoreProductID: entity.StoreProductID,
		ProductName:    entity.ProductName,
		ProductID:      entity.ProductID,
		Quantity:       entity.Quantity,
		PurchaseUnit:   entity.PurchaseUnit,
		UnitPrice:      entity.UnitPrice,
		Subtotal:       entity.Subtotal,
		Status:         entity.Status,
	}

	if entity.SupplierOptions != nil {
		suppliers := make([]models.SupplierOption, len(entity.SupplierOptions))
		for i, sup := range entity.SupplierOptions {
			suppliers[i] = models.SupplierOption{
				SupplierID: sup.SupplierID,
				Price:      sup.Price,
			}
		}

		model.SupplierOptions = suppliers
	}

	return model
}

func (r *PurchaseRepo) toModelHistory(entity *entities.EntityPurchaseHistory) *models.ModelPurchaseHistory {
	if entity == nil {
		return nil
	}

	model := &models.ModelPurchaseHistory{
		ID:          entity.ID,
		PurchaseID:  entity.PurchaseID,
		NewStatus:   entity.NewStatus,
		Observation: entity.Observation,
		ChangedAt:   entity.ChangedAt,
	}

	return model
}

func (r *PurchaseRepo) toModelHierarchyList(entities []entities.EntityPurchaseHierarchy) []models.ModelPurchaseHierarchy {
	modelsPurchaseHierarchy := make([]models.ModelPurchaseHierarchy, len(entities))
	for i, e := range entities {
		model := models.ModelPurchaseHierarchy{
			PurchaseChildID:        e.ChildPurchaseID,
			PurchaseChildDisplayID: e.ChildDisplayID,
		}
		modelsPurchaseHierarchy[i] = model
	}

	return modelsPurchaseHierarchy
}

func (r *PurchaseRepo) toModelItemList(entities []entities.EntityPurchaseItem) []models.ModelPurchaseItem {
	if entities == nil {
		return nil
	}

	modelItems := make([]models.ModelPurchaseItem, len(entities))
	for i, entity := range entities {
		model := r.toModelItem(&entity)
		if model != nil {
			modelItems[i] = *model
		}
	}

	return modelItems
}

func (r *PurchaseRepo) toModelHistoryList(entities []entities.EntityPurchaseHistory) []models.ModelPurchaseHistory {
	if entities == nil {
		return nil
	}

	modelItems := make([]models.ModelPurchaseHistory, len(entities))
	for i, entity := range entities {
		model := r.toModelHistory(&entity)
		if model != nil {
			modelItems[i] = *model
		}
	}

	return modelItems
}
