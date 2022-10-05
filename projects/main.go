package main

import (
	"io"
	"net/http"
	"os"
	"projects/config"

	"github.com/labstack/echo/v4"
)

type HelloWorld struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	//e.POST("/user", saveUser)
	e.GET("user/:id", getUser)
	//e.PUT("/user/:id", updateUser)
	//e.DELETE("/user/:id", deleteUser)
	e.GET("/", Greetings)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/save2", save2)
	e.Logger.Fatal(e.Start(config.Port))
}

func Greetings(c echo.Context) error {
	return c.JSON(http.StatusOK, HelloWorld{
		Message: "Hello World",
	})
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email+"\n")
}

func save2(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}
