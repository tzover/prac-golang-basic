package book1Lesson4

import (
	"fmt"
)

func shortLoop() {
	fmt.Println("----------")
	fmt.Println("shortLoop.go")
	fmt.Println("----------")

	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
	// fmt.Println(count) // Error
}
