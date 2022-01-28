package book1Lesson27

import (
	"fmt"
)

type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}

func guard() {
	fmt.Println("----------")
	fmt.Println("guard.go")
	fmt.Println("----------")

	var nobody *person
	fmt.Println(nobody)
	nobody.birthday()

}
