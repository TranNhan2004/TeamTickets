package app

import (
	"github.com/trannhanlv2004/team-tickets/internal/configs"
	"github.com/trannhanlv2004/team-tickets/internal/database"
	"github.com/trannhanlv2004/team-tickets/internal/features/users"
	"github.com/trannhanlv2004/team-tickets/internal/health"
	"github.com/trannhanlv2004/team-tickets/internal/logger"
)

type Dependencies struct {
	HealthHandler *health.Handler
	UserHandler   *users.UserHandler
}

func NewDependencies(dbConfig *configs.DBConfig, logger *logger.Logger) *Dependencies {
	db, err := database.ConnectPostgres(*dbConfig)
	if err != nil {
		logger.Error.Printf("failed to connect postgres: %v", err)
		panic(err)
	}

	// txManager := shared.NewTransactionManager(db)

	healthDatabaseChecker := health.NewDatabaseChecker(db)
	healthService := health.NewService(healthDatabaseChecker)
	healthHandler := health.NewHandler(healthService)

	userRepository := users.NewUserRepository(db)
	userService := users.NewUserService(userRepository)
	userHandler := users.NewUserHandler(userService)

	return &Dependencies{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
	}
}
