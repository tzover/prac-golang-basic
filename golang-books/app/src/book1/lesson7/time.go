package book1Lesson7

import (
	"fmt"
	"time"
)

func pracTime() {
	fmt.Println("----------")
	fmt.Println("time.go")
	fmt.Println("----------")

	future := time.Unix(12622780800, 0)
	fmt.Println(future)
}
