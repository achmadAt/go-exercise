package model

import (
	"net/http"
	"strconv"
	"todo/entities"

	"github.com/labstack/echo/v4"
)

// ----------
// Handlers
// ----------
func CreateUser(c echo.Context) error {
	u := &entities.User{
		ID: entities.Seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	entities.Users[u.ID] = u
	entities.Seq++
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, entities.Users[id])
}

func UpdateUser(c echo.Context) error {
	u := new(entities.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	entities.Users[id].Name = u.Name
	return c.JSON(http.StatusOK, entities.Users[id])
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(entities.Users, id)
	return c.NoContent(http.StatusNoContent)
}

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, entities.Users)
}
