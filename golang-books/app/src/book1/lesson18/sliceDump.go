package book1Lesson18

import (
	"fmt"
)

func sliceDump(label string, slice []string) {
	fmt.Println("----------")
	fmt.Println("sliceDump.go")
	fmt.Println("----------")

	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", label, len(slice), cap(slice), slice)
}
