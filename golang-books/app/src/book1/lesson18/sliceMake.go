package book1Lesson18

import (
	"fmt"
)

func sliceMake() {
	fmt.Println("----------")
	fmt.Println("sliceMake.go")
	fmt.Println("----------")

	dwarfs := make([]string, 0, 10)

	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "dwarfs", len(dwarfs), cap(dwarfs), dwarfs)

	dwarfs = append(dwarfs, "aaa", "bbb")
	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "dwarfs", len(dwarfs), cap(dwarfs), dwarfs)

	dwarfs = append(dwarfs, "test1", "test2", "test3", "test1", "test2", "test3", "test1", "test2", "test3", "test1", "test2", "test3")
	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "dwarfs", len(dwarfs), cap(dwarfs), dwarfs)

}
