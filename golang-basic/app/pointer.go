package main

import "fmt"

func pointer() {
	var (
		a       int  = 5
		pointer *int = &a
	)
	fmt.Println("-----pointer.go-----")
	fmt.Println(pointer)
	fmt.Println(*pointer)
}
