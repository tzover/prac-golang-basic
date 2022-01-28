package book1Lesson19

import (
	"fmt"
	"math"
)

func group() {
	fmt.Println("----------")
	fmt.Println("group.go")
	fmt.Println("----------")

	temperature := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	group := make(map[float64][]float64)

	for _, t := range temperature {
		g := math.Trunc(t/10) * 10
		group[g] = append(group[g], t)
	}

	for g, temp := range group {
		fmt.Printf("%v : %v \n", g, temp)
	}

	// ヒストグラムとかに使えそう

}
