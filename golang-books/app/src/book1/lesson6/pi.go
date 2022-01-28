package book1Lesson6

import (
	"fmt"
	"math"
)

func pi() {
	fmt.Println("----------")
	fmt.Println("pi.go")
	fmt.Println("----------")

	var (
		pi64         = math.Pi
		pi32 float32 = math.Pi
	)

	fmt.Println("float64 : ", pi64)
	fmt.Println("float32 : ", pi32)
}
