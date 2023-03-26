package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, My name is Muhammad Rafi Andeo Praja")
	})

	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "POST request received")
	})

	e.PUT("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "PUT request received")
	})

	e.DELETE("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "DELETE request received")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
