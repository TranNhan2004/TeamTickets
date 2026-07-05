package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GenerateAccessTokenRequestModel struct {
	UserID string
	Email  string
}

type GenerateRefreshTokenRequestModel struct {
	UserID    string
	SessionID string
}

type AccessTokenClaims struct {
	Email string `json:"email"`

	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	SessionID string `json:"sid"`

	jwt.RegisteredClaims
}

type TokenPairModel struct {
	AccessToken            string
	RefreshToken           string
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
}
