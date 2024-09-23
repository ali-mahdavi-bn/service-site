package e2e_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ali-mahdavi-bn/service-site/src"
	"github.com/ali-mahdavi-bn/service-site/src/organization/adapter/data_model"
	"github.com/ali-mahdavi-bn/service-site/test/e2e"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type E2ESuite struct {
	suite.Suite
	echo *echo.Echo
}

func (suite *E2ESuite) SetupSuite() {
	suite.echo = echo.New()
	suite.echo.Use(middleware.Logger())
	db := data_model.InitDB(&data_model.Config{Debug: true, AutoMigrate: true})

	app := src.App(suite.echo, db)
	go app()
}

func (suite *E2ESuite) TearDownSuite() {
	suite.echo.Shutdown(context.Background())
}

func (suite *E2ESuite) addStock(ref, sku, qty, eta interface{}) {
	res, _ := json.Marshal(map[string]interface{}{
		"reference":          ref,
		"sku":                sku,
		"purchased_quantity": qty,
		"eta":                eta,
	})

	http.Post("http://localhost:8080/stock", "application/json", bytes.NewBuffer(res))
}

func (suite *E2ESuite) TestApiReturns201AndAllocatedBatch() {
	sku := e2e.RandomSku()
	othersku := e2e.RandomSku("other")
	earlybatch := e2e.RandomBatchref("1")
	laterbatch := e2e.RandomBatchref("2")
	otherbatch := e2e.RandomBatchref("3")

	suite.addStock(laterbatch, sku, 100, time.Date(2011, 1, 2, 0, 0, 0, 0, time.UTC))
	suite.addStock(earlybatch, sku, 100, time.Date(2011, 1, 1, 0, 0, 0, 0, time.UTC))
	suite.addStock(otherbatch, othersku, 100, nil)

	resp, err := http.PostForm("http://localhost:8080/allocate",
		url.Values{"orderid": {e2e.RandomOrderid()}, "sku": {sku}, "qty": {"3"}},
	)
	assert.NoError(suite.T(), err)

	defer resp.Body.Close()

	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(suite.T(), 201, resp.StatusCode)
	assert.Equal(suite.T(), earlybatch, response["batchref"])
}

func (suite *E2ESuite) TestUnhappy400AndErrorMessage() {
	unknownSku := e2e.RandomSku()
	orderid := e2e.RandomOrderid()
	data := url.Values{"orderid": {orderid}, "sku": {unknownSku}, "qty": {"20"}}

	resp, err := http.PostForm("http://localhost:8080/allocate", data)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 400, resp.StatusCode)
	var response map[string]string
	json.NewDecoder(resp.Body).Decode(&response)
	assert.Equal(suite.T(), fmt.Sprintf("Invalid SKU %s", unknownSku), response["message"])
}

func TestE2ESuite(t *testing.T) {
	suite.Run(t, new(E2ESuite))
}
