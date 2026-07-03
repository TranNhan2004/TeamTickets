package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/trannhanlv2004/team-tickets/internal/app"
	"github.com/trannhanlv2004/team-tickets/internal/features/users"
	"github.com/trannhanlv2004/team-tickets/internal/health"
)

func RegisterV1(engine *gin.Engine, deps *app.Dependencies) {
	v1 := engine.Group("/api/v1")

	health.RegisterRoutes(v1, deps.HealthHandler)
	users.RegisterRoutes(v1, deps.UserHandler)
}
