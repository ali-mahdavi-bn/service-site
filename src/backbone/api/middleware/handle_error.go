package middleware

import (
	"fmt"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/api/traslator"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/container"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

var (
	ErrorMap           = make(map[string]interface{})
	key, lang, subject string
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := RunNext(c, next)
		if ep := c.Get("ErrorPanic"); ep != nil {
			errorEp := ep.(error)
			epp := strings.Split(errorEp.Error(), ":")
			if len(epp) == 4 {
				key = epp[0]
				subject = epp[1]
				lang = epp[2]
				statusCode, _ := strconv.Atoi(epp[3])
				if v, ok := traslator.MapTranslate[subject][lang][key]; ok {
					return c.JSON(statusCode, map[string]string{
						"message": v,
					})
				}
			}
			container.Logger.Error(ep)

			c.Error(errorEp)
		}

		return err
	}
}

func RunNext(c echo.Context, next echo.HandlerFunc) error {
	defer func() {
		if r := recover(); r != nil {
			err, ok := r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
			c.Set("ErrorPanic", err)
		}
	}()
	return next(c)
}
