package config

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	mongo_uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongo_uri))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	//[] todo must use when service already cleared/finished
	//[] todo need to move/fix this
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			fmt.Printf("client.Disconnect error: %s", err)
			return
		}

		fmt.Println("mongodb disconnected")
	}()
	fmt.Println("Connected")
	return client
}

// client instance
var DB *mongo.Client = ConnectDB()

// get collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangApi").Collection(collectionName)
	return collection
}
