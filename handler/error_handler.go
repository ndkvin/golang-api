package handler

import (
	"fmt"
	"log"
	"net/http"
	"shortlink/dto"
	customError "shortlink/error"
	"strings"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	var code = http.StatusInternalServerError
	var message interface{} = "Internal Server Error"

	switch e := err.(type) {
	case *customError.BadRequest:
		code = http.StatusBadRequest
		message = e.Message

	case *customError.NotFound:
		code = http.StatusNotFound
		message = e.Message

	case *echo.HTTPError:
		msg := fmt.Sprintf("%v", e.Message)
		if strings.Contains(strings.ToLower(msg), "jwt") { 
			code = http.StatusUnauthorized
			message = e.Message
		} else {
			code = e.Code
			message = e.Message
		}

	default:
		log.Printf("internal error: %+v", err)
	}

	_ = c.JSON(code, dto.WebResponse{
		Code:   code,
		Status: http.StatusText(code),
		Data:   message,
	})
}
