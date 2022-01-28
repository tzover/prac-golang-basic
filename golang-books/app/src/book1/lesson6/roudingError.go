package book1Lesson6

import "fmt"

func roudingError() {
	fmt.Println("----------")
	fmt.Println("roudingError.go")
	fmt.Println("----------")

	celsius := 21.0

	fmt.Println((celsius/5.0*9.0)+32, "°F")
	fmt.Println((9.0/5.0*celsius)+32, "°F")

}
