package adapter

import (
	"context"
	"gorm.io/gorm"
)

type IRepository[T IBaseEntity] interface {
	FindById(ctx context.Context, id uint) (T, error)
	FindByFiled(ctx context.Context, field string, value any) (T, error)
	Save(ctx context.Context, model T) error
}

type BaseRepository[T IBaseEntity] struct {
	db *gorm.DB
}

func NewBaseRepository[T IBaseEntity](db *gorm.DB) IRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (c *BaseRepository[T]) FindById(ctx context.Context, id uint) (T, error) {
	return c.FindByFiled(ctx, "id", id)
}

func (c *BaseRepository[T]) FindByFiled(ctx context.Context, field string, value any) (T, error) {
	var t T
	err := c.db.WithContext(ctx).Model(&t).Where(field+"=?", value).First(&t).Error
	return t, err
}

func (c *BaseRepository[T]) Save(ctx context.Context, model T) error {
	return c.db.WithContext(ctx).Model(&model).Save(model).Error
}
