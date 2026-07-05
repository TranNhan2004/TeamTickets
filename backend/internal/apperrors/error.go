package apperrors

import "net/http"

type ErrorCode string

type AppError struct {
	Code       ErrorCode
	Message    string
	StatusCode int
	Details    any
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}

	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(code ErrorCode, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

func BadRequest(code ErrorCode, message string) *AppError {
	return New(code, message, http.StatusBadRequest)
}

func Unauthorized(code ErrorCode, message string) *AppError {
	return New(code, message, http.StatusUnauthorized)
}

func Forbidden(code ErrorCode, message string) *AppError {
	return New(code, message, http.StatusForbidden)
}

func NotFound(code ErrorCode, message string) *AppError {
	return New(code, message, http.StatusNotFound)
}

func Conflict(code ErrorCode, message string) *AppError {
	return New(code, message, http.StatusConflict)
}

func Internal(code ErrorCode, message string) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}
}

func (err *AppError) WithError(e error) *AppError {
	err.Err = e
	return err
}

func (err *AppError) WithDetails(details any) *AppError {
	err.Details = details
	return err
}
