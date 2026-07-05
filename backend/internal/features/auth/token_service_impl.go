package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/trannhanlv2004/team-tickets/internal/apperrors"
	"github.com/trannhanlv2004/team-tickets/internal/configs"
	"github.com/trannhanlv2004/team-tickets/internal/features/shared"
)

type jwtTokenService struct {
	accessTokenSecret  []byte
	refreshTokenSecret []byte
	accessTokenTTL     time.Duration
	refreshTokenTTL    time.Duration
	issuer             string
	audience           []string
}

func NewJWTTokenService(cfg *configs.AppConfig) TokenService {
	return &jwtTokenService{
		accessTokenSecret:  []byte(cfg.JWT.AccessTokenSecret),
		refreshTokenSecret: []byte(cfg.JWT.RefreshTokenSecret),
		accessTokenTTL:     cfg.JWT.AccessTokenExpiration,
		refreshTokenTTL:    cfg.JWT.RefreshTokenExpiration,
		issuer:             cfg.JWT.Issuer,
		audience:           cfg.JWT.Audience,
	}
}

func (s *jwtTokenService) GenerateAccessToken(model GenerateAccessTokenRequestModel) (string, *apperrors.AppError) {
	userID, err := uuid.Parse(model.UserID)
	if err != nil {
		return "", NewErrorInvalidTokenSubject()
	}

	now := time.Now().UTC()

	claims := AccessTokenClaims{
		Email: model.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			Issuer:    s.issuer,
			Audience:  s.audience,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.accessTokenTTL)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(s.accessTokenSecret)
	if err != nil {
		return "", NewErrorTokenSignFailed()
	}

	return signedToken, nil
}

func (s *jwtTokenService) GenerateRefreshToken(model GenerateRefreshTokenRequestModel) (string, *apperrors.AppError) {
	userID, err := uuid.Parse(model.UserID)
	if err != nil {
		return "", NewErrorInvalidTokenSubject()
	}

	if model.SessionID == "" {
		return "", NewErrorInvalidSession()
	}

	now := time.Now().UTC()

	claims := RefreshTokenClaims{
		SessionID: model.SessionID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			Issuer:    s.issuer,
			Audience:  s.audience,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.refreshTokenTTL)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(s.refreshTokenSecret)
	if err != nil {
		return "", NewErrorTokenSignFailed()
	}

	return signedToken, nil
}

func (s *jwtTokenService) VerifyAccessToken(tokenString string) (*shared.CurrentUserModel, *apperrors.AppError) {
	if tokenString == "" {
		return nil, NewErrorInvalidAccessToken()
	}

	claims := &AccessTokenClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}

			return s.accessTokenSecret, nil
		},
		jwt.WithExpirationRequired(),
		jwt.WithIssuer(s.issuer),
		jwt.WithIssuedAt(),
	)

	if err != nil {
		return nil, NewErrorInvalidAccessToken()
	}

	if token == nil || !token.Valid {
		return nil, NewErrorInvalidAccessToken()
	}

	if !hasRequiredAudience(claims.Audience, s.audience) {
		return nil, NewErrorInvalidAccessToken()
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, NewErrorInvalidAccessToken()
	}

	if claims.Email == "" {
		return nil, NewErrorInvalidAccessToken()
	}

	return &shared.CurrentUserModel{
		UserID: userID,
		Email:  claims.Email,
	}, nil
}

func hasRequiredAudience(tokenAudiences jwt.ClaimStrings, requiredAudiences []string) bool {
	if len(requiredAudiences) == 0 {
		return true
	}

	if len(tokenAudiences) == 0 {
		return false
	}

	tokenAudienceSet := make(map[string]struct{}, len(tokenAudiences))
	for _, audience := range tokenAudiences {
		tokenAudienceSet[audience] = struct{}{}
	}

	for _, requiredAudience := range requiredAudiences {
		if _, exists := tokenAudienceSet[requiredAudience]; !exists {
			return false
		}
	}

	return true
}
