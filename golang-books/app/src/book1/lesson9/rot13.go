package book1Lesson9

import (
	"fmt"
)

func rot13() {
	fmt.Println("----------")
	fmt.Println("rot13.go")
	fmt.Println("----------")

	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ {

		c := message[i]

		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c \n", c)
	}

}
