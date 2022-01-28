package book1Lesson4

import (
	"fmt"
)

func loop() {
	fmt.Println("----------")
	fmt.Println("loop.go")
	fmt.Println("----------")

	var (
		count int = 0
	)

	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count)
}
