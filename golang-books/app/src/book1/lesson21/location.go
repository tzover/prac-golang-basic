package book1Lesson21

import (
	"fmt"
)

func location() {
	fmt.Println("----------")
	fmt.Println("location.go")
	fmt.Println("----------")

	type curiosity struct {
		lat  float64
		long float64
	}

	var spirit curiosity
	spirit.lat = -14.5684
	spirit.long = 175.472636

	fmt.Println(spirit.lat, spirit.long)
	fmt.Println(spirit)

	var opportunity curiosity
	opportunity.lat = -1.9462
	opportunity.long = 354.4734

	fmt.Println(opportunity.lat, opportunity.long)
	fmt.Println(opportunity)

}
