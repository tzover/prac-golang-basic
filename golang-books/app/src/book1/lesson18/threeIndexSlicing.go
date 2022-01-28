package book1Lesson18

import (
	"fmt"
)

func threeIndexSlicing() {
	fmt.Println("----------")
	fmt.Println("threeIndexSlicing.go")
	fmt.Println("----------")

	planets := []string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "planets", len(planets), cap(planets), planets)

	tutorial := planets[0:4:4]
	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "tutorial", len(tutorial), cap(tutorial), tutorial)

	worlds := append(tutorial, "testData")
	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "worlds", len(worlds), cap(worlds), worlds)

	fmt.Println(planets)
}
