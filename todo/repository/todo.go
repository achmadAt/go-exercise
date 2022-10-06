package repository

import (
	"net/http"
	"strconv"
	"todo/model"

	"github.com/labstack/echo/v4"
)

func GetTodo(c echo.Context) error {
	result := model.GetTodo()
	return c.JSON(http.StatusOK, result)
}
func PostTodos(c echo.Context) error {
	name := c.FormValue("name")
	result := model.PostTodos(name)
	return c.JSON(http.StatusOK, result)
}
func UpdateTodos(c echo.Context) error {
	id := c.FormValue("id")
	name := c.FormValue("name")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	result := model.UpdateTodos(name, conv_id)
	return c.JSON(http.StatusAccepted, result)
}
func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	conv_id, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	result := model.DeleteTodo(conv_id)
	return c.JSON(http.StatusAccepted, result)
}
