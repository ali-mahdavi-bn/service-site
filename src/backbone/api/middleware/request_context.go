package middleware

import (
	"github.com/labstack/echo/v4"
)

var (
	Request echo.Context
)

func AddRequestContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		Request = c
		err := next(c)
		return err
	}
}
