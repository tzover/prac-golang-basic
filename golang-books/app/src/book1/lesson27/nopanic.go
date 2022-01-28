package book1Lesson27

import (
	"fmt"
)

func noPanic() {
	fmt.Println("----------")
	fmt.Println("noPanic.go")
	fmt.Println("----------")

	var nowhere *int
	fmt.Println(nowhere)

	if nowhere != nil {
		fmt.Println(*nowhere)
	}

}
