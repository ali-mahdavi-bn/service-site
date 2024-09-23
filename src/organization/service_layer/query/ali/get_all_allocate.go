package ali

import (
	"fmt"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/api"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/helper/utils"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/unit_of_work"
	"github.com/ali-mahdavi-bn/service-site/src/organization/adapter/data_model"
	"net/http"
)

func GetAllAllocate(uow *unit_of_work.UnitOfWork, search *api.BaseSearch) error {
	var user *data_model.Batches
	a := *search
	fmt.Println("============================")
	a.Filter()
	fmt.Println("============================")
	uow.Transaction(func(u *unit_of_work.UnitOfWork) error {
		//user = u.User.FindById()
		//u.DB.Delete(&user)
		return nil
	})
	v, _ := utils.TypeConverter[map[string]interface{}](user)
	return api.BaseResponse(http.StatusOK, v)

}
