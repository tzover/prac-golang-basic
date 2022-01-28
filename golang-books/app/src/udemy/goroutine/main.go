package goroutine

import (
	"fmt"
	"sync"
)

func normal(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func goroute(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func Main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go normal("World", &wg)
	goroute("Hello")
	wg.Wait()

	pracChannel()
	buffChannel()
	pracChannel2()
	fanOutFanIn()
	// pracChannel3()
	defaultSelection()
	syncMutex()
}
