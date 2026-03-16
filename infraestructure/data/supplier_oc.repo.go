package data

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

type SupplierOCRepo struct {
	db *sqlx.DB
}

func NewSupplierOCRepo(db *sqlx.DB) ports.PortSupplierOC {
	return &SupplierOCRepo{
		db: db,
	}
}

// GetSupplierOC implements ports.SupplierOCPort.
func (s *SupplierOCRepo) GetSupplierOC(hash string) (*models.ModelSupplierToken, error) {
	query := `
		SELECT 
		*
		FROM supplier_token
		WHERE token_hash = $1
	`
	var supplierOC entities.EntitySupplierOC
	if err := s.db.QueryRowx(query, hash).StructScan(
		&supplierOC,
	); err != nil {
		return nil, types.ThrowData("error al obtener el token del proveedor")
	}

	return s.toModel(&supplierOC), nil
}

func (s *SupplierOCRepo) CreateSupplierOC(supplierOC *models.ModelSupplierToken) (*models.ModelSupplierToken, error) {
	query := `
		INSERT INTO supplier_token (
			purchase_id, token_hash, exp_time, used
		) VALUES ($1, $2, $3, $4)
		RETURNING *
	`
	var createdEntity entities.EntitySupplierOC
	if err := s.db.QueryRow(query,
		supplierOC.PurchaseID,
		supplierOC.TokenHash,
		supplierOC.Exp,
		supplierOC.Used,
	).Scan(
		&createdEntity.ID,
		&createdEntity.PurchaseID,
		&createdEntity.TokenHash,
		&createdEntity.Exp,
		&createdEntity.Used,
	); err != nil {
		return nil, types.ThrowData("error al crear el token del proveedor")
	}

	return s.toModel(&createdEntity), nil
}

// UpdateSupplierOC implements ports.SupplierOCPort.
func (s *SupplierOCRepo) UpdateSupplierOC(supplierOC *models.ModelSupplierToken) (*models.ModelSupplierToken, error) {
	query := `
		UPDATE supplier_token
		SET used = true
		WHERE id = $1
		RETURNING *
	`

	var updatedEntity entities.EntitySupplierOC
	if err := s.db.QueryRow(query,
		supplierOC.ID,
	).Scan(
		&updatedEntity.ID,
		&updatedEntity.PurchaseID,
		&updatedEntity.TokenHash,
		&updatedEntity.Exp,
		&updatedEntity.Used,
	); err != nil {
		return nil, types.ThrowData("error al actualizar el token del proveedor")
	}

	return s.toModel(&updatedEntity), nil

}

func (s *SupplierOCRepo) toModel(entity *entities.EntitySupplierOC) *models.ModelSupplierToken {
	if entity == nil {
		return nil
	}

	return &models.ModelSupplierToken{
		ID:         entity.ID,
		PurchaseID: entity.PurchaseID,
		TokenHash:  entity.TokenHash,
		Exp:        entity.Exp,
		Used:       entity.Used,
	}
}
