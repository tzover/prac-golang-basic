package book1Lesson6

import (
	"fmt"
	"math"
)

func float() {
	fmt.Println("----------")
	fmt.Println("float.go")
	fmt.Println("----------")

	third := 1.0 / 3.0

	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2

	fmt.Println(piggyBank)
	fmt.Println(piggyBank == 0.3)                 //false
	fmt.Println(math.Abs(piggyBank-0.3) < 0.0001) // true

}
