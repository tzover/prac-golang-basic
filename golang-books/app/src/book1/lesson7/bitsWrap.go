package book1Lesson7

import "fmt"

func bitsWrap() {
	fmt.Println("----------")
	fmt.Println("bitsWrap.go")
	fmt.Println("----------")

	var (
		blue uint8 = 255
	)

	fmt.Printf("%08b\n", blue)

	blue++

	fmt.Printf("%08b\n", blue)

}
