package book1Lesson26

import "fmt"

func pracType() {
	fmt.Println("----------")
	fmt.Println("type.go")
	fmt.Println("----------")

	answer := 40
	address := &answer

	fmt.Printf("addressの型は %T . \n", address)

}
