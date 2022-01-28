package main

// goroutine : 並行処理
// channel :

import (
	"fmt"
	"time"
)

func task1(result chan string) {
	time.Sleep(time.Second * 5)
	result<- "task1 result"
	fmt.Println("task1 finished!")
	
}
func task2() {
	fmt.Println("task2 finished!")
}

func routine() {
	fmt.Println("-----routine.go-----")

	result := make(chan string)

	// 並行処理
	go task1(result)
	go task2()

	fmt.Println(<-result)

	time.Sleep(time.Second * 10)
}