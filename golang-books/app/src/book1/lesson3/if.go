package book1Lesson3

import (
	"fmt"
	"strings"
)

func pracIf() {
	fmt.Println("----------")
	fmt.Println("if.go")
	fmt.Println("----------")

	var (
		command = "go east"
	)

	if command == "go east" {
		fmt.Println("command == go east : ", command == "go east")
	} else if command == "go inside" {
		fmt.Println("command == go inside : ", command == "go inside")
	} else {
		fmt.Println("I don't know")
	}

	fmt.Println("if strings.Contains(command, go)")

	if strings.Contains(command, "go") {
		fmt.Println("output : ", strings.Contains(command, "go"))
	} else {
		fmt.Println("output : ", strings.Contains(command, "go"))
	}
}
