package book1Lesson12

import (
	"fmt"
)

func kelvin() {
	fmt.Println("----------")
	fmt.Println("kelvin.go")
	fmt.Println("----------")

	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "°Kは、", celsius, "°Cです")

}

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
