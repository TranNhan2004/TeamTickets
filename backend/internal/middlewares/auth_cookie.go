package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/trannhanlv2004/team-tickets/internal/apperrors"
	"github.com/trannhanlv2004/team-tickets/internal/configs"
	"github.com/trannhanlv2004/team-tickets/internal/features/auth"
	"github.com/trannhanlv2004/team-tickets/internal/features/shared"
)

func CookieAuth(verifier auth.TokenVerifier, cfg *configs.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie(cfg.HTTP.AccessTokenCookieName)
		if err != nil || accessToken == "" {
			missingAccessTokenError := auth.NewErrorMissingAccessToken()
			apperrors.RespondAbort(c, missingAccessTokenError)
			return
		}

		authUser, appErr := verifier.VerifyAccessToken(accessToken)
		if appErr != nil {
			apperrors.RespondAbort(c, appErr)
			return
		}

		shared.SetCurrentUser(c, authUser)
		c.Next()
	}
}
