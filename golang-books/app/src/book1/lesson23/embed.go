package book1Lesson23

import "fmt"

// type report struct {
// 	sol int
// 	temperature
// 	location
// }

func embed() {
	fmt.Println("----------")
	fmt.Println("embed.go")
	fmt.Println("----------")

	report := report{
		sol:         15,
		location:    location{lat: -4.5895, long: 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}

	fmt.Printf("Average : %v °C\n", report.average())

	fmt.Printf("high : %v °C\n", report.high)

	report.high = 90

	fmt.Printf("high : %v °C\n", report.temperature.high)

}
