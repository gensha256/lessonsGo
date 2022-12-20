package main

import (
	stuff "example/app/mypackage"
	"log"
	"reflect"
)

func main() {
	log.Println("Hello", stuff.Name)
	intArr := []int{2, 3, 4, 6, 8}
	strArr := stuff.IntArrToStrArr(intArr)
	log.Println(strArr)
	log.Println(reflect.TypeOf(strArr))
}
