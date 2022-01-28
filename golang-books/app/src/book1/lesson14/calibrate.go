package book1Lesson14

import "fmt"

func calibrate(s sensorFunc, offset kelvin) sensorFunc {
	fmt.Println("----------")
	fmt.Println("calibrate.go")
	fmt.Println("----------")

	return func() kelvin {
		return s() + offset
	}
}
