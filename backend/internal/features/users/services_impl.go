package users

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/trannhanlv2004/team-tickets/internal/apperrors"
	"github.com/trannhanlv2004/team-tickets/internal/features/shared"
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
	audit := shared.AuditModel{
		AuditTime: time.Now().UTC(),
	}

	model.FirstName = strings.TrimSpace(model.FirstName)
	model.LastName = strings.TrimSpace(model.LastName)
	model.DisplayName = strings.TrimSpace(model.DisplayName)
	model.Email = normalizeEmail(model.Email)

	if model.FirstName == "" ||
		model.LastName == "" ||
		model.DisplayName == "" ||
		model.Email == "" ||
		model.Password == "" {
		return nil, NewErrorInvalidInput()
	}

	hashedPassword, err := hashPassword(model.Password)
	if err != nil {
		return nil, apperrors.Internal("HASH_PASSWORD_FAILED", "hash password failed")
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

	if err := s.userRepository.Create(ctx, user, &audit); err != nil {
		return nil, NewErrorCreateFailed(err)
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
		return nil, NewErrorInvalidInput()
	}

	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, NewErrorGetFailed(err)
	}

	if user == nil {
		return nil, NewErrorNotFound()
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
	audit := shared.AuditModel{
		AuditTime: time.Now().UTC(),
	}

	if id == uuid.Nil {
		return nil, NewErrorInvalidInput()
	}

	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, NewErrorGetFailed(err)
	}

	if user == nil {
		return nil, NewErrorNotFound()
	}

	model.FirstName = strings.TrimSpace(model.FirstName)
	model.LastName = strings.TrimSpace(model.LastName)
	model.DisplayName = strings.TrimSpace(model.DisplayName)

	if model.FirstName == "" ||
		model.LastName == "" ||
		model.DisplayName == "" {
		return nil, NewErrorInvalidInput()
	}

	user.FirstName = model.FirstName
	user.LastName = model.LastName
	user.DisplayName = model.DisplayName
	user.AvatarURL = model.AvatarURL

	if err := s.userRepository.Update(ctx, user, &audit); err != nil {
		return nil, NewErrorUpdateFailed(err)
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
	audit := shared.AuditModel{
		AuditTime: time.Now().UTC(),
	}

	if id == uuid.Nil {
		return NewErrorInvalidInput()
	}

	user, err := s.userRepository.FindByID(ctx, id)
	if err != nil {
		return NewErrorGetFailed(err)
	}

	if user == nil {
		return NewErrorNotFound()
	}

	s.userRepository.Delete(ctx, id, &audit)
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
