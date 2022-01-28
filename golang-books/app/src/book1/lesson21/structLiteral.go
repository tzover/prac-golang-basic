package book1Lesson21

import (
	"fmt"
)

func structLiteral() {
	fmt.Println("----------")
	fmt.Println("structLiteral.go")
	fmt.Println("----------")

	type location struct {
		lat, long float64
	}

	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity.lat, opportunity.long)
	fmt.Println(opportunity)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight.lat, insight.long)
	fmt.Println(insight)

	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit.lat, spirit.long)
	fmt.Println(spirit)

	fmt.Printf("opportunity : %+v \n", opportunity)
	fmt.Printf("insight     : %+v \n", insight)
	fmt.Printf("spirit      : %+v \n", spirit)

}
