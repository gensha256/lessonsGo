package main

import "log"

func main() {

	samp := 1
	changeVar := func() {
		samp += 1
	}
	changeVar()
	log.Println("Samp =", samp)

	newValue := incrementor()
	log.Println(newValue())
	log.Println(newValue())
}

func incrementor() func() int {
	inc := 0
	return func() int {
		inc++
		return inc
	}
}
