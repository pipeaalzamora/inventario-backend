package facades

import (
	"context"
	"sofia-backend/api/v1/dto"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/config"
	"sofia-backend/domain/external"
	"sofia-backend/domain/models"
	"sofia-backend/domain/services"
	"sofia-backend/shared"
	"sofia-backend/types"
)

type UserFacade struct {
	services *services.ServiceContainer
	external *external.ServiceContainer
	baseUrl  string
	secret   string
}

func NewUserFacade(
	services *services.ServiceContainer,
	external *external.ServiceContainer,
	config *config.Config,
) *UserFacade {
	return &UserFacade{
		services: services,
		external: external,
		baseUrl:  config.BaseUrl,
		secret:   config.Secret,
	}
}

func (f *UserFacade) GetAllUsers(ctx context.Context) ([]dto.UserWithProfilesDTO, error) {
	users, err := f.services.UserService.GetUsers(ctx)
	if err != nil {
		return nil, err
	}

	usersPowers, err := f.services.ProfileService.GetAllProfilesWithUserId(ctx)
	if err != nil {
		return nil, err
	}

	// for perfomance, we can use a map to group profiles by user ID
	usersMap := make(map[string][]models.ProfileAccountModel)
	for _, profile := range usersPowers {
		profileModel := models.ProfileAccountModel{
			ID:          profile.ID,
			ProfileName: profile.ProfileName,
			Description: profile.Description,
		}
		usersMap[profile.UserID] = append(usersMap[profile.UserID], profileModel)
	}

	usersWithProfiles := make([]dto.UserWithProfilesDTO, len(users))
	for index, user := range users {
		profiles := usersMap[user.ID]
		if usersMap[user.ID] == nil {
			profiles = []models.ProfileAccountModel{}
		}
		userDto := dto.UserWithProfilesDTO{
			ID:           user.ID,
			UserEmail:    user.UserEmail,
			UserName:     user.UserName,
			IsNewAccount: user.IsNewAccount,
			Available:    user.Available,
			Description:  user.Description,
			DeletedAt:    user.DeletedAt,
			Profiles:     profiles,
			CreatedAt:    user.CreatedAt,
			UpdatedAt:    user.UpdatedAt,
		}

		usersWithProfiles[index] = userDto
	}

	return usersWithProfiles, nil
}

func (f *UserFacade) GetUserById(ctx context.Context, userId string) (*dto.UserWithProfilesDTO, error) {
	user, err := f.services.UserService.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, types.ThrowMsg("Usuario no encontrado")
	}

	profiles, err := f.services.ProfileService.GetProfilesByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	userDto := &dto.UserWithProfilesDTO{
		ID:           user.ID,
		UserEmail:    user.UserEmail,
		UserName:     user.UserName,
		IsNewAccount: user.IsNewAccount,
		Available:    user.Available,
		DeletedAt:    user.DeletedAt,
		Description:  user.Description,
		Profiles:     profiles,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}

	return userDto, nil
}

func (f *UserFacade) CreateUser(ctx context.Context, recipe *recipe.UserRecipe) (*dto.UserDTO, error) {
	user, err := f.services.UserService.GetUserByEmail(ctx, recipe.UserEmail)

	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, types.ThrowMsg("El correo electrónico ya está en uso")
	}

	generatePassword, err := shared.CreateRandomString(10)
	if err != nil {
		return nil, err
	}

	userPassword := shared.CreatePassword(generatePassword, f.secret)
	recipe.Password = userPassword

	userCreated, err := f.services.UserService.CreateUser(ctx, recipe)
	if err != nil {
		return nil, err
	}

	if userCreated == nil {
		return nil, types.ThrowMsg("No se pudo crear el usuario")
	}

	err = f.external.EmailService.SendWelcomeEmail(userCreated.UserEmail, generatePassword, f.baseUrl)
	if err != nil {
		return nil, err
	}

	userDto := &dto.UserDTO{
		ID:           userCreated.ID,
		UserEmail:    userCreated.UserEmail,
		UserName:     userCreated.UserName,
		Description:  userCreated.Description,
		Available:    userCreated.Available,
		IsNewAccount: userCreated.IsNewAccount,
		DeletedAt:    userCreated.DeletedAt,
		CreatedAt:    userCreated.CreatedAt,
		UpdatedAt:    userCreated.UpdatedAt,
	}

	return userDto, nil
}

func (f *UserFacade) UpdateUser(ctx context.Context, userId string, recipe *recipe.UserRecipe) (*dto.UserDTO, error) {
	user, err := f.services.UserService.GetUserByID(ctx, userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, types.ThrowMsg("Usuario no encontrado")
	}

	userUpdated, err := f.services.UserService.UpdateUser(ctx, userId, recipe)
	if err != nil {
		return nil, err
	}

	err = f.services.AuthService.DeletePersistedUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	userDto := &dto.UserDTO{
		ID:           userUpdated.ID,
		UserEmail:    userUpdated.UserEmail,
		UserName:     userUpdated.UserName,
		IsNewAccount: userUpdated.IsNewAccount,
		Description:  userUpdated.Description,
		Available:    userUpdated.Available,
		CreatedAt:    userUpdated.CreatedAt,
		UpdatedAt:    userUpdated.UpdatedAt,
	}

	return userDto, nil
}

func (f *UserFacade) DeactivateUser(ctx context.Context, userId string) error {
	err := f.services.UserService.DeactivateUser(ctx, userId)
	if err != nil {
		return err
	}

	err = f.services.AuthService.DeletePersistedUser(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}

func (f *UserFacade) ActivateUser(ctx context.Context, userId string) error {
	err := f.services.UserService.ActivateUser(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}
