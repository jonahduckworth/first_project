package handler

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"first_project/model"
)

func CreateUser(db *sql.DB) echo.HandlerFunc {
  return func(c echo.Context) error {
    u := new(model.User)
    if err := c.Bind(u); err != nil {
      return err
    }

    // Insert a new user into the database
    result, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", u.Name, u.Email, u.Password)
    if err != nil {
      return err
    }

    id, err := result.LastInsertId()
    if err != nil {
      return err
    }

    u.ID = int(id)
    return c.JSON(http.StatusCreated, u)
  }
}

func GetUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		var u model.User
		err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	}
}

func UpdateUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			return err
		}
		id := c.Param("id")
		_, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", u.Name, u.Email, id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, u)
	}
}

func DeleteUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		_, err := db.Exec("DELETE FROM users WHERE id = ?", id)
		if err != nil {
			return err
		}
		return c.NoContent(http.StatusNoContent)
	}
}
