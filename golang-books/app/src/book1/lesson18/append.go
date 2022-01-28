package book1Lesson18

import (
	"fmt"
)

func pracAppend() {
	fmt.Println("----------")
	fmt.Println("append.go")
	fmt.Println("----------")

	// 配列にappendはできない
	// dwarfs := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(dwarfs)
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)

	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs)
}
