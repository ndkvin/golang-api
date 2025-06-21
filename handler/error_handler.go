package handler

import (
	"log"
	"net/http"
	"shortlink/dto"
	customError "shortlink/error"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	switch e := err.(type) {

	case *customError.BadRequest:
		c.JSON(http.StatusBadRequest, dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   e.Message,
		})

	default:
		log.Printf("internal error: %+v", err)
		c.JSON(http.StatusInternalServerError, dto.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Internal Server Error",
			Data:   "Internal Server Error",
		})
	}
}
