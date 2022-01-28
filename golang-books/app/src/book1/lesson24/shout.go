package book1Lesson24

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	fmt.Println("----------")
	fmt.Println("shout.go")
	fmt.Println("----------")
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}
