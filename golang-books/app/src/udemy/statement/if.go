package statement

import (
	"fmt"
	"math/rand"
	"time"
)

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "ng"
	}
}

func ifElse() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(1000)
	fmt.Println("num :", num)
	if num%2 == 0 {
		fmt.Println("by2")
	} else if num%3 == 0 {
		fmt.Println("by3")
	} else {
		fmt.Println("else")
	}

	x, y, z := 10, 10, 20

	fmt.Println("x, y, z", x, y, z)

	if x == 10 && y == 10 {
		fmt.Println("&&")
	}
	if x == 10 || z == 10 {
		fmt.Println("||")
	}

	// function
	result := by2(num)
	if result == "ok" {
		fmt.Println("Great")
	} else {
		fmt.Println("boo...")
	}

	if result2 := by2(num); result2 == "ok" {
		fmt.Println("Great2")
	} else {
		fmt.Println("boo...boo...")
	}
}
