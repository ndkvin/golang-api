package handler

import (
	"net/http"
	"shortlink/dto"
	"shortlink/dto/auth"
	"shortlink/service"
	"shortlink/util"
	customError "shortlink/error"

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
		return &customError.BadRequest{Message: "Invalid request body"}
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:  util.ParseValidationError(err),
		})
	}

	user, err := a.authService.Register(req)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, dto.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   auth.RegisterResponse{Email: user.Email},
	})

	return nil
}
