package book1Lesson23

import "fmt"

type sol int

type report struct {
	sol
	temperature
	location
}

func (s sol) days(s2 sol) int {
	fmt.Println("----------")
	fmt.Println("sol.go")
	fmt.Println("----------")

	days := int(s2 - s)

	if days < 0 {
		days = -days
	}
	return days
}
