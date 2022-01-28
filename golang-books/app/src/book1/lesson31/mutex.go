package book1Lesson31

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func mutex() {
	fmt.Println("----------")
	fmt.Println("mutex.go")
	fmt.Println("----------")

	mu.Lock()

	defer mu.Unlock()

}
