package book1Lesson18

import (
	"fmt"
)

func sliceAppend() {
	fmt.Println("----------")
	fmt.Println("sliceAppend.go")
	fmt.Println("----------")

	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "dwarfs", len(dwarfs), cap(dwarfs), dwarfs)

	dwarfs2 := append(dwarfs, "test1")
	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "dwarfs2", len(dwarfs2), cap(dwarfs2), dwarfs2)

	dwarfs3 := append(dwarfs2, "test2", "test3", "test4", "test5", "test6")
	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "dwarfs3", len(dwarfs3), cap(dwarfs3), dwarfs2)
}
