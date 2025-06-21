package router

import (
	"shortlink/handler"

	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo, handler *handler.Handler) {
	e.POST("/register", handler.AuthHandler.Register)
	e.POST("/login", handler.AuthHandler.Login)
}