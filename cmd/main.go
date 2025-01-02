package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	SetupRoutes(e)

	e.Logger.Fatal(e.Start(":42069"))
}
