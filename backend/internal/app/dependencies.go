package app

import (
	"github.com/gin-gonic/gin"
	"github.com/trannhanlv2004/team-tickets/internal/configs"
	"github.com/trannhanlv2004/team-tickets/internal/database"
	"github.com/trannhanlv2004/team-tickets/internal/features/users"
	"github.com/trannhanlv2004/team-tickets/internal/health"
	"github.com/trannhanlv2004/team-tickets/internal/logger"
)

type Dependencies struct {
	HealthHandler *health.Handler
	UserHandler   *users.UserHandler

	AuthMiddleware gin.HandlerFunc
	CSRFMiddleware gin.HandlerFunc
}

func NewDependencies(cfg *configs.AppConfig, logger *logger.Logger) *Dependencies {
	db, err := database.ConnectPostgres(*cfg.DB)
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
