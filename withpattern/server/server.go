package server

import (
	"context"
	"withpattern/service"

	"github.com/labstack/echo/v4"
)

func Server(ctx context.Context, e *echo.Echo, service service.Service) {
	root := e.Group("todo")
	root.GET("", func(c echo.Context) error {
		res := service.GetTodo(ctx)
		return res
	})
}
