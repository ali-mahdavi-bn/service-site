package application

import (
	"github.com/ali-mahdavi-bn/service-site/src/organization/adapter/repository"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func (h *Handler) InitRepo(ctx echo.Context) repository.Repository {
	db := ctx.Get("db")
	if db == nil {
		db = h.DB
	}
	return &repository.gormUserRepository{DB: db.(*gorm.DB)}
}
