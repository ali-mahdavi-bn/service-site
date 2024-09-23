package ali

import (
	"fmt"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/api"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/unit_of_work"
	"github.com/ali-mahdavi-bn/service-site/src/organization/adapter/data_model"
	"github.com/ali-mahdavi-bn/service-site/src/organization/domain/commands"
	"net/http"
)

func CreateAliHandler(command *commands.AliMahCommand, uow *unit_of_work.UnitOfWork) error {
	fmt.Println("command: ", uow)
	var user *data_model.Batches

	uow.Transaction(func(u *unit_of_work.UnitOfWork) error {
		//user = u.User.Find(10)
		//u.DB.Delete(&user)
		return nil
	})
	fmt.Println("user: ", user)
	fmt.Println("unit_of_work: ", "unit_of_work")
	return api.BaseResponse(http.StatusOK, map[string]interface{}{"LL": "kkllk::::::::"})
}
