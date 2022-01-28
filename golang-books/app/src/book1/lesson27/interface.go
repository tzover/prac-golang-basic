package book1Lesson27

import (
	"fmt"
)

func pracInterface() {
	fmt.Println("----------")
	fmt.Println("interface.go")
	fmt.Println("----------")

	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	var p *int
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	fmt.Printf("%#v\n", v)

}
