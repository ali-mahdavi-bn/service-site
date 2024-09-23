package api

import (
	"github.com/ali-mahdavi-bn/service-site/src/backbone/container"
	"net/http"
)

func ShouldBind(i interface{}) error {
	r := *Request

	if err := r.Bind(i); err != nil {
		container.Logger.Error(err)
		return r.String(http.StatusUnprocessableEntity, "bad request")
	}

	if err := r.Validate(i); err != nil {
		container.Logger.Error(err)
		return r.JSON(http.StatusUnprocessableEntity, err)
	}
	return nil
}
