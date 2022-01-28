package book1Lesson4

import (
	"fmt"
	"math/rand"
)

var (
	era = "AD" // Can be used with all packages
)

func scopeRules() {
	fmt.Println("----------")
	fmt.Println("scopeRules.go")
	fmt.Println("----------")

	var (
		year int = 2018
	)

	switch month := rand.Intn(12) + 1; month {

	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
}
