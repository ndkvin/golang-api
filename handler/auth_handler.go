package handler

import (

	"net/http"
	"shortlink/dto/auth"
	"shortlink/service"
	"shortlink/util"

	"github.com/labstack/echo/v4"
)

type AuthHandlerInterface interface {
	Register(c echo.Context) error
}

type authHandler struct {
	authService service.AuthServiceInterface
}

func NewAuthHandler(authService service.AuthServiceInterface) AuthHandlerInterface {
	return &authHandler{authService}
}



func (a *authHandler) Register(c echo.Context) error {
	var req auth.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"errors": util.ParseValidationError(err),
		})
	}

	c.JSON(http.StatusOK, map[string]string{"message": "success"})

	return nil
}
