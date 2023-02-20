package main

import (
	"fmt"

	"first_project/db"
	"first_project/server"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	err = server.StartServer(db)
	if err != nil {
		fmt.Println(err)
		return
	}
}
