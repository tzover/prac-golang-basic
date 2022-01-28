package statement

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func pracSwitch() {

	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac os !!")
	case "windows":
		fmt.Println("Windows...")
	default:
		fmt.Println("boo...", os)
	}

	t := time.Now()
	fmt.Println(t)

	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}
