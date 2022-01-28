package typecast

import (
	"fmt"
	"strconv"
)

func Main() {
	var x int = 1
	xx := float64(x)

	fmt.Printf("%T , %v , %f \n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)

	fmt.Printf("%T , %v , %d \n", yy, yy, yy)

	var s string = "15"
	// z := int(s) // Error

	// i, _ := strconv.Atoi(s) // Errorハンドリングを無視
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Printf("%T , %v , %d \n", i, i, i)

	h := "Hello"
	fmt.Println("h[0] : ", string(h[0]))

}
