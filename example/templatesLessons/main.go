package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type NewsPage struct {
	Title string
	News  string
}

func indexHandler(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "<h1>Whoa, Go is neat</h1>")
}

func newNewsHandler(writer http.ResponseWriter, req *http.Request) {
	var news NewsPage
	news = NewsPage{Title: "Amazing News", News: "some news"}
	temp, err := template.ParseFiles("basic.html")
	if err != nil {
		log.Fatal(err)
	}
	err = temp.Execute(writer, news)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/news", newNewsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
