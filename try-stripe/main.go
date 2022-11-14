package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"try-stripe/repository"
	"try-stripe/server"
	"try-stripe/service"

	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error load env with %s", err)
	}
	port := os.Getenv("PORT")
	uri := os.Getenv("MONGO_URI")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("stripe")
	repo := repository.NewRepository(database.Collection("pay"))
	service := service.NewService(repo)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	server.Server(ctx, e, service)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello",
		})
	})

	e.Logger.Fatal(e.Start(port))
}
