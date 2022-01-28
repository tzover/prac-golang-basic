package goroutine

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (ch *Counter) Inc(key string) {
	ch.mux.Lock()
	defer ch.mux.Unlock()
	ch.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func syncMutex() {

	ch := Counter{v: make(map[string]int)}

	go func() {
		for i := 0; i < 10; i++ {
			ch.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			ch.Inc("Key")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(ch, ch.Value("Key"))

}
