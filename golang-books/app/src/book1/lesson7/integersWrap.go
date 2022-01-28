package book1Lesson7

import "fmt"

func integersWrap() {
	fmt.Println("----------")
	fmt.Println("integersWrap.go")
	fmt.Println("----------")

	var (
		red  uint8 = 255
		blue int8  = 127
	)

	red++
	fmt.Println(red) // 0

	blue++
	fmt.Println(blue) // -128

}
