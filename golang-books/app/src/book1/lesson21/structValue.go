package book1Lesson21

import (
	"fmt"
)

func structValue() {
	fmt.Println("----------")
	fmt.Println("structValue.go")
	fmt.Println("----------")

	type location struct {
		lat, long float64
	}

	insight := location{lat: 4.5, long: 135.9}
	insight2 := insight
	fmt.Printf("%-15v: %+v \n", "insight", insight)
	fmt.Printf("%-15v: %+v \n", "insight2", insight2)

	insight2.long = 123.456
	fmt.Printf("%-15v: %+v \n", "insight", insight)
	fmt.Printf("%-15v: %+v \n", "insight2", insight2)

}
