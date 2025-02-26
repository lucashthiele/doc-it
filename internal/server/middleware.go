package server

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/internal/middlewares"
)

func SetupMiddlewares(e *echo.Echo) {
	e.Use(middlewares.CorsMiddleware())
	e.Use(middlewares.LoggerMiddleware())
}
