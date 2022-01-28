package book1Lesson16

import "fmt"

func terraform() {
	fmt.Println("----------")
	fmt.Println("terraform.go")
	fmt.Println("----------")

	planets := [...]string{
		"水星",
		"金星",
		"地球",
		"火星",
		"木星",
		"土星",
		"天王星",
		"海王星",
	}

	test(planets)
	fmt.Println(planets)

	// dwarfs := [5]string{"ケレス", "冥王星", "ハウメア", "マケマケ", "エリス"}
	// test(dwarfs) // Errorは吐いてくれる -> [8]string でないため

}

// functionは値渡しなので配列(複数の値)は渡せない -> 意味がない
func test(planets [8]string) {
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}
