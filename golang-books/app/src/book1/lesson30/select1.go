package book1Lesson30

import (
	"fmt"
	"time"
)

func select1() {
	fmt.Println("----------")
	fmt.Println("select1.go")
	fmt.Println("----------")

	timeout := time.After(2 * time.Second)

	c := make(chan int)

	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, "done")
		case <-timeout:
			fmt.Println("over...")
			return
		}
	}

}
