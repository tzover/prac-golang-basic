package book1Lesson26

import (
	"fmt"
	"strings"
)

type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew ", int(*l))
}

func pracInterface() {
	fmt.Println("----------")
	fmt.Println("interface.go")
	fmt.Println("----------")

	pew := laser(2)
	shout(&pew)

}
