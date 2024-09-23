package app

import (
	"github.com/ali-mahdavi-bn/service-site/src/organization/entrypoints/project"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Group) {
	// Project
	project.InitRoute(e)
}
