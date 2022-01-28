package book1Lesson19

import (
	"fmt"
)

func whoops() {
	fmt.Println("----------")
	fmt.Println("whoops.go")
	fmt.Println("----------")

	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	planetsMark2 := planets
	planets["Earth"] = "whoops"

	fmt.Println("planets : ", planets)
	fmt.Println("planetsMark2", planetsMark2)

	delete(planets, "Earth")

	fmt.Println("planets : ", planets)
	fmt.Println("planetsMark2", planetsMark2)

}
