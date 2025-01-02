package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/controller/page"
)

func SetupRoutes(e *echo.Echo) {
	indexHandler := page.IndexHandler{}

	e.GET("/", indexHandler.Show)
}
