package route

import (
	"echo-mongo/controller"

	"github.com/labstack/echo/v4"
)

func TodoRoute(e *echo.Echo) {
	e.POST("/todo", controller.CreateTodo)
	e.GET("/todo", controller.GetTodo)
	e.GET("/todo/name", controller.GetByName)
	e.PUT("/todo", controller.UpdateTodo)
	e.DELETE("/todo", controller.DeleteTodo)
}
