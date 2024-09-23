package unit_of_work

import (
	"github.com/ali-mahdavi-bn/service-site/src/backbone/container"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/unit_of_work/reposetory"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	reposetory.Repositories
	DB *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	return &UnitOfWork{Repositories: reposetory.NewRepositories(db), DB: db}
}

func (u *UnitOfWork) Transaction(fn func(uow *UnitOfWork) error) {
	db := u.DB.Begin()
	if err := fn(u); err != nil {
		container.Logger.Error(err)
		db.Rollback()
	} else {
		defer db.Commit()
	}
}
