package book1Lesson16

import "fmt"

func pracArray() {
	fmt.Println("----------")
	fmt.Println("array.go")
	fmt.Println("----------")

	var planets [8]string

	planets[0] = "水星"
	planets[1] = "金星"
	planets[2] = "地球"

	fmt.Println(planets)

	earth := planets[2]

	fmt.Println(earth)

	fmt.Println(len(planets))
	fmt.Println(planets[5] == "")
}
