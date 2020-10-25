package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func assetsRoutes(router *mux.Router, prefix string) {
	fileServer := http.FileServer(http.Dir("./public/assets/"))

	r := router.PathPrefix(prefix)

	r.Handler(http.StripPrefix(prefix, fileServer))
}
