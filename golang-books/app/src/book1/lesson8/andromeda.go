package book1Lesson8

import (
	"fmt"
	"math/big"
)

func andromeda() {
	fmt.Println("----------")
	fmt.Println("andromeda.go")
	fmt.Println("----------")

	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400) // 1day

	distance := new(big.Int)

	distance.SetString("24000000000000000000", 10)

	fmt.Println("Distance : ", distance, "km")

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("lightSpeed : ", days, "day")

}
