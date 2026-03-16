package data

import (
	"context"
	"database/sql"
	"encoding/json"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserAccountRepo struct {
	db *sqlx.DB
}

func NewUserAccountRepo(db *sqlx.DB) ports.PortUser {
	return &UserAccountRepo{
		db: db,
	}
}

func (r *UserAccountRepo) GetUserByEmail(email string) (*models.UserAccountModel, error) {
	var user entities.UserAccount
	err := r.db.Get(&user, "SELECT * FROM user_accounts WHERE user_email = $1", email)
	if err == sql.ErrNoRows {
		return nil, nil // User not found
	}
	if err != nil {
		return nil, types.ThrowData("Error al obtener el usuario por email")
	}
	return r.toUserModel(&user), nil
}

func (r *UserAccountRepo) GetUserByID(id string) (*models.UserAccountModel, error) {
	var user entities.UserAccount
	err := r.db.Get(&user, "SELECT * FROM user_accounts WHERE id = $1", id)
	if err != nil {
		return nil, types.ThrowData("Error al obtener el usuario por ID")
	}
	return r.toUserModel(&user), nil
}

func (r *UserAccountRepo) GetAllUsers(page int, size int, filter *map[string]interface{}) ([]models.UserAccountModel, int, error) {
	var users []entities.UserAccount
	query := `SELECT ua.*
		FROM user_accounts ua`

	if filter != nil {
		_, err := json.Marshal(*filter)
		if err != nil {
			return nil, 0, types.ThrowData("Error al procesar los filtros")
		}
	}

	err := r.db.SelectContext(context.TODO(), &users, query+" ORDER BY ua.created_at DESC LIMIT $2 OFFSET $3", size, (page-1)*size)
	if err != nil {
		return nil, 0, types.ThrowData("Error al obtener la lista de usuarios")
	}

	var total int
	err = r.db.GetContext(context.TODO(), &total, "SELECT COUNT(*) FROM user_accounts")

	if err != nil {
		return nil, 0, types.ThrowData("Error al contar los usuarios")
	}

	userModels := r.toUserModelList(users)

	return userModels, total, nil
}

func (r *UserAccountRepo) GetUsers() ([]models.UserAccountModel, error) {
	var users []entities.UserAccount
	query := `SELECT * FROM user_accounts ORDER BY created_at DESC`

	err := r.db.SelectContext(context.TODO(), &users, query)
	if err != nil {
		return nil, types.ThrowData("Error al obtener la lista de usuarios")
	}

	userModels := r.toUserModelList(users)

	return userModels, nil
}

func (r *UserAccountRepo) GetUsersByProfileID(page int, size int, filter *map[string]interface{}, profileID string) ([]models.UserAccountModel, int, error) {
	var users []entities.UserAccount
	query := `SELECT ua.*
		FROM user_accounts ua
		JOIN user_account_per_profiles uap ON ua.id = uap.user_account_id
		WHERE uap.profile_account_id = $1`

	if filter != nil {
		_, err := json.Marshal(*filter)
		if err != nil {
			return nil, 0, types.ThrowData("Error al procesar los filtros")
		}
	}
	err := r.db.SelectContext(context.TODO(), &users, query+" ORDER BY ua.created_at DESC LIMIT $2 OFFSET $3", profileID, size, (page-1)*size)
	if err != nil {
		return nil, 0, types.ThrowData("Error al obtener la lista de usuarios")
	}

	var total int
	err = r.db.GetContext(context.TODO(), &total, `SELECT COUNT(*) FROM user_account_per_profiles WHERE profile_account_id = $1`, profileID)
	if err != nil {
		return nil, 0, types.ThrowData("Error al contar los usuarios")
	}

	userModels := r.toUserModelList(users)

	return userModels, total, nil
}

func (r *UserAccountRepo) CreateUser(user *models.UserAccountModel) (*models.UserAccountModel, error) {
	query := `
		INSERT INTO user_accounts (user_name, user_email, description, user_password, is_new_account)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id,
			user_name,
			user_email,
			description,
			user_password,
			available,
			is_new_account,
			deleted_at,
			created_at,
			updated_at
	`

	var createdUser entities.UserAccount
	err := r.db.QueryRow(
		query,
		user.UserName,
		user.UserEmail,
		user.Description,
		user.UserPassword,
	).Scan(
		&createdUser.ID,
		&createdUser.UserName,
		&createdUser.UserEmail,
		&createdUser.Description,
		&createdUser.UserPassword,
		&createdUser.Available,
		&createdUser.IsNewAccount,
		&createdUser.DeletedAt,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)
	if err != nil {
		return nil, types.ThrowData("Error al crear el usuario")
	}

	createdUserModel := r.toUserModel(&createdUser)

	return createdUserModel, nil
}

func (r *UserAccountRepo) CreateUserWithProfiles(input *models.UserAccountModel, profileIDs []string) (*models.UserAccountModel, error) {
	// begin transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción")
	}
	defer tx.Rollback() // Ensure rollback on error

	query := `INSERT INTO user_accounts (user_name, user_email, description, user_password)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			user_name,
			user_email,
			description,
			user_password,
			available,
			is_new_account,
			deleted_at,
			created_at,
			updated_at
	`
	var createdUser entities.UserAccount
	err = tx.QueryRow(
		query,
		input.UserName,
		input.UserEmail,
		input.Description,
		input.UserPassword,
	).Scan(
		&createdUser.ID,
		&createdUser.UserName,
		&createdUser.UserEmail,
		&createdUser.Description,
		&createdUser.UserPassword,
		&createdUser.Available,
		&createdUser.IsNewAccount,
		&createdUser.DeletedAt,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)
	if err != nil {
		return nil, types.ThrowData("Error al crear el usuario")
	}
	// Insert user profiles
	for _, profileID := range profileIDs {
		_, err = tx.Exec(
			`INSERT INTO user_account_per_profiles (user_account_id, profile_account_id) VALUES ($1, $2)`,
			createdUser.ID,
			profileID,
		)
		if err != nil {
			return nil, types.ThrowData("Error al asociar el usuario con el perfil")
		}
	}
	// Commit transaction
	if err = tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	createdUserModel := r.toUserModel(&createdUser)

	return createdUserModel, nil
}

func (r *UserAccountRepo) UpdateUserWithProfiles(userId string, input *models.UserAccountModel, profileIDs []string) (*models.UserAccountModel, error) {
	// begin transaction
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción")
	}
	defer tx.Rollback() // Ensure rollback on error

	query := `UPDATE user_accounts
		SET user_name = $1, user_email = $2, description = $3, is_new_account = $4
		WHERE id = $5
		RETURNING
			id,
			user_name,
			user_email,
			description,
			user_password,
			available,
			is_new_account,
			deleted_at,
			created_at,
			updated_at
	`
	var updatedUser entities.UserAccount
	err = tx.QueryRow(
		query,
		input.UserName,
		input.UserEmail,
		input.Description,
		input.IsNewAccount,
		userId,
	).Scan(
		&updatedUser.ID,
		&updatedUser.UserName,
		&updatedUser.UserEmail,
		&updatedUser.Description,
		&updatedUser.UserPassword,
		&updatedUser.Available,
		&updatedUser.IsNewAccount,
		&updatedUser.DeletedAt,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)
	if err != nil {
		return nil, types.ThrowData("Error al actualizar el usuario")
	}

	// Delete existing user profiles
	_, err = tx.Exec(`DELETE FROM user_account_per_profiles WHERE user_account_id = $1`, userId)
	if err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al eliminar los perfiles existentes del usuario")
	}
	// Insert user profiles
	for _, profileID := range profileIDs {
		_, err = tx.Exec(
			`INSERT INTO user_account_per_profiles (user_account_id, profile_account_id) VALUES ($1, $2)`,
			updatedUser.ID,
			profileID,
		)
		if err != nil {
			tx.Rollback()
			return nil, types.ThrowData("Error al asociar el usuario con el perfil")
		}
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	updatedUserModel := r.toUserModel(&updatedUser)

	return updatedUserModel, nil
}

func (r *UserAccountRepo) UpdateUserPassword(userId string, password string) error {
	query := `
		UPDATE user_accounts
		SET user_password = $1, updated_at = NOW(), is_new_account = false
		WHERE id = $2
	`
	_, err := r.db.Exec(query, password, userId)
	if err != nil {
		return types.ThrowData("Error al actualizar la contraseña del usuario")
	}
	return nil
}

func (r *UserAccountRepo) UpdateUser(id string, user *models.UserAccountModel) (*models.UserAccountModel, error) {
	query := `
		UPDATE user_accounts
		SET user_name = $1, user_email = $2, description = $3, is_new_account = $4, updated_at = NOW()
		WHERE id = $5
		RETURNING
			id,
			user_name,
			user_email,
			description,
			user_password,
			available,
			is_new_account,
			deleted_at,
			created_at,
			updated_at
	`

	var updatedUser entities.UserAccount
	err := r.db.QueryRow(
		query,
		user.UserName,
		user.UserEmail,
		user.Description,
		user.IsNewAccount,
		id,
	).Scan(
		&updatedUser.ID,
		&updatedUser.UserName,
		&updatedUser.UserEmail,
		&updatedUser.Description,
		&updatedUser.UserPassword,
		&updatedUser.Available,
		&updatedUser.IsNewAccount,
		&updatedUser.DeletedAt,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)
	if err != nil {
		return nil, types.ThrowData("Error al actualizar el usuario")
	}

	updatedUserModel := r.toUserModel(&updatedUser)

	return updatedUserModel, nil
}

func (r *UserAccountRepo) DeactivateUser(id string) error {
	query := `
		UPDATE user_accounts
		SET deleted_at = NOW(), available = false
		WHERE id = $1
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return types.ThrowData("Error al desactivar el usuario")
	}
	return nil
}

// ActivateUser implements ports.PortUser.
func (userRepo *UserAccountRepo) ActivateUser(id string) error {
	query := `
		UPDATE user_accounts
		SET deleted_at = NULL, available = true
		WHERE id = $1
	`
	_, err := userRepo.db.Exec(query, id)
	if err != nil {
		return types.ThrowData("Error al activar el usuario")
	}
	return nil
}

func (r *UserAccountRepo) GetUsersByProfileIDs(profileIDs []string) ([]models.UserAccountModel, error) {
	var users []entities.UserAccount
	query := `SELECT DISTINCT ua.*
		FROM user_accounts ua
		JOIN user_account_per_profiles uap ON ua.id = uap.user_account_id
		WHERE uap.profile_account_id = ANY($1)`

	err := r.db.SelectContext(context.TODO(), &users, query, pq.Array(profileIDs))
	if err != nil {
		return nil, types.ThrowData("Error al obtener la lista de usuarios")
	}

	userModels := r.toUserModelList(users)

	return userModels, nil
}

func (userRepo *UserAccountRepo) toUserModel(u *entities.UserAccount) *models.UserAccountModel {
	return &models.UserAccountModel{
		ID:           u.ID,
		UserName:     u.UserName,
		UserEmail:    u.UserEmail,
		Description:  u.Description,
		UserPassword: u.UserPassword,
		Available:    u.Available,
		IsNewAccount: u.IsNewAccount,
		DeletedAt:    u.DeletedAt,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func (userRepo *UserAccountRepo) toUserModelList(us []entities.UserAccount) []models.UserAccountModel {
	result := make([]models.UserAccountModel, len(us))
	for i, u := range us {
		result[i] = *userRepo.toUserModel(&u)
	}
	return result
}
