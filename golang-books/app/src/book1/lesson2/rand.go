package book1Lesson2

import (
	"fmt"
	"math/rand"
	"time"
)

func pracRand() {
	fmt.Println("----------")
	fmt.Println("rand.go")
	fmt.Println("----------")

	rand.Seed(time.Now().UnixNano())

	var num = rand.Intn(10) + 1
	fmt.Println("Rand 1 : ", num)

	num = rand.Intn(10) + 1
	fmt.Println("Rand 2 : ", num)

}
