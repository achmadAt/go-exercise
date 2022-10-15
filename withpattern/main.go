package main

import (
	"context"
	"log"
	"net/http"
	"withpattern/repository"
	"withpattern/server"
	"withpattern/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://root:root@cluster0.66asf7z.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
		return
	}
	database := client.Database("golangApi")
	repo := repository.NewRepository(database.Collection("todo"))
	service := service.NewService(repo)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
			return
		}
	}()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
		return
	}
	server.Server(ctx, e, service)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello",
		})
	})
	e.Logger.Fatal(e.Start(":8000"))
}
