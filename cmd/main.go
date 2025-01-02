package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<p>Hello World.</p>")
	})

	e.Logger.Fatal(e.Start(":42069"))
}
