package routes

import "github.com/gorilla/mux"

// Register all routes to router
func Register(router *mux.Router) {
	webRoutes(router, "/")
	apiRoutes(router, "/api")
	assetsRoutes(router, "/assets")
}
