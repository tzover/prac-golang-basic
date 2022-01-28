package book1Lesson2

import "fmt"

func shortcut() {
	fmt.Println("----------")
	fmt.Println("shortcut.go")
	fmt.Println("----------")

	var (
		weight = 176.0
		age    = 28
	)
	weight *= 0.3783
	age++
	age--

	fmt.Println(weight, age)
}
