package book1Lesson9

import (
	"fmt"
)

func index() {
	fmt.Println("----------")
	fmt.Println("index.go")
	fmt.Println("----------")

	message := "HelloWorld"
	c := message[5]
	// message[5] = "d"

	fmt.Printf("%c \n", c)

}
