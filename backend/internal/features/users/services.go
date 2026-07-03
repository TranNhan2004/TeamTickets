package users

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/trannhanlv2004/team-tickets/internal/apperrors"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidUserInput  = errors.New("invalid user input")
	ErrEmailAlreadyInUse = errors.New("email already in use")
)

type UserService interface {
	Create(ctx context.Context, model CreateUserRequestModel) (*CreateUserResponseModel, *apperrors.AppError)
	FindByID(ctx context.Context, id uuid.UUID) (*GetUserResponseModel, *apperrors.AppError)
	Update(ctx context.Context, id uuid.UUID, model UpdateUserRequestModel) (*UpdateUserResponseModel, *apperrors.AppError)
	Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError
}
