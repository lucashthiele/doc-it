package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/handlers"
)

func SetupRoutes(e *echo.Echo) {
	githubHandler := handlers.GitHub{}
	e.POST("/github", githubHandler.Post)
}
