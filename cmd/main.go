package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/env"
	"github.com/lucashthiele/doc-it/logger"
	"github.com/lucashthiele/doc-it/middlewares"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
func main() {
	logger := logger.Get()
	env.Load()

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middlewares.CORS())
	e.Use(middlewares.Logger())

	SetupRoutes(e)

	logger.Error().AnErr("ERROR", e.Start(":42069"))
}
