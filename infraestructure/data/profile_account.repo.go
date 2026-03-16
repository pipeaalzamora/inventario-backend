package data

import (
	"fmt"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/infraestructure/entities"
	"sofia-backend/types"

	"github.com/jmoiron/sqlx"
)

type ProfileAccountRepo struct {
	db *sqlx.DB
}

func NewProfileAccountRepo(db *sqlx.DB) ports.PortProfile {
	return &ProfileAccountRepo{
		db: db,
	}
}

func (r *ProfileAccountRepo) GetProfilesByUserID(userId string) ([]models.ProfileAccountModel, error) {
	var profiles []entities.ProfileAccount

	query := `
		SELECT pa.*
		FROM profile_accounts pa
		JOIN user_account_per_profiles uapp ON uapp.profile_account_id = pa.id
		WHERE uapp.user_account_id = $1
	`

	err := r.db.Select(&profiles, query, userId)
	if err != nil {
		return nil, types.ThrowData("Error al obtener los perfiles por ID de usuario")
	}

	profileModels := r.toProfileAccountModelList(profiles)

	return profileModels, nil
}

func (r *ProfileAccountRepo) GetProfileByID(profileId string) (*models.ProfileAccountModel, error) {
	var profile entities.ProfileAccount

	query := `
		SELECT *
		FROM profile_accounts
		WHERE id = $1
	`

	err := r.db.Get(&profile, query, profileId)
	if err != nil {
		return nil, types.ThrowData("Error al obtener el perfil por ID")
	}

	profileModel := r.toProfileAccountModel(&profile)

	return profileModel, nil
}

func (r *ProfileAccountRepo) GetProfiles() ([]models.ProfileAccountModel, error) {
	var profiles []entities.ProfileAccount

	query := `
		SELECT *
		FROM profile_accounts
	`

	err := r.db.Select(&profiles, query)
	if err != nil {
		return nil, types.ThrowData("Error al obtener todos los perfiles")
	}
	return r.toProfileAccountModelList(profiles), nil
}

func (r *ProfileAccountRepo) CreateProfile(input *models.ProfileAccountModel) (*models.ProfileAccountModel, error) {
	query := `
		INSERT INTO profile_accounts (profile_name, description)
		VALUES ($1, $2)
		RETURNING id, profile_name, description
	`
	var profile entities.ProfileAccount
	err := r.db.QueryRowx(query, input.ProfileName, input.Description).Scan(&profile.ID, &profile.ProfileName, &profile.Description)
	if err != nil {
		return nil, types.ThrowData("Error al crear el perfil")
	}

	profileModel := r.toProfileAccountModel(&profile)

	return profileModel, nil
}

func (r *ProfileAccountRepo) UpdateProfile(id string, input *models.ProfileAccountModel) (*models.ProfileAccountModel, error) {
	query := `		UPDATE profile_accounts
		SET profile_name = $1, description = $2
		WHERE id = $3
		RETURNING id, profile_name, description
	`
	var updatedProfile entities.ProfileAccount
	err := r.db.QueryRowx(query, input.ProfileName, input.Description, id).Scan(&updatedProfile.ID, &updatedProfile.ProfileName, &updatedProfile.Description)
	if err != nil {
		return nil, types.ThrowData("Error al actualizar el perfil")
	}
	return r.toProfileAccountModel(&updatedProfile), nil
}

func (r *ProfileAccountRepo) CreateProfileWithPowers(input *models.ProfileAccountModel, powerIDs []string) (*models.ProfileAccountModel, error) {
	query := `
		INSERT INTO profile_accounts (profile_name, description)
		VALUES ($1, $2)
		RETURNING id, profile_name, description
	`
	var profile entities.ProfileAccount
	// With tx
	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción")
	}
	defer tx.Rollback()

	err = tx.QueryRowx(query, input.ProfileName, input.Description).Scan(
		&profile.ID,
		&profile.ProfileName,
		&profile.Description,
	)
	if err != nil {
		fmt.Printf("Error creating profile: %v\n", err)
		tx.Rollback()
		return nil, types.ThrowData("Error al crear el perfil")
	}

	for _, powerID := range powerIDs {
		powerQuery := `
			INSERT INTO profile_account_per_power_accounts (profile_account_id, power_account_id)
			VALUES ($1, $2)
		`
		_, err = tx.Exec(powerQuery, profile.ID, powerID)
		if err != nil {
			tx.Rollback()
			fmt.Printf("Error associating powers: %v\n", err)
			return nil, types.ThrowData("Error al asociar los permisos con el perfil")
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	return r.toProfileAccountModel(&profile), nil
}

func (r *ProfileAccountRepo) UpdateProfileWithPowers(id string, input *models.ProfileAccountModel, powerIDs []string) (*models.ProfileAccountModel, error) {

	// With tx
	query := `		UPDATE profile_accounts
		SET profile_name = $1, description = $2
		WHERE id = $3
		RETURNING id, profile_name, description
	`

	tx, err := r.db.Beginx()
	if err != nil {
		return nil, types.ThrowData("Error al iniciar la transacción")
	}
	defer tx.Rollback()

	var updatedProfile entities.ProfileAccount

	err = tx.QueryRowx(query, input.ProfileName, input.Description, id).Scan(&updatedProfile.ID, &updatedProfile.ProfileName, &updatedProfile.Description)
	if err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al actualizar el perfil")
	}

	// Clear existing powers
	clearQuery := `		DELETE FROM profile_account_per_power_accounts
		WHERE profile_account_id = $1
	`
	_, err = tx.Exec(clearQuery, updatedProfile.ID)
	if err != nil {
		tx.Rollback()
		return nil, types.ThrowData("Error al limpiar los permisos existentes")
	}

	for _, powerID := range powerIDs {
		// Insert new powers
		powerQuery := `INSERT INTO profile_account_per_power_accounts (profile_account_id, power_account_id)
			VALUES ($1, $2)
		`
		_, err = tx.Exec(powerQuery, updatedProfile.ID, powerID)
		if err != nil {
			tx.Rollback()
			fmt.Printf("Error inserting new power: %v\n", err)
			return nil, types.ThrowData("Error al insertar los nuevos permisos")
		}
	}

	if err = tx.Commit(); err != nil {
		return nil, types.ThrowData("Error al confirmar la transacción")
	}

	updatedProfileModel := r.toProfileAccountModel(&updatedProfile)

	return updatedProfileModel, nil
}

func (r *ProfileAccountRepo) DeleteProfile(id string) error {
	query := `DELETE FROM profile_accounts WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return types.ThrowData("Error al eliminar la cuenta de perfil")
	}
	return nil
}

func (r *ProfileAccountRepo) GetAllProfilesWithUserId() ([]models.ProfileAccountUserAccountModel, error) {
	query := `
		SELECT 
			pa.id, 
			pa.profile_name, 
			pa.description, 
			ua.id AS user_id
		FROM profile_accounts pa
		JOIN user_account_per_profiles uap ON pa.id = uap.profile_account_id
		JOIN user_accounts ua ON ua.id = uap.user_account_id
	`
	var profiles []entities.ProfileAccountUserAccount
	err := r.db.Select(&profiles, query)
	if err != nil {
		fmt.Printf("Error fetching profiles with user IDs: %v\n", err)
		return nil, types.ThrowData("Error al obtener los perfiles con IDs de usuario")
	}

	profilesModels := make([]models.ProfileAccountUserAccountModel, len(profiles))
	for index, profile := range profiles {
		profilesModels[index] = models.ProfileAccountUserAccountModel{
			ID:          profile.ID,
			ProfileName: profile.ProfileName,
			Description: profile.Description,
			UserID:      profile.UserID,
		}
	}

	return profilesModels, nil
}

func (r *ProfileAccountRepo) GetProfilesByPowerID(powerName string) ([]models.ProfileAccountModel, error) {
	var profiles []entities.ProfileAccount
	query := `
		SELECT pa.*
		FROM profile_accounts pa
		JOIN profile_account_per_power_accounts pppa ON pppa.profile_account_id = pa.id
		JOIN power_accounts pwa ON pwa.id = pppa.power_account_id
		WHERE pwa.power_name = $1
	`
	err := r.db.Select(&profiles, query, powerName)
	if err != nil {
		fmt.Printf("Error fetching profiles by power ID: %v\n", err)
		return nil, types.ThrowData("Error al obtener los perfiles por ID de permiso")
	}
	return r.toProfileAccountModelList(profiles), nil
}

func (r *ProfileAccountRepo) toProfileAccountModelList(profiles []entities.ProfileAccount) []models.ProfileAccountModel {
	profilesModels := make([]models.ProfileAccountModel, len(profiles))

	for index, profile := range profiles {
		profilesModels[index] = models.ProfileAccountModel{
			ID:          profile.ID,
			ProfileName: profile.ProfileName,
			Description: profile.Description,
		}
	}
	return profilesModels
}

func (r *ProfileAccountRepo) toProfileAccountModel(profile *entities.ProfileAccount) *models.ProfileAccountModel {
	if profile == nil {
		return nil
	}
	return &models.ProfileAccountModel{
		ID:          profile.ID,
		ProfileName: profile.ProfileName,
		Description: profile.Description,
	}
}
