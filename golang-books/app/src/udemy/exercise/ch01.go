package exercise

import "fmt"

func ch01() {
	fmt.Println("********** Packege exercise -> exercise.go **********")
	f := 1.11
	fmt.Printf("%-15v 型Check: %T, Value: %v, Value(f): %f, int型に変換: %d \n", "ch01 Q1", f, f, f, int(f))
	fmt.Printf("%-15v %v \n", "ch01 Q2", "[5, 6]")

	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Printf("%-15v %v : %v \n", "ch01 Q2", "Answer", s[2:4])

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}

	fmt.Printf("%-15v %T %v \n", "ch01 Q3", m, m)

}
