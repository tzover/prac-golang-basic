package book1Lesson17

import "fmt"

func slicingDefault() {
	fmt.Println("----------")
	fmt.Println("slicingDefault.go")
	fmt.Println("----------")

	planets := [...]string{
		"水星",  // 0
		"金星",  // 1
		"地球",  // 2
		"火星",  // 3
		"木星",  // 4
		"土星",  // 5
		"天王星", // 6
		"海王星", // 7
	}

	// 先頭と最後はデフォルトで良しなにやってくれる
	terrestrial := planets[:4] // 0 1 2 3
	gasGiants := planets[:6]   // 4 5
	iceGiants := planets[6:]   // 6 7
	allPlanets := planets[:]

	fmt.Println(terrestrial, gasGiants, iceGiants, allPlanets)
	fmt.Println(planets)
	fmt.Println(gasGiants[0])

	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]

	fmt.Println(giants, gas, ice)

	iceGiantsMark2 := iceGiants

	iceGiants[1] = "test"

	fmt.Println("planets : ", planets)

	fmt.Println(iceGiants, iceGiantsMark2, ice)
}
