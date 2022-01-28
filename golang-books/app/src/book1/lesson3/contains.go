package book1Lesson3

import (
	"fmt"
	"strings"
)

func contains() {
	fmt.Println("----------")
	fmt.Println("contains.go")
	fmt.Println("----------")

	fmt.Println("You are in a dim cave (薄暗い洞窟にいる)")

	var (
		command = "walk outside"
		exit    = strings.Contains(command, "outside")
	)

	fmt.Println("Leave the cave : ", exit)
}
