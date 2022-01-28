package book1Lesson16

import "fmt"

func arrayValue() {
	fmt.Println("----------")
	fmt.Println("arrayValue.go")
	fmt.Println("----------")

	planets := [...]string{
		"水星",
		"金星",
		"地球",
		"火星",
		"木星",
		"土星",
		"天王星",
		"海王星",
	}
	fmt.Println(planets)

	copyPlanets := planets

	planets[2] = "test"

	fmt.Println(planets)
	fmt.Println(copyPlanets)

}
