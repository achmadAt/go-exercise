package repository

import (
	"context"
	"net/http"
	"withpattern/model"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetTodo(ctx context.Context) error
}
type BaseRepository struct {
	collect *mongo.Collection
}

func NewRepository(collect *mongo.Collection) Repository {
	return &BaseRepository{collect: collect}
}

func (base *BaseRepository) GetTodo(ctx context.Context) error {
	result, err := base.collect.Find(ctx, bson.M{})
	if err != nil {
		return echo.New().AcquireContext().JSON(http.StatusInternalServerError, &echo.Map{"err": err})
	}
	defer result.Close(ctx)
	var todos []model.Todo
	for result.Next(ctx) {
		var todo model.Todo
		if err := result.Decode(&todo); err != nil {
			return echo.New().AcquireContext().JSON(http.StatusInternalServerError, &echo.Map{"error": err})
		}
		todos = append(todos, todo)
	}
	return echo.New().AcquireContext().JSON(http.StatusAccepted, &echo.Map{"data": todos})
}
