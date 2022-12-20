package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		numByt, err := fmt.Fprint(w, "Hello Browser")
		if err != nil {
			log.Println(err)
		}
		log.Println("Bytes Written :", numByt)
	})
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
