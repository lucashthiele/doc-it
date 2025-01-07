package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/handler/page"
)

func SetupRoutes(e *echo.Echo) {
	indexHandler := page.IndexHandler{}
	loginHandler := page.LoginHandler{}

	// static assets
	e.Static("/public", "view/assets")

	// pages
	e.GET("/", indexHandler.Show)
	e.GET("/login", loginHandler.Show)
}
