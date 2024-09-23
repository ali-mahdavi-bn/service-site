package container

import "github.com/labstack/echo/v4"

var (
	Logger echo.Logger
)

func NewLogger(l echo.Logger) {
	Logger = l
}
