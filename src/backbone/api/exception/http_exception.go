package exception

import (
	"github.com/ali-mahdavi-bn/service-site/src/backbone/container"
	"net/http"
	"strconv"
)

func NotFoundException(key string, lang string) {
	status := http.StatusNotFound
	subject := "error"

	if lang == "" {
		lang = "fa"
	}
	message := generateErrorMessage(key, lang, subject, status)
	panic(message)
}

func BadRequestException(key string, lang string) {
	status := http.StatusBadRequest
	subject := "error"

	if lang == "" {
		lang = "fa"
	}

	message := generateErrorMessage(key, lang, subject, status)
	panic(message)
}

func UnauthorizedException(key string, lang string) {
	status := http.StatusUnauthorized
	subject := "error"

	if lang == "" {
		lang = "fa"
	}

	message := generateErrorMessage(key, lang, subject, status)
	panic(message)
}

func UnsupportedMediaTypeException(key string, lang string) {
	status := http.StatusUnsupportedMediaType
	subject := "error"
	if lang == "" {
		lang = "fa"
	}
	message := generateErrorMessage(key, lang, subject, status)
	panic(message)
}

func ForbiddenException(key string, lang string) {
	status := http.StatusForbidden
	subject := "error"

	if lang == "" {
		lang = "fa"
	}

	message := generateErrorMessage(key, lang, subject, status)
	panic(message)
}

func ConflictException(key string, lang string) {
	status := http.StatusConflict
	subject := "error"

	if lang == "" {
		lang = "fa"
	}

	message := generateErrorMessage(key, lang, subject, status)
	panic(message)
}

func InternalServerException(key string, lang string) {
	status := http.StatusInternalServerError
	subject := "error"

	if lang == "" {
		lang = "fa"
	}
	message := generateErrorMessage(key, lang, subject, status)
	panic(message)
}
func generateErrorMessage(key string, lang string, subject string, status int) string {
	message := key + ":" + subject + ":" + lang + ":" + strconv.Itoa(status)
	container.Logger.Error(message)
	return message
}
