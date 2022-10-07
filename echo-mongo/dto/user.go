package dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id   primitive.ObjectID `json:"id,omitempty"`
	Name string             `json:"name,omitempty"`
}
