package main

import (
	"log"
	"net/http"
	"url-shortener/internal/controllers"
)

func main() {
	http.HandleFunc("GET /", controllers.ShowIndex)
	http.HandleFunc("GET /shorten", controllers.Shorten)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
