package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/internal/middlewares"
	"github.com/lucashthiele/doc-it/pkg/env"
	"github.com/lucashthiele/doc-it/pkg/logger"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func configServer(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middlewares.CORS())
	e.Use(middlewares.Logger())

	SetupRoutes(e)
}

func main() {
	logger := logger.Get()
	env.Load()

	e := echo.New()
	configServer(e)
	logger.Error().AnErr("ERROR", e.Start(":42069"))
}
