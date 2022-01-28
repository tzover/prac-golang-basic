package book1Lesson10

import (
	"fmt"
)

func launch() {
	fmt.Println("----------")
	fmt.Println("launch.go")
	fmt.Println("----------")

	launch := false
	launchText := fmt.Sprintf("%v \n", launch)
	fmt.Println("Ready for launch : ", launchText)

	var yesNo string

	if launch {
		yesNo = "yes"
	} else {
		yesNo = "No"
	}
	fmt.Println("Ready for launch : ", yesNo)

}
