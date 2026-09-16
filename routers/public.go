package routers

import (
	"connect/core/response"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GeneratePublicApiRoute(router *mux.Router)  {
	pubApi := router.PathPrefix("/api/v1").Subrouter()

	pubApi.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("METHOD: %s\n", r.Method)
		response.JSON(w, 200, true, "You are at public root dir what you need", nil)
	}, ).Methods(http.MethodGet)
}