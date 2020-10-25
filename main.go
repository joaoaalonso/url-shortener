package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joaoaalonso/url-shortener/routes"
	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()
	routes.Register(router)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8000"
	}

	log.Println("Server is running on port " + port)
	http.ListenAndServe(":"+port, router)
}
