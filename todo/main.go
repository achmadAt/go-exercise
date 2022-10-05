package main

import (
	"fmt"
	"log"

	"os"

	"todo/model"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println(port)
	// Routes
	e.GET("/users", model.GetAllUsers)
	e.POST("/users", model.CreateUser)
	e.GET("/users/:id", model.GetUser)
	e.PUT("/users/:id", model.UpdateUser)
	e.DELETE("/users/:id", model.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(port))
}
