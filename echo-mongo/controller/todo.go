package controller

import (
	"echo-mongo/config"
	"echo-mongo/dto"
	"echo-mongo/entities/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

var todoCollection *mongo.Collection = config.GetCollection(config.DB, "todo")

func CreateTodo(c echo.Context) error {
	var todo dto.Todo
	//todo.Name = c.FormValue("name")
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	if err := utils.Validate(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"struct": err}})
	}
	newTodo := dto.Todo{
		Id:   primitive.NewObjectID(),
		Name: todo.Name,
	}
	result, err := todoCollection.InsertOne(context.Background(), newTodo)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, dto.Response{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func GetTodo(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var todos []*dto.Todo
	result, err := todoCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer result.Close(ctx)
	for result.Next(ctx) {
		var singleTodo dto.Todo
		if err := result.Decode(&singleTodo); err != nil {
			log.Fatal(err)
		}
		todos = append(todos, &singleTodo)
	}
	return c.JSON(http.StatusCreated, dto.Response{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": todos}})
}
func GetByName(c echo.Context) error {
	Name := c.FormValue("name")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var todo dto.Todo
	filter := bson.D{{Key: "name", Value: Name}}
	err := todoCollection.FindOne(ctx, filter).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, dto.Response{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": todo}})
}

func UpdateTodo(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	id := c.FormValue("id")
	name := c.FormValue("name")
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "_id", Value: hexId}}
	payload := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: name}}}}
	result, err := todoCollection.UpdateOne(ctx, filter, payload, opts)
	if err != nil {
		log.Fatal(err)
	}
	if result.MatchedCount != 0 {
		fmt.Println("match and replace an existing document")
		return c.JSON(http.StatusAccepted, dto.Response{Status: http.StatusAccepted, Message: "ok", Data: &echo.Map{"updated": result.MatchedCount}})
	}
	if result.UpsertedCount != 0 {
		return c.JSON(http.StatusAccepted, dto.Response{Status: http.StatusAccepted, Message: "ok", Data: &echo.Map{"upserted id": result.UpsertedID}})
	}
	return nil
}
func DeleteTodo(c echo.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	id := c.FormValue("id")
	hexId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})
	payload := bson.D{{Key: "_id", Value: hexId}}
	res, err := todoCollection.DeleteOne(ctx, payload, opts)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusAccepted, dto.Response{Status: http.StatusAccepted, Message: "ok", Data: &echo.Map{"deleted": res.DeletedCount}})
}

//default valu for pointer is nil, default value for primitive bool is false
//
