package auth

import "github.com/trannhanlv2004/team-tickets/internal/apperrors"

const (
	// Auth + JWT
	ErrUnauthorized        apperrors.ErrorCode = "AUTH_UNAUTHORIZED"
	ErrMissingAccessToken  apperrors.ErrorCode = "AUTH_MISSING_ACCESS_TOKEN"
	ErrInvalidAccessToken  apperrors.ErrorCode = "AUTH_INVALID_ACCESS_TOKEN"
	ErrInvalidTokenSubject apperrors.ErrorCode = "AUTH_INVALID_TOKEN_SUBJECT"
	ErrInvalidSession      apperrors.ErrorCode = "AUTH_INVALID_SESSION"
	ErrTokenSignedFailed   apperrors.ErrorCode = "AUTH_TOKEN_SIGNED_FAILED"

	// CSRF
	ErrMissingCSRFTokenHeader apperrors.ErrorCode = "AUTH_MISSING_CSRF_TOKEN_HEADER"
	ErrMissingCSRFTokenCookie apperrors.ErrorCode = "AUTH_MISSING_CSRF_TOKEN_COOKIE"
	ErrInvalidCSRFToken       apperrors.ErrorCode = "AUTH_INVALID_CSRF_TOKEN"
)

func NewErrorMissingAccessToken() *apperrors.AppError {
	return apperrors.Unauthorized(ErrMissingAccessToken, "missing access token")
}

func NewErrorInvalidAccessToken() *apperrors.AppError {
	return apperrors.Unauthorized(ErrInvalidAccessToken, "invalid or expired access token")
}

func NewErrorUnauthorized() *apperrors.AppError {
	return apperrors.Unauthorized(ErrUnauthorized, "unauthorized")
}

func NewErrorInvalidTokenSubject() *apperrors.AppError {
	return apperrors.Unauthorized(ErrInvalidTokenSubject, "invalid token subject")
}

func NewErrorInvalidSession() *apperrors.AppError {
	return apperrors.Unauthorized(ErrInvalidSession, "invalid session")
}

func NewErrorTokenSignFailed() *apperrors.AppError {
	return apperrors.Internal(ErrTokenSignedFailed, "failed to sign token")
}

func NewErrorMissingCSRFTokenHeader() *apperrors.AppError {
	return apperrors.Unauthorized(ErrMissingCSRFTokenHeader, "missing CSRF token header")
}

func NewErrorMissingCSRFTokenCookie() *apperrors.AppError {
	return apperrors.Unauthorized(ErrMissingCSRFTokenCookie, "missing CSRF token cookie")
}

func NewErrorInvalidCSRFToken() *apperrors.AppError {
	return apperrors.Unauthorized(ErrInvalidCSRFToken, "invalid CSRF token")
}
