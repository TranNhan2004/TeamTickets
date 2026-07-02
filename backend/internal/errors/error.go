package errors

import "net/http"

type AppError struct {
	Code       string
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

func New(code string, message string, statusCode int) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
	}
}

func BadRequest(code string, message string) *AppError {
	return New(code, message, http.StatusBadRequest)
}

func Unauthorized(code string, message string) *AppError {
	return New(code, message, http.StatusUnauthorized)
}

func Forbidden(code string, message string) *AppError {
	return New(code, message, http.StatusForbidden)
}

func NotFound(code string, message string) *AppError {
	return New(code, message, http.StatusNotFound)
}

func Conflict(code string, message string) *AppError {
	return New(code, message, http.StatusConflict)
}

func Internal(code string, message string, err error) *AppError {
	return &AppError{
		Code:       code,
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Err:        err,
	}
}

func WithDetails(err *AppError, details any) *AppError {
	err.Details = details
	return err
}
