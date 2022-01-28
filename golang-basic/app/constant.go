package main

import "fmt"

func constant() {
	const (
		sun = iota // 0
		mon
		tue
	)

	fmt.Println("-----constant.go-----")
	fmt.Println(sun, mon, tue)
}
