package src

import (
	customeMiddleware "github.com/ali-mahdavi-bn/service-site/src/backbone/api/middleware"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func App(e *echo.Echo) func() {
	e.Use(middleware.Logger())
	e.Use(customeMiddleware.AddRequestContextMiddleware)
	e.Use(customeMiddleware.ErrorHandler)
	e.Use(middleware.RemoveTrailingSlash())

	baseUrl := e.Group("/api/v1")

	app.InitRoutes(baseUrl)

	return func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server", err)
		}
	}
}
