package main

import "fmt"

func calc() {
	var (
		x       int
		s       string
		isBoolA bool
		isBoolB bool
	)
	// number
	x = 10 % 3
	x += 3
	x++

	// string
	s = "Hello" + "World"

	// bool
	isBoolA = true
	// isBoolB = false

	// output
	fmt.Println("-----calc.go-----")
	fmt.Printf("int : %d \n", x)
	fmt.Printf("string : %s \n", s)
	fmt.Println("isBoolA && isBoolB :", isBoolA && isBoolB)
	fmt.Println("isBoolA || isBoolB :", isBoolA || isBoolB)
	fmt.Println("!isBoolA :", !isBoolA)
}
