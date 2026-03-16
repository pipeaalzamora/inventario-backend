package data

import (
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/shared"
	"sofia-backend/types"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type SupplierRepo struct {
	db *sqlx.DB
}

func NewSupplierRepo(db *sqlx.DB) ports.PortSupplier {
	return &SupplierRepo{
		db: db,
	}
}

// GetSuppliersByTemplateProductId implements ports.PortSupplier.
func (r *SupplierRepo) GetSuppliersByTemplateProductId(companyId string, templateProductId string) ([]models.ModelSupplier, error) {
	query := `
		SELECT
			supplier.*
		FROM supplier
		JOIN supplier_product ON supplier_product.supplier_id = supplier.id
		JOIN product ON product.id = supplier_product.product_id
		JOIN supplier_per_company ON supplier_per_company.supplier_id = supplier.id
		WHERE 
			supplier_per_company.company_id = $1
			AND product.id = $2
	`
	var supEntities []entities.EntitySupplier
	if err := r.db.Select(&supEntities, query, companyId, templateProductId); err != nil {
		fmt.Println(err)
		return nil, types.ThrowData("Error al obtener los proveedores por producto plantilla")
	}
	return r.toModelMap(supEntities), nil
}

//////////////// SUPPLIER CRUD ///////////////////

func (r *SupplierRepo) GetAllSuppliers() ([]models.ModelSupplier, error) {
	// Implementation for getting all suppliers
	var supEntities []entities.EntitySupplier
	query := `SELECT * FROM supplier`
	if err := r.db.Select(&supEntities, query); err != nil {
		return nil, types.ThrowData("Error al obtener la lista de proveedores")
	}
	return r.toModelMap(supEntities), nil
}

func (r *SupplierRepo) GetSupplierByID(id string) (*models.ModelSupplier, error) {
	// Implementation for getting a supplier by ID
	var supEntity entities.EntitySupplier

	query := `SELECT * FROM supplier WHERE id = $1`
	if err := r.db.QueryRowx(query, id).StructScan(&supEntity); err != nil {
		fmt.Println(err)
		return nil, types.ThrowData("Error al obtener el proveedor")
	}
	return r.toModel(&supEntity, true)
}

func (r *SupplierRepo) GetSuppliersByStoreProductId(storeID string, productIDs []string) ([]models.ModelSupplierStoreProduct, error) {
	var supEntities []entities.EntitySupplierStoreProduct

	query := `
        SELECT 
			spsp.id,
			pps.id AS store_product_id,
			s.id AS supplier_id,
			s.supplier_name,
			pps.product_id,
			pps.product_name,
			spsp.priority,
			spsp.created_at,
			sp.unit_price,
			mu.abbreviation AS purchase_unit
		FROM supplier_per_store_product spsp
		JOIN product_per_store pps ON pps.id = spsp.store_product_id
		JOIN supplier s ON s.id = spsp.supplier_id
		JOIN product p ON p.id = pps.product_id
		JOIN supplier_product sp ON sp.product_id = p.id AND sp.supplier_id = s.id
		JOIN measurement_unit mu ON mu.id = sp.purchase_unit_id
		WHERE pps.store_id = $1 AND pps.product_id = ANY($2)
		ORDER BY spsp.priority ASC;
    `

	if err := r.db.Select(&supEntities, query, storeID, pq.Array(productIDs)); err != nil {
		return nil, err
	}
	return r.toModelMapFromStoreProducts(supEntities), nil
}

func (r *SupplierRepo) GetSupplierEmail(fiscalDataID string) (string, error) {
	var email string
	query := `SELECT email FROM fiscal_data WHERE id = $1`
	if err := r.db.QueryRow(query, fiscalDataID).Scan(&email); err != nil {
		return "", types.ThrowData("Error al obtener el email del proveedor")
	}
	return email, nil
}

func (r *SupplierRepo) ExistsSupplierInCompany(supplierID, companyID string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM supplier_per_company WHERE supplier_id = $1 AND company_id = $2)`
	if err := r.db.QueryRow(query, supplierID, companyID).Scan(&exists); err != nil {
		return false, types.ThrowData("Error al verificar el proveedor en la compañía")
	}
	return exists, nil
}

func (r *SupplierRepo) GetSupplierByFiscalIDAndCountry(fiscalID string, countryID int) (*models.ModelSupplier, error) {
	query := `
		SELECT *
		FROM supplier
		JOIN fiscal_data ON supplier.fiscal_data_id = fiscal_data.id
		WHERE fiscal_data.id_fiscal = $1 AND supplier.country_id = $2
	`
	var supEntity entities.EntitySupplier
	if err := r.db.Get(query, fiscalID, countryID); err != nil {
		return nil, types.ThrowData("Error al obtener información fiscal")
	}
	return r.toModel(&supEntity, false)
}

func (r *SupplierRepo) GetCompanyByFiscalNameAndCountry(fiscalName string, countryID int) (*models.ModelSupplier, error) {
	fiscalName = strings.TrimSpace(fiscalName)
	query := `
		SELECT 
			supplier.*,
			fd.*
		FROM supplier
		JOIN fiscal_data fd ON supplier.fiscal_data_id = fd.id
		WHERE LOWER(fd.fiscal_name) = LOWER($1) AND supplier.country_id = $2
	`

	var supEntity entities.EntitySupplier
	if err := r.db.Get(&supEntity, query, fiscalName, countryID); err != nil {
		return nil, types.ThrowData("Error al obtener el proveedor por nombre fiscal y país")
	}

	return r.toModel(&supEntity, false)
}

func (r *SupplierRepo) CreateSupplier(supplier *models.ModelSupplier) (*models.ModelSupplier, error) {
	// Implementation for creating a supplier

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción")
	}
	defer tx.Rollback()

	queryFiscalData := `
		INSERT INTO fiscal_data 
			(id_fiscal,raw_fiscal_id, fiscal_name, fiscal_address, fiscal_state, fiscal_city, email)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	var fiscalDataID string
	err = tx.QueryRowx(
		queryFiscalData,
		supplier.FiscalData.IDFiscal,
		supplier.FiscalData.RawFiscalID,
		supplier.FiscalData.FiscalName,
		supplier.FiscalData.FiscalAddress,
		supplier.FiscalData.FiscalState,
		supplier.FiscalData.FiscalCity,
		supplier.FiscalData.Email,
	).Scan(&fiscalDataID)

	if err != nil {
		return nil, types.ThrowData("Error al crear los datos fiscales")
	}

	querySupplier := `
		INSERT INTO supplier 
			(fiscal_data_id, country_id, description, supplier_name, available)
		VALUES 
			($1, $2, $3, $4, $5)
		RETURNING *
	`

	var supEntity entities.EntitySupplier
	if err = tx.QueryRowx(
		querySupplier,
		fiscalDataID,
		supplier.CountryID,
		supplier.Description,
		supplier.SupplierName,
		supplier.Available,
	).Scan(
		&supEntity.ID,
		&supEntity.FiscalDataID,
		&supEntity.CountryID,
		&supEntity.Description,
		&supEntity.SupplierName,
		&supEntity.Available,
		&supEntity.CreatedAt,
		&supEntity.UpdatedAt,
	); err != nil {
		return nil, types.ThrowData("Error al crear el proveedor")
	}

	for _, contact := range supplier.Contacts {
		queryContact := `
			INSERT INTO supplier_contact
				(supplier_id, contact_name, description, email, phone)
			VALUES
				($1, $2, $3, $4, $5)
		`
		if _, err := tx.Exec(
			queryContact,
			supEntity.ID,
			contact.Name,
			contact.Description,
			contact.Email,
			contact.Phone,
		); err != nil {
			return nil, types.ThrowData("Error al crear contactos para el proveedor")
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return r.toModel(&supEntity, true)
}

func (r *SupplierRepo) UpdateSupplier(supplier *models.ModelSupplier, ogSupplier *models.ModelSupplier) (*models.ModelSupplier, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción")
	}
	defer tx.Rollback()

	queryFiscalData := `
		UPDATE fiscal_data
		SET 
			fiscal_name = $1,
			fiscal_address = $2,
			fiscal_state = $3,
			fiscal_city = $4,
			email = $5
		WHERE id = $6
		RETURNING *
	`
	var fiscalDataEntity entities.EntityFiscalData
	if err := tx.QueryRowx(queryFiscalData,
		supplier.FiscalData.FiscalName,
		supplier.FiscalData.FiscalAddress,
		supplier.FiscalData.FiscalState,
		supplier.FiscalData.FiscalCity,
		supplier.FiscalData.Email,
		supplier.FiscalData.ID,
	).StructScan(&fiscalDataEntity); err != nil {
		return nil, types.ThrowData("Error al actualizar los datos fiscales")
	}

	querySupplier := `
		UPDATE supplier
		SET 
			supplier_name = $1,
			description = $2,
			available = $3
		WHERE id = $4
		RETURNING *
	`
	var supEntity entities.EntitySupplier
	if err := tx.QueryRowx(querySupplier,
		supplier.SupplierName,
		supplier.Description,
		supplier.Available,
		supplier.ID,
	).StructScan(&supEntity); err != nil {
		return nil, types.ThrowData("Error al actualizar el proveedor")
	}

	//
	newContactsMapId := make(map[string]int)
	for i, contact := range supplier.Contacts {
		hash, _ := shared.HashMapGeneric(map[string]interface{}{
			"name":        contact.Name,
			"description": contact.Description,
			"email":       contact.Email,
			"phone":       contact.Phone,
		})
		newContactsMapId[hash] = i
	}

	oldContactsMapId := make(map[string]string)
	for _, contact := range ogSupplier.Contacts {
		hash, _ := shared.HashMapGeneric(map[string]interface{}{
			"name":        contact.Name,
			"description": contact.Description,
			"email":       contact.Email,
			"phone":       contact.Phone,
		})
		oldContactsMapId[hash] = contact.ID
	}

	newContactHashes := make([]string, 0)
	oldContactHashes := make([]string, 0)

	for hash := range newContactsMapId {
		newContactHashes = append(newContactHashes, hash)
	}

	for hash := range oldContactsMapId {
		oldContactHashes = append(oldContactHashes, hash)
	}

	sliceUtil := shared.NewSliceUtils()

	contactsToAdd := sliceUtil.DifferenceString(newContactHashes, oldContactHashes)
	contactsToRemove := sliceUtil.DifferenceString(oldContactHashes, newContactHashes)

	queryDelete := `
		DELETE FROM supplier_contact
		WHERE id = $1
	`

	for _, hash := range contactsToRemove {
		contactID := oldContactsMapId[hash]
		if _, err := tx.Exec(queryDelete, contactID); err != nil {
			return nil, types.ThrowData("Error al eliminar contactos")
		}
	}

	queryAdd := `
		INSERT INTO supplier_contact
			(supplier_id, contact_name, description, email, phone)
		VALUES
			($1, $2, $3, $4, $5)

	`
	for _, hash := range contactsToAdd {
		i := newContactsMapId[hash]
		contact := supplier.Contacts[i]

		if _, err := tx.Exec(
			queryAdd,
			supplier.ID,
			contact.Name,
			contact.Description,
			contact.Email,
			contact.Phone,
		); err != nil {
			return nil, types.ThrowData("Error al agregar nuevos contactos")
		}

	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return r.toModel(&supEntity, true)

}

func (r *SupplierRepo) DeleteSupplier(id string) error {
	// Implementation for deleting a supplier
	query := `DELETE FROM supplier WHERE id = $1`
	if _, err := r.db.Exec(query, id); err != nil {
		return types.ThrowData("Error al eliminar el proveedor")
	}
	return nil
}

func (r *SupplierRepo) EnableDisableSupplier(id string, available bool) error {
	query := `UPDATE supplier SET available = $1 WHERE id = $2`
	if _, err := r.db.Exec(query, available, id); err != nil {
		return types.ThrowData("Error al actualizar la disponibilidad del proveedor")
	}
	return nil
}

func (r *SupplierRepo) EnableDisableSupplierStore(supplierID, storeID string, available bool) error {
	return types.ThrowData("Endpoint deprecado. Use la asignación de proveedores por empresa")
}

////////////////// SUPPLIER PRODUCTS CRUD ///////////////////

func (r *SupplierRepo) GetSupplierProducts(supplierID string) ([]models.ModelSupplierProduct, error) {
	return r.getSupplierProducts(supplierID)
}

func (r *SupplierRepo) GetSupplierProductById(supplierID, productID string) (*models.ModelSupplierProduct, error) {
	products, err := r.getSupplierProducts(supplierID)
	if err != nil {
		return nil, err
	}

	for _, prod := range products {
		if prod.ID == productID {
			return &prod, nil
		}
	}

	return nil, types.ThrowData("Producto no encontrado para el proveedor dado")
}

func (r *SupplierRepo) GetSupplierProductBySku(supplierID, sku string) (*models.ModelSupplierProduct, error) {
	products, err := r.getSupplierProducts(supplierID)
	if err != nil {
		return nil, err
	}

	for _, prod := range products {
		if prod.SKU == sku {
			return &prod, nil
		}
	}

	return nil, types.ThrowData("Producto no encontrado para el proveedor dado")
}

func (r *SupplierRepo) AddProductToSupplier(supplierID string, product *models.ModelSupplierProduct) (*models.ModelSupplierProduct, error) {
	query := `
		INSERT INTO supplier_product (
			supplier_id,
			product_id,
			product_name,
			description,
			sku,
			unit_price,
			purchase_unit_id,
			available
		)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING *
	`

	var supProdEntity entities.EntitySupplierProduct
	if err := r.db.QueryRowx(
		query,
		supplierID,
		product.ProductID,
		product.Name,
		product.Description,
		product.SKU,
		product.Price,
		product.PurchaseUnit.UnitID,
		product.Available,
	).StructScan(&supProdEntity); err != nil {
		return nil, types.ThrowData("Error al agregar el producto al proveedor")
	}

	return r.toModelSupplierProduct(&supProdEntity), nil
}

func (r *SupplierRepo) UpdateSupplierProductsPrices(supplierID string, products []models.ModelSupplierProduct) ([]models.ModelSupplierProduct, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar transacción de actualización de precios")
	}
	defer tx.Rollback()

	query := `
		UPDATE supplier_product
		SET unit_price = $1
		WHERE supplier_id = $2 AND product_id = $3
		RETURNING *
	`

	updatedProducts := make([]models.ModelSupplierProduct, 0)
	for _, product := range products {
		var supProdEntity entities.EntitySupplierProduct
		if err := tx.QueryRowx(
			query,
			product.Price,
			supplierID,
			product.ProductID,
		).StructScan(&supProdEntity); err != nil {
			return nil, types.ThrowData("Error al actualizar el precio del producto para el proveedor: " + err.Error())
		}
		updatedProduct := r.toModelSupplierProduct(&supProdEntity)
		updatedProducts = append(updatedProducts, *updatedProduct)
	}

	if err = tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción de actualización de precios")
	}

	return updatedProducts, nil

}

func (r *SupplierRepo) UpdateSupplierProduct(supplierID string, product *models.ModelSupplierProduct) (*models.ModelSupplierProduct, error) {
	query := `
		UPDATE supplier_product
		SET 
			product_name = $1,
			description = $2,
			sku = $3,
			unit_price = $4,
			purchase_unit_id = $5,
			available = $6
		WHERE supplier_id = $7 AND product_id = $8
		RETURNING *
	`

	var supProdEntity entities.EntitySupplierProduct
	if err := r.db.QueryRowx(
		query,
		product.Name,
		product.Description,
		product.SKU,
		product.Price,
		product.PurchaseUnit.UnitID,
		product.Available,
		supplierID,
		product.ProductID,
	).StructScan(&supProdEntity); err != nil {
		return nil, types.ThrowData("Error al actualizar el producto del proveedor")
	}

	return r.toModelSupplierProduct(&supProdEntity), nil
}

func (r *SupplierRepo) DeleteSupplierProduct(supplierID, productID string) (*models.ModelSupplierProduct, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción para eliminar el producto del proveedor")
	}
	defer tx.Rollback()

	query := `
		DELETE FROM supplier_product
		WHERE supplier_id = $1 AND id = $2
		RETURNING *
	`
	var supProdEntity entities.EntitySupplierProduct
	if err := tx.QueryRowx(
		query,
		supplierID,
		productID,
	).StructScan(&supProdEntity); err != nil {
		return nil, types.ThrowData("Error al eliminar el producto del proveedor: " + err.Error())
	}

	if err = tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción para eliminar el producto del proveedor")
	}

	return r.toModelSupplierProduct(&supProdEntity), nil
}

//////////////// PRIVATE METHODS ///////////////////

func (r *SupplierRepo) getFiscalData(fiscalDataId string) (*entities.EntityFiscalData, error) {
	var fiscalEntity entities.EntityFiscalData
	query := `
		SELECT 
			id,
			id_fiscal,
			raw_fiscal_id,
			fiscal_name,
			fiscal_address,
			fiscal_state,
			fiscal_city,
			email
		FROM fiscal_data WHERE id = $1
	`
	if err := r.db.QueryRowx(query,
		fiscalDataId,
	).StructScan(
		&fiscalEntity,
	); err != nil {
		return nil, types.ThrowData("Error al traer los datos fiscales")
	}
	return &fiscalEntity, nil
}

func (r *SupplierRepo) toModel(entity *entities.EntitySupplier, withDetails bool) (*models.ModelSupplier, error) {
	fiscalData, err := r.getFiscalData(entity.FiscalDataID)
	if err != nil {
		return nil, err
	}

	contacts := make([]models.ModelSupplierContact, 0)
	products := make([]models.ModelSupplierProduct, 0)
	if withDetails {
		contacts, err = r.getSuplierContacts(entity.ID)
		if err != nil {
			return nil, err
		}

		products, err = r.getSupplierProducts(entity.ID)
		if err != nil {
			return nil, err
		}
	}

	return &models.ModelSupplier{
		ID: entity.ID,
		FiscalData: models.ModelFiscalData{
			ID:            fiscalData.ID,
			IDFiscal:      fiscalData.IDFiscal,
			RawFiscalID:   fiscalData.RawFiscalID,
			FiscalName:    fiscalData.FiscalName,
			FiscalAddress: fiscalData.FiscalAddress,
			FiscalState:   fiscalData.FiscalState,
			FiscalCity:    fiscalData.FiscalCity,
			Email:         fiscalData.Email,
		},
		CountryID:    entity.CountryID,
		SupplierName: entity.SupplierName,
		Description:  entity.Description,
		Available:    entity.Available,
		Contacts:     contacts,
		Products:     products,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}, nil
}

func (r *SupplierRepo) toModelMap(entities []entities.EntitySupplier) []models.ModelSupplier {
	models := make([]models.ModelSupplier, len(entities))
	for i, entity := range entities {
		pModel, _ := r.toModel(&entity, false)
		models[i] = *pModel
	}
	return models
}

func (r *SupplierRepo) toModelFromStoreProduct(entity *entities.EntitySupplierStoreProduct) *models.ModelSupplierStoreProduct {
	// Handle nullable price from LEFT JOIN with supplier_product
	var price float32 = 0
	if entity.UnitPrice != nil {
		price = float32(*entity.UnitPrice)
	}

	var purchaseUnit string = ""
	if entity.PurchaseUnit != nil {
		purchaseUnit = *entity.PurchaseUnit
	}

	return &models.ModelSupplierStoreProduct{
		ID:           entity.ID,
		SupplierID:   entity.SupplierID,
		SupplierName: entity.SupplierName,
		ProductID:    entity.ProductID,
		ProductName:  entity.ProductName,
		Priority:     entity.Priority,
		Price:        price,
		PurchaseUnit: purchaseUnit,
	}
}

func (r *SupplierRepo) toModelMapFromStoreProducts(entities []entities.EntitySupplierStoreProduct) []models.ModelSupplierStoreProduct {
	models := make([]models.ModelSupplierStoreProduct, len(entities))
	for i, entity := range entities {
		models[i] = *r.toModelFromStoreProduct(&entity)
	}
	return models
}

func (r *SupplierRepo) toModelSupplierProduct(entity *entities.EntitySupplierProduct) *models.ModelSupplierProduct {
	return &models.ModelSupplierProduct{
		ID:          entity.ID,
		SupplierID:  entity.SupplierID,
		ProductID:   entity.ProductID,
		Name:        entity.Name,
		Description: entity.Description,
		SKU:         entity.SKU,
		Price:       entity.Price,
		PurchaseUnit: models.ModelProductPurchaseUnit{
			UnitID:  entity.PurchaseUnitID,
			UnitAbv: entity.PurchaseUnit,
		},
		Available: entity.Available,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func (r *SupplierRepo) getSuplierContacts(supplierID string) ([]models.ModelSupplierContact, error) {
	var contactEntities []entities.EntitySupplierContact
	query := `SELECT * FROM supplier_contact WHERE supplier_id = $1`
	if err := r.db.Select(&contactEntities, query, supplierID); err != nil {
		return nil, types.ThrowData("Error al obtener los contactos del proveedor")
	}

	contacts := make([]models.ModelSupplierContact, len(contactEntities))
	for i, entity := range contactEntities {
		contacts[i] = models.ModelSupplierContact{
			ID:          entity.ID,
			Name:        entity.ContactName,
			Description: entity.Description,
			Email:       entity.Email,
			Phone:       entity.Phone,
		}
	}

	return contacts, nil
}

func (r *SupplierRepo) getSupplierProducts(supplierID string) ([]models.ModelSupplierProduct, error) {
	query := `
		SELECT
			sp.id,
			sp.supplier_id,
			sp.product_id,
			sp.product_name,
			sp.description,
			sp.sku,
			sp.unit_price,
			sp.purchase_unit_id,
			mu.abbreviation AS purchase_unit,
			mu.unit_name AS purchase_unit_name,
			sp.available,
			sp.created_at,
			sp.updated_at
		FROM supplier_product sp
		JOIN measurement_unit mu ON sp.purchase_unit_id = mu.id
		WHERE supplier_id = $1
	`

	var supProdEntities []entities.EntitySupplierProduct
	if err := r.db.Select(&supProdEntities, query, supplierID); err != nil {
		return nil, types.ThrowData("Error al obtener los productos del proveedor")
	}

	products := make([]models.ModelSupplierProduct, len(supProdEntities))
	for i, entity := range supProdEntities {
		products[i] = models.ModelSupplierProduct{
			ID:          entity.ID,
			SupplierID:  entity.SupplierID,
			ProductID:   entity.ProductID,
			Name:        entity.Name,
			Description: entity.Description,
			SKU:         entity.SKU,
			Price:       entity.Price,
			Available:   entity.Available,
			PurchaseUnit: models.ModelProductPurchaseUnit{
				UnitID:   entity.PurchaseUnitID,
				UnitAbv:  entity.PurchaseUnit,
				UnitName: entity.PurchaseUnitName,
			},
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		}
	}

	return products, nil
}

// GetSupplierProductsByProductID obtiene todos los supplier_product asociados a un product_id base.
func (r *SupplierRepo) GetSupplierProductsByProductID(productID string) ([]models.ModelSupplierProduct, error) {
	query := `
		SELECT 
			sp.id,
			sp.supplier_id,
			sp.product_id,
			sp.product_name,
			sp.description,
			sp.sku,
			sp.unit_price AS price,
			sp.purchase_unit_id,
			mu.abbreviation AS purchase_unit,
			mu.unit_name AS purchase_unit_name,
			sp.available,
			sp.created_at,
			sp.updated_at
		FROM supplier_product sp
		JOIN measurement_unit mu ON sp.purchase_unit_id = mu.id
		WHERE sp.product_id = $1 AND sp.available = true
	`

	var supProdEntities []entities.EntitySupplierProduct
	if err := r.db.Select(&supProdEntities, query, productID); err != nil {
		return nil, types.ThrowData("Error al obtener los productos del proveedor por producto base")
	}

	products := make([]models.ModelSupplierProduct, len(supProdEntities))
	for i, entity := range supProdEntities {
		products[i] = models.ModelSupplierProduct{
			ID:          entity.ID,
			SupplierID:  entity.SupplierID,
			ProductID:   entity.ProductID,
			Name:        entity.Name,
			Description: entity.Description,
			SKU:         entity.SKU,
			Price:       entity.Price,
			Available:   entity.Available,
			PurchaseUnit: models.ModelProductPurchaseUnit{
				UnitID:   entity.PurchaseUnitID,
				UnitAbv:  entity.PurchaseUnit,
				UnitName: entity.PurchaseUnitName,
			},
			CreatedAt: entity.CreatedAt,
			UpdatedAt: entity.UpdatedAt,
		}
	}

	return products, nil
}

// UpsertSupplierProductPerStore inserta o actualiza los proveedores asignados a un store product.
func (r *SupplierRepo) UpsertSupplierProductPerStore(storeID string, suppliers []models.ModelSupplierStoreProduct) error {
	if len(suppliers) == 0 {
		return nil
	}

	query := `
		INSERT INTO supplier_per_store_product (
			supplier_product_id, store_id, priority, preferred, 
			service_zone, price, min_order_quantity, payment_terms, lead_time_days, available
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (supplier_product_id, store_id) 
		DO UPDATE SET
			priority = EXCLUDED.priority,
			preferred = EXCLUDED.preferred,
			service_zone = EXCLUDED.service_zone,
			price = EXCLUDED.price,
			min_order_quantity = EXCLUDED.min_order_quantity,
			payment_terms = EXCLUDED.payment_terms,
			lead_time_days = EXCLUDED.lead_time_days,
			available = EXCLUDED.available,
			updated_at = NOW()
	`

	for _, sup := range suppliers {
		available := true
		_, err := r.db.Exec(query,
			sup.SupplierProductID,
			storeID,
			sup.Priority,
			sup.Preferred,
			sup.ServiceZone,
			sup.Price,
			sup.MinOrderQuantity,
			sup.PaymentTerms,
			sup.LeadTimeDays,
			available,
		)
		if err != nil {
			return types.ThrowData("Error al upsert proveedor del producto de tienda: " + err.Error())
		}
	}

	return nil
}

// DeleteSupplierProductPerStoreByIDs elimina proveedores asignados que ya no están en el request.
func (r *SupplierRepo) DeleteSupplierProductPerStoreByIDs(storeID string, supplierProductIDs []string) error {
	if len(supplierProductIDs) == 0 {
		return nil
	}

	query := `
		DELETE FROM supplier_per_store_product 
		WHERE store_id = $1 AND supplier_product_id = ANY($2)
	`

	_, err := r.db.Exec(query, storeID, pq.Array(supplierProductIDs))
	if err != nil {
		return types.ThrowData("Error al eliminar proveedores del producto de tienda: " + err.Error())
	}

	return nil
}
