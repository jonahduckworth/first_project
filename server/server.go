package server

import (
	"database/sql"
	"first_project/handler"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

// StartServer starts an HTTP server and handles incoming requests
func StartServer(db *sql.DB) error {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.GET("/users/:id", handler.GetUser(db))
	e.POST("/users", handler.CreateUser(db))
	e.PUT("/users/:id", handler.UpdateUser(db))
	e.DELETE("/users/:id", handler.DeleteUser(db))
	e.POST("/login", handler.Login(db))
	e.PUT("/updatePassword", handler.UpdatePassword(db))

	return e.Start(":8081")
}
