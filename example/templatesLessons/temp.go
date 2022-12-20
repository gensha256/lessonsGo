package main

import (
	"log"
	"os"
	"text/template"
)

type User struct {
	Name    string
	Surname string
	Age     int
	Address string
}

func main() {
	var users User
	users = User{Name: "Djohn", Surname: "Matrix", Age: 25}

	temp, err := template.New("users").Parse("The user is {{.Name}} {{.Surname}} and age {{.Age}}")
	if err != nil {
		log.Fatal(err)
	}
	err = temp.Execute(os.Stdin, users)
	if err != nil {
		log.Fatal(err)
	}
}
