package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"fetchlogger/apis"
	"fetchlogger/conf"
	"fetchlogger/middleware"
	"fetchlogger/services/logger_service"
)

func main() {

	e := echo.New()
	e.Use(middleware.RecoverMiddleware())

	config := conf.ReadConfig()

	loggerService := logger_service.InitLoggerService(config.Logger)

	loggerApiGroup := e.Group("/")
	apis.InitRoute(config, loggerService).RegisterRoute(loggerApiGroup)

	address := fmt.Sprintf(":%d", config.Server.Port)

	logrus.Info("Starting server at port: %s", address)
	e.Logger.Fatal(e.Start(address))
}
