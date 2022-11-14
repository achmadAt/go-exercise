package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Models struct {
	ID           primitive.ObjectID
	Amount       int64
	ReceiptEmail string
}
