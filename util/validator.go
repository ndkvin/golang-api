package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseValidationError(err error) map[string]string {
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		return map[string]string{"error": err.Error()}
	}

	errors := make(map[string]string)
	for _, ve := range validationErrors {
		field := strings.ToLower(ve.Field())
		tag := ve.Tag()

		var message string
		switch tag {
		case "required":
			message = fmt.Sprintf("%s is required", field)
		case "email":
			message = fmt.Sprintf("%s must be a valid email address", field)
		case "min":
			message = fmt.Sprintf("%s must be at least %s characters", field, ve.Param())
		case "max":
			message = fmt.Sprintf("%s must be at most %s characters", field, ve.Param())
		case "url":
			message = fmt.Sprintf("%s must be a valid url", field)
		case "alphanum":
			message = fmt.Sprintf("%s must be a valid alphanumeric", field)
		default:
			message = fmt.Sprintf("%s is not valid", field)
		}

		errors[field] = message
	}
	return errors
}
