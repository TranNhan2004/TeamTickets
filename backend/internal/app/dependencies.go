package app

import (
	"github.com/trannhanlv2004/team-tickets/internal/health"
	"gorm.io/gorm"
)

type Dependencies struct {
	HealthHandler *health.Handler
}

func NewDependencies(db *gorm.DB) *Dependencies {
	healthDatabaseChecker := health.NewDatabaseChecker(db)
	healthService := health.NewService(healthDatabaseChecker)
	healthHandler := health.NewHandler(healthService)

	return &Dependencies{
		HealthHandler: healthHandler,
	}
}
