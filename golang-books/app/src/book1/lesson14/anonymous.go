package book1Lesson14

import (
	"fmt"
	"time"
)

func anonymous() {
	fmt.Println("----------")
	fmt.Println("anonymous.go")
	fmt.Println("----------")
	time.Sleep(time.Second)
	func() {
		fmt.Println("Functions anonymous")
	}()
}
