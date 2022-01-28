package book1Lesson30

import (
	"fmt"
)

func sleepyGophers2(id int, c chan int) {
	// time.Sleep(3 * time.Second)
	select2()
	fmt.Println("...", id, " snore ...")
	c <- id
}

func simpleChan() {
	fmt.Println("----------")
	fmt.Println("simpleChan.go")
	fmt.Println("----------")

	c := make(chan int)

	for i := 0; i < 5; i++ {
		go sleepyGophers2(i, c)
	}

	for i := 0; i < 5; i++ {
		gopherID := <-c
		fmt.Println("gopher ", gopherID, "done")
	}
}
