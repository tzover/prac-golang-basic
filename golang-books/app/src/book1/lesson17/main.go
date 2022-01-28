package book1Lesson17

import (
	"fmt"
	"strings"
)

func Main() {
	fmt.Println("******* Package book1Lesson17 *******")

	slicing()
	slicingDefault()
	dwarfSlice()

	// 配列 Error -> hyperSpace(planets)
	// planets := [...]string{" Venus ", " Earth ", " Mars "}
	// Slice OK -> hyperSpace(planets)
	planets := []string{" Venus ", " Earth ", " Mars "}

	hyperSpace(planets)
	fmt.Println(strings.Join(planets, ""))

	pracSort()

}
