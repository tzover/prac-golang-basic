package array

import "fmt"

func Main() {
	fmt.Println("********** Packege array -> array.go **********")

	var a [2]int
	a[0] = 100
	a[1] = 200

	fmt.Println("a: ", a)

	// var b [2]int = [2]int{100, 200}
	b := [2]int{100, 200}
	// b = append(b, 300) // Error
	fmt.Println("b: ", b)

	// slice
	c := []int{100, 200, 300}
	c = append(c, 400)
	fmt.Println("c: ", c)

	slice()
	makeAndMap()
	pracByte()

}
