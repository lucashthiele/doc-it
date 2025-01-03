package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/handler/page"
)

func SetupRoutes(e *echo.Echo) {
	indexHandler := page.IndexHandler{}

	// pages
	e.GET("/", indexHandler.Show)
}
