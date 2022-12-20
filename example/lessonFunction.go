package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var humArr = [7]string{

	" +---+\n" +
		"0     |\n" +
		"     |\n" +
		"     |\n" +
		"   ===\n",
	" +---+\n" +
		" 0   |\n" +
		" |    |\n" +
		"     |\n" +
		"   ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|    |\n" +
		"     |\n" +
		"   ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\   |\n" +
		"     |\n" +
		"   ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\    |\n" +
		"/     |\n" +
		"   ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\    |\n" +
		"/ \\    |\n" +
		"   ===\n",
}

var wordsArr = [7]string{
	"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER", "ZODIAC", "ZOMBIE", "FLUFF",
}

var randWord string
var guessedLetter string
var correctLetter []string
var wrongGuesses []string

func getRandomWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randWord = wordsArr[rand.Intn(7)]
	correctLetter = make([]string, len(randWord))
	return randWord
}

func showBoard() {
	fmt.Println(humArr[len(wrongGuesses)])
	fmt.Print("Secret Word:")
	for _, v := range correctLetter {
		if v == "" {
			fmt.Println("_")
		} else {
			fmt.Print(v)
		}
	}
	log.Print("\nIncorrect Guesses :")
	if len(wrongGuesses) > 0 {
		for _, value := range wrongGuesses {
			fmt.Print(value + " ")
		}
	}
}

func getUserLetter() string {
	reader := bufio.NewReader(os.Stdin)
	var guess string
	for true {
		fmt.Println("\nGuess a Letter :")
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.ToUpper(guess)
		guess = strings.TrimSpace(guess)
		var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

		if len(guess) != 1 {
			log.Println("Please Enter Only One Letter")
		} else if !isLetter(guess) {
			log.Println("Please Enter a Letter")
		} else if strings.Contains(guessedLetter, guess) {
			fmt.Println("Please Enter a Letter you havent guessed")
		} else {
			return guess
		}
	}
	return guess
}

func getAllIndex(theStr, subStr string) (indices []int) {
	if len(subStr) == 0 || (len(theStr) == 0) {
		return indices
	}
	offset := 0
	for {
		i := strings.Index(theStr[offset:], subStr)
		if i == -1 {
			return indices
		}
		offset += i
		indices = append(indices, offset)
		offset += len(subStr)
	}
}

func updateCorrectLetters(letter string) {
	indexMatches := getAllIndex(randWord, letter)
	for _, value := range indexMatches {
		correctLetter[value] = letter
	}
}

func sliceHasEmptys(theSlice []string) bool {
	for _, value := range theSlice {
		if len(value) == 0 {
			return true
		}
	}
	return false
}

func main() {
	log.Println(getRandomWord())

	for true {
		showBoard()
		guess := getUserLetter()
		if strings.Contains(randWord, guess) {
			updateCorrectLetters(guess)
			if sliceHasEmptys(correctLetter) {
				log.Println("More letters to Guess")
			} else {
				log.Println("Yes the secret word is", randWord)
				break

			}
		} else {
			guessedLetter += guess
			wrongGuesses = append(wrongGuesses, guess)

			if len(wrongGuesses) >= 6 {
				log.Println("Sorry you dead! The word is", randWord)
				break
			}
		}
	}

}

var nums = []int{12, 43, 6, 7, 90}

func getArraySum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
