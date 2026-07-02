package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Env  string
	Port string
	DB   *DBConfig
}

func LoadAppConfig() (*AppConfig, error) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	_ = godotenv.Load("env/.env")

	envFile := fmt.Sprintf("env/.env.%s", appEnv)
	if err := godotenv.Overload(envFile); err != nil {
		return nil, fmt.Errorf("failed to load env file %s: %w", envFile, err)
	}

	dbConfig, err := LoadDBConfig()
	if err != nil {
		return nil, err
	}

	return &AppConfig{
		Env:  appEnv,
		Port: os.Getenv("APP_PORT"),
		DB:   dbConfig,
	}, nil
}
