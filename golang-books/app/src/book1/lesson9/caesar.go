package book1Lesson9

import (
	"fmt"
)

func caesar() {
	fmt.Println("----------")
	fmt.Println("caesar.go")
	fmt.Println("----------")

	c := 'a'
	c = c + 3 // Error
	fmt.Printf("%c \n", c)

}
