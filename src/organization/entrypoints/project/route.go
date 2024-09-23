package project

import (
	"github.com/ali-mahdavi-bn/service-site/src/backbone/api"
	"github.com/ali-mahdavi-bn/service-site/src/organization"
	"github.com/ali-mahdavi-bn/service-site/src/organization/service_layer/query/ali"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" form:"name" validate:"required,min=2"`
	Email string `json:"email" form:"email" validate:"required,email"`
}

func InitRoute(e *echo.Group) {
	mb := organization.Bootstrap()
	g := e.Group("/project")

	g.POST("/allocate", func(c echo.Context) error {
		//domain = c.QueryParam("domain")
		s := api.NewBaseSearch(c.QueryParam("domain"))

		//u := new(entrypoints.UserRequestModel)
		//if err := api.ShouldBind(u); err != nil {
		//	return err
		//}
		//
		//command := &commands.AliMahCommand{}
		//return mb.Handle(command)
		return ali.GetAllAllocate(mb.Uow, s)
	})
}
