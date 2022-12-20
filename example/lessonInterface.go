package main

import (
	"log"
	"math"
)

type Animal interface {
	AngrySound()
	HappySound()
}

type Volume interface {
	Volume() float64
}

type value struct {
	radius float64
	height float64
}

type Cat string

func (v value) Volume() float64 {
	return math.Pi * v.radius * v.radius * v.height
}

func (c Cat) Attack() {
	log.Println("Cat attack its Prey")
}

func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	log.Println("Cat says Hissss")
}

func (c Cat) HappySound() {
	log.Println("Cat says Purrr")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	var sKitty Cat = kitty.(Cat)
	sKitty.Attack()
	log.Println("Cats name :", sKitty.Name())

	var mathOper Volume
	mathOper = value{12, 4}
	log.Println("Volume =", mathOper.Volume())
}
