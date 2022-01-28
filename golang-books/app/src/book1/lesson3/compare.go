package book1Lesson3

import "fmt"

func compare() {
	fmt.Println("----------")
	fmt.Println("compare.go")
	fmt.Println("----------")

	var (
		age   = 28
		minor = age < 18
	)

	fmt.Printf("%v の僕は未成年か？ ---> %v \n", age, minor)

}
