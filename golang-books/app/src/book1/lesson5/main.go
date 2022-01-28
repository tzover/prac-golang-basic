package book1Lesson5

import (
	"fmt"
	"math/rand"
	"time"
)

func Main() {
	rand.Seed(time.Now().UnixNano())
	var (
		spaceline string
		tripType  string
		money     string = "$"
		distance  int    = 62100000
		maxSpeed  int    = 30
		minSpeed  int    = 16
		maxPrice  int    = 50
		minPrice  int    = 36
	)
	fmt.Println("******* Package book1Lesson5 *******")

	fmt.Printf("%-20v %-5v %20v %10v \n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("===============================================================================")
	for count := 0; count < 10; count++ {
		var (
			numName int = rand.Intn(100) + 1
			numType int = rand.Intn(100) + 1
			speed   int = rand.Intn(maxSpeed-minSpeed) + minSpeed
			date    int = distance / speed / 86400
			price   int = rand.Intn(maxPrice-minPrice) + minPrice
		)
		switch {
		case numName%3 == 0:
			spaceline = "aaaaa"
		case numName%3 == 1:
			spaceline = "bbbbb"
		default:
			spaceline = "ccccc"
		}

		if numType%2 == 0 {
			tripType = "Round-trip"
			price *= 2
		} else {
			tripType = "One-way"
		}

		fmt.Printf("%-20v %2v %20v %10v %v \n", spaceline, date, tripType, money, price)
	}

}
