package routes

import (
	"github.com/joaoaalonso/url-shortener/factories"

	"github.com/gorilla/mux"
)

func webRoutes(router *mux.Router, prefix string) {
	urlController := factories.CreateURLController()

	r := router.PathPrefix(prefix).Subrouter()

	r.HandleFunc("/", urlController.Home).Methods("GET")
	r.HandleFunc("/{alias}", urlController.Redirect).Methods("GET")
}
