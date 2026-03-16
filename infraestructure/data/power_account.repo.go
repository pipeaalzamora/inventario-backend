package data

import (
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

type PowerAccountRepo struct {
	db *sqlx.DB
}

func NewPowerAccountRepo(db *sqlx.DB) ports.PortPower {
	return &PowerAccountRepo{
		db: db,
	}
}

func (r *PowerAccountRepo) GetPowersByProfile(profileId string) ([]models.PowerAccountModel, error) {
	var powers []entities.PowerAccount

	query := `
		SELECT p.*
		FROM power_accounts p
		JOIN profile_account_per_power_accounts pap ON pap.power_account_id = p.id
		WHERE pap.profile_account_id = $1
	`

	err := r.db.Select(&powers, query, profileId)
	if err != nil {
		return nil, types.ThrowData("Error al obtener los permisos por ID de perfil")
	}
	return r.toPowerAccountModelList(powers), nil
}

func (r *PowerAccountRepo) GetPowerCategories() ([]models.PowerAccountCategoryModel, error) {
	var categories []entities.PowerAccountCategory
	query := `SELECT id, category_name, description, ownable FROM power_account_categories`
	err := r.db.Select(&categories, query)
	if err != nil {
		return nil, types.ThrowData("Error al obtener las categorías de permisos")
	}

	categoryModels := r.toPowerAccountCategoryModelList(categories)
	return categoryModels, nil
}

func (r *PowerAccountRepo) GetPowerCategoryByID(categoryId string) (*models.PowerAccountCategoryModel, error) {
	var category entities.PowerAccountCategory
	query := `SELECT id, category_name, description, ownable FROM power_account_categories WHERE id = $1`
	err := r.db.Get(&category, query, categoryId)
	if err != nil {
		return nil, types.ThrowData("Error al obtener la categoría de permisos por ID")
	}
	categoryModel := r.toPowerAccountCategoryModel(&category)
	return categoryModel, nil
}

func (r *PowerAccountRepo) GetPowerAccountsByCategory(categoryId string) ([]models.PowerAccountModel, error) {
	var powers []entities.PowerAccount
	query := `
		SELECT p.*
		FROM power_accounts p
		WHERE p.power_account_category_id = $1
	`
	err := r.db.Select(&powers, query, categoryId)
	if err != nil {
		return nil, types.ThrowData("Error al obtener los permisos por ID de categoría")
	}
	powerModels := r.toPowerAccountModelList(powers)
	return powerModels, nil
}

func (r *PowerAccountRepo) CreatePowerCategory(input *models.PowerAccountCategoryModel) (*models.PowerAccountCategoryModel, error) {
	query := `
		INSERT INTO power_account_categories (category_name, description, ownable)
		VALUES ($1, $2, $3)
		RETURNING id, category_name, description, ownable
	`
	var category entities.PowerAccountCategory
	err := r.db.QueryRow(
		query,
		input.CategoryName,
		input.Description,
		input.Ownable,
	).Scan(&category.ID, &category.CategoryName, &category.Description, &category.Ownable)
	if err != nil {
		return nil, types.ThrowData("Error al crear la categoría de permisos")
	}

	categoryModel := r.toPowerAccountCategoryModel(&category)
	return categoryModel, nil
}

func (r *PowerAccountRepo) CreatePowerAccount(input *models.PowerAccountModel) (*models.PowerAccountModel, error) {
	query := `
		INSERT INTO power_accounts (power_name, description, power_account_category_id)
		VALUES ($1, $2, $3)
		RETURNING id, power_name, description, power_account_category_id
	`
	var power entities.PowerAccount
	err := r.db.QueryRow(
		query,
		input.PowerName,
		input.Description,
		input.CategoryID,
	).Scan(&power.ID, &power.PowerName, &power.Description, &power.CategoryID)
	if err != nil {
		return nil, types.ThrowData("Error al crear el permiso")
	}

	powerModel := r.toPowerAccountModel(&power)
	return powerModel, nil
}

func (r *PowerAccountRepo) GetPowers() ([]models.PowerAccountModel, error) {
	var powers []entities.PowerAccount
	query := `SELECT id, power_name, description, power_display, power_account_category_id FROM power_accounts`
	err := r.db.Select(&powers, query)
	if err != nil {
		return nil, types.ThrowData("Error al obtener los permisos")
	}

	powerModels := r.toPowerAccountModelList(powers)

	return powerModels, nil
}

func (r *PowerAccountRepo) GetPowersByUserID(userId string) ([]models.PowerAccountModel, error) {
	var powers []entities.PowerAccount
	query := `
		SELECT DISTINCT ON (p.id) p.*
		FROM power_accounts p
		JOIN profile_account_per_power_accounts pap ON pap.power_account_id = p.id
		JOIN user_account_per_profiles u ON u.profile_account_id = pap.profile_account_id
		WHERE u.user_account_id = $1
	`
	err := r.db.Select(&powers, query, userId)
	if err != nil {
		return nil, types.ThrowData("Error al obtener los permisos por ID de usuario")
	}
	powerModels := r.toPowerAccountModelList(powers)
	return powerModels, nil
}

func (r *PowerAccountRepo) GetAllPowersWithProfileId() ([]models.PowerAccountProfileModel, error) {
	query := `
		SELECT 
			p.id, 
			p.power_name, 
			p.power_display,
			p.description, 
			p.power_account_category_id,
			pa.id AS profile_id
		FROM power_accounts p
		JOIN profile_account_per_power_accounts pap ON pap.power_account_id = p.id
		JOIN profile_accounts pa ON pa.id = pap.profile_account_id
	`
	var powers []entities.PowerAccountProfile
	err := r.db.Select(&powers, query)
	if err != nil {
		return nil, types.ThrowData("Error al obtener todos los permisos con IDs de perfil")
	}

	powerModels := r.toPowerAccountProfileModelList(powers)

	return powerModels, nil
}

func (r *PowerAccountRepo) AddOwnPowerToProfileTx(tx *sqlx.Tx, profileIds []string, power *models.PowerAccountModel) error {
	query := `
		INSERT INTO profile_account_per_power_accounts (profile_account_id, power_account_id)
		VALUES ($1, $2)
	`
	for _, profileId := range profileIds {
		_, err := tx.Exec(query, profileId, power.ID)
		if err != nil {
			return types.ThrowData("Error al agregar permiso propio al perfil")
		}
	}
	return nil
}

////////////// Private methods ///////////////

func (r *PowerAccountRepo) toPowerAccountCategoryModel(category *entities.PowerAccountCategory) *models.PowerAccountCategoryModel {
	if category == nil {
		return nil
	}
	return &models.PowerAccountCategoryModel{
		ID:           category.ID,
		CategoryName: category.CategoryName,
		Description:  category.Description,
		Ownable:      category.Ownable,
	}
}

func (r *PowerAccountRepo) toPowerAccountCategoryModelList(categories []entities.PowerAccountCategory) []models.PowerAccountCategoryModel {
	categoryModels := make([]models.PowerAccountCategoryModel, len(categories))
	for index, category := range categories {
		categoryModels[index] = *r.toPowerAccountCategoryModel(&category)
	}
	return categoryModels
}

func (r *PowerAccountRepo) toPowerAccountModel(power *entities.PowerAccount) *models.PowerAccountModel {
	if power == nil {
		return nil
	}
	return &models.PowerAccountModel{
		ID:          power.ID,
		PowerName:   power.PowerName,
		DisplayName: power.DisplayName,
		Description: power.Description,
		CategoryID:  power.CategoryID,
	}
}

func (r *PowerAccountRepo) toPowerAccountModelList(powers []entities.PowerAccount) []models.PowerAccountModel {
	powerModels := make([]models.PowerAccountModel, len(powers))
	for index, power := range powers {
		powerModels[index] = *r.toPowerAccountModel(&power)
	}
	return powerModels
}

func (r *PowerAccountRepo) toPowerAccountProfileModelList(entities []entities.PowerAccountProfile) []models.PowerAccountProfileModel {
	if len(entities) == 0 {
		return []models.PowerAccountProfileModel{}
	}
	powers := make([]models.PowerAccountProfileModel, len(entities))
	for i, entity := range entities {
		powers[i] = models.PowerAccountProfileModel{
			ID:          entity.ID,
			PowerName:   entity.PowerName,
			DisplayName: entity.DisplayName,
			Description: entity.Description,
			CategoryID:  entity.CategoryID,
			ProfileID:   entity.ProfileID,
		}
	}
	return powers
}
