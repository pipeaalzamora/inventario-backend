package services

import (
	"context"
	"fmt"
	"sofia-backend/api/v1/recipe"
	"sofia-backend/domain/models"
	"sofia-backend/domain/ports"
	"sofia-backend/shared"
	"sofia-backend/types"
)

type UserService struct {
	PowerChecker
	userAccountRepo    ports.PortUser
	profileAccountRepo ports.PortProfile
}

func NewUserService(repo ports.PortUser, profileRepo ports.PortProfile) *UserService {
	return &UserService{
		userAccountRepo:    repo,
		profileAccountRepo: profileRepo,
	}
}

func (s *UserService) GetAllUsers(ctx context.Context, page int, size int, filter *map[string]interface{}) (data []models.UserAccountModel, total int, err error) {
	userAccounts, total, err := s.userAccountRepo.GetAllUsers(page, size, filter)
	if err != nil {
		return nil, 0, err
	}
	return userAccounts, total, nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]models.UserAccountModel, error) {
	userAccounts, err := s.userAccountRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	return userAccounts, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.UserAccountModel, error) {
	user, err := s.userAccountRepo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.UserAccountModel, error) {
	user, err := s.userAccountRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil // User not found
	}

	return user, nil
}

func (s *UserService) CreateUser(ctx context.Context, recipe *recipe.UserRecipe) (*models.UserAccountModel, error) {
	if ok := s.EveryPower(ctx, PowerUserCreate); !ok {
		return nil, shared.PowerError{Message: fmt.Sprintf("User does not have %s power", PowerUserCreate)}
	}

	inputUser := &models.UserAccountModel{
		UserName:     recipe.UserName,
		UserEmail:    recipe.UserEmail,
		Description:  recipe.Description,
		UserPassword: recipe.Password,
	}

	createdUser, err := s.userAccountRepo.CreateUserWithProfiles(inputUser, recipe.ProfileIDs)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id string, recipe *recipe.UserRecipe) (*models.UserAccountModel, error) {
	currUser := ctx.Value(shared.UserIdKey())

	// Chequeamos si el es un usuario iniciado sesión o no
	if currUser != nil {
		if !s.EveryPower(ctx, PowerUserUpdate) {
			return nil, shared.PowerError{Message: fmt.Sprintf("User does not have %s power", PowerUserUpdate)}
		}
	}

	inputUser := &models.UserAccountModel{
		UserName:    recipe.UserName,
		UserEmail:   recipe.UserEmail,
		Description: recipe.Description,
	}

	// This method does not allow updating the password
	updatedUser, err := s.userAccountRepo.UpdateUserWithProfiles(id, inputUser, recipe.ProfileIDs)
	if err != nil {
		return nil, err
	}

	// update password if provided
	if recipe.Password != "" {
		err = s.userAccountRepo.UpdateUserPassword(id, recipe.Password)
		if err != nil {
			return nil, err
		}
	}

	return updatedUser, nil
}

func (s *UserService) UpdateUserNotAuth(ctx context.Context, id string, recipe *recipe.UserRecipe) (*models.UserAccountModel, error) {
	inputUser := &models.UserAccountModel{
		UserName:     recipe.UserName,
		UserEmail:    recipe.UserEmail,
		IsNewAccount: recipe.IsNewAccount,
		Description:  recipe.Description,
	}

	// This method does not allow updating the password
	updatedUser, err := s.userAccountRepo.UpdateUser(id, inputUser)
	if err != nil {
		return nil, err
	}

	// update password if provided
	if recipe.Password != "" {
		err = s.userAccountRepo.UpdateUserPassword(id, recipe.Password)
		if err != nil {
			return nil, err
		}
	}

	return updatedUser, nil
}

func (s *UserService) DeactivateUser(ctx context.Context, id string) error {
	if ok := s.EveryPower(ctx, PowerUserDelete); !ok {
		return shared.PowerError{Message: fmt.Sprintf("User does not have %s power", PowerUserDelete)}
	}
	return s.userAccountRepo.DeactivateUser(id)
}

func (s *UserService) ActivateUser(ctx context.Context, id string) error {
	if ok := s.EveryPower(ctx, PowerUserDelete); !ok {
		return types.ThrowPower(fmt.Sprintf("El usuario no tiene el permiso %s", PowerUserDelete))
	}

	return s.userAccountRepo.ActivateUser(id)
}

func (s *UserService) GetUsersByProfileIDs(ctx context.Context, profileIDs []string) ([]models.UserAccountModel, error) {
	return s.userAccountRepo.GetUsersByProfileIDs(profileIDs)
}
