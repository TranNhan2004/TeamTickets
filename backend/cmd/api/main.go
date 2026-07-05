package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/trannhanlv2004/team-tickets/internal/app"
	"github.com/trannhanlv2004/team-tickets/internal/configs"
	"github.com/trannhanlv2004/team-tickets/internal/logger"
	"github.com/trannhanlv2004/team-tickets/internal/routes"
)

func main() {
	appConfig, err := configs.LoadAppConfig()
	if err != nil {
		panic(err)
	}

	appLogger, cleanup, err := logger.New(appConfig.LogDir)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	appLogger.App.Printf(
		"App running in %s mode on port %s",
		appConfig.Env,
		appConfig.Port,
	)

	r := gin.Default()

	dependencies := app.NewDependencies(appConfig.DB, appLogger)
	routes.RegisterV1(r, dependencies)

	if err := r.Run(":" + appConfig.Port); err != nil {
		appLogger.Error.Printf("failed to run server: %v", err)
		panic(err)
	}
}
