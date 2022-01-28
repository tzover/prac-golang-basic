package book1Lesson19

import (
	"fmt"
)

func frequency() {
	fmt.Println("----------")
	fmt.Println("frequency.go")
	fmt.Println("----------")

	temperature := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	println("temperature", temperature)

	frequency := make(map[float64]int)
	println("frequency", len(frequency))

	for _, t := range temperature {
		frequency[t]++ // intの初期値が0のため出現回数分インクリメントされる
		println("t", t)
	}
	println("frequency", len(frequency), frequency)

	for t, num := range frequency {
		fmt.Printf("%+.2f の出現は %d 回です。\n", t, num)
	}

}
