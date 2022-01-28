package book1Lesson26

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func pracStruct() {
	fmt.Println("----------")
	fmt.Println("struct.go")
	fmt.Println("----------")

	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	timmy.superpower = "flying"

	fmt.Printf("%+v \n", timmy)

}
