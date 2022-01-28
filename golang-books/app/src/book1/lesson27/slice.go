package book1Lesson27

import (
	"fmt"
)

func pracSlice() {
	fmt.Println("----------")
	fmt.Println("slice.go")
	fmt.Println("----------")

	var soup []string
	fmt.Println(soup == nil)

	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	fmt.Println(len(soup))

	soup = append(soup, "onion", "carrot", "celery")

	fmt.Println(soup)

}
