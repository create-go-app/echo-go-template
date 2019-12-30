package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Define server
	s := echo.New()

	// Middleware
	s.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "\n[${method}] ${uri} ${status}",
	}))

	// Router
	s.GET("/", indexHandler)
	s.GET("/user", userHandler)
	s.GET("/find/user/:name", findUserHandler)

	// Start server
	s.Logger.Fatal(s.Start(":8080"))
}
