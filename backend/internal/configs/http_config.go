package configs

import "os"

type HTTPConfig struct {
	AccessTokenCookieName  string
	RefreshTokenCookieName string
	CSRFTokenCookieName    string
	CSRFTokenHeaderName    string
}

func LoadHTTPConfig() (*HTTPConfig, error) {
	return &HTTPConfig{
		AccessTokenCookieName:  os.Getenv("ACCESS_TOKEN_COOKIE_NAME"),
		RefreshTokenCookieName: os.Getenv("REFRESH_TOKEN_COOKIE_NAME"),
		CSRFTokenCookieName:    os.Getenv("CSRF_TOKEN_COOKIE_NAME"),
		CSRFTokenHeaderName:    os.Getenv("CSRF_TOKEN_HEADER_NAME"),
	}, nil
}
