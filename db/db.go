package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/mydb")
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %v", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), email VARCHAR(255))")
	if err != nil {
		return nil, fmt.Errorf("Error creating table: %v", err)
	}
	return db, nil
}
