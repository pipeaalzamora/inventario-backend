package data

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

type StoreRepo struct {
	db            *sqlx.DB
	powerRepo     ports.PortPower
	warehouseRepo ports.PortWarehouse
}

func NewStoreRepo(
	db *sqlx.DB,
	powerRepo ports.PortPower,
	warehouseRepo ports.PortWarehouse,
) ports.PortStore {
	return &StoreRepo{
		db:            db,
		powerRepo:     powerRepo,
		warehouseRepo: warehouseRepo,
	}
}

func (r *StoreRepo) GetStores() ([]models.StoreModel, error) {
	var storeEntities []entities.EntityStore

	query := `SELECT * FROM store`
	if err := r.db.Select(&storeEntities, query); err != nil {
		return nil, types.ThrowData("ocurrio un error al obtener las tiendas")
	}

	return r.toModelMap(storeEntities, false), nil
}

func (r *StoreRepo) GetStoreByID(id string) (*models.StoreModel, error) {
	var storeEntity entities.EntityStore
	query := `SELECT * FROM store WHERE id = $1`
	if err := r.db.Get(&storeEntity, query, id); err != nil {
		return nil, types.ThrowData("Ocurrió un error al obtener la tienda")
	}

	//return r.toModel(&storeEntity, nil), nil
	return r.toModel(&storeEntity, true), nil
}

// GetStoreByCompanyID implements ports.PortStore.
func (r *StoreRepo) GetStoreByCompanyID(companyID string) (*models.StoreModel, error) {
	var storeEntity entities.EntityStore
	query := `SELECT * FROM store WHERE company_id = $1`
	if err := r.db.Get(&storeEntity, query, companyID); err != nil {
		return nil, types.ThrowData("Ocurrió un error al obtener la tienda")
	}
	return r.toModel(&storeEntity, false), nil
}

func (r *StoreRepo) CreateStore(
	store *models.StoreModel,
	profiles []models.ProfileAccountModel,
) (*models.StoreModel, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("ocurrio un error al iniciar la creación de la tienda")
	}
	defer tx.Rollback()

	var createdStoreEntity entities.EntityStore

	queryNewStore := `
		INSERT INTO store (
			company_id,
			store_name,
			store_address,
			description,
			id_cost_center,
			created_at,
			updated_at
		) VALUES (
			$1, $2, $3, $4, $5, NOW(), NOW()
		) RETURNING *		
	`
	err = tx.QueryRowx(
		queryNewStore,
		store.CompanyID,
		store.StoreName,
		store.StoreAddress,
		store.Description,
		store.IDCostCenter,
	).StructScan(&createdStoreEntity)

	if err != nil {
		return nil, types.ThrowData("ocurrio un error al crear la tienda")
	}

	ownPowerModel := &models.PowerAccountModel{
		PowerName:   "store:" + createdStoreEntity.ID,
		DisplayName: "Propiedad tienda " + createdStoreEntity.StoreName,
		Description: "Propiedad que otorga permisos exlusivos sobre la tienda " + createdStoreEntity.StoreName,
		CategoryID:  "0e496576-580c-4a9d-8009-1855b28d8f96", //hardcodeado arreglar para usar un entorno TODO:*
	}

	queryPower := `
		INSERT INTO power_accounts
			(power_name, power_display, description, power_account_category_id)
		VALUES
			($1, $2, $3, $4)
		RETURNING id
	`

	err = tx.QueryRowx(queryPower,
		ownPowerModel.PowerName,
		ownPowerModel.DisplayName,
		ownPowerModel.Description,
		ownPowerModel.CategoryID,
	).Scan(&ownPowerModel.ID)
	if err != nil {
		return nil, types.ThrowData("Error al crear el poder propio para la empresa")
	}

	profileIds := []string{}
	for _, profile := range profiles {
		profileIds = append(profileIds, profile.ID)
	}

	err = r.powerRepo.AddOwnPowerToProfileTx(tx, profileIds, ownPowerModel)
	if err != nil {
		return nil, err
	}

	//debemos crear la bodega principal y la de traspaso de la tienda
	//storeModel := r.toModel(&createdStoreEntity, nil)
	storeModel := r.toModel(&createdStoreEntity, true)
	err = r.warehouseRepo.CreateFirstWarehouse(tx, storeModel)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, types.ThrowData("ocurrio un error al finalizar la creación de la tienda")
	}

	return storeModel, nil

}

func (r *StoreRepo) UpdateStore(id string, store *models.StoreModel) (*models.StoreModel, error) {
	query := `UPDATE store SET  store_name=$1, store_address=$2, description=$3, id_cost_center=$4, updated_at=NOW()
			  WHERE id=$5 RETURNING *`

	var storeEntity entities.EntityStore
	err := r.db.QueryRowx(
		query,
		store.StoreName,
		store.StoreAddress,
		store.Description,
		store.IDCostCenter,
		id,
	).StructScan(&storeEntity)
	if err != nil {
		return nil, types.ThrowData("Error al actualizar la tienda")
	}

	//return r.toModel(&storeEntity, nil), nil
	return r.toModel(&storeEntity, false), nil
}

func (r *StoreRepo) DeleteStore(id string) error {
	query := `DELETE FROM store WHERE id = $1`
	if _, err := r.db.Exec(query, id); err != nil {
		return types.ThrowData("Error al eliminar la tienda")
	}
	return nil
}

func (r *StoreRepo) GetStoresByCompanyID(companyID string) ([]models.StoreModel, error) {
	var storeEntities []entities.EntityStore
	query := `SELECT * FROM store WHERE company_id = $1`
	if err := r.db.Select(&storeEntities, query, companyID); err != nil {
		return make([]models.StoreModel, 0), types.ThrowData("ocurrio un error al traer las tiendas")

	}

	return r.toModelMap(storeEntities, false), nil
}

func (r *StoreRepo) toModel(
	entity *entities.EntityStore,
	//supplierApplieds []entities.EntitySupplierApplied,
	withWarehouses bool,
) *models.StoreModel {
	if entity == nil {
		return nil
	}

	_warehouses := make([]models.ModelWarehouse, 0)
	if withWarehouses {
		warehouses, err := r.warehouseRepo.GetWarehousesByStoreId(entity.ID)
		if err == nil {
			_warehouses = warehouses
		}
	}
	//_suppliers := make([]models.StoreSupplierModel, len(supplierApplieds))
	/*
		for i, sa := range supplierApplieds {
			_suppliers[i] = models.StoreSupplierModel{
				SupplierID:   sa.SupplierID,
				SupplierName: sa.SupplierName,
				Available:    sa.Available,
			}
		}
	*/

	return &models.StoreModel{
		ID:           entity.ID,
		CompanyID:    entity.CompanyID,
		StoreName:    entity.StoreName,
		StoreAddress: entity.StoreAddress,
		Description:  entity.Description,
		IDCostCenter: entity.IDCostCenter,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
		WareHouses:   _warehouses,
		//SupplierApplied: _suppliers,
	}
}

func (r *StoreRepo) toModelMap(
	entities []entities.EntityStore,
	withWarehouses bool,
) []models.StoreModel {
	models := make([]models.StoreModel, len(entities))
	for i, entity := range entities {
		//models[i] = *r.toModel(&entity, nil)
		models[i] = *r.toModel(&entity, withWarehouses)
	}
	return models
}

func (r *StoreRepo) UpdateStoreSuppliers(storeID string, supplierIDs []string) error {
	return types.ThrowData("Endpoint deprecado. Use POST /companies/:id/suppliers")
}
