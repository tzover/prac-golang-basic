package main

import "fmt"

func variable() {
	/*
		basic
		string		"Hello"
		int				10
		float64		10.0
		bool			false true
		nil

		initial value
		var str string -> ""
		var num string -> 0
		var isBool bool -> false
	*/

	var (
		a int     = 10
		b float64 = 12.3
		c string  = "tojima"
		d bool
	)
	fmt.Println("-----variable.go-----")
	fmt.Printf("a: %d, b: %f, c: %s, d: %t \n", a, b, c, d)
}
