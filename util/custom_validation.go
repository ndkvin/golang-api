package util

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validate *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	return &CustomValidator{validate: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validate.Struct(i)
}