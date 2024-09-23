package organization

import (
	"fmt"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/container"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/service_layer"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/unit_of_work"
	"github.com/ali-mahdavi-bn/service-site/src/organization/domain"
	"github.com/ali-mahdavi-bn/service-site/src/organization/service_layer/command_handlers"
)

func Bootstrap() *service_layer.MessageBus {
	dependency := map[string]interface{}{
		"uow": unit_of_work.NewUnitOfWork(container.DB),
	}
	commands := domain.CommandIndex
	handlers := command_handlers.HandlersIndex

	injected_command_handlers := make(map[string]func(message interface{}) error)
	for _, handler := range handlers {
		c, h := service_layer.InjectedCommandHandlers(handler, dependency, commands)
		injected_command_handlers[c] = h
	}
	fmt.Println(dependency["uow"])
	return service_layer.NewMessageBus(
		dependency["uow"].(*unit_of_work.UnitOfWork),
		injected_command_handlers,
	)

}
