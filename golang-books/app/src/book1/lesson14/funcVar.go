package book1Lesson14

import "fmt"

func funcVar() {
	fmt.Println("----------")
	fmt.Println("funcVar.go")
	fmt.Println("----------")

	f := func(message string) {
		fmt.Println(message)
	}

	f("testtesttest")
}
