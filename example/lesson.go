package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	log.Println("What is your age : ")

	age, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	age = strings.TrimSpace(age)
	iAge, err := strconv.Atoi(age)
	if err != nil {
		log.Fatal(err)
	}

	if iAge < 5 {
		log.Println("Too young for school")
	} else if iAge == 5 {
		log.Println("Go to kindergarten")
	} else if (iAge > 5) && (iAge <= 17) {
		grade := iAge - 5
		log.Printf("Go to grade %d\n", grade)
	} else {
		log.Println("Go to college")
	}
}
