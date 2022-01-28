package book1Lesson9

import (
	"fmt"
)

func spanishRange() {
	fmt.Println("----------")
	fmt.Println("spanishRange.go")
	fmt.Println("----------")

	question := "¿átʃe"

	for i, c := range question {
		fmt.Printf("%v %c \n", i, c)
	}

}
