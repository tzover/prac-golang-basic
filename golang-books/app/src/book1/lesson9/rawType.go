package book1Lesson9

import (
	"fmt"
)

func rawType() {
	fmt.Println("----------")
	fmt.Println("rawType.go")
	fmt.Println("----------")

	fmt.Printf("%v は %[1]T 型です。\n", "文字列リテラル")
	fmt.Printf("%v は %[1]T 型です。\n", `生の文字列リテラル`)

}
