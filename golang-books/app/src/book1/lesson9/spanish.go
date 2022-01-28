package book1Lesson9

import (
	"fmt"
	"unicode/utf8"
)

func spanish() {
	fmt.Println("----------")
	fmt.Println("spanish.go")
	fmt.Println("----------")

	question := "¿átʃe"
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")
	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes \n", c, size)

}
