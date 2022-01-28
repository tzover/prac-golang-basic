package book1Lesson3

import (
	"fmt"
	"time"
)

func countdown() {
	fmt.Println("----------")
	fmt.Println("countdown.go")
	fmt.Println("----------")

	var (
		count int = 10
	)

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	fmt.Println("Fly away!")
}
