package initialize

import (
	"crypto-dashboard/gw-example/infrastructure/core"
	"crypto-dashboard/gw-example/infrastructure/global"
	"crypto-dashboard/pkg/appLogger"
	"crypto-dashboard/pkg/common"
	"crypto-dashboard/pkg/database/caching"
	database "crypto-dashboard/pkg/database/postgres"

	"crypto-dashboard/pkg/response"
)

func must[T any](value T, err *response.AppError) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Run() {
	global.AppConfig = must(common.LoadConfig[global.Config]())
	global.Log = must(appLogger.NewLogger(global.AppConfig.Logger, global.AppConfig.Server))
	global.Cache = must(caching.NewRedisClient(global.AppConfig.Cache))
	global.SQLDB = must(database.NewConnection(global.AppConfig.SQL))
	http := must(core.NewHttpServer())
	global.App = http.App

	http.Start()
}
