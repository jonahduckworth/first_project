package handler

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"cloud.google.com/go/storage"

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
  
	  _, err = db.Exec("INSERT INTO passwords (user_id, password) VALUES (?, ?)", id, u.Password)
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
		err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Email)
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
  
	  // If the user exists, return userExists as true and the user's id
	  return c.JSON(http.StatusOK, map[string]interface{}{"userExists": true, "id": id})
	}
  }

func UpdatePassword(db *sql.DB) echo.HandlerFunc {
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

		// Check if the password meets the requirements
		if len(password) < 8 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Password must contain at least 8 characters"})
		}
		if !regexp.MustCompile(`[a-zA-Z]`).MatchString(password) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Password must contain at least one letter"})
		}
		if !regexp.MustCompile(`\d`).MatchString(password) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Password must contain at least one digit"})
		}
		if !regexp.MustCompile(`[!"#$%&'()*+,\-./]`).MatchString(password) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Password must contain at least one of the following special characters: ! \" # $ % & ' ( ) * + , - . /"})
		}

		// Query the database to check if the user exists
		var id int
		err := db.QueryRow("SELECT id FROM users WHERE email = ?", email).Scan(&id)
		if err == sql.ErrNoRows {
			// If the user doesn't exist, return an error
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid email"})
		} else if err != nil {
			// If there's another error, return an internal server error
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred while trying to reset password"})
		}

		// Query the database to get the list of previous passwords used by the user
		var previousPasswords []string
		rows, err := db.Query("SELECT password FROM passwords WHERE user_id = ?", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred while trying to reset password"})
		}
		defer rows.Close()
		for rows.Next() {
			var prevPassword string
			if err := rows.Scan(&prevPassword); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "An error occurred while trying to reset password"})
			}
			previousPasswords = append(previousPasswords, prevPassword)
		}
		
		// Check if the password contains a sequence of 4 consecutive characters that is also found in any of the previous passwords
		for _, prevPassword := range previousPasswords {
			for i := 0; i < len(password)-3; i++ {
				if strings.Contains(prevPassword, password[i:i+4]) {
					return c.JSON(http.StatusBadRequest, map[string]string{"error": "Password cannot contain a sequence of 4 consecutive characters that is also found in any of the previous passwords used by this user"})
				}
			}
		}
		
		_, err = db.Exec("UPDATE users SET password = ? WHERE email = ?", password, email)
		if err != nil {
			return err
		}

		_, err = db.Exec("INSERT INTO passwords (user_id, password) VALUES (?, ?)", id, password)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]bool{"passwordUpdated": true})
	}
}

func storeImageInGCS(image io.Reader, objectName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	bucketName := "my-bucket"
	bucket := client.Bucket(bucketName)

	object := bucket.Object(objectName)
	w := object.NewWriter(ctx)
	if _, err := io.Copy(w, image); err != nil {
		return fmt.Errorf("failed to store image: %v", err)
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %v", err)
	}

	return nil
}

func handleImageUpload(w http.ResponseWriter, r *http.Request) {
	image, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "failed to get image from request", http.StatusBadRequest)
		return
	}
	defer image.Close()

	if err := storeImageInGCS(image, header.Filename); err != nil {
		http.Error(w, "failed to store image", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Image uploaded successfully"))
}