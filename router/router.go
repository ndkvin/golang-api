package router

import (
	"os"
	"shortlink/handler"
	"shortlink/util"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo, handler *handler.Handler) {
	godotenv.Load()

	api := e.Group("/api")

	api.POST("/register", handler.AuthHandler.Register)
	api.POST("/login", handler.AuthHandler.Login)

	link := api.Group("/link")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(util.JwtClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	link.Use(echojwt.WithConfig(config))

	link.POST("", handler.LinkHandler.CreateLink)

	e.GET("/:name", handler.LinkHandler.VisitLink)
}
