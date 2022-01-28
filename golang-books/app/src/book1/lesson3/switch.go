package book1Lesson3

import "fmt"

func pracSwitch() {
	fmt.Println("----------")
	fmt.Println("switch.go")
	fmt.Println("----------")

	var (
		room string = "lake"
	)

	switch {
	case room == "cave":
		fmt.Println("case : room == cave")
	case room == "lake":
		fmt.Println("case : room == lake")
		fallthrough
	case room == "underwater":
		fmt.Println("case : room == underwater")
	}

}
