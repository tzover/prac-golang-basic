package book1Lesson27

import (
	"fmt"
)

type number struct {
	value int
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

func (n number) String() string {
	if !n.valid {
		return "none"
	}
	return fmt.Sprintf("%d", n.value)
}

func valid() {
	fmt.Println("----------")
	fmt.Println("valid.go")
	fmt.Println("----------")

	n := newNumber(27)
	fmt.Println(n)

	e := number{}
	fmt.Println(e)

}
