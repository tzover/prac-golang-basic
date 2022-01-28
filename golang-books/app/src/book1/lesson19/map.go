package book1Lesson19

import (
	"fmt"
)

func pracMap() {
	fmt.Println("----------")
	fmt.Println("map.go")
	fmt.Println("----------")

	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	fmt.Println(temperature)

	temp := temperature["Earth"]

	fmt.Printf("Average : %v  in the Earth\n ", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("Average : %v  in the Moon\n ", moon)
	} else {
		fmt.Println("Where is the Moon?")
	}

}
