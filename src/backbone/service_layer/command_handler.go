package service_layer

import (
	"fmt"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/container"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/unit_of_work"
	"reflect"
)

type MessageBus struct {
	CommandHandlers map[string]func(message interface{}) error
	Uow             *unit_of_work.UnitOfWork
}

func NewMessageBus(uow *unit_of_work.UnitOfWork, commandHandlers map[string]func(message interface{}) error) *MessageBus {
	return &MessageBus{CommandHandlers: commandHandlers, Uow: uow}
}

func (m *MessageBus) Handle(message interface{}) error {
	commandName := reflect.TypeOf(message).Elem().Name()
	handler, ok := m.CommandHandlers[commandName]
	if !ok {
		container.Logger.Error(fmt.Sprint(commandName, " Command Not Register"))
		panic("Command Not Register")
	} else {
		container.Logger.Info(fmt.Sprint(commandName, " Command Handling"))
		return handler(message)
	}
}

func InjectedCommandHandlers(constructor interface{}, dependency map[string]interface{}, commands []interface{}) (string, func(message interface{}) error) {
	var commandNames string
	commandMap := make(map[string]interface{})
	c := reflect.TypeOf(constructor)
	numOut := c.NumIn()
	usCase := reflect.ValueOf(constructor)
	args := make([]reflect.Value, numOut)
	if numOut == 0 {
		panic("I Need At Least One Dependency")
	}
	for _, cmd := range commands {
		commandMap[reflect.TypeOf(cmd).Name()] = cmd
	}

	clearDependency := generateNewDependency(dependency)

	for i := 0; i < numOut; i++ {

		commandNames = generateArgs(c, i, clearDependency, &args, commandMap, commandNames)
	}

	if len(commandNames) <= 0 {
		panic("I Need One Command")

	}

	return commandNames, func(message interface{}) error {
		messageName := reflect.TypeOf(message).Elem().Name()
		if _, ok := commandMap[messageName]; !ok {
			panic(messageName + " Command Not Register")
		}
		args[0] = reflect.ValueOf(message)
		result := defaultInvoker(usCase, args)

		if len(result) > 0 && !result[0].IsNil() && result[0].CanInterface() {
			return result[0].Interface().(error)
		}
		return nil
	}
}

func generateArgs(c reflect.Type, i int, clearDependency map[string]interface{}, args *[]reflect.Value, commandMap map[string]interface{}, commandNames string) string {
	t := c.In(i)
	el := t.Elem()
	name := el.Name()

	arg := *args
	if v, ok := clearDependency[name]; ok {
		arg[i] = reflect.ValueOf(v)

	} else if v, ok = commandMap[name]; ok {
		commandNames = name
	} else {
		panic(name + " Command Not Register")
	}
	return commandNames
}

func generateNewDependency(dependency map[string]interface{}) map[string]interface{} {
	newNameDependency := make(map[string]interface{})
	for _, v := range dependency {
		newNameDependency[reflect.TypeOf(v).Elem().Name()] = v
	}

	return newNameDependency
}

func defaultInvoker(fn reflect.Value, args []reflect.Value) []reflect.Value {
	return fn.Call(args)

}
