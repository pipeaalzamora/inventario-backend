package data

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type DeliveryPurchaseNoteRepo struct {
	db *sqlx.DB
}

func NewDeliveryPurchaseNoteRepo(db *sqlx.DB) ports.PortDeliveryPurchaseNote {
	return &DeliveryPurchaseNoteRepo{
		db: db,
	}
}

func (r *DeliveryPurchaseNoteRepo) CreateDeliveryPurchaseNote(deliveryPurchaseNote *models.ModelDeliveryPurchaseNote) (*models.ModelDeliveryPurchaseNote, error) {
	entity := &entities.EntityDeliveryPurchaseNote{
		SupplierID:  deliveryPurchaseNote.SupplierID,
		StoreID:     deliveryPurchaseNote.StoreID,
		WarehouseID: deliveryPurchaseNote.WarehouseID,
		PurchaseID:  deliveryPurchaseNote.PurchaseID,
		DueDate:     deliveryPurchaseNote.DueDate,
		Comment:     deliveryPurchaseNote.Comment,
		Total:       deliveryPurchaseNote.Total,
		Status:      deliveryPurchaseNote.Status,
		UserID:      deliveryPurchaseNote.UserID,
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Hubo un error al iniciar la operación")
	}
	defer tx.Rollback()

	var created entities.EntityDeliveryPurchaseNote
	err = tx.QueryRowx(`
		INSERT INTO delivery_purchase_note (
			supplier_id,
			store_id,
			warehouse_id,
			purchase_id,
			due_date,
			comment,
			total,
			note_status,
			user_id,
			folio_invoice,
			folio_guide
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		) RETURNING *
		`,
		entity.SupplierID,
		entity.StoreID,
		entity.WarehouseID,
		entity.PurchaseID,
		entity.DueDate,
		entity.Comment,
		entity.Total,
		entity.Status,
		entity.UserID,
		entity.FolioInvoice,
		entity.FolioGuide,
	).StructScan(&created)

	if err != nil {
		// Capturar violación de constraint único en purchase_id
		if strings.Contains(err.Error(), "unique constraint") || strings.Contains(err.Error(), "duplicate key") {
			return nil, types.ThrowMsg("Esta orden de compra ya tiene una nota de entrega asociada.")
		}
		return nil, types.ThrowData("Error al crear la nota de entrega de compra")
	}

	for _, item := range deliveryPurchaseNote.Items {
		entityItem := &entities.EntityDeliveryPurchaseNoteItem{
			DeliveryPurchaseNoteID: created.ID,
			StoreProductID:         item.StoreProductID,
			Quantity:               item.Quantity,
			Difference:             item.Difference,
			UnitPrice:              item.UnitPrice,
			TaxTotal:               item.TaxTotal,
			Status:                 item.Status,
		}

		_, err := tx.Exec(
			`INSERT INTO delivery_purchase_note_item (
				delivery_purchase_note_id,
				store_product_id,
				quantity,
				difference,
				unit_price,
				tax_total,
				item_status
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
			entityItem.DeliveryPurchaseNoteID,
			entityItem.StoreProductID,
			entityItem.Quantity,
			entityItem.Difference,
			entityItem.UnitPrice,
			entityItem.TaxTotal,
			entityItem.Status,
		)

		if err != nil {
			return nil, types.ThrowData("Error al crear el ítem de la nota de entrega de compra")
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, types.ThrowData("Error al confirmar la creación de la nota de entrega de compra")
	}

	return r.toModel(&created), nil
}

func (r *DeliveryPurchaseNoteRepo) UpdateDeliveryPurchaseNote(id string, deliveryPurchaseNote *models.ModelDeliveryPurchaseNote) (*models.ModelDeliveryPurchaseNote, error) {
	entity := &entities.EntityDeliveryPurchaseNote{
		ID:          id,
		DisplayID:   deliveryPurchaseNote.DisplayID,
		SupplierID:  deliveryPurchaseNote.SupplierID,
		StoreID:     deliveryPurchaseNote.StoreID,
		WarehouseID: deliveryPurchaseNote.WarehouseID,
		PurchaseID:  deliveryPurchaseNote.PurchaseID,
		DueDate:     deliveryPurchaseNote.DueDate,
		Comment:     deliveryPurchaseNote.Comment,
		Status:      deliveryPurchaseNote.Status,
		Total:       deliveryPurchaseNote.Total,
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`
		UPDATE delivery_purchase_note
		SET
			display_id=$1,
			supplier_id=$2,
			store_id=$3,
			warehouse_id=$4,
			due_date=$5,
			purchase_id=$6,
			comment=$7,
			total=$8,
			note_status=$9
		WHERE id=$10`,
		entity.DisplayID,
		entity.SupplierID,
		entity.StoreID,
		entity.WarehouseID,
		entity.DueDate,
		entity.PurchaseID,
		entity.Comment,
		entity.Total,
		entity.Status,
		entity.ID,
	)
	if err != nil {
		return nil, err
	}

	var updated entities.EntityDeliveryPurchaseNote
	err = tx.QueryRowx(`
		SELECT id, display_id, supplier_id, store_id, warehouse_id, purchase_id, due_date, comment, total
		FROM delivery_purchase_note
		WHERE id=$1
		`,
		entity.ID,
	).Scan(
		&updated.ID,
		&updated.DisplayID,
		&updated.SupplierID,
		&updated.StoreID,
		&updated.WarehouseID,
		&updated.PurchaseID,
		&updated.DueDate,
		&updated.Comment,
		&updated.Total,
	)

	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(`DELETE FROM delivery_purchase_note_item WHERE delivery_purchase_note_id=$1`, entity.ID)
	if err != nil {
		return nil, err
	}

	for _, item := range deliveryPurchaseNote.Items {
		entityItem := &entities.EntityDeliveryPurchaseNoteItem{
			DeliveryPurchaseNoteID: entity.ID,
			StoreProductID:         item.StoreProductID,
			Quantity:               item.Quantity,
			Difference:             item.Difference,
			UnitPrice:              item.UnitPrice,
			Subtotal:               item.Subtotal,
			TaxTotal:               item.TaxTotal,
			Status:                 item.Status,
		}

		_, err = tx.NamedExec(`
			INSERT INTO delivery_purchase_note_item (
				delivery_purchase_note_id,
				store_product_id,
				quantity,
				difference,
				unit_price,
				subtotal,
				tax_total,
				item_status
			)
			VALUES (
				:delivery_purchase_note_id,
				:store_product_id,
				:quantity,
				:difference,
				:unit_price,
				:subtotal,
				:tax_total,
				:item_status
			)
		`, entityItem)
		if err != nil {
			return nil, types.ThrowData("Error al crear el ítem de la nota de entrega de compra")
		}
	}

	return r.toModel(&updated), nil
}

func (r *DeliveryPurchaseNoteRepo) GetAllDeliveryPurchaseNotes(storeID string, page int, size int, filter *map[string]interface{}) ([]models.ModelDeliveryPurchaseNote, int, error) {
	var entities []entities.EntityDeliveryPurchaseNote
	var total int

	query := `
		SELECT
			p.id,
			p.display_id,
			p.supplier_id, 
			s.supplier_name as supplier_name, 
			p.folio_invoice as folio_invoice, 
			p.folio_guide as folio_guide,
			p.store_id, 
			st.store_name as store_name,
			p.warehouse_id, 
			w.warehouse_name as warehouse_name, 
			p.due_date, 
			p.user_id, 
			u.user_name as user_name,
			p.comment, 
			p.note_status,
			p.total, 
			p.created_at, 
			p.updated_at
		FROM delivery_purchase_note p
		JOIN supplier s ON p.supplier_id = s.id
		JOIN store st ON p.store_id = st.id
		JOIN warehouse w ON p.warehouse_id = w.id
		JOIN user_accounts u ON p.user_id = u.id
		WHERE p.store_id=$1
	`
	countQuery := "SELECT COUNT(*) FROM delivery_purchase_note WHERE store_id=$1"
	args := []interface{}{storeID}
	argIndex := 2

	err := r.db.Get(&total, countQuery, args...)
	if err != nil {
		return nil, 0, types.ThrowData("Error al contar las compras")
	}

	query += " ORDER BY p.created_at DESC LIMIT $" + strconv.Itoa(argIndex) +
		" OFFSET $" + strconv.Itoa(argIndex+1)
	args = append(args, size, (page-1)*size)

	err = r.db.Select(&entities, query, args...)
	if err != nil {
		return nil, 0, types.ThrowData("Error al obtener las compras")
	}

	models := make([]models.ModelDeliveryPurchaseNote, len(entities))
	for i, entity := range entities {
		model := *r.toModel(&entity)
		models[i] = model
	}

	return models, total, nil
}

func (r *DeliveryPurchaseNoteRepo) GetDetailDeliveryPurchaseNote(id string) (*models.ModelDeliveryPurchaseNote, error) {
	var entity entities.EntityDeliveryPurchaseNote
	err := r.db.Get(&entity, `
		SELECT 
			dpn.id,
			dpn.display_id,
			dpn.supplier_id,
			dpn.folio_invoice,
			dpn.folio_guide,
			dpn.store_id,
			dpn.warehouse_id,
			dpn.purchase_id,
			dpn.due_date,
			dpn.comment,
			dpn.note_status,
			dpn.total,
			dpn.user_id,
			dpn.created_at,
			dpn.updated_at,
			p.display_id as purchase_display_id,
			dpn.folio_invoice as folio_invoice,
			dpn.folio_guide as folio_guide,
			s.supplier_name as supplier_name,
			st.company_id as company_id,
			st.store_name as store_name,
			w.warehouse_name as warehouse_name,
			u.user_name as user_name
		FROM delivery_purchase_note dpn
		JOIN purchase p ON dpn.purchase_id = p.id
		JOIN supplier s ON dpn.supplier_id = s.id
		JOIN store st ON dpn.store_id = st.id
		JOIN warehouse w ON dpn.warehouse_id = w.id
		JOIN user_accounts u ON dpn.user_id = u.id
		WHERE dpn.id=$1`, id)
	if err != nil {
		return nil, err
	}

	var itemEntities []entities.EntityDeliveryPurchaseNoteItem
	err = r.db.Select(&itemEntities, `
		SELECT
			pi.*,
			pps.product_name as product_name,
			mu.abbreviation as purchase_unit
		FROM delivery_purchase_note_item pi
		JOIN product_per_store pps ON pi.store_product_id = pps.id
		JOIN measurement_unit mu ON mu.id = pps.unit_inventory_id
		WHERE delivery_purchase_note_id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}

	var fileEntities []entities.EntityFile
	err = r.db.Select(&fileEntities, `
			SELECT
				f.id as id,
				f.file_type as file_type,
				f.file_url as file_url
			FROM file_per_entity fpe
			JOIN file f ON fpe.file_id = f.id
			WHERE
				fpe.entity_id=$1
			AND
				fpe.entity_name='delivery_purchase_note'
		`,
		id,
	)

	if err != nil {
		return nil, err
	}

	model := *r.toModel(&entity)

	model.Items = make([]models.ModelDeliveryPurchaseNoteItem, len(itemEntities))
	for i, itemEntity := range itemEntities {
		model.Items[i] = *r.toModelItem(&itemEntity)
	}

	model.Files = make([]models.ModelFile, len(fileEntities))
	for i, fileEntity := range fileEntities {
		model.Files[i] = models.ModelFile{
			ID:       fileEntity.ID,
			FileType: fileEntity.FileType,
			FileURL:  fileEntity.FileURL,
		}
	}

	return &model, nil
}

func (r *DeliveryPurchaseNoteRepo) CompleteDeliveryPurchaseNote(id string, status entities.DeliveryPurchaseNoteStatus, invoiceFolio string, invoiceGuide string) error {

	_, err := r.db.Exec(`
		UPDATE delivery_purchase_note
		SET note_status=$1,
		    folio_invoice=$2,
		    folio_guide=$3
		WHERE id=$4`,
		status,
		invoiceFolio,
		invoiceGuide,
		id,
	)
	if err != nil {
		return types.ThrowData("Error al completar el estado de la nota de entrega de compra")
	}

	return nil
}

func (r *DeliveryPurchaseNoteRepo) AddFileToDeliveryPurchaseNote(deliveryPurchaseNoteID string, file *models.ModelFile) error {
	entityFile := &entities.EntityFile{
		FileType: file.FileType,
		FileURL:  file.FileURL,
	}

	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	var fileID string
	err = tx.QueryRowx(`
		INSERT INTO file (
			file_type,
			file_url
		)
		VALUES ($1, $2) RETURNING id
	`, entityFile.FileType, entityFile.FileURL).Scan(&fileID)
	if err != nil {
		return types.ThrowData("Error al insertar el archivo")
	}

	_, err = tx.Exec(`
		INSERT INTO file_per_entity (
			file_id,
			entity_name,
			entity_id
		)
		VALUES ($1, $2, $3)
	`, fileID, "delivery_purchase_note", deliveryPurchaseNoteID)
	if err != nil {
		return types.ThrowData("Error al vincular el archivo a la nota de entrega de compra")
	}

	return nil
}

func (r *DeliveryPurchaseNoteRepo) RemoveFileFromDeliveryPurchaseNote(fileID string) error {
	_, err := r.db.Exec(`DELETE FROM file WHERE id=$1`, fileID)
	if err != nil {
		return types.ThrowData("Error al eliminar el archivo")
	}
	return nil
}

func (r *DeliveryPurchaseNoteRepo) GetDetailDeliveryPurchaseNoteByOC(purchaseID string) ([]models.ModelDeliveryPurchaseNote, error) {
	var entities []entities.EntityDeliveryPurchaseNote
	err := r.db.Select(&entities, `
		SELECT 
			dpn.*
		FROM delivery_purchase_note_per_purchase dpp
		JOIN delivery_purchase_note dpn
			ON dpp.delivery_purchase_note_id = dpn.id
		JOIN supplier s ON dpn.supplier_id = s.id
		JOIN store st ON dpn.store_id = st.id
		JOIN warehouse w ON dpn.warehouse_id = w.id
		WHERE dpp.purchase_id=$1
	`, purchaseID)

	if err != nil {
		return nil, types.ThrowData("Error al obtener las notas de entrega de compra por ID de compra")
	}

	models := make([]models.ModelDeliveryPurchaseNote, len(entities))
	for i, entity := range entities {
		model := *r.toModel(&entity)
		models[i] = model
	}

	return models, nil
}

func (r *DeliveryPurchaseNoteRepo) GetFileByID(fileID string) (*models.ModelFile, error) {
	var entity entities.EntityFile
	err := r.db.Get(&entity, `SELECT id, file_type, file_url FROM file WHERE id=$1`, fileID)
	if err != nil {
		return nil, types.ThrowData("Error al obtener el archivo por ID")
	}

	model := &models.ModelFile{
		ID:       entity.ID,
		FileType: entity.FileType,
		FileURL:  entity.FileURL,
	}

	return model, nil
}

func (r *DeliveryPurchaseNoteRepo) UpdateDeliveryPurchaseNoteStatus(id string, status entities.DeliveryPurchaseNoteStatus) error {
	_, err := r.db.Exec(`
		UPDATE delivery_purchase_note
		SET note_status=$1
		WHERE id=$2`,
		status,
		id,
	)
	if err != nil {
		return types.ThrowData("Error al actualizar el estado de la nota de entrega de compra")
	}
	return nil
}

func (r *DeliveryPurchaseNoteRepo) toModel(entity *entities.EntityDeliveryPurchaseNote) *models.ModelDeliveryPurchaseNote {
	return &models.ModelDeliveryPurchaseNote{
		ID:           entity.ID,
		DisplayID:    entity.DisplayID,
		SupplierID:   entity.SupplierID,
		SupplierName: entity.SupplierName,
		CompanyID:    entity.CompanyID,
		StoreID:      entity.StoreID,
		//CompanyID:         entity.CompanyID,
		FolioInvoice:      entity.FolioInvoice,
		FolioGuide:        entity.FolioGuide,
		StoreName:         entity.StoreName,
		WarehouseID:       entity.WarehouseID,
		WarehouseName:     entity.WarehouseName,
		PurchaseID:        entity.PurchaseID,
		PurchaseDisplayID: entity.PurchaseDisplayID,
		Status:            entity.Status,
		DueDate:           entity.DueDate,
		Comment:           entity.Comment,
		UserID:            entity.UserID,
		UserName:          entity.UserName,
		Total:             entity.Total,
		CreatedAt:         entity.CreatedAt,
		UpdatedAt:         entity.UpdatedAt,
	}
}

func (r *DeliveryPurchaseNoteRepo) toModelItem(entity *entities.EntityDeliveryPurchaseNoteItem) *models.ModelDeliveryPurchaseNoteItem {
	return &models.ModelDeliveryPurchaseNoteItem{
		ID:                     entity.ID,
		DeliveryPurchaseNoteID: entity.DeliveryPurchaseNoteID,
		StoreProductID:         entity.StoreProductID,
		ProductName:            entity.ProductName,
		Quantity:               entity.Quantity,
		PurchaseUnit:           entity.PurchaseUnit,
		Difference:             entity.Difference,
		Status:                 entity.Status,
		UnitPrice:              entity.UnitPrice,
		Subtotal:               entity.Subtotal,
		TaxTotal:               entity.TaxTotal,
	}
}
