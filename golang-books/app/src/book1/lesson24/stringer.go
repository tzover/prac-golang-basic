package book1Lesson24

import "fmt"

type location struct {
	lat, long float64
}

func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func stringer() {
	fmt.Println("----------")
	fmt.Println("stringer.go")
	fmt.Println("----------")

	curiosity := location{lat: -4.5895, long: 137.4417}

	fmt.Println(curiosity)
}
