package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trannhanlv2004/team-tickets/internal/apperrors"
	"github.com/trannhanlv2004/team-tickets/internal/configs"
	"github.com/trannhanlv2004/team-tickets/internal/features/auth"
)

func CSRF(cfg *configs.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		if isSafeMethod(c.Request.Method) {
			c.Next()
			return
		}

		cookieToken, err := c.Cookie(cfg.HTTP.CSRFTokenCookieName)
		if err != nil || cookieToken == "" {
			missingCSRFTokenCookieError := auth.NewErrorMissingCSRFTokenCookie()
			apperrors.RespondAbort(c, missingCSRFTokenCookieError)
			return
		}

		headerToken := c.GetHeader(cfg.HTTP.CSRFTokenHeaderName)
		if headerToken == "" {
			missingCSRFTokenHeaderError := auth.NewErrorMissingCSRFTokenHeader()
			apperrors.RespondAbort(c, missingCSRFTokenHeaderError)
			return
		}

		if headerToken != cookieToken {
			invalidCSRFTokenError := auth.NewErrorInvalidCSRFToken()
			apperrors.RespondAbort(c, invalidCSRFTokenError)
			return
		}

		c.Next()
	}
}

func isSafeMethod(method string) bool {
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		return true
	default:
		return false
	}
}
