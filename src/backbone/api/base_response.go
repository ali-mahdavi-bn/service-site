package api

func BaseResponse(code int, i interface{}) error {
	r := *Request
	return r.JSON(code, i)
}
