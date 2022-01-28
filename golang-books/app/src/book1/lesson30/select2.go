package book1Lesson30

import (
	"fmt"
	"math/rand"
	"time"
)

func select2() {
	fmt.Println("----------")
	fmt.Println("select2.go")
	fmt.Println("----------")
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
}
