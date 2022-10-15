package server

import (
	"context"
	"fmt"
	"net/http"
	"withpattern/model"
	"withpattern/service"

	validator "github.com/go-playground/validator/v10"
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
	root.POST("", func(c echo.Context) error {
		var todo model.Todo
		if err := c.Bind(&todo); err != nil {
			return err
		}
		validate := validator.New()
		err := validate.Struct(&todo)
		if err != nil {
			return err
		}
		fmt.Println(todo.Name)
		res, err := service.AddTodo(ctx, todo.Name)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, &echo.Map{"data": res})
	})
	root.PUT("", func(c echo.Context) error {
		var todo model.Todo
		if err := c.Bind(&todo); err != nil {
			return err
		}
		validate := validator.New()
		err := validate.Struct(&todo)
		if err != nil {
			return err
		}
		res, err := service.UpdateTodo(ctx, todo.Id.Hex(), todo.Name)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusAccepted, &echo.Map{"data": res})
	})
	root.DELETE("", func(c echo.Context) error {
		var todo model.Todo
		if err := c.Bind(&todo); err != nil {
			return err
		}
		validate := validator.New()
		err := validate.Struct(&todo)
		if err != nil {
			return err
		}
		res, err := service.DeleteTodo(ctx, todo.Id.Hex())
		if err != nil {
			return err
		}
		return c.JSON(http.StatusAccepted, &echo.Map{"data": res})
	})
}
