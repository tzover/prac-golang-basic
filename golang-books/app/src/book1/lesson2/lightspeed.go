package book1Lesson2

import "fmt"

func lightspeed() {
	fmt.Println("----------")
	fmt.Println("lightspeed.go")
	fmt.Println("----------")

	var (
		lightspeed = 299792   // km/sec
		distance   = 56000000 // km
	)

	fmt.Println(distance/lightspeed, "sec")

	distance = 401000000
	fmt.Println(distance/lightspeed, "sec")

}
