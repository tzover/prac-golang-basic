package book1Lesson10

import (
	"fmt"
)

func marsAge() {
	fmt.Println("----------")
	fmt.Println("marsAge.go")
	fmt.Println("----------")

	age := 28
	marsAge := float64(age)
	marsDays := 687.0
	earthDays := 365.2425

	marsAge = marsAge * earthDays / marsDays

	fmt.Println("私は火星では", int(marsAge), "歳です。")

}
