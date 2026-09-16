package main

import (
	"connect/core/config"
	"connect/core/db"
	"connect/core/middleware"
	"connect/routers"
	"fmt"

	"github.com/gorilla/mux"
)

func main() {

	db, err := db.InitPostgreSQLDB()

	if err != nil {
		fmt.Printf("Error occured during database processing: %v\n", err)
	}

	_ = db

	router := mux.NewRouter()

	router.Use(middleware.CORS)
	router.Use(middleware.LoggingMiddleware)

	routers.GeneratePublicApiRoute(router)

	routers.GenerateApiRoute(router)

	config.StartServer(router)
	
}