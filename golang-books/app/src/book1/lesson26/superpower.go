package book1Lesson26

import "fmt"

func superpower() {
	fmt.Println("----------")
	fmt.Println("superpower.go")
	fmt.Println("----------")

	superpowers := &[3]string{"flight", "invisibility", "super strength"}

	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])

}
