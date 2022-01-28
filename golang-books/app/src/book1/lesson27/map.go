package book1Lesson27

import (
	"fmt"
)

func pracMap() {
	fmt.Println("----------")
	fmt.Println("map.go")
	fmt.Println("----------")

	var soup map[string]int
	fmt.Println(soup == nil)

	measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}

	for ingredient, measurement := range soup {
		fmt.Println(ingredient, measurement)
	}

}
