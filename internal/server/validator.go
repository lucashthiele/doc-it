package server

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i any) error {
	return cv.validator.Struct(i)
}

func SetupValidator(e *echo.Echo) {
	e.Validator = &customValidator{validator: validator.New()}
}
