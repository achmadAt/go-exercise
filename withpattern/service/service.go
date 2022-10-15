package service

import (
	"context"
	"withpattern/model"
	"withpattern/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	GetTodo(ctx context.Context) ([]*model.Todo, error)
	AddTodo(ctx context.Context, name string) (*mongo.InsertOneResult, error)
	UpdateTodo(ctx context.Context, id string, name string) (*mongo.UpdateResult, error)
	DeleteTodo(ctx context.Context, id string) (*mongo.DeleteResult, error)
}

type BaseService struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return &BaseService{repo: r}
}

func (b *BaseService) GetTodo(ctx context.Context) ([]*model.Todo, error) {
	return b.repo.GetTodo(ctx)
}
func (b *BaseService) AddTodo(ctx context.Context, name string) (*mongo.InsertOneResult, error) {
	return b.repo.AddTodo(ctx, name)
}
func (b *BaseService) UpdateTodo(ctx context.Context, id string, name string) (*mongo.UpdateResult, error) {
	return b.repo.UpdateTodo(ctx, id, name)
}
func (b *BaseService) DeleteTodo(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	return b.repo.DeleteTodo(ctx, id)
}
