package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name string             `bson:"name" json:"name,omitempty"`
}
