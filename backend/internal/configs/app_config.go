package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env                string
	Port               string
	LogDir             string
	CORSAllowedOrigins []string
	DB                 *DBConfig
	Keycloak           *KeycloakConfig
	CSRF               *CSRFConfig
}

func LoadAppConfig() (*AppConfig, error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	_ = godotenv.Load("env/.env.example")

	envFile := fmt.Sprintf("env/.env.%s", appEnv)
	if err := godotenv.Overload(envFile); err != nil {
		return nil, fmt.Errorf("failed to load env file %s: %w", envFile, err)
	}

	dbConfig, err := LoadDBConfig()
	if err != nil {
		return nil, err
	}

	keycloakConfig, err := LoadKeycloakConfig()
	if err != nil {
		return nil, err
	}

	csrfConfig, err := LoadCSRFConfig()
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		Env:                appEnv,
		Port:               os.Getenv("APP_PORT"),
		LogDir:             os.Getenv("LOG_DIR"),
		CORSAllowedOrigins: parseCommaSeparatedEnv(os.Getenv("CORS_ALLOWED_ORIGINS")),
		DB:                 dbConfig,
		Keycloak:           keycloakConfig,
		CSRF:               csrfConfig,
	}, nil
}
