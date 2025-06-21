package main

import (
	"shortlink/handler"
	"shortlink/router"
	"shortlink/util"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = util.NewCustomValidator()
	e.HTTPErrorHandler = handler.ErrorHandler

	handler, err := InitializeWire()
	if err != nil {
		panic(err)
	}

	router.RegisterRouter(e, handler)

	e.Logger.Fatal(e.Start(":8000"))
}
