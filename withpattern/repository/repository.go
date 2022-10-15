package repository

import (
	"context"
	"withpattern/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetTodo(ctx context.Context) ([]*model.Todo, error)
}
type BaseRepository struct {
	collect *mongo.Collection
}

func NewRepository(collect *mongo.Collection) Repository {
	return &BaseRepository{collect: collect}
}

func (base *BaseRepository) GetTodo(ctx context.Context) ([]*model.Todo, error) {
	result, err := base.collect.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)
	var todos []*model.Todo
	for result.Next(ctx) {
		var todo model.Todo
		if err := result.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}
