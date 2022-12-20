package main

import (
	"log"
	"strings"
	"unicode/utf8"
)

func main() {
	runeStr := "abcdefg"
	log.Println("Rune count: ", utf8.RuneCountInString(runeStr))
	for i, runeValue := range runeStr {
		log.Printf("%d : %#U : %c\n", i, runeValue, runeValue)
	}
}

func usingString() {
	str := "A world"
	replacer := strings.NewReplacer("A", "Another")
	repStr := replacer.Replace(str)
	log.Println(repStr)

	log.Println("Length: ", len(repStr))
	log.Println("Contains Another: ", strings.Contains(repStr, "Another"))
	log.Println("o index: ", strings.Index(repStr, "o"))
	log.Println("o index: ", strings.Replace(repStr, "o", "0", -1))

	strSp := "\nSome words "
	strSp = strings.TrimSpace(strSp)
	log.Println(strSp)
	log.Println("Split: ", strings.Split("a-b-c-d", "-"))
	log.Println("Prefix: ", strings.HasPrefix("coconut", "coco"))
	log.Println("Suffix: ", strings.HasSuffix("coconut", "nut"))
}
