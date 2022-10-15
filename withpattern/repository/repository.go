package repository

import (
	"context"
	"withpattern/model"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetTodo(ctx context.Context) ([]*model.Todo, error)
	AddTodo(ctx context.Context, name string) (*mongo.InsertOneResult, error)
	UpdateTodo(ctx context.Context, id string, name string) (*mongo.UpdateResult, error)
	DeleteTodo(ctx context.Context, id string) (*mongo.DeleteResult, error)
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
		log.Warn(err)
		return nil, err
	}
	defer result.Close(ctx)
	var todos []*model.Todo
	for result.Next(ctx) {
		var todo model.Todo
		if err := result.Decode(&todo); err != nil {
			log.Warn(err)
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}
func (b *BaseRepository) AddTodo(ctx context.Context, name string) (*mongo.InsertOneResult, error) {
	addTodo := model.Todo{
		Id:   primitive.NewObjectID(),
		Name: name,
	}
	result, err := b.collect.InsertOne(ctx, addTodo)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return result, nil
}

func (b *BaseRepository) UpdateTodo(ctx context.Context, id string, name string) (*mongo.UpdateResult, error) {
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	filter := bson.D{{Key: "_id", Value: hexId}}
	payload := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}}}}
	result, err := b.collect.UpdateOne(ctx, filter, payload)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return result, nil
}

func (b *BaseRepository) DeleteTodo(ctx context.Context, id string) (*mongo.DeleteResult, error) {
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	payload := bson.D{{Key: "_id", Value: hexId}}
	result, err := b.collect.DeleteOne(ctx, payload)
	if err != nil {
		log.Warn(err)
		return nil, err
	}
	return result, nil
}
