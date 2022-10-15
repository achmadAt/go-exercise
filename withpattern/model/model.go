package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id   primitive.ObjectID `bson:"_id" json:"id"`
	Name string             `bson:"name" json:"name"`
}
