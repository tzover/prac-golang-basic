package book1Lesson4

import (
	"fmt"
	"math/rand"
	"time"
)

func shortIf() {
	fmt.Println("----------")
	fmt.Println("shortIf.go")
	fmt.Println("----------")

	rand.Seed(time.Now().UnixNano())

	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}
}
