package book1Lesson16

import "fmt"

func arrayLoop() {
	fmt.Println("----------")
	fmt.Println("arrayLoop.go")
	fmt.Println("----------")

	dwarfs := [5]string{"ケレス", "冥王星", "ハウメア", "マケマケ", "エリス"}

	fmt.Println(dwarfs)

	for i := 0; i < len(dwarfs); i++ {
		fmt.Println(dwarfs[i])
	}

}
