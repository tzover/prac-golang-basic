package book1Lesson8

import (
	"fmt"
)

func constant() {
	fmt.Println("----------")
	fmt.Println("constant.go")
	fmt.Println("----------")

	const (
		distance      = 24000000000000000000
		lightSpeed    = 299792
		secondsPerDay = 86400 // 1day
		days          = distance / lightSpeed / secondsPerDay
	)

	// fmt.Println(distance) // overflow
	fmt.Println("lightSpeed : ", days, "day")

}
