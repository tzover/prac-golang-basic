package book1Lesson13

import (
	"fmt"
)

func celsius() {
	fmt.Println("----------")
	fmt.Println("celsius.go")
	fmt.Println("----------")

	type celsius float64

	var temperature celsius = 20
	fmt.Println(temperature)

}
