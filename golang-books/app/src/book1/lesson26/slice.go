package book1Lesson26

import "fmt"

func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func pracSlice() {
	fmt.Println("----------")
	fmt.Println("slice.go")
	fmt.Println("----------")

	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune", "Pluto",
	}
	reclassify(&planets)
	fmt.Printf("%+v \n", planets)

}
