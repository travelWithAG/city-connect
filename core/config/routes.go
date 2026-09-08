package config

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateAppRoutes() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		
	})

	return router
}