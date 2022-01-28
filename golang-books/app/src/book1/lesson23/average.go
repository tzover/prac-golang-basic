package book1Lesson23

import "fmt"

func (t temperature) average() celsius {
	fmt.Println("----------")
	fmt.Println("average.go")
	fmt.Println("----------")

	return (t.high + t.low) / 2
}
