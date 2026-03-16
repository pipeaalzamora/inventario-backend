package data

import (
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

type WarehouseRepo struct {
	db *sqlx.DB
}

func NewWarehouseRepo(db *sqlx.DB) ports.PortWarehouse {
	return &WarehouseRepo{
		db: db,
	}
}

/*
func (r *WarehouseRepo) GetAll() ([]models.ModelWarehouse, error) {
	query := `
	SELECT
		id,
		store_id,
		warehouse_name,
		description,
		warehouse_address,
		delivery_instructions,
		working_hours,
		working_timezone,
		created_at
	FROM warehouse
	`
	var warehouses []entities.EntityWarehouse
	if err := r.db.Select(&warehouses, query); err != nil {
		return nil, types.ThrowData("error al obtener las bodegas")
	}
	return r.toModelMap(warehouses), nil
}
*/

func (r *WarehouseRepo) GetWarehouseByID(id string) (*models.ModelWarehouse, error) {
	query := `
		SELECT
			*
		FROM warehouse
		WHERE id = $1
	`
	var warehouse entities.EntityWarehouse
	if err := r.db.Get(&warehouse, query, id); err != nil {
		return nil, types.ThrowData("error al obtener la bodega" + err.Error())
	}
	return r.toModel(&warehouse), nil
}

func (r *WarehouseRepo) GetWarehousesByStoreId(id string) ([]models.ModelWarehouse, error) {
	query := `
		SELECT 
			*
		FROM warehouse
		WHERE store_id = $1 AND is_momevent_warehouse = false
	`

	var warehouses []entities.EntityWarehouse
	if err := r.db.Select(&warehouses, query, id); err != nil {
		return nil, types.ThrowData("error al obtener las bodegas de la tienda")
	}

	return r.toModelMap(warehouses), nil
}

func (r *WarehouseRepo) GetTransitionWarehouseByStoreID(storeID string) (*models.ModelWarehouse, error) {
	query := `
		SELECT 
			*
		FROM warehouse
		WHERE store_id = $1 AND is_momevent_warehouse = true
		LIMIT 1
	`

	var warehouse entities.EntityWarehouse
	if err := r.db.Get(&warehouse, query, storeID); err != nil {
		return nil, types.ThrowData("error al obtener la bodega de transición de la tienda" + err.Error())
	}

	return r.toModel(&warehouse), nil
}

func (r *WarehouseRepo) CreateWarehouse(warehouse *models.ModelWarehouse) (*models.ModelWarehouse, error) {
	query := `
		INSERT INTO warehouse (
			store_id,
			warehouse_name,
			description,
			warehouse_address,
			warehouse_phone,
			created_at
		) VALUES (
			$1, $2, $3, $4, $5, NOW()
		) RETURNING *
	`
	var createdWarehouse entities.EntityWarehouse
	err := r.db.QueryRowx(query,
		warehouse.StoreId,
		warehouse.WarehouseName,
		warehouse.Description,
		warehouse.WarehouseAddress,
		warehouse.WarehousePhone,
	).StructScan(&createdWarehouse)

	if err != nil {
		fmt.Println(err)
		return nil, types.ThrowData("error al crear la bodega")
	}

	return r.toModel(&createdWarehouse), nil
}

func (r *WarehouseRepo) UpdateWarehouse(warehouse *models.ModelWarehouse) (*models.ModelWarehouse, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("error al iniciar la transacción")
	}
	defer tx.Rollback()

	query := `
		UPDATE warehouse SET
			store_id = $1,
			warehouse_name = $2,
			description = $3,
			warehouse_address = $4,
			warehouse_phone = $5
		WHERE id = $6
		RETURNING *
	`

	var warehouseEntity entities.EntityWarehouse
	err = tx.QueryRowx(query,
		warehouse.StoreId,
		warehouse.WarehouseName,
		warehouse.Description,
		warehouse.WarehouseAddress,
		warehouse.WarehousePhone,
		warehouse.ID,
	).StructScan(&warehouseEntity)
	if err != nil {
		return nil, types.ThrowData("error al actualizar la bodega")
	}

	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("error al confirmar la transacción")
	}

	return r.toModel(&warehouseEntity), nil
}

func (r *WarehouseRepo) CreateFirstWarehouse(
	tx *sqlx.Tx, store *models.StoreModel,
) error {
	// Insert bodega principal
	queryMain := `
		INSERT INTO warehouse (
			store_id,
			warehouse_name,
			description,
			warehouse_address,
			warehouse_phone,
			is_momevent_warehouse,
			created_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, NOW()
		)
	`
	_, err := tx.Exec(queryMain,
		store.ID,
		"Bodega principal - "+store.StoreName,
		store.Description,
		store.StoreAddress,
		nil, // warehouse_phone NULL para bodega principal
		false,
	)
	if err != nil {
		return types.ThrowData("error al crear la bodega principal")
	}

	// Insert bodega de traspaso
	queryMovement := `
	INSERT INTO warehouse (
		store_id,
		warehouse_name,
		description,
		warehouse_address,
		warehouse_phone,
		is_momevent_warehouse,
		created_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, NOW()
	)
	`
	_, err = tx.Exec(queryMovement,
		store.ID,
		"Bodega de traspaso - "+store.StoreName,
		nil,  // description NULL
		nil,  // warehouse_address NULL
		nil,  // warehouse_phone NULL
		true, // is_momevent_warehouse = true
	)
	if err != nil {
		return types.ThrowData("error al crear la bodega de traspaso")
	}

	return nil
}

func (r *WarehouseRepo) DeleteWarehouse(id string) error {
	query := `DELETE FROM warehouse WHERE id = $1`
	if _, err := r.db.Exec(query, id); err != nil {
		return types.ThrowData("error al eliminar la bodega")
	}
	return nil
}

func (r *WarehouseRepo) toModel(warehouse *entities.EntityWarehouse) *models.ModelWarehouse {

	if warehouse.Description == nil {
		warehouse.Description = new(string)
	}
	if warehouse.Address == nil {
		warehouse.Address = new(string)
	}
	if warehouse.Phone == nil {
		warehouse.Phone = new(string)
	}

	return &models.ModelWarehouse{
		ID:                  warehouse.ID,
		StoreId:             warehouse.StoreId,
		WarehouseName:       warehouse.Name,
		Description:         *warehouse.Description,
		WarehouseAddress:    *warehouse.Address,
		WarehousePhone:      *warehouse.Phone,
		IsMomeventWarehouse: warehouse.IsMomeventWarehouse,
		CreatedAt:           warehouse.CreatedAt,
	}
}

func (r *WarehouseRepo) toModelMap(warehouse []entities.EntityWarehouse) []models.ModelWarehouse {
	var result []models.ModelWarehouse

	for _, w := range warehouse {
		result = append(result, *r.toModel(&w))
	}
	return result
}
