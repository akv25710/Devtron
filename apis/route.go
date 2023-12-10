package apis

import (
	"fetchlogger/conf"
	"fetchlogger/services/logger_service"
	"github.com/labstack/echo/v4"
)

type Route struct {
	Conf   conf.Configuration
	Logger logger_service.LoggerService
}

func (r Route) RegisterRoute(group *echo.Group) {
	group.POST("/getLogs", r.getLogs)
}

func InitRoute(conf conf.Configuration, logger logger_service.LoggerService) *Route {

	return &Route{
		Conf:   conf,
		Logger: logger,
	}
}
