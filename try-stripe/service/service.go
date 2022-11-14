package service

import (
	"context"
	"try-stripe/model"
	"try-stripe/repository"

	"go.mongodb.org/mongo-driver/mongo"
)

type Service interface {
	AddPayment(ctx context.Context, addPayment *model.Models) (*mongo.InsertOneResult, error)
}

type BaseService struct {
	repo repository.Repository
}

func NewService(r repository.Repository) Service {
	return &BaseService{repo: r}
}

func (b *BaseService) AddPayment(ctx context.Context, addPayment *model.Models) (*mongo.InsertOneResult, error) {
	return b.repo.AddPayment(ctx, addPayment)
}
