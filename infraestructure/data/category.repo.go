package data

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"
	"strings"

	"github.com/jmoiron/sqlx"
)

type categoryRepo struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) ports.PortCategory {
	return &categoryRepo{
		db: db,
	}
}

// GetById implements ports.PortCategory.
func (c *categoryRepo) GetById(id string) (*models.ModelProductCategory, error) {
	var categoryEntity entities.EntityProductCategory

	query := "SELECT * FROM product_category WHERE id = $1"
	if err := c.db.Get(&categoryEntity, query, id); err != nil {
		return nil, err
	}
	return c.toModel(&categoryEntity), nil
}

// GetByName implements ports.PortCategory.
func (c *categoryRepo) GetByName(name string) (*models.ModelProductCategory, error) {
	var categoryEntity entities.EntityProductCategory

	name = strings.TrimSpace(name)

	query := "SELECT * FROM product_category WHERE LOWER(name) = LOWER($1)"
	if err := c.db.Get(&categoryEntity, query, name); err != nil {
		return nil, err
	}
	return c.toModel(&categoryEntity), nil
}

// GetAll implements ports.PortCategory.
func (c *categoryRepo) GetAll() ([]models.ModelProductCategory, error) {
	var categoryEntities []entities.EntityProductCategory

	query := "SELECT * FROM product_category"
	if err := c.db.Select(&categoryEntities, query); err != nil {
		return nil, err
	}
	return c.toModelMap(categoryEntities), nil
}

// GetAllByProductId implements ports.PortCategory.
func (c *categoryRepo) GetAllByProductId(productId string) ([]models.ModelProductCategory, error) {
	var categoryEntities []entities.EntityProductCategory

	query := `
		SELECT 
			product_category.* 
		FROM product_category
		JOIN product_per_category ON product_per_category.category_id = product_category.id
		WHERE product_per_category.product_id = $1
	`
	if err := c.db.Select(&categoryEntities, query, productId); err != nil {
		return nil, err
	}
	return c.toModelMap(categoryEntities), nil
}

// Create implements ports.PortCategory.
func (c *categoryRepo) Create(category *models.ModelProductCategory) (*models.ModelProductCategory, error) {
	query := `
		INSERT INTO product_category (category_name, description, available)
		VALUES ($1, $2, $3)
		RETURNING *
	`
	var categoryEntity entities.EntityProductCategory
	if err := c.db.QueryRowx(
		query,
		category.Name,
		category.Description,
		category.Available,
	).StructScan(&categoryEntity); err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"product_category_category_name_key\"" {
			return nil, types.ThrowRecipe("Ya existe una categoría con ese nombre", "name")
		}
		return nil, err
	}

	return c.toModel(&categoryEntity), nil
}

// Update implements ports.PortCategory.
func (c *categoryRepo) Update(category *models.ModelProductCategory) (*models.ModelProductCategory, error) {
	query := `
		UPDATE product_category
		SET
			category_name = $1,
			description = $2
		WHERE id = $3
	`

	if _, err := c.db.Exec(
		query,
		category.Name,
		category.Description,
		category.ID,
	); err != nil {
		return nil, err
	}

	return category, nil
}

// EnableDisable implements ports.PortCategory.
func (c *categoryRepo) EnableDisable(category *models.ModelProductCategory, value bool) (*models.ModelProductCategory, error) {
	query := "UPDATE product_category SET available = $1 WHERE id = $2"
	if _, err := c.db.Exec(
		query,
		value,
		category.ID,
	); err != nil {
		return nil, err
	}

	category.Available = value

	return category, nil
}

func (c *categoryRepo) toModel(entity *entities.EntityProductCategory) *models.ModelProductCategory {
	return &models.ModelProductCategory{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Available:   entity.Available,
	}
}

func (c *categoryRepo) toModelMap(entities []entities.EntityProductCategory) []models.ModelProductCategory {
	models := make([]models.ModelProductCategory, 0)
	for _, e := range entities {
		models = append(models, *c.toModel(&e))
	}
	return models
}
