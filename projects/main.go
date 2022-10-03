package main

import (
	"net/http"
	"projects/config"

	"github.com/labstack/echo/v4"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	e.GET("/", Greetings)
	e.Logger.Fatal(e.Start(config.Port))
}

func Greetings(c echo.Context) error {
	return c.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World",
	})
}
