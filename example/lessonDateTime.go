package main

import (
	"log"
	"time"
)

func main() {
	now := time.Now()
	log.Println(now.Year(), now.Month(), now.Day())
	log.Println(now.Hour(), now.Minute(), now.Second())
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Time in New York", now.In(loc))

	birthDate := time.Date(2004, time.September, 20, 11, 25, 30, 40, time.Local)
	log.Println(birthDate)
	diff := now.Sub(birthDate)
	log.Printf("Days alive : %d days", int(diff.Hours()/24))
	log.Printf("Hours alive : %d hours", int(diff.Hours()))

}
