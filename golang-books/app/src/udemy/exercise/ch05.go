package exercise

import (
	"fmt"
)

func goroutine(s []string, ch chan string) {
	sum := ""
	for _, v := range s {
		sum += v
		ch <- sum
	}
	close(ch)
}

func ch05() {

	words := []string{"test1!", "test2!", "test3!", "test4!"}
	ch := make(chan string)
	go goroutine(words, ch)
	for w := range ch {
		fmt.Println(w)
	}

}
