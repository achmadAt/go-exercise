package repository

import (
	"net/http"
	"todo/model"

	"github.com/labstack/echo/v4"
)

func GetTodo(c echo.Context) error {
	result := model.GetTodo()
	return c.JSON(http.StatusOK, result)
}
func PostTodos(c echo.Context) error {
	name := c.FormValue("name")
	return c.JSON(http.StatusOK, name)
}
