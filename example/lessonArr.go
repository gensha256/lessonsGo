package main

import (
	"log"
)

func main() {
	aStr := "abcde"
	runeArr := []rune(aStr)
	for _, value := range runeArr {
		log.Println("Rune array:", value)
	}
	byteArr := []byte{'a', 'b', 'c'}
	byteStr := string(byteArr)
	log.Println("Byte array strings", byteStr)

}

func useArrays() {
	var arr [5]int
	arr[0] = 1

	sArr := [5]int{1, 3, 5, 6, 7}
	log.Println("Index 0: ", sArr[0])
	log.Println("Length arr", len(sArr))
	for i := 0; i < len(sArr); i++ {
		log.Println(sArr[i])
	}
	for i, value := range sArr {
		log.Println("index:", i, value)
	}
	dimenArr := [2][2]int{
		{1, 2},
		{3, 6},
	}
	log.Println(dimenArr)
	for iter1 := 0; iter1 < 2; iter1++ {
		for iter2 := 0; iter2 < 2; iter2++ {
			log.Println(dimenArr[iter1][iter2])
		}
	}
}

func useSlices() {
	slArray := make([]string, 6)
	slArray[0] = "Jack"
	slArray[1] = "America"
	slArray[2] = "Casino"
	slArray[3] = "Universe"
	slArray[4] = "Simulated"
	slArray[5] = "The"

	for i := 0; i < len(slArray); i++ {
		log.Println("Slices value:", slArray[i])
	}
	slc := [5]int{1, 3, 6, 7, 8}
	slcArr := slc[:3]
	slc[0] = 10
	log.Println(slcArr)

	slcAppend := append(slcArr, 45)
	log.Println("Its slice ", slcAppend)
	log.Println("Its main array ", slc)

}
