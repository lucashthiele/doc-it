package server

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/internal/handlers"
)

func SetupRoutes(e *echo.Echo) {
	githubHandler := handlers.GitHub{}
	e.POST("/github", githubHandler.Post)
}
