package repository

import (
	"net/http"
	"strconv"
	"todo/model"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func GetTodo(c echo.Context) error {
	result := model.GetTodo()
	return c.JSON(http.StatusOK, result)
}
func PostTodos(c echo.Context) error {
	name := c.FormValue("name")
	result, err := model.PostTodos(name)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Tracef("post todo with name %s", name)
	return c.JSON(http.StatusOK, result)
}
func UpdateTodos(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err)
	}
	result, err := model.UpdateTodos(name, conv_id)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Tracef("update todo with id %s", id)
	return c.JSON(http.StatusAccepted, result)
}
func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		logrus.Error(err)
	}
	result, err := model.DeleteTodo(conv_id)
	if err != nil {
		logrus.Error(err)
	}
	logrus.Tracef("delete todo with id %s", id)
	return c.JSON(http.StatusAccepted, result)
}
