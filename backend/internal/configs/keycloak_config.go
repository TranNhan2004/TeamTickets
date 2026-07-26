package configs

import (
	"os"
)

type KeycloakConfig struct {
	BaseURL               string
	PublicURL             string
	Realm                 string
	ClientID              string
	ClientSecret          string
	RedirectURI           string
	PostLogoutRedirectURI string
	FrontendURL           string
}

func LoadKeycloakConfig() (*KeycloakConfig, error) {
	return &KeycloakConfig{
		BaseURL:               os.Getenv("KEYCLOAK_BASE_URL"),
		PublicURL:             os.Getenv("KEYCLOAK_PUBLIC_URL"),
		Realm:                 os.Getenv("KEYCLOAK_REALM"),
		ClientID:              os.Getenv("KEYCLOAK_CLIENT_ID"),
		ClientSecret:          os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		RedirectURI:           os.Getenv("KEYCLOAK_REDIRECT_URI"),
		PostLogoutRedirectURI: os.Getenv("KEYCLOAK_POST_LOGOUT_REDIRECT_URI"),
		FrontendURL:           os.Getenv("FRONTEND_URL"),
	}, nil
}
