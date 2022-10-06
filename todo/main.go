package main

import (
	//"fmt"
	"log"

	"os"

	//"database/sql"

	"todo/repository"

	_ "github.com/go-sql-driver/mysql"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")
	//router
	e.POST("/todo", repository.PostTodos)
	e.GET("/", repository.GetTodo)
	e.PUT("/todo", repository.UpdateTodos)
	e.DELETE("/todo/:id", repository.DeleteTodo)
	// Start server
	e.Logger.Fatal(e.Start(port))
}
