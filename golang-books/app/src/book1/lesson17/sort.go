package book1Lesson17

import (
	"fmt"
	"sort"
)

type StringSlice []string

func pracSort() {
	fmt.Println("----------")
	fmt.Println("sort.go")
	fmt.Println("----------")

	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	sort.StringSlice(planets).Sort()

	fmt.Println(planets)

}
