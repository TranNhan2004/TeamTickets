package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trannhanlv2004/team-tickets/internal/configs"
)

func GenerateCSRFToken() (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}

func SetAuthCookies(
	c *gin.Context,
	cfg *configs.AppConfig,
	accessToken string,
	refreshToken string,
) error {
	csrfToken, err := GenerateCSRFToken()
	if err != nil {
		return err
	}

	secure := cfg.Env == "production"

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     cfg.HTTP.AccessTokenCookieName,
		Value:    accessToken,
		Path:     "/",
		MaxAge:   15 * 60,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     cfg.HTTP.RefreshTokenCookieName,
		Value:    refreshToken,
		Path:     "/api/v1/auth/refresh",
		MaxAge:   7 * 24 * 60 * 60,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     cfg.HTTP.CSRFTokenCookieName,
		Value:    csrfToken,
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60,
		HttpOnly: false,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})

	return nil
}

func ClearAuthCookies(c *gin.Context, cfg *configs.AppConfig) {
	secure := cfg.Env == "production"

	clearCookie(c, cfg.HTTP.AccessTokenCookieName, "/", true, secure)
	clearCookie(c, cfg.HTTP.RefreshTokenCookieName, "/api/v1/auth/refresh", true, secure)
	clearCookie(c, cfg.HTTP.CSRFTokenCookieName, "/", false, secure)
}

func clearCookie(
	c *gin.Context,
	name string,
	path string,
	httpOnly bool,
	secure bool,
) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     path,
		MaxAge:   -1,
		HttpOnly: httpOnly,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}
