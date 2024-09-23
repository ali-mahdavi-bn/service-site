package application

type LineSerializer struct {
	Orderid string `json:"orderid" form:"orderid"`
	Sku     string `json:"sku" form:"sku"`
	Qty     int    `json:"qty" form:"qty"`
}

//func (h *Handler) AllocateEndpoint(c echo.Context) error {
//
//
//	//return c.JSON(http.StatusCreated, map[string]string{"batchref": "batchref"})
//}
