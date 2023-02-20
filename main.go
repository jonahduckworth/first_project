package main

import (
	"fmt"

	"first_project/db"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
}
