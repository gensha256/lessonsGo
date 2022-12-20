package main

import "log"

func main() {
	var heroes map[string]string

	heroes = make(map[string]string)
	villians := make(map[string]string)

	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clack Kent"
	heroes["The Flash"] = "Barry Allen"
	villians["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{1: "Krypto", 2: "Bat Hound"}

	log.Println("Batman is :", heroes["Batman"])

	log.Println("Chip :", superPets[3])

	_, ok := superPets[3]
	log.Println("Is there a 3rd pet :", ok)

	for key, value := range heroes {
		log.Printf("%s is %s\n", key, value)
	}
	delete(heroes, "The Flash")
}
