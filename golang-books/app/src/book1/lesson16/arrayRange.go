package book1Lesson16

import "fmt"

func arrayRange() {
	fmt.Println("----------")
	fmt.Println("arrayRange.go")
	fmt.Println("----------")

	dwarfs := [5]string{"ケレス", "冥王星", "ハウメア", "マケマケ", "エリス"}

	fmt.Println(dwarfs)

	for i, dwarf := range dwarfs {
		fmt.Println(i, dwarf)
	}

}
