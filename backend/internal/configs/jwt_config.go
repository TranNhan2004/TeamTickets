package configs

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type JWTConfig struct {
	Issuer                 string
	Audience               []string
	AccessTokenSecret      string
	RefreshTokenSecret     string
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
}

func LoadJWTConfig() (*JWTConfig, error) {
	jwtAccessTokenExpiration, err := strconv.Atoi(os.Getenv("JWT_ACCESS_TOKEN_EXPIRATION"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_ACCESS_TOKEN_EXPIRATION: %w", err)
	}

	jwtRefreshTokenExpiration, err := strconv.Atoi(os.Getenv("JWT_REFRESH_TOKEN_EXPIRATION"))
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_REFRESH_TOKEN_EXPIRATION: %w", err)
	}

	return &JWTConfig{
		Issuer:                 os.Getenv("JWT_ISSUER"),
		Audience:               strings.Split(os.Getenv("JWT_AUDIENCE"), ","),
		AccessTokenSecret:      os.Getenv("JWT_ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret:     os.Getenv("JWT_REFRESH_TOKEN_SECRET"),
		AccessTokenExpiration:  time.Duration(jwtAccessTokenExpiration) * time.Second,
		RefreshTokenExpiration: time.Duration(jwtRefreshTokenExpiration) * time.Second,
	}, nil
}
