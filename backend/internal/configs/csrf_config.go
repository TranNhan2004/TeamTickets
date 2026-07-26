package configs

import "os"

type CSRFConfig struct {
	TokenCookieName string
	TokenHeaderName string
}

func LoadCSRFConfig() (*CSRFConfig, error) {
	return &CSRFConfig{
		TokenCookieName: os.Getenv("CSRF_TOKEN_COOKIE_NAME"),
		TokenHeaderName: os.Getenv("CSRF_TOKEN_HEADER_NAME"),
	}, nil
}
