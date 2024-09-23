package command_handlers

import (
	"github.com/ali-mahdavi-bn/service-site/src/organization/service_layer/command_handlers/ali"
)

var (
	HandlersIndex = []interface{}{
		ali.CreateAliHandler,
	}
)
