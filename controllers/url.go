package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaoaalonso/url-shortener/entities"
	"github.com/joaoaalonso/url-shortener/services"
)

// URLController handle all url requests for url entities
type URLController struct {
	URLService services.URLService
}

type errorResponse struct {
	Message string `json:"message"`
}

// Home returns a home view
func (urlController *URLController) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, "./public/views/index.html")
}

// Create a new shorten url
func (urlController *URLController) Create(w http.ResponseWriter, r *http.Request) {
	var url entities.URL
	_ = json.NewDecoder(r.Body).Decode(&url)

	w.Header().Set("Content-type", "application/json")

	url, err := urlController.URLService.Create(url.LongURL, url.Alias)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errorResponse{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(url)
}

// Redirect to long url
func (urlController *URLController) Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	alias := vars["alias"]

	url := urlController.URLService.GetURLFromAlias(alias)

	http.Redirect(w, r, url, 302)
}
