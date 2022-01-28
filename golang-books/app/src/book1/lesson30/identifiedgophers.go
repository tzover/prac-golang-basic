package book1Lesson30

import (
	"fmt"
	"time"
)

func identifiedgophers() {
	fmt.Println("----------")
	fmt.Println("identifiedgophers.go")
	fmt.Println("----------")

	for i := 0; i < 5; i++ {
		go sleepyGophers(i)
	}
	time.Sleep(4 * time.Second)
}

func sleepyGophers(id int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...", id, " snore ...")
}
