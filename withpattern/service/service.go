package service

import (
	"context"
	"withpattern/repository"
)

type Service interface {
	GetTodo(ctx context.Context) error
}

type BaseService struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return &BaseService{repo: r}
}

func (b *BaseService) GetTodo(ctx context.Context) error {
	return b.repo.GetTodo(ctx)
}
