package book1Lesson30

import (
	"fmt"
)

// func sourceGopher(downstream chan string) {
// 	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
// 		downstream <- v
// 	}
// 	downstream <- ""
// }

// func filterGopher(upstream, downstream chan string) {
// 	for {
// 		item := <-upstream

// 		if item == "" {
// 			downstream <- ""
// 			return
// 		}
// 		if !strings.Contains(item, "bad") {
// 			downstream <- item
// 		}
// 	}
// }

// func printGopher(upstream chan string) {
// 	for {
// 		v := <-upstream
// 		if v == "" {
// 			return
// 		}
// 		fmt.Println(v)
// 	}
// }

func pipeline1() {
	fmt.Println("----------")
	fmt.Println("pipeline1.go")
	fmt.Println("----------")

	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
