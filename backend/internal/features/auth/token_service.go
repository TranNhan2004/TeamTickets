package auth

import (
	"github.com/trannhanlv2004/team-tickets/internal/apperrors"
	"github.com/trannhanlv2004/team-tickets/internal/features/shared"
)

type TokenVerifier interface {
	VerifyAccessToken(token string) (*shared.CurrentUserModel, *apperrors.AppError)
}

type TokenIssuer interface {
	GenerateAccessToken(model GenerateAccessTokenRequestModel) (string, *apperrors.AppError)
	GenerateRefreshToken(model GenerateRefreshTokenRequestModel) (string, *apperrors.AppError)
}

type TokenService interface {
	TokenVerifier
	TokenIssuer
}