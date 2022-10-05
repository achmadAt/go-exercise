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
