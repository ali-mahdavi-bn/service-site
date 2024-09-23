package application

import (
	"github.com/ali-mahdavi-bn/service-site/src/organization/domain/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type BatchSerializer struct {
	Reference         string    `json:"reference" form:"reference"`
	SKU               string    `json:"sku" form:"sku"`
	PurchasedQuantity int       `json:"purchased_quantity" form:"purchased_quantity"`
	ETA               time.Time `json:"eta" form:"eta"`
}

func (h *Handler) AddBatch(c echo.Context) error {
	repo := h.InitRepo(c)

	data := new(BatchSerializer)
	if err := c.Bind(data); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Logger().Errorf("%+v", data)

	batchId := services.AddBatch(data.Reference, data.SKU, data.PurchasedQuantity, data.ETA, repo)
	return c.JSON(http.StatusCreated, map[string]uint{"id": batchId})
}
