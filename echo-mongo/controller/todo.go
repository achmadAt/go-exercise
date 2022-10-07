package controller

import (
	"echo-mongo/config"
	"echo-mongo/dto"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var todoCollection *mongo.Collection = config.GetCollection(config.DB, "todo")
var validate = validator.New()

func CreateTodo(c echo.Context) error {
	var todo dto.Todo
	todo.Name = c.FormValue("name")
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, dto.Response{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}
	if err := validate.Struct(&todo); err != nil {
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
	fmt.Println(todo.Name)
	return c.JSON(http.StatusCreated, dto.Response{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func GetTodo() {

}
func UpdateTodo() {

}
func DeleteTodo() {

}
