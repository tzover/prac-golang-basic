package book1Lesson21

import (
	"fmt"
)

func pracStruct() {
	fmt.Println("----------")
	fmt.Println("struct.go")
	fmt.Println("----------")

	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long)

	fmt.Println(curiosity)

}
