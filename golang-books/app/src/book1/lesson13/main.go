package book1Lesson13

import (
	"fmt"
)

type celsiusType float64
type kelvinType float64

func Main() {
	fmt.Println("******* Package book1Lesson13 *******")

	celsius()
	celsiusAddition()

	var k kelvinType = 294.0
	c := temperatureTypes(k)
	fmt.Println(k, "°Kは、", c, "°Cです。")

}
