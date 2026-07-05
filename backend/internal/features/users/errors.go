package users

import "github.com/trannhanlv2004/team-tickets/internal/apperrors"

const (
	ErrNotFound     apperrors.ErrorCode = "USER_NOT_FOUND"
	ErrEmailExisted apperrors.ErrorCode = "USER_EMAIL_EXISTED"
	ErrInvalidInput apperrors.ErrorCode = "USER_INVALID_INPUT"
	ErrCreateFailed apperrors.ErrorCode = "USER_CREATE_FAILED"
	ErrGetFailed    apperrors.ErrorCode = "USER_GET_FAILED"
	ErrUpdateFailed apperrors.ErrorCode = "USER_UPDATE_FAILED"
	ErrDeleteFailed apperrors.ErrorCode = "USER_DELETE_FAILED"
)

func NewErrorNotFound() *apperrors.AppError {
	return apperrors.NotFound(ErrNotFound, "user not found")
}

func NewErrorEmailExisted() *apperrors.AppError {
	return apperrors.Conflict(ErrEmailExisted, "email already exists")
}

func NewErrorInvalidInput() *apperrors.AppError {
	return apperrors.BadRequest(ErrInvalidInput, "invalid input")
}

func NewErrorCreateFailed(e error) *apperrors.AppError {
	return apperrors.
		Internal(ErrCreateFailed, "failed to create user").
		WithError(e)
}

func NewErrorGetFailed(e error) *apperrors.AppError {
	return apperrors.
		Internal(ErrGetFailed, "failed to get user").
		WithError(e)
}

func NewErrorUpdateFailed(e error) *apperrors.AppError {
	return apperrors.
		Internal(ErrUpdateFailed, "failed to update user").
		WithError(e)
}

func NewErrorDeleteFailed(e error) *apperrors.AppError {
	return apperrors.
		Internal(ErrDeleteFailed, "failed to delete user").
		WithError(e)
}
