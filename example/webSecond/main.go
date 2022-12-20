package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/divide", divideHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

func write(writer http.ResponseWriter, msg string) {
	_, err := writer.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemp, err := template.ParseFiles("./templates/" + tmpl)
	errorCheck(err)
	err = parsedTemp.Execute(w, nil)
	errorCheck(err)

}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func divideHandler(writer http.ResponseWriter, request *http.Request) {
	value, err := getQuotient(5, 4)
	if err != nil {
		write(writer, "can't divide by zero\n")
	} else {
		output := fmt.Sprintf("5 / 4 =%.2f\n", value)
		write(writer, output)
	}
}

func homeHandler(w http.ResponseWriter, request *http.Request) {
	renderTemplate(w, "homePage.tmpl")
}

func getQuotient(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	} else {
		return x / y, nil
	}
}
