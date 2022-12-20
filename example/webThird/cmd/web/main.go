package main

import (
	"log"
	"net/http"
	"webThird/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
