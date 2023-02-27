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

    result, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", u.Name, u.Email, u.Password)
    if err != nil {
      return err
    }

    id, err := result.LastInsertId()
    if err != nil {
      return err
    }

    u.ID = int(id)
	return c.JSON(http.StatusOK, map[string]bool{"userCreated": true})
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

func Login(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
	  type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	  }
  
	  var req Request
	  if err := c.Bind(&req); err != nil {
		return err
	  }
  
	  email := req.Email
	  password := req.Password
  
	  // Query the database to check if the user exists
	  var id int
	  err := db.QueryRow("SELECT id FROM users WHERE email = ? AND password = ?", email, password).Scan(&id)
	  if err == sql.ErrNoRows {
		// If the user doesn't exist, return an error
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid email or password"})
	  } else if err != nil {
		// If there's another error, return an internal server error
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred while trying to log in"})
	  }
  
	  // If the user exists, return success with the user's id
	  return c.JSON(http.StatusOK, map[string]bool{"userExists": true})
	}
  }
  
