package data_model

import (
	"github.com/ali-mahdavi-bn/service-site/src/organization/domain/entities"
)

type Batches struct {
	entities.Batch
	Allocations []OrderLines `gorm:"many2many:allocations;"`
}
