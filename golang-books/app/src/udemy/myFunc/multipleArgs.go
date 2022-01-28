package myFunc

import "fmt"

func multipleArgs(params ...int) {
	fmt.Println(len(params), params)

	for _, param := range params {
		fmt.Println(param)
	}
}
