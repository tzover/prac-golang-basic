package book1Lesson26

import "fmt"

func method() {
	fmt.Println("----------")
	fmt.Println("method.go")
	fmt.Println("----------")

	terry := &person{
		name: "Terry",
		age:  15,
	}

	terry.birthday()

	fmt.Printf("%+v \n", terry)

	nathan := person{
		name: "Nathan",
		age:  17,
	}

	nathan.birthday()

	fmt.Printf("%+v \n", nathan)
}

func (p *person) birthday() {
	p.age++
}
