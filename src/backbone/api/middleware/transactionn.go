package middleware

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) Transaction(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db := h.DB.Begin()
		c.Set("db", db)
		if err := next(c); err != nil {
			c.Error(err)
			db.Rollback()
		} else {
			defer db.Commit()
		}

		return nil
	}
}
