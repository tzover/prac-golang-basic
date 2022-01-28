package goroutine

import (
	"fmt"
	"time"
)

func goroutine4(ch chan string) {
	for {
		ch <- "Packet from 1"
		time.Sleep(1 * time.Second)
	}
}
func goroutine5(ch chan string) {
	for {
		ch <- "Packet from 2"
		time.Sleep(1 * time.Second)
	}
}

func pracChannel3() {
	c1 := make(chan string)
	c2 := make(chan string)

	go goroutine4(c1)
	go goroutine5(c2)

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
