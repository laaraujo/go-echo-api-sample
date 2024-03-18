package main

import (
	"context"
	"golang-echo-api/api"
	"golang-echo-api/config"
	"golang-echo-api/db"
	"golang-echo-api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	ctx := context.Background()
	connPool := db.GetConnPool(ctx)
	defer connPool.Close(ctx)

	e.Use(middleware.LoggerWithConfig(config.CustomLoggerConfig))

	app := handler.NewApp(connPool)
	api.Router(e, app)

	e.Logger.Fatal(e.Start(config.GetServerString()))
}
