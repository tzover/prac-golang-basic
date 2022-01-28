package book1Lesson24

import (
	"fmt"
)

type startship struct {
	laser
}

func stardate(t stardater) float64 {
	fmt.Println("----------")
	fmt.Println("startShip.go")
	fmt.Println("----------")

	doy := float64(t.YearDay())
	h := float64(t.Hour())

	return 1000 + doy + h

}
