package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/internal/server"
	"github.com/lucashthiele/doc-it/pkg/env"
	"github.com/lucashthiele/doc-it/pkg/logger"
)

func configServer(e *echo.Echo) {
	server.SetupMiddlewares(e)
	server.SetupValidator(e)
	server.SetupRoutes(e)
}

func main() {
	logger := logger.Get()
	env.Load()

	e := echo.New()
	configServer(e)
	logger.Error().AnErr("ERROR", e.Start(":42069"))
}
