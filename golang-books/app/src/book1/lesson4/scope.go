package book1Lesson4

import (
	"fmt"
	"math/rand"
	"time"
)

func scope() {
	fmt.Println("----------")
	fmt.Println("scope.go")
	fmt.Println("----------")

	rand.Seed(time.Now().UnixNano())

	var (
		count int = 0
	)

	for count < 10 {
		var num int = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	}
}
