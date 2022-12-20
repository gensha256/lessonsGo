package main

import (
	"log"
)

func main() {
	num := 10
	log.Println(num)
	log.Println("Address in memory num: ", &num)
	var numPoint *int = &num

	log.Println("Address in memory numPoint: ", numPoint)
	log.Println("Value numPoint: ", *numPoint)
	*numPoint = 24
	log.Println("numPoint new value: ", *numPoint)
	log.Println("Before changeValue: ", num)
	changeValue(&num)
	log.Println("After changeValue: ", num)

	arr := [4]int{1, 4, 6, 8}
	arrValues(&arr)
	log.Println("Changed array: ", arr)

	iSlice := []float64{11, 13, 17}
	log.Println("Average :", getAverage(iSlice...))
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, value := range nums {
		sum += value
	}
	return sum / numSize
}

func arrValues(arr *[4]int) {
	for x := 0; x < len(arr); x++ {
		arr[x] = arr[x] * 2
	}
}

func changeValue(n *int) {
	*n = 14
}
