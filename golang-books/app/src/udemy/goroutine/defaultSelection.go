package goroutine

import (
	"fmt"
	"time"
)

func defaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

OuterLoop2:
	for {
		select {
		case t := <-tick:
			fmt.Println("tick.", t)
		case <-boom:
			fmt.Println("BOOM!")
			break OuterLoop2
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
