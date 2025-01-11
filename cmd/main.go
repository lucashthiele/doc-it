package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucashthiele/doc-it/env"
)

func main() {
	e := echo.New()

	err := env.Load()
	if err != nil {
		panic("Error loading .env file. This should never happen.")
	}

	e.Use(middleware.CORS())

	SetupRoutes(e)

	e.Logger.Fatal(e.Start(":42069"))
}
