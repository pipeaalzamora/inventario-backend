package data

import (
	"encoding/json"
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type InventoryReportRepo struct {
	db *sqlx.DB
}

func NewInventoryCountRepo(db *sqlx.DB) ports.PortInventoryCount {
	return &InventoryReportRepo{db: db}
}

// GetAll implements ports.PortInventoryReport.
func (r *InventoryReportRepo) GetAll() ([]models.ModelInventoryCount, error) {
	query := `
		SELECT ic.*, 
            store.store_name,
            user_created.user_name as created_by_name,
            user_assgnd.user_name as assigned_to_name,
            wh.warehouse_name
        FROM inventory_count ic
        JOIN store ON store.id = ic.store_id
        JOIN user_accounts user_created ON user_created.id = ic.created_by
        LEFT JOIN user_accounts user_assgnd ON user_assgnd.id = ic.assigned_to
        JOIN warehouse wh ON wh.id = ic.warehouse_id
	`
	var reports []entities.EntityInventoryCount

	if err := r.db.Select(&reports, query); err != nil {
		return nil, types.ThrowData("Error al obtener los datos")
	}

	return r.toModelList(reports), nil
}

func (r *InventoryReportRepo) GetByID(id string) (*models.ModelInventoryCount, error) {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	query := `
		SELECT ic.*, 
            store.store_name,
            user_created.user_name as created_by_name,
            user_assgnd.user_name as assigned_to_name,
            wh.warehouse_name
        FROM inventory_count ic
        JOIN store ON store.id = ic.store_id
        JOIN user_accounts user_created ON user_created.id = ic.created_by
        LEFT JOIN user_accounts user_assgnd ON user_assgnd.id = ic.assigned_to
        JOIN warehouse wh ON wh.id = ic.warehouse_id
		WHERE ic.id = $1
	`

	var entity entities.EntityInventoryCount
	if err := tx.Get(&entity, query, id); err != nil {
		return nil, types.ThrowData("Error al el conteo")
	}

	queryItems := `
		SELECT 
			ici.*,
			p.product_name,
			p.sku as product_sku,
			p.image as product_image,
			pps.unit_inventory_id as product_base_unit_id,
			mu.abbreviation as product_base_unit_abv
		FROM inventory_count_item ici
		JOIN product_per_store pps ON pps.id = ici.store_product_id
		JOIN product p ON p.id = pps.product_id
		JOIN measurement_unit mu ON mu.id = pps.unit_inventory_id
		WHERE inventory_count_id = $1
	`
	var items []entities.EntityInventoryCountItem
	if err := tx.Select(&items, queryItems, id); err != nil {

		return nil, types.ThrowData("Error al obtener los datos del conteo")
	}

	if err := tx.Commit(); err != nil {

		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	entity.CountItems = items

	return r.toModel(&entity), nil
}

func (r *InventoryReportRepo) GetCompletedByID(itemsMetadata []models.ModelInventoryCountMetadata) ([]models.ModelInventoryCountItem, error) {
	tx := r.db.MustBegin()

	query := `
		SELECT
			pc.id as product_company_id,
			wpp.warehouse_id,
			p.product_name,
			pc.sku as product_sku,
			p.image as product_image,
			mu.id as product_base_unit_id,
			mu.abbreviation as product_base_unit_abv
		FROM product_company pc
		JOIN product p ON p.id = pc.product_id
		JOIN warehouse_per_product wpp ON wpp.product_company_id = pc.id
		JOIN measurement_unit mu ON mu.id = pc.unit_inventory_id
		where pc.id = $1
	`
	var items []entities.EntityInventoryCountItem
	for _, item := range itemsMetadata {
		var entity entities.EntityInventoryCountItem
		if err := tx.Get(&entity,
			query,
			item.ProductID,
		); err != nil {
			tx.Rollback()
			return nil, types.ThrowData("Error al obtener los datos")
		}
		items = append(items, entity)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toModelItems(items), nil

}

func (r *InventoryReportRepo) GetAllByUserId(userId string) ([]models.ModelInventoryCount, error) {
	query := `
		SELECT ic.*, 
            store.store_name,
            user_created.user_name as created_by_name,
            user_assgnd.user_name as assigned_to_name,
            wh.warehouse_name
        FROM inventory_count ic
        JOIN store ON store.id = ic.store_id
        JOIN user_accounts user_created ON user_created.id = ic.created_by
        LEFT JOIN user_accounts user_assgnd ON user_assgnd.id = ic.assigned_to
        JOIN warehouse wh ON wh.id = ic.warehouse_id
		WHERE assigned_to = $1 OR created_by = $1
	`
	var reports []entities.EntityInventoryCount
	if err := r.db.Select(&reports, query, userId); err != nil {
		return nil, types.ThrowData("Error al obtener los datos")
	}

	return r.toModelList(reports), nil

}

func (r *InventoryReportRepo) GetItemsByInventoryCountID(id string) ([]models.ModelInventoryCountItem, error) {
	query := `
		SELECT *
		FROM inventory_count_item
		WHERE inventory_count_id = $1
	`

	var entities []entities.EntityInventoryCountItem
	if err := r.db.Select(&entities, query, id); err != nil {
		return nil, types.ThrowData("Error al obtener los datos")
	}

	return r.toModelItems(entities), nil
}

func (r *InventoryReportRepo) Create(report *models.ModelInventoryCount) (*models.ModelInventoryCount, error) {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	queryCreate := `
		INSERT INTO inventory_count (
			store_id,
			company_id,
			warehouse_id,
			created_by,
			assigned_to,
			status,
			scheduled_at,
			movement_track_id,
			metadata
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING *
	`

	var createdEntity entities.EntityInventoryCount
	if err := tx.QueryRowx(queryCreate,
		report.StoreID,
		report.CompanyID,
		report.WarehouseID,
		report.CreatedBy,
		report.AssignedTo,
		report.Status,
		report.ScheduledAt,
		report.MovementTrackId,
		r.toMetadataEntity(report.Metadata),
	).StructScan(&createdEntity); err != nil {
		return nil, types.ThrowData("Error al crear el conteo de inventario")
	}

	queryItems := `
		INSERT INTO inventory_count_item 
		(
			inventory_count_id,
			store_product_id,
			warehouse_id,
			scheduled_at
		) VALUES 
		($1, $2, $3, $4)
		RETURNING *
	`

	var createdItems []entities.EntityInventoryCountItem
	for _, item := range report.CountItems {
		var createdItem entities.EntityInventoryCountItem
		if err := tx.QueryRowx(queryItems,
			createdEntity.ID,
			item.ProductID,
			createdEntity.WarehouseID,
			item.ScheduledAt,
		).StructScan(&createdItem); err != nil {
			fmt.Println(err)
			return nil, types.ThrowData("Un item ya ha sido agregado a una lista de conteo en la misma fecha")
		}

		createdItems = append(createdItems, createdItem)
	}

	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.GetByID(createdEntity.ID)
}

func (r *InventoryReportRepo) Update(model *models.ModelInventoryCount) (*models.ModelInventoryCount, error) {

	tx := r.db.MustBegin()
	defer tx.Rollback()

	query := `
		UPDATE inventory_count
		SET
			assigned_to = $1,
			scheduled_at = $2,
			metadata = $3,
			updated_at = NOW()
		WHERE id = $4
		RETURNING *
	`

	var entity entities.EntityInventoryCount
	if err := tx.QueryRowx(query,
		model.AssignedTo,
		model.ScheduledAt,
		r.toMetadataEntity(model.Metadata),
		model.ID,
	).StructScan(&entity); err != nil {
		return nil, types.ThrowData("Error al actualizar el conteo de inventario")
	}

	queryItems := `
		INSERT INTO inventory_count_item 
		(
			inventory_count_id,
			store_product_id,
			warehouse_id,
			scheduled_at
		) VALUES 
		($1, $2, $3, $4)
		ON CONFLICT (store_product_id, warehouse_id, scheduled_at) DO NOTHING
	`

	for _, item := range model.CountItems {
		//insertar cada item sin retornar valor, solo retornar un error en caso de que falle
		if _, err := tx.Exec(queryItems,
			entity.ID,
			item.ProductID,
			entity.WarehouseID,
			item.ScheduledAt,
		); err != nil {
			return nil, types.ThrowData("Error al crear el item del conteo de inventario")
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.GetByID(entity.ID)
}

func (r *InventoryReportRepo) Delete(id string) error {

	tx := r.db.MustBegin()
	defer tx.Rollback()

	query2 := `
		DELETE
		FROM inventory_count_item
		WHERE inventory_count_id = $1
	`
	_, err := tx.Exec(query2, id)
	if err != nil {
		return types.ThrowData("Error al eliminar los items del conteo de inventario")
	}

	query1 := `
		DELETE
		FROM inventory_count
		WHERE id = $1
	`
	_, err = tx.Exec(query1, id)
	if err != nil {
		fmt.Println(err)
		return types.ThrowData("Error al eliminar el conteo de inventario")
	}

	if err := tx.Commit(); err != nil {
		return types.ThrowData("Error al confirmar la transacción")
	}

	return nil
}

func (r *InventoryReportRepo) Commit(model *models.ModelInventoryCount) (*models.ModelInventoryCount, error) {
	tx := r.db.MustBegin()

	queryCommit := `
		UPDATE inventory_count
		SET
			status = $1,
			completed_at = NOW(),
			updated_at = NOW(),
			movement_track_id = $2,
			metadata = $3
		WHERE id = $4
		RETURNING *
	`

	var updatedEntity entities.EntityInventoryCount
	if err := tx.QueryRow(queryCommit,
		model.Status,
		model.MovementTrackId,
		r.toMetadataEntity(model.Metadata),
		model.ID,
	).Scan(
		&updatedEntity.ID,
		&updatedEntity.DisplayID,
		&updatedEntity.StoreID,
		&updatedEntity.CompanyID,
		&updatedEntity.WarehouseID,
		&updatedEntity.CreatedBy,
		&updatedEntity.AssignedTo,
		&updatedEntity.Status,
		&updatedEntity.ScheduledAt,
		&updatedEntity.CompletedAt,
		&updatedEntity.CreatedAt,
		&updatedEntity.UpdatedAt,
		&updatedEntity.MovementTrackId,
		&updatedEntity.MetaData,
	); err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al confirmar el conteo de inventario")
	}

	/* queryDelete := `
		DELETE
		FROM inventory_count_item
		WHERE inventory_count_id = $1
	`

	if _, err := tx.Exec(queryDelete, id); err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al eliminar los ítems existentes del conteo de inventario")
	} */

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.GetByID(updatedEntity.ID)
}

func (r *InventoryReportRepo) ChangeAssigned(id string, newAssignedTo *string) (*models.ModelInventoryCount, error) {

	query := `
		UPDATE inventory_count
		SET
			assigned_to = $1, 
			updated_at = NOW()
		WHERE id = $2
		RETURNING *
	`
	var entity entities.EntityInventoryCount
	if err := r.db.QueryRow(query,
		newAssignedTo, id,
	).Scan(
		&entity.ID,
		&entity.DisplayID,
		&entity.StoreID,
		&entity.CompanyID,
		&entity.WarehouseID,
		&entity.CreatedBy,
		&entity.AssignedTo,
		&entity.Status,
		&entity.ScheduledAt,
		&entity.CompletedAt,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.MovementTrackId,
		&entity.MetaData,
	); err != nil {
		return nil, types.ThrowData("Error al actualizar el usuario asignado")
	}

	return r.GetByID(entity.ID)
}

func (r *InventoryReportRepo) ChangeState(id string, newState string) (*models.ModelInventoryCount, error) {

	query := `
		UPDATE inventory_count
		SET
			status = $1, 
			updated_at = NOW()
		WHERE id = $2
		RETURNING *
	`
	var entity entities.EntityInventoryCount
	if err := r.db.QueryRowx(query,
		newState, id,
	).StructScan(&entity); err != nil {
		return nil, types.ThrowData("Error al actualizar el estado del conteo de inventario")
	}

	return r.GetByID(entity.ID)
}

func (r *InventoryReportRepo) ChangeDate(id string, newDate time.Time) (*models.ModelInventoryCount, error) {
	query := `
		UPDATE inventory_count
		SET
			scheduled_at = $1, 
			updated_at = NOW()
		WHERE id = $2
		RETURNING *
	`

	var entity entities.EntityInventoryCount
	if err := r.db.QueryRow(query,
		newDate, id,
	).Scan(
		&entity.ID,
		&entity.DisplayID,
		&entity.StoreID,
		&entity.CompanyID,
		&entity.WarehouseID,
		&entity.CreatedBy,
		&entity.AssignedTo,
		&entity.Status,
		&entity.ScheduledAt,
		&entity.CompletedAt,
		&entity.CreatedAt,
		&entity.UpdatedAt,
		&entity.MovementTrackId,
		&entity.MetaData,
	); err != nil {
		return nil, types.ThrowData("Error al actualizar la fecha del conteo de inventario")
	}

	return r.GetByID(entity.ID)
}

func (r *InventoryReportRepo) Reject(model *models.ModelInventoryCount) (*models.ModelInventoryCount, error) {

	tx := r.db.MustBegin()

	queryStatus := `
		UPDATE inventory_count
		SET
			status = 'rejected',
			updated_at = NOW(),
			metadata = $1
		WHERE id = $2
		RETURNING *
	`

	var updatedEntity entities.EntityInventoryCount
	if err := tx.QueryRow(queryStatus,
		r.toMetadataEntity(model.Metadata),
		model.ID,
	).Scan(
		&updatedEntity.ID,
		&updatedEntity.DisplayID,
		&updatedEntity.StoreID,
		&updatedEntity.CompanyID,
		&updatedEntity.WarehouseID,
		&updatedEntity.CreatedBy,
		&updatedEntity.AssignedTo,
		&updatedEntity.Status,
		&updatedEntity.ScheduledAt,
		&updatedEntity.CompletedAt,
		&updatedEntity.CreatedAt,
		&updatedEntity.UpdatedAt,
		&updatedEntity.MovementTrackId,
		&updatedEntity.MetaData,
	); err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al rechazar el conteo de inventario")
	}

	/* queryItems := `
		INSERT INTO inventory_count_item
		(
			inventory_count_id,
			product_company_id,
			warehouse_id,
			scheduled_at
		) VALUES
		($1, $2, $3, $4)
		RETURNING *
	`

	for _, item := range model.CountItems {
		var createdItem entities.EntityInventoryCountItem
		if err := tx.QueryRow(queryItems,
			updatedEntity.ID,
			item.ProductID,
			updatedEntity.WarehouseID,
			updatedEntity.ScheduledAt,
		).Scan(
			&createdItem.ID,
			&createdItem.InventoryCountID,
			&createdItem.ProductID,
			&createdItem.WarehouseID,
			&createdItem.ScheduledAt,
		); err != nil {
			tx.Rollback()
			return nil, types.ThrowData("Error al crear el ítem del conteo de inventario: Un ítem ya ha sido agregado a una lista de conteo en la misma fecha")
		}
	} */

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.GetByID(updatedEntity.ID)
}

func (r *InventoryReportRepo) GetIncidenceByProduct(countId string, productId string) (*models.ModelInventoryCountItem, error) {
	query := `
		SELECT 
			ici.*,
			p.product_name,
			p.sku as product_sku,
			p.image as product_image,
			pps.unit_inventory_id as product_base_unit_id,
			mu.abbreviation as product_base_unit_abv
		FROM inventory_count_item ici
		JOIN product_per_store pps ON pps.id = ici.store_product_id
		JOIN product p ON p.id = pps.product_id
		JOIN measurement_unit mu ON mu.id = pps.unit_inventory_id
		WHERE inventory_count_id = $1
		AND store_product_id = $2
		LIMIT 1
	`

	var entity entities.EntityInventoryCountItem
	if err := r.db.Get(&entity, query, countId, productId); err != nil {
		// Si no existe, retornar nil sin error
		return nil, nil
	}

	return r.toModelItem(&entity), nil
}

func (r *InventoryReportRepo) SaveIncidence(countId string, productId string, imageUrl *string, observation *string) error {
	// Construir query dinámicamente según qué campos se actualizan
	var setParts []string
	var args []interface{}
	argIndex := 1

	if imageUrl != nil {
		setParts = append(setParts, fmt.Sprintf("incidence_image_url = $%d", argIndex))
		args = append(args, *imageUrl)
		argIndex++
	}

	if observation != nil {
		setParts = append(setParts, fmt.Sprintf("incidence_observation = $%d", argIndex))
		args = append(args, *observation)
		argIndex++
	}

	if len(setParts) == 0 {
		return types.ThrowData("No hay campos para actualizar")
	}

	query := fmt.Sprintf(`
		UPDATE inventory_count_item
		SET %s
		WHERE inventory_count_id = $%d
		AND store_product_id = $%d
	`, strings.Join(setParts, ", "), argIndex, argIndex+1)

	args = append(args, countId, productId)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return types.ThrowData("Error al guardar la incidencia")
	}

	return nil
}

func (r *InventoryReportRepo) DeleteIncidenceImage(countId string, productId string, observation *string) error {
	// Construir query dinámicamente
	var setParts []string
	var args []interface{}
	argIndex := 1

	// Siempre establecer incidence_image_url a NULL
	setParts = append(setParts, "incidence_image_url = NULL")

	// Si se proporciona observación, actualizarla también
	if observation != nil {
		setParts = append(setParts, fmt.Sprintf("incidence_observation = $%d", argIndex))
		args = append(args, *observation)
		argIndex++
	}

	query := fmt.Sprintf(`
		UPDATE inventory_count_item
		SET %s
		WHERE inventory_count_id = $%d
		AND store_product_id = $%d
	`, strings.Join(setParts, ", "), argIndex, argIndex+1)

	args = append(args, countId, productId)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return types.ThrowData("Error al eliminar la imagen de la incidencia")
	}

	return nil
}

/////////// Mappers ////////////////

func (r *InventoryReportRepo) toModel(entity *entities.EntityInventoryCount) *models.ModelInventoryCount {
	return &models.ModelInventoryCount{
		ID:              entity.ID,
		DisplayID:       entity.DisplayID,
		StoreID:         entity.StoreID,
		StoreName:       entity.StoreName,
		CompanyID:       entity.CompanyID,
		WarehouseID:     entity.WarehouseID,
		WarehouseName:   entity.WarehouseName,
		CreatedBy:       entity.CreatedBy,
		CreatedByName:   entity.CreatedByName,
		AssignedTo:      entity.AssignedTo,
		AssignedToName:  entity.AssignedToName,
		Status:          entity.Status,
		ScheduledAt:     entity.ScheduledAt,
		CompletedAt:     entity.CompletedAt,
		CreatedAt:       entity.CreatedAt,
		UpdatedAt:       entity.UpdatedAt,
		MovementTrackId: entity.MovementTrackId,
		Metadata:        r.toMetadataModel(entity.MetaData),
		CountItems:      r.toModelItems(entity.CountItems),
	}
}

func (r *InventoryReportRepo) toModelList(entitiesList []entities.EntityInventoryCount) []models.ModelInventoryCount {
	modelsList := make([]models.ModelInventoryCount, len(entitiesList))
	if len(entitiesList) == 0 {
		return modelsList
	}

	for i, entity := range entitiesList {
		modelsList[i] = *r.toModel(&entity)
	}
	return modelsList
}

func (r *InventoryReportRepo) toModelItem(entity *entities.EntityInventoryCountItem) *models.ModelInventoryCountItem {
	return &models.ModelInventoryCountItem{
		ProductID:            entity.ProductID,
		ProductName:          entity.ProductName,
		ProductSKU:           entity.ProductSKU,
		ProductImage:         entity.ProductImage,
		UnitId:               entity.ProductBaseUnitID,
		UnitAbv:              entity.ProductBaseUnitAbv,
		ScheduledAt:          entity.ScheduledAt,
		IncidenceImageURL:    entity.IncidenceImageURL,
		IncidenceObservation: entity.IncidenceObservation,
	}
}

func (r *InventoryReportRepo) toModelItems(entitiesList []entities.EntityInventoryCountItem) []models.ModelInventoryCountItem {
	modelsList := make([]models.ModelInventoryCountItem, len(entitiesList))
	if len(entitiesList) == 0 {
		return modelsList
	}

	for i, entity := range entitiesList {
		modelsList[i] = *r.toModelItem(&entity)
	}
	return modelsList
}

func (r *InventoryReportRepo) toMetadataModel(entity map[string][]entities.EntityInventoryCountMetadata) []models.ModelInventoryCountMetadata {
	var metadataModel []models.ModelInventoryCountMetadata
	for productId, metadataEntities := range entity {

		completed := false
		var unitsCount []models.ModelInventoryUnitsCount
		for _, me := range metadataEntities {
			completed = me.Completed

			unitsCount = append(unitsCount, models.ModelInventoryUnitsCount{
				UnitId:  me.UnitId,
				UnitAbv: me.UnitAbv,
				Count:   me.Count,
				Factor:  me.Factor,
			})
		}

		metadataModel = append(metadataModel, models.ModelInventoryCountMetadata{
			ProductID:  productId,
			Completed:  completed,
			UnitsCount: unitsCount,
		})
	}
	return metadataModel
}

func (r *InventoryReportRepo) toMetadataEntity(metadata []models.ModelInventoryCountMetadata) *string {
	metadataEntity := make(map[string][]entities.EntityInventoryCountMetadata, 0)

	for _, product := range metadata {
		var metadata []entities.EntityInventoryCountMetadata
		for _, uc := range product.UnitsCount {
			metadata = append(metadata, entities.EntityInventoryCountMetadata{
				Completed: product.Completed,
				UnitId:    uc.UnitId,
				UnitAbv:   uc.UnitAbv,
				Count:     uc.Count,
				Factor:    uc.Factor,
			})
		}
		metadataEntity[product.ProductID] = metadata

	}

	jsonData, err := json.Marshal(metadataEntity)
	if err != nil {
		return nil
	}

	// to avoid unused variable warning if needed

	jsonString := string(jsonData)
	return &jsonString
}
