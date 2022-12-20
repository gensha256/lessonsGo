package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	log.Println("Enter first number")
	firstNum, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	firstNum = strings.TrimSpace(firstNum)
	iFirstNum, err := strconv.Atoi(firstNum)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Enter second number")
	secondNum, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	secondNum = strings.TrimSpace(secondNum)
	iSecondNum, err := strconv.Atoi(secondNum)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s + %s = %d\n", firstNum, secondNum, iFirstNum+iSecondNum)
}

func useMath() {
	log.Println("Abs(-10) =", math.Abs(-10))
	log.Println("Pow(4, 2) =", math.Pow(4, 2))
	log.Println("Cbrt(8) =", math.Cbrt(8))
	log.Println("Ceil(4.4) =", math.Ceil(4.4))
	log.Println("Floor(4.4) =", math.Floor(4.4))
	log.Println("Round(15.2) =", math.Round(15.2))
	log.Println("Max(15, 40) =", math.Max(15, 40))

}
