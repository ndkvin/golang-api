package main

import (
	"shortlink/handler"
	"shortlink/router"
	"shortlink/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Validator = util.NewCustomValidator()
	e.HTTPErrorHandler = handler.ErrorHandler
	e.Pre(middleware.RemoveTrailingSlash()) 

	handler, err := InitializeWire()
	if err != nil {
		panic(err)
	}

	router.RegisterRouter(e, handler)

	e.Logger.Fatal(e.Start(":8000"))
}
