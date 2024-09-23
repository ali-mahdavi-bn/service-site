package container

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	errorsMap := make(map[string]string)
	if err := cv.Validator.Struct(i); err != nil {
		for _, fe := range err.(validator.ValidationErrors) {
			errorsMap[fe.Field()] = fmt.Sprint("Field validation for '",
				fe.Field(),
				"' failed on the '",
				fe.Tag(),
				"' tag")
		}
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errorsMap)

	}

	return nil
}
