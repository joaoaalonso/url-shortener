package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaoaalonso/url-shortener/entities"
	"github.com/joaoaalonso/url-shortener/services"
)

// URLController handle all url requests for url entities
type URLController struct {
	URLService services.URLService
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

	url, err := urlController.URLService.Create(url.LongURL, url.Alias)

	if err != nil {
		fmt.Println("vixi")
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
