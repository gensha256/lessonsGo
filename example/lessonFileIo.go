package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	iPrimeArr := []int{2, 5, 7, 3, 11}
	var sPrimeArr []string

	for _, value := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(value))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		log.Println("Prime :", scan.Text())
	}

	_, err = os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		log.Println("File doesn't exist")
	} else {
		f, err = os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer f.Close()
	if _, err := f.WriteString("13\n"); err != nil {
		log.Fatal(err)
	}
}
