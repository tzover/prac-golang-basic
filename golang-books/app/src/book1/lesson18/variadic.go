package book1Lesson18

import (
	"fmt"
)

func variadic(prefix string, worlds ...string) []string {
	fmt.Println("----------")
	fmt.Println("variadic.go")
	fmt.Println("----------")

	newWorlds := make([]string, len(worlds))

	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "worlds", len(worlds), cap(worlds), worlds)
	fmt.Printf("%v : 長さ %v , 容量 %v %v\n", "newWorlds", len(newWorlds), cap(newWorlds), newWorlds)

	for i := range worlds {
		newWorlds[i] = prefix + "_" + worlds[i]
	}
	return newWorlds
}
