package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

// Index ...
type Index struct {
	Hello string `json:"hello" query:"hello"`
}

// User ...
type User struct {
	Name    string `json:"name" query:"name"`
	Surname string `json:"surname" query:"surname"`
	Email   string `json:"email" query:"email"`
}

// indexHandler ...
func indexHandler(c echo.Context) error {
	index := &Index{
		Hello: "world",
	}
	return c.JSON(http.StatusOK, index)
}

// userHandler ...
func userHandler(c echo.Context) error {
	user := &User{
		Name:    "John",
		Surname: "Doe",
		Email:   "john@doe.com",
	}
	return c.JSON(http.StatusOK, user)
}

// findUserHandler ...
func findUserHandler(c echo.Context) error {
	param := c.Param("name")
	user := &User{
		Name:    "Jane",
		Surname: "Doe",
		Email:   "jane@doe.com",
	}

	if strings.ToLower(param) == strings.ToLower(user.Name) {
		return c.JSON(http.StatusOK, user)
	}

	return c.String(http.StatusNotFound, "Not Found")
}
