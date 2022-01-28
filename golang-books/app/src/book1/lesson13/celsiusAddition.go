package book1Lesson13

import (
	"fmt"
)

func celsiusAddition() {
	fmt.Println("----------")
	fmt.Println("celsiusAddition.go")
	fmt.Println("----------")

	type celsius float64

	const degrees = 20

	var temperature celsius = degrees

	temperature += 10

	var test float64 = 10

	temperature += celsius(test) // celsius型にキャストしなければいけない 自動補完されるけど

	fmt.Println(temperature)

}
