package book1Lesson6

import "fmt"

func temperature() {
	fmt.Println("----------")
	fmt.Println("temperature.go")
	fmt.Println("----------")

	celsius := 21.0

	fahrenheit := (celsius * 9.0 / 5.0) + 32

	fmt.Println(fahrenheit, "°F")

}
