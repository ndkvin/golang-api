package handler

import (
	"net/http"
	"shortlink/dto"
	dtoLink "shortlink/dto/link"
	customError "shortlink/error"
	"shortlink/service"
	"shortlink/util"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type LinkHandlerInterface interface {
	CreateLink(c echo.Context) error
}

type linkHandler struct {
	linkService service.LinkServiceInterface
}

func NewLinkHandler(linkService service.LinkServiceInterface) LinkHandlerInterface {
	return &linkHandler{linkService}
}

func (l *linkHandler) CreateLink(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims, err := util.VerifyJWT(user)

	if err != nil {
		return err
	}

	var req dtoLink.CreateLinkRequest
	if err := c.Bind(&req); err != nil {
		return &customError.BadRequest{Message: "Invalid request body"}
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   util.ParseValidationError(err),
		})
	}

	link, err := l.linkService.Create(req, claims.Id)

	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, dto.WebResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   dtoLink.CreateLinkResponse{Name: link.Name, Url: link.Url},
	})

	return nil
}
