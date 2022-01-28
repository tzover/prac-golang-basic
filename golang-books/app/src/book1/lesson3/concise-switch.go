package book1Lesson3

import "fmt"

func conciseSwitch() {
	fmt.Println("----------")
	fmt.Println("concise-switch.go")
	fmt.Println("----------")

	var (
		command string = "go inside"
	)

	switch command {
	case "go east":
		fmt.Println("case : go east")
	case "enter cave", "go inside":
		fmt.Println("case : enter cave || go inside")
	case "read sign":
		fmt.Println("case : read sign")
	default:
		fmt.Println("case : default")
	}

}
