package routers

import (
	"connect/core/config"
	"connect/core/middleware"
	"connect/core/response"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func GenerateApiRoute(router *mux.Router)  {
	api := router.PathPrefix("/api/v1").Subrouter()

	scr := config.LoadAppConfig("").JWTSecret

	protected := api.PathPrefix("").Subrouter()

	protected.Use(middleware.AuthMiddleware(scr))

	protected.HandleFunc("/add-product", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("METHOD: %s\n", r.Method)
		response.JSON(w, http.StatusFound, true, "The server you are looking contain the api you requested", nil)
	})
}