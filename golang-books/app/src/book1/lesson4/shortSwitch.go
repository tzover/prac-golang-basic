package book1Lesson4

import (
	"fmt"
	"math/rand"
	"time"
)

func shortSwitch() {
	fmt.Println("----------")
	fmt.Println("shortSwitch.go")
	fmt.Println("----------")

	rand.Seed(time.Now().UnixNano())

	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}

}
