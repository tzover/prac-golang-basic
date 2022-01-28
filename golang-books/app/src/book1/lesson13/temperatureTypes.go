package book1Lesson13

import (
	"fmt"
)

func temperatureTypes(k kelvinType) celsiusType {
	fmt.Println("----------")
	fmt.Println("temperatureTypes.go")
	fmt.Println("----------")

	return celsiusType(k - 273.15)
}
