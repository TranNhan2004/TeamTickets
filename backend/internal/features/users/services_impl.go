package users

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/trannhanlv2004/team-tickets/internal/apperrors"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepository UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) Create(ctx context.Context, model CreateUserRequestModel) (*CreateUserResponseModel, *apperrors.AppError) {
	model.FirstName = strings.TrimSpace(model.FirstName)
	model.LastName = strings.TrimSpace(model.LastName)
	model.DisplayName = strings.TrimSpace(model.DisplayName)
	model.Email = normalizeEmail(model.Email)

	if model.FirstName == "" ||
		model.LastName == "" ||
		model.DisplayName == "" ||
		model.Email == "" ||
		model.Password == "" {
		return nil, apperrors.BadRequest("INVALID_USER_INPUT", "invalid user input")
	}

	hashedPassword, err := hashPassword(model.Password)
	if err != nil {
		return nil, apperrors.Internal("HASH_PASSWORD_FAILED", "hash password failed", err)
	}

	user := &User{
		FirstName:       model.FirstName,
		LastName:        model.LastName,
		DisplayName:     model.DisplayName,
		Email:           model.Email,
		HashedPassword:  hashedPassword,
		IsEmailVerified: model.IsEmailVerified,
		IsActive:        model.IsActive,
		AvatarURL:       model.AvatarURL,
	}

	if err := s.userRepository.Create(ctx, user); err != nil {
		return nil, apperrors.Conflict("CREATE_USER_CONFLICT", "registration failed, please try again later")
	}

	return &CreateUserResponseModel{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

func (s *userService) FindByID(ctx context.Context, id uuid.UUID) (*GetUserResponseModel, *apperrors.AppError) {
	if id == uuid.Nil {
		return nil, apperrors.BadRequest("INVALID_USER_INPUT", "invalid user input")
	}

	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, apperrors.Internal("FIND_USER_ERROR", "failed to find user", err)
	}

	if user == nil {
		return nil, apperrors.NotFound("USER_NOT_FOUND", "user not found")
	}

	return &GetUserResponseModel{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		AvatarURL:   user.AvatarURL,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

func (s *userService) Update(ctx context.Context, id uuid.UUID, model UpdateUserRequestModel) (*UpdateUserResponseModel, *apperrors.AppError) {
	if id == uuid.Nil {
		return nil, apperrors.BadRequest("INVALID_USER_INPUT", "invalid user input")
	}

	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, apperrors.Internal("FIND_USER_ERROR", "failed to find user", err)
	}

	if user == nil {
		return nil, apperrors.NotFound("USER_NOT_FOUND", "user not found")
	}

	model.FirstName = strings.TrimSpace(model.FirstName)
	model.LastName = strings.TrimSpace(model.LastName)
	model.DisplayName = strings.TrimSpace(model.DisplayName)

	if model.FirstName == "" ||
		model.LastName == "" ||
		model.DisplayName == "" {
		return nil, apperrors.BadRequest("INVALID_USER_INPUT", "invalid user input")
	}

	user.FirstName = model.FirstName
	user.LastName = model.LastName
	user.DisplayName = model.DisplayName
	user.AvatarURL = model.AvatarURL

	if err := s.userRepository.Update(ctx, user); err != nil {
		return nil, apperrors.Internal("UPDATE_USER_ERROR", "failed to update user", err)
	}

	return &UpdateUserResponseModel{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		DisplayName: user.DisplayName,
		Email:       user.Email,
		AvatarURL:   user.AvatarURL,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	if id == uuid.Nil {
		return apperrors.BadRequest("INVALID_USER_INPUT", "invalid user input")
	}

	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return apperrors.Internal("FIND_USER_ERROR", "failed to find user", err)
	}

	if user == nil {
		return apperrors.NotFound("USER_NOT_FOUND", "user not found")
	}

	s.userRepository.Delete(ctx, id)
	return nil
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedBytes), nil
}
