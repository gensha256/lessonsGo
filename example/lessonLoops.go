package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1

	for true {
		log.Println("Guess number between 1 and 50 : ")
		log.Println("Random number is : ", randNum)

		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if iGuess > randNum {
			log.Println("Guess something lower")
		} else if iGuess < randNum {
			log.Println("Guess something higher")
		} else {
			log.Println("You Guessed it")
			break
		}
	}
}

func looping() {
	for i := 1; i <= 5; i++ {
		log.Println(i)
	}
	num := 1
	for num <= 5 {
		log.Println(num)
		num++
	}

	numsArray := []int{2, 3, 4, 6}
	for _, numbers := range numsArray {
		log.Println(numbers)
	}

	xValue := 1
	for true {
		if xValue == 4 {
			break
		}
		log.Println(xValue)
		xValue++
	}
}
