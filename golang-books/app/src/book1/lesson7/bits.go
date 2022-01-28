package book1Lesson7

import "fmt"

func bits() {
	fmt.Println("----------")
	fmt.Println("bits.go")
	fmt.Println("----------")

	var (
		green uint8 = 3
	)

	fmt.Printf("%08b\n", green)

	green++

	fmt.Printf("%08b\n", green)

}
