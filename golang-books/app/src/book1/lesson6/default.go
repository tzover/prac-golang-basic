package book1Lesson6

import (
	"fmt"
)

func pracDefault() {
	fmt.Println("----------")
	fmt.Println("default.go")
	fmt.Println("----------")

	var (
		price float64
	)

	fmt.Println("float64 : ", price)
	price = 0.0
	fmt.Println("float64 : ", price)
}
