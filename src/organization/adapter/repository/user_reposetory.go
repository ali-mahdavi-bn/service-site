package repository

import (
	"github.com/ali-mahdavi-bn/service-site/src/backbone/adapter"
	"github.com/ali-mahdavi-bn/service-site/src/organization/domain/entities"
	"gorm.io/gorm"
)

type IUserRepository interface {
	adapter.IRepository[*entities.OrderLine]
}

type UserRepository struct {
	adapter.IRepository[*entities.OrderLine]
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{adapter.NewBaseRepository[*entities.OrderLine](db)}
}
