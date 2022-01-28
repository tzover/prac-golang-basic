package book1Lesson10

import (
	"fmt"
	"math"
)

func ariane() {
	fmt.Println("----------")
	fmt.Println("ariane.go")
	fmt.Println("----------")

	var (
		bh float64 = 32768 // 32767
		b          = int16(bh)
	)

	// Check
	if bh < math.MinInt16 || bh > math.MaxInt16 {
		fmt.Println("Error")
	}

	fmt.Println(b)
}
