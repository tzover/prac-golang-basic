package book1Lesson18

import (
	"fmt"
)

func Main() {
	fmt.Println("******* Package book1Lesson18 *******")
	pracAppend()

	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	sliceDump("dwarfs", dwarfs)
	sliceDump("dwards[1:2]", dwarfs[1:2])
	dwarfs = append(dwarfs, "test")
	sliceDump("dwarfs", dwarfs)

	sliceAppend()
	threeIndexSlicing()
	sliceMake()

	twoWorlds := variadic("New", "test1", "test2", "asdsa", "dcc")
	fmt.Println(twoWorlds)

	newDwarfs := variadic("New", dwarfs...)
	fmt.Println(newDwarfs)
}
