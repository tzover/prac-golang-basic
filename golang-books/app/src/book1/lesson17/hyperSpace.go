package book1Lesson17

import (
	"fmt"
	"strings"
)

func hyperSpace(worlds []string) {
	fmt.Println("----------")
	fmt.Println("hyperSpace.go")
	fmt.Println("----------")

	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}

}
