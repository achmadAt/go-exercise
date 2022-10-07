package main

import (
	"echo-mongo/config"
	"echo-mongo/route"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	config.ConnectDB()
	//Router
	route.TodoRoute(e)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message":  "Hello",
			"message2": "Succes connected to mongodb",
		})
	})

	e.Logger.Fatal(e.Start(port))
}
