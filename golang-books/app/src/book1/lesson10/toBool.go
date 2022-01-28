package book1Lesson10

import (
	"fmt"
)

func toBool() {
	fmt.Println("----------")
	fmt.Println("toBool.go")
	fmt.Println("----------")

	yesNo := "no"
	launch := (yesNo == "yes")

	fmt.Println("Ready for launch : ", launch)

}
