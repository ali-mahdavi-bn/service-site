package main

import (
	"context"
	"github.com/ali-mahdavi-bn/service-site/src"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/container"
	"github.com/ali-mahdavi-bn/service-site/src/backbone/infrastructr/database"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"os"
	"os/signal"
	"time"
)

func main() {
	e := echo.New()
	container.NewLogger(e.Logger)
	database.InitDB(&database.Config{Debug: true, AutoMigrate: true})
	e.Validator = &container.CustomValidator{Validator: validator.New()}
	server := src.App(e)
	server()
	//go server()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		container.Logger.Fatal(err)
	}
}
