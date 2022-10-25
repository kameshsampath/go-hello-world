package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const greeting = "Hello, World!"

func main() {
	e := echo.New()

	e.GET("/", sayHello)
	e.Logger.Fatal(e.Start(":8080"))
}

func sayHello(c echo.Context) error {
	return c.String(http.StatusOK, greeting)
}
