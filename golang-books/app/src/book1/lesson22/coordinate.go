package book1Lesson22

import (
	"fmt"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

func (c coordinate) decimal() float64 {
	fmt.Println("----------")
	fmt.Println("coordinate.go")
	fmt.Println("----------")

	sign := 1.0

	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}

	return sign * (c.d + c.m/60 + c.s/3600)

}
