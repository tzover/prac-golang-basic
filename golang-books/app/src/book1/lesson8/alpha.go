package book1Lesson8

import (
	"fmt"
)

func alpha() {
	fmt.Println("----------")
	fmt.Println("alpha.go")
	fmt.Println("----------")

	const (
		lightSpeed    = 299792
		secondsPerDay = 86400 // 1day
	)

	var distance int64 = 41.3e12

	fmt.Println("Distance : ", distance, "km")

	days := distance / lightSpeed / secondsPerDay

	fmt.Println("lightSpeed : ", days, "day")

}
