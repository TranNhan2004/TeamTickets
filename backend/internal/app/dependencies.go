package app

import (
	"github.com/trannhanlv2004/team-tickets/internal/features/users"
	"github.com/trannhanlv2004/team-tickets/internal/health"
	"gorm.io/gorm"
)

type Dependencies struct {
	HealthHandler *health.Handler
	UserHandler   *users.UserHandler
}

func NewDependencies(db *gorm.DB) *Dependencies {
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
