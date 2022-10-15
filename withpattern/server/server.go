package server

import (
	"context"
	"net/http"
	"withpattern/service"

	"github.com/labstack/echo/v4"
)

func Server(ctx context.Context, e *echo.Echo, service service.Service) {
	root := e.Group("todo")
	root.GET("", func(c echo.Context) error {
		res, err := service.GetTodo(ctx)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusAccepted, &echo.Map{"data": res})
	})
}
