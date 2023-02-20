package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db/mydatabase.db")
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %v", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT)")
	if err != nil {
		return nil, fmt.Errorf("Error creating table: %v", err)
	}
	return db, nil
}
