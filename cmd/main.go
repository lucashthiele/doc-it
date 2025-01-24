package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	Logger.Info().Msg("Starting application configuration")
	e := echo.New()

	e.Use(middleware.CORS())

	Logger.Error().AnErr("ERROR", e.Start(":42069"))
}
