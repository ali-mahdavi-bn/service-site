package reposetory

import (
	organization_rep "github.com/ali-mahdavi-bn/service-site/src/organization/adapter/repository"
	"gorm.io/gorm"
)

type Repositories struct {
	User organization_rep.IUserRepository
}

func NewRepositories(DB *gorm.DB) Repositories {
	return Repositories{
		User: organization_rep.NewUserRepository(DB),
	}
}
