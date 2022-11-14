package repository

import (
	"context"
	"try-stripe/model"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	AddPayment(ctx context.Context, addPayment *model.Models) (*mongo.InsertOneResult, error)
}
type BaseRepository struct {
	collect *mongo.Collection
}

func NewRepository(collect *mongo.Collection) Repository {
	return &BaseRepository{collect: collect}
}

func (b *BaseRepository) AddPayment(ctx context.Context, addPayment *model.Models) (*mongo.InsertOneResult, error) {
	result, err := b.collect.InsertOne(ctx, addPayment)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return result, nil
}
