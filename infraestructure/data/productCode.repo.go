package data

import (
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

type productCodeRepo struct {
	db *sqlx.DB
}

func NewProductCodeRepo(db *sqlx.DB) ports.PortProductCode {
	return &productCodeRepo{
		db: db,
	}
}

// GetAll implements ports.PortProductCode.
func (p *productCodeRepo) GetAll() ([]models.ModelProductCodeKind, error) {
	var entities []entities.EntityCodeKind
	query := "SELECT * FROM code_kind"
	if err := p.db.Select(&entities, query); err != nil {
		return nil, types.ThrowData(fmt.Sprintf("Error al obtener los tipos de código de producto: %v", err))
	}

	return p.toModelMap(entities), nil
}

// GetAllByProductId implements ports.PortProductCode.
func (p *productCodeRepo) GetAllByProductId(productId string) ([]models.ModelProductCode, error) {
	var _entities []entities.EntityProductCodeWithKind
	query := `
		SELECT 
			product_code.kind_id, 
			code_kind.code_name, 
			code_kind.description, 
			product_code.code_value
		FROM product_code
		JOIN code_kind ON code_kind.id = product_code.kind_id
		WHERE product_code.product_id = $1
	`

	if err := p.db.Select(&_entities, query, productId); err != nil {
		return nil, types.ThrowData(fmt.Sprintf("Error al obtener los códigos de producto para el ID %s: %v", productId, err))
	}

	_models := make([]models.ModelProductCode, 0, len(_entities))
	for _, entity := range _entities {
		_models = append(_models, models.ModelProductCode{
			Kind: p.toModel(&entities.EntityCodeKind{
				ID:          entity.ID,
				Name:        entity.Name,
				Description: entity.Description,
			}),
			Value: entity.Value,
		})
	}
	return _models, nil
}

func (p *productCodeRepo) toModelMap(entities []entities.EntityCodeKind) []models.ModelProductCodeKind {
	_models := make([]models.ModelProductCodeKind, 0, len(entities))
	for _, entity := range entities {
		_models = append(_models, p.toModel(&entity))
	}
	return _models
}

func (p *productCodeRepo) toModel(entity *entities.EntityCodeKind) models.ModelProductCodeKind {
	return models.ModelProductCodeKind{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
	}
}
