package main

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v7"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Greeting struct {
	Prefix  string `env:"GREETING_PREFIX" envDefault:"Hello"`
	Name    string `query:"name"`
	Message string
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", sayHello)
	e.Logger.Fatal(e.Start(":8080"))
}

func sayHello(c echo.Context) error {
	var greeting Greeting
	err := env.Parse(&greeting)
	if err != nil {
		return err
	}

	err = c.Bind(&greeting)
	if err != nil {
		return err
	}
	if greeting.Name == "" {
		greeting.Name = "Anonymous"
	}

	greeting.Message = fmt.Sprintf("%s,%s", greeting.Prefix, greeting.Name)

	return c.JSON(http.StatusOK, greeting)
}
