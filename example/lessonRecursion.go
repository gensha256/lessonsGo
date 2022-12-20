package main

import "log"

func main() {
	log.Println("Factorial number 10 =", factorial(10))
	log.Println("Factorial number 5 =", factorial(5))
	log.Println("Factorial number 8 =", factorial(8))

}

func factorial(num uint64) uint64 {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}
