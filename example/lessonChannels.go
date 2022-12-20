package main

import "log"

func main() {
	getChannels()
}

func printTo15() {
	for i := 1; i <= 15; i++ {
		log.Println("Func 1 :", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		log.Println("Func 2 :", i)
	}
}

func firstNums(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func secondNums(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func getChannels() {
	firstChannel := make(chan int)
	secondChannel := make(chan int)
	go firstNums(firstChannel)
	go secondNums(secondChannel)
	log.Println(<-firstChannel)
	log.Println(<-firstChannel)
	log.Println(<-firstChannel)
	log.Println(<-secondChannel)
	log.Println(<-secondChannel)
	log.Println(<-secondChannel)
}
