package book1Lesson27

import (
	"fmt"
)

func fn() {
	fmt.Println("----------")
	fmt.Println("fn.go")
	fmt.Println("----------")

	var fn func(a, b int) int
	fmt.Println(fn == nil)
}
