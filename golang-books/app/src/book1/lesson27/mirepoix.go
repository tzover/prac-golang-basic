package book1Lesson27

import (
	"fmt"
)

func funcMirepoix(ingredients []string) []string {
	return append(ingredients, "onion")
}

func mirepoix() {
	fmt.Println("----------")
	fmt.Println("mirepoix.go")
	fmt.Println("----------")

	soup := funcMirepoix(nil)
	fmt.Println(soup)

}
