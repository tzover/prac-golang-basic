package book1Lesson30

import (
	"fmt"
	"time"
)

func sleepygopher() {
	fmt.Println("----------")
	fmt.Println("sleepygopher.go")
	fmt.Println("----------")

	for i := 0; i < 5; i++ {
		go sleepyGopher()
	}
	time.Sleep(4 * time.Second)
}

func sleepyGopher() {
	time.Sleep(3 * time.Second)
	fmt.Println("... snore ...")
}
