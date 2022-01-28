package book1Lesson22

import (
	"fmt"
)

func Main() {
	fmt.Println("******* Package book1Lesson22 *******")

	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())

	curiosity := location{lat.decimal(), long.decimal()}
	fmt.Println(curiosity)

	test := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println("test", test)

	spirit := location{lat: -14.5684, long: 175.472636}
	opportunity := location{lat: -1.9462, long: 354.4734}

	dist := mars.distance(spirit, opportunity)

	fmt.Printf("%.2f km \n", dist)
}
