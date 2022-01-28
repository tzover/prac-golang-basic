package book1Lesson7

import "fmt"

func inspect() {
	fmt.Println("----------")
	fmt.Println("inspect.go")
	fmt.Println("----------")

	year := 2018
	days := 365.2425

	fmt.Printf("%T 型 : %v \n", year, year)
	fmt.Printf("%T 型 : %[1]v \n", days)
}
