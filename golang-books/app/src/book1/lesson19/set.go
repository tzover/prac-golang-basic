package book1Lesson19

import (
	"fmt"
	"sort"
)

func set() {
	fmt.Println("----------")
	fmt.Println("set.go")
	fmt.Println("----------")

	temperature := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)

	for _, t := range temperature {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("member!")
	}

	fmt.Println(set)

	unique := make([]float64, 0, len(set))

	for t := range set {
		unique = append(unique, t)
	}

	fmt.Printf("before sort : %10v \n", unique)
	sort.Float64s(unique)
	fmt.Printf("after  sort : %10v \n", unique)

}
