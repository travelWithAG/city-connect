package main

import (
	"connect/core/db"
	"fmt"
)

func main() {

	db, err := db.InitPostgreSQLDB()

	if err != nil {
		fmt.Printf("Error occured during database processing: %v\n", err)
	}

	_ = db
	
}