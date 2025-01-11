package main

import (
	"github.com/labstack/echo/v4"
	"github.com/lucashthiele/doc-it/handler/page"
)

func SetupRoutes(e *echo.Echo) {

	// static assets
	e.Static("/public", "view/assets")

	// pages
	indexHandler := page.IndexHandler{}
	loginHandler := page.LoginHandler{}
	loginCallbackHandler := page.LoginCallbackHandler{}

	e.GET("/", indexHandler.Show)
	e.GET("/login", loginHandler.Show)
	e.GET("/login/callback", loginCallbackHandler.Get)
}
