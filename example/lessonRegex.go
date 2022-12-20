package main

import (
	"log"
	"regexp"
)

func main() {
	firstStr := "The ape was the apex"
	fMatch, _ := regexp.MatchString("ape[^ ]", firstStr)
	log.Println(fMatch)

	secondStr := "cat fat rat mat pat"
	sMatch, _ := regexp.Compile("[crmpf]at")
	log.Println("MatchString :", sMatch.MatchString(secondStr))
	log.Println("FindString :", sMatch.FindString(secondStr))
	log.Println("Index :", sMatch.FindStringIndex(secondStr))
	log.Println("All Matches :", sMatch.FindAllString(secondStr, -1))

	mailStr := "djonyeek@gmail.com"

	mailM := regexp.MustCompile("^[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-z]{2,4}$")
	log.Println(mailM.MatchString(mailStr))

}
