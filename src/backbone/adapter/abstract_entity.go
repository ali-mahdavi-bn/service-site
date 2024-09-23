package adapter

import (
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	gorm.Model
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type IBaseEntity interface{}
