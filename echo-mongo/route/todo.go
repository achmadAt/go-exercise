package route

import (
	"echo-mongo/controller"

	"github.com/labstack/echo/v4"
)

func TodoRoute(e *echo.Echo) {
	e.POST("/todo", controller.CreateTodo)
}
