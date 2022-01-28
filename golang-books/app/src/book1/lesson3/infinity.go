package book1Lesson3

import (
	"fmt"
	"math/rand"
)

func infinity() {
	fmt.Println("----------")
	fmt.Println("infinity.go")
	fmt.Println("----------")

	var (
		count int = 0
	)

	for { // 無限ループ
		fmt.Println(count)
		count++

		if count >= 360 {
			count = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}
