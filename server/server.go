package server

import (
	"database/sql"
	"first_project/handler"

	"github.com/labstack/echo/v4"
)

// StartServer starts an HTTP server and handles incoming requests
func StartServer(db *sql.DB) error {
	e := echo.New()
	e.GET("/users/:id", handler.GetUser(db))
	e.POST("/users", handler.CreateUser(db))
	e.PUT("/users/:id", handler.UpdateUser(db))
	e.DELETE("/users/:id", handler.DeleteUser(db))

	return e.Start(":8080")
}
