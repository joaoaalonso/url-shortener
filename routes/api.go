package routes

import (
	"github.com/gorilla/mux"
	"github.com/joaoaalonso/url-shortener/factories"
)

func apiRoutes(router *mux.Router, prefix string) {
	urlController := factories.CreateURLController()

	r := router.PathPrefix(prefix).Subrouter()

	r.HandleFunc("", urlController.Create).Methods("POST")
}
