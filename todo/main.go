package main

import (
	"log"

	"os"

	"database/sql"
	"todo/model"

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
	//db
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/todos")
	if err != nil {
		panic(err.Error())
	}
	db.Ping()
	// Routes
	e.GET("/users", model.GetAllUsers)
	e.POST("/users", model.CreateUser)
	e.GET("/users/:id", model.GetUser)
	e.PUT("/users/:id", model.UpdateUser)
	e.DELETE("/users/:id", model.DeleteUser)

	// Start server
	e.Logger.Fatal(e.Start(port))
}
