package service

import (
	"context"
	"withpattern/model"
	"withpattern/repository"
)

type Service interface {
	GetTodo(ctx context.Context) ([]*model.Todo, error)
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
