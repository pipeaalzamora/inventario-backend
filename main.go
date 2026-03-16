package main

import (
	"fmt"
	"sofia-backend/api/v1/server"
	"sofia-backend/config"
	"sofia-backend/database"
	"sofia-backend/domain/external"
	"sofia-backend/domain/facades"
	"sofia-backend/domain/services"
	"sofia-backend/infraestructure/data"
	externalservices "sofia-backend/infraestructure/external-services"
)

func main() {
	//appCtx := context.Background()
	appCfg := config.Get()

	// mongoDb := database.NewMongo(appCfg)
	redisDb := database.NewRedis(appCfg)
	redisDb.Conn()
	postgresDb := database.NewPostgres(appCfg)

	// Infra
	appExternalServices := externalservices.Build(appCfg, redisDb)
	appData := data.Build(postgresDb)

	// Domain
	appService := services.Build(appCfg, appData, appExternalServices)
	appExternal := external.Build(appExternalServices, appCfg)
	appFacade := facades.Build(appCfg, appService, appExternal)

	// Api
	apiServer := server.NewV1Server(appCfg, appFacade)
	// Start the API server
	apiServer.Run(fmt.Sprintf(":%d", appCfg.Port))
}
