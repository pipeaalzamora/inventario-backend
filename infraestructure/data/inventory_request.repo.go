package data

import (
	"context"
	"encoding/json"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

type InventoryRequestRepo struct {
	db *sqlx.DB
}

func NewInventoryRequestRepo(db *sqlx.DB) ports.PortInventoryRequest {
	return &InventoryRequestRepo{db: db}
}

// Listar solicitudes por local, con filtro de estado + paginación
func (r *InventoryRequestRepo) GetInventoryRequestsByStore(storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelInventoryRequest, int, error) {
	query := `
		SELECT 
			inventory_request.id as id, 
			inventory_request.display_id as display_id,
			inventory_request.store_id as store_id,
			inventory_request.warehouse_id as warehouse_id,
			inventory_request.status as status,
			inventory_request.request_type as request_type,
			inventory_request.requester_id as requester_id,
			inventory_request.created_at as created_at,
			inventory_request.updated_at as updated_at,
			user_accounts.user_name AS requester_name,
			warehouse.warehouse_name AS warehouse_name,
			store.store_name AS store_name
		FROM inventory_request
		LEFT JOIN user_accounts ON user_accounts.id = inventory_request.requester_id
		LEFT JOIN warehouse ON warehouse.id = inventory_request.warehouse_id
		LEFT JOIN store ON store.id = inventory_request.store_id
	`

	queryFilter := " WHERE inventory_request.store_id = $1"
	if filter != nil {
		jsonData, err := json.Marshal(*filter)
		if err != nil {
			return nil, 0, types.ThrowData("Error al serializar los datos del filtro")
		}
		var filter entities.FilterInventoryRequest
		if err := json.Unmarshal(jsonData, &filter); err != nil {
			return nil, 0, types.ThrowData("Error al deserializar los datos del filtro")
		}

		if filter.Status != nil {
			prevStatus := filter.Status.ToString()
			queryFilter += " AND status = " + prevStatus
		}
	}

	query += queryFilter + " ORDER BY created_at DESC LIMIT $2 OFFSET $3"

	var requests []entities.EntityInventoryRequest
	err := r.db.SelectContext(context.TODO(), &requests, query, storeID, size, (page-1)*size)
	if err != nil {

		return nil, 0, types.ThrowData("Error al obtener las solicitudes de inventario")
	}

	var total int
	err = r.db.GetContext(context.TODO(), &total, "SELECT COUNT(*) FROM inventory_request"+queryFilter, storeID)
	if err != nil {
		return nil, 0, types.ThrowData("Error al contar las solicitudes de inventario")
	}

	models := make([]models.ModelInventoryRequest, len(requests))
	for i, req := range requests {
		models[i] = *r.toInventoryRequestModel(&req)
	}

	return models, total, nil
}

// Editar orden de solicitud de compra (cabecera + ítems)
func (r *InventoryRequestRepo) UpdateInventoryRequest(id string, model *models.ModelInventoryRequest, obs *string, userID string) (*models.ModelInventoryRequest, error) {
	tx, err := r.db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	entity := r.fromInventoryRequestModel(model)
	entity.ID = id

	var updated entities.EntityInventoryRequest

	// Actualizar cabecera con QueryRow + Scan
	err = tx.QueryRow(
		`UPDATE inventory_request
		 SET warehouse_id = $1, request_type = $2, status = $3, updated_at = NOW()
		 WHERE id = $4
		 RETURNING id, store_id, warehouse_id, status, request_type, requester_id, created_at, updated_at, display_id`,
		entity.WarehouseID,
		entity.RequestType,
		entity.Status,
		entity.ID,
	).Scan(
		&updated.ID,
		&updated.StoreID,
		&updated.WarehouseID,
		&updated.Status,
		&updated.RequestType,
		&updated.RequesterID,
		&updated.CreatedAt,
		&updated.UpdatedAt,
		&updated.DisplayID,
	)
	if err != nil {
		return nil, types.ThrowData("Error al actualizar la solicitud")
	}

	// Eliminar ítems previos
	if _, err := tx.Exec("DELETE FROM inventory_request_item WHERE inventory_request_id = $1", id); err != nil {
		return nil, types.ThrowData("Error al eliminar los ítems antiguos")
	}

	// Insertar ítems nuevos
	for _, item := range model.Items {
		_, err := tx.Exec(`
			INSERT INTO inventory_request_item (inventory_request_id, store_product_id, quantity)
			VALUES ($1, $2, $3)`,
			id,
			item.StoreProductID,
			item.Quantity,
		)
		if err != nil {
			return nil, types.ThrowData("Error al insertar el ítem de la solicitud")
		}
	}
	// Insertar en historial de ediciones
	if obs != nil {
		_, err = tx.Exec(`
			INSERT INTO inventory_request_history (inventory_request_id, new_status, observation, changed_at, changed_by)
			VALUES ($1, $2, $3, NOW(), $4)
		`, id, updated.Status, *obs, userID)
		if err != nil {
			return nil, types.ThrowData("Error al insertar el historial de estado")
		}
	} else {
		var oldStatus entities.RequestStatus
		err := tx.QueryRow("SELECT status FROM inventory_request WHERE id = $1", id).Scan(&oldStatus)
		if err != nil {
			return nil, types.ThrowData("Error al obtener el estado anterior")
		}

		if oldStatus != updated.Status {
			_, err = tx.Exec(`
				INSERT INTO inventory_request_history (inventory_request_id, new_status, observation, changed_at, changed_by)
				VALUES ($1, $2, $3, NOW(), $4)
			`, id, updated.Status, userID)
			if err != nil {
				return nil, types.ThrowData("Error al insertar el historial de estado")
			}
		}

	}

	// Commit
	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toInventoryRequestModel(&updated), nil
}

// Cambiar de estado y guardar en historial
func (r *InventoryRequestRepo) ChangeInventoryRequestStatus(id string, newStatus entities.RequestStatus, obs *string, userID string) (*models.ModelInventoryRequest, error) {
	tx, err := r.db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var oldStatus entities.RequestStatus
	err = tx.QueryRow("SELECT status FROM inventory_request WHERE id = $1", id).Scan(&oldStatus)
	if err != nil {
		return nil, types.ThrowData("Error al obtener el estado actual")
	}
	newStatusEntity := entities.RequestStatus(newStatus)

	var updated entities.EntityInventoryRequest
	err = tx.QueryRow(`
		UPDATE inventory_request SET
			status = $1,
			updated_at = NOW()
		WHERE id = $2
		RETURNING
			id, store_id, warehouse_id, status, request_type, requester_id, created_at, updated_at, display_id`, newStatusEntity, id,
	).Scan(
		&updated.ID,
		&updated.StoreID,
		&updated.WarehouseID,
		&updated.Status,
		&updated.RequestType,
		&updated.RequesterID,
		&updated.CreatedAt,
		&updated.UpdatedAt,
		&updated.DisplayID,
	)
	if err != nil {
		return nil, types.ThrowData("Error al actualizar el estado de la solicitud")
	}

	// Insertar en historial de estados
	if obs != nil {
		_, err = tx.Exec(`
			INSERT INTO inventory_request_history (inventory_request_id, new_status, observation, changed_at, changed_by)
			VALUES ($1, $2, $3, NOW(), $4)
		`, id, newStatus, *obs, userID)
	} else {
		_, err = tx.Exec(`
			INSERT INTO inventory_request_history (inventory_request_id, new_status, observation, changed_at, changed_by)
			VALUES ($1, $2, NULL, NOW(), $3)
		`, id, newStatus, userID)
	}

	if err != nil {
		return nil, types.ThrowData("Error al insertar el historial de estado")
	}

	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toInventoryRequestModel(&updated), nil
}

// Crear solicitud de compra (con ítems)
func (r *InventoryRequestRepo) CreateInventoryRequest(model *models.ModelInventoryRequest, obs *string) (*models.ModelInventoryRequest, error) {
	tx, err := r.db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	entity := r.fromInventoryRequestModel(model)

	var created entities.EntityInventoryRequest

	// Insert cabecera con QueryRow + Scan
	err = tx.QueryRow(
		`INSERT INTO inventory_request 
			(store_id, warehouse_id, status, request_type, requester_id, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		 RETURNING id, store_id, warehouse_id, status, request_type, requester_id, created_at, updated_at, display_id`,
		entity.StoreID,
		entity.WarehouseID,
		entity.Status,
		entity.RequestType,
		entity.RequesterID,
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
	for _, item := range model.Items {
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
	if obs != nil {
		_, err = tx.Exec(`
			INSERT INTO inventory_request_history (inventory_request_id, new_status, changed_at, observation, changed_by)
			VALUES ($1, $2, NOW(), $3, $4)
		`, created.ID, created.Status, *obs, created.RequesterID)
		if err != nil {
			return nil, types.ThrowData("Error al insertar el estado inicial")
		}
	} else {
		_, err = tx.Exec(`
			INSERT INTO inventory_request_history (inventory_request_id, new_status, changed_at, changed_by)
			VALUES ($1, $2, NOW(), $3)
		`, created.ID, created.Status, created.RequesterID)
		if err != nil {
			return nil, types.ThrowData("Error al insertar el estado inicial")
		}
	}

	// Confirmar la transacción
	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toInventoryRequestModel(&created), nil
}

func (r *InventoryRequestRepo) GetInventoryRequestByID(id string) (*models.ModelInventoryRequest, error) {
	query := `
		SELECT
			inventory_request.id,
			inventory_request.display_id,
			store.company_id,
			inventory_request.store_id,
			inventory_request.warehouse_id,
			inventory_request.status,
			inventory_request.request_type,
			inventory_request.requester_id,
			inventory_request.created_at,
			inventory_request.updated_at,
			user_accounts.user_name AS requester_name,
			warehouse.warehouse_name AS warehouse_name,
			store.store_name AS store_name
		FROM inventory_request
		LEFT JOIN user_accounts ON user_accounts.id = inventory_request.requester_id
		LEFT JOIN warehouse ON warehouse.id = inventory_request.warehouse_id
		LEFT JOIN store ON store.id = inventory_request.store_id
		WHERE inventory_request.id = $1
	`
	var entity entities.EntityInventoryRequest
	if err := r.db.Get(&entity, query, id); err != nil {
		return nil, types.ThrowData("Error al obtener la solicitud")
	}

	model := r.toInventoryRequestModel(&entity)

	// Cargar ítems
	itemsQuery := `
		SELECT
			iri.id,
			iri.inventory_request_id,
			pps.id AS store_product_id,
			pps.product_id AS product_id,
			iri.quantity,
			mu.abbreviation AS purchase_unit
		FROM inventory_request_item iri
		JOIN inventory_request ir ON ir.id = iri.inventory_request_id
		JOIN store s ON s.id = ir.store_id
		JOIN product_per_store pps ON pps.id = iri.store_product_id AND pps.store_id = s.id
		JOIN measurement_unit mu ON mu.id = pps.unit_inventory_id
		WHERE iri.inventory_request_id = $1;
	`
	var items []entities.EntityInventoryRequestItem
	if err := r.db.Select(&items, itemsQuery, id); err != nil {
		return nil, types.ThrowData("Error al obtener los ítems de la solicitud")
	}

	for _, item := range items {
		model.Items = append(model.Items, *r.toInventoryRequestItemModel(&item))
	}

	historyQuery := `
		SELECT
			irh.id,
			inventory_request_id,
			new_status,
			observation,
			changed_at,
			changed_by,
			user_accounts.user_name AS changed_by_name
		FROM inventory_request_history irh
		JOIN user_accounts ON user_accounts.id = irh.changed_by
		WHERE inventory_request_id = $1
		ORDER BY changed_at DESC
	`
	var history []entities.EntityInventoryHistoryStatus
	if err := r.db.Select(&history, historyQuery, id); err != nil {
		return nil, types.ThrowData("Error al obtener el historial de la solicitud")
	}
	for _, status := range history {
		model.RequestHistory = append(model.RequestHistory, *r.toInventoryHistoryStatusModel(&status))
	}

	return model, nil
}

func (r *InventoryRequestRepo) ApproveAndUpdateRequest(id string, model *models.ModelInventoryRequest, obs *string, userID string) (*models.ModelInventoryRequest, error) {
	tx, err := r.db.BeginTxx(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	entity := r.fromInventoryRequestModel(model)

	var updateApproveRequest entities.EntityInventoryRequest

	// Insert cabecera con QueryRow + Scan
	err = tx.QueryRow(
		`UPDATE inventory_request
		 SET store_id = $1, warehouse_id = $2, status = $3, request_type = $4, requester_id = $5, updated_at = NOW()
		 WHERE id = $6
		 RETURNING id, store_id, warehouse_id, status, request_type, requester_id, created_at, updated_at, display_id`,
		entity.StoreID,
		entity.WarehouseID,
		entity.Status,
		entity.RequestType,
		entity.RequesterID,
		id,
	).Scan(
		&updateApproveRequest.ID,
		&updateApproveRequest.StoreID,
		&updateApproveRequest.WarehouseID,
		&updateApproveRequest.Status,
		&updateApproveRequest.RequestType,
		&updateApproveRequest.RequesterID,
		&updateApproveRequest.CreatedAt,
		&updateApproveRequest.UpdatedAt,
		&updateApproveRequest.DisplayID,
	)
	if err != nil {
		return nil, types.ThrowData("Error al insertar la solicitud")
	}
	// Eliminar ítems previos
	if _, err := tx.Exec("DELETE FROM inventory_request_item WHERE inventory_request_id = $1", id); err != nil {
		return nil, types.ThrowData("Error al eliminar los ítems antiguos")
	}

	// Insert ítems
	for _, item := range model.Items {
		_, err := tx.Exec(`
			INSERT INTO inventory_request_item (inventory_request_id, store_product_id, quantity)
			VALUES ($1, $2, $3)`,
			updateApproveRequest.ID,
			item.StoreProductID,
			item.Quantity,
		)
		if err != nil {
			return nil, types.ThrowData("Error al insertar el ítem de la solicitud")
		}
	}

	var oldStatus entities.RequestStatus
	err = tx.QueryRow("SELECT status FROM inventory_request WHERE id = $1", id).Scan(&oldStatus)
	if err != nil {
		return nil, types.ThrowData("Error al obtener el estado actual")
	}

	if obs != nil {
		_, err = tx.Exec(`
			INSERT INTO inventory_request_history (inventory_request_id, new_status, changed_at, observation, changed_by)
			VALUES ($1, $2, NOW(), $3, $4)
		`, updateApproveRequest.ID, updateApproveRequest.Status, *obs, userID)
		if err != nil {
			return nil, types.ThrowData("Error al insertar el estado inicial")
		}
	} else {
		_, err = tx.Exec(`
			INSERT INTO inventory_request_history (inventory_request_id, new_status, changed_at, changed_by)
			VALUES ($1, $2, NOW(), $3)
		`, updateApproveRequest.ID, updateApproveRequest.Status, userID)
		if err != nil {
			return nil, types.ThrowData("Error al insertar el estado inicial")
		}
	}

	// Confirmar la transacción
	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toInventoryRequestModel(&updateApproveRequest), nil
}

func (r *InventoryRequestRepo) toInventoryHistoryStatusModel(e *entities.EntityInventoryHistoryStatus) *models.ModelInventoryHistoryStatus {
	if e == nil {
		return nil
	}

	modelHis := &models.ModelInventoryHistoryStatus{
		ID:                 e.ID,
		InventoryRequestID: e.InventoryRequestID,
		NewStatus:          e.NewStatus.ToString(),
		ChangedAt:          e.ChangedAt,
		ChangedByName:      e.ChangedByName,
	}

	if e.Observation != nil {
		modelHis.Observation = e.Observation
	}

	return modelHis
}

func (r *InventoryRequestRepo) toInventoryRequestItemModel(e *entities.EntityInventoryRequestItem) *models.ModelInventoryRequestItem {
	if e == nil {
		return nil
	}
	return &models.ModelInventoryRequestItem{
		ID:                 e.ID,
		InventoryRequestID: e.InventoryRequestID,
		StoreProductID:     e.StoreProductID,
		ProductID:          e.ProductID,
		Quantity:           e.Quantity,
		PurchaseUnit:       e.PurchaseUnit,
	}
}

func (r *InventoryRequestRepo) toInventoryRequestModel(e *entities.EntityInventoryRequest) *models.ModelInventoryRequest {
	if e == nil {
		return nil
	}
	return &models.ModelInventoryRequest{
		ID:            e.ID,
		DisplayID:     e.DisplayID,
		StoreID:       e.StoreID,
		CompanyID:     e.CompanyID,
		WarehouseID:   e.WarehouseID,
		Status:        e.Status,
		RequestType:   e.RequestType,
		RequesterID:   e.RequesterID,
		RequesterName: e.RequesterName, // Solo aparece cuando vienen del repo
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		WarehouseName: e.WarehouseName,
		StoreName:     e.StoreName,
		// Items deben llenarse aparte
		// History debe llenarse aparte
	}
}

// Conversión desde domain model
func (r *InventoryRequestRepo) fromInventoryRequestModel(m *models.ModelInventoryRequest) *entities.EntityInventoryRequest {
	if m == nil {
		return nil
	}
	return &entities.EntityInventoryRequest{
		ID:          m.ID,
		StoreID:     m.StoreID,
		WarehouseID: m.WarehouseID,
		Status:      entities.RequestStatus(m.Status),
		RequestType: entities.RequestType(m.RequestType),
		RequesterID: m.RequesterID,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}
