package book1Lesson26

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

func mainMartian() {
	fmt.Println("----------")
	fmt.Println("martian.go")
	fmt.Println("----------")

	shout(martian{})
	shout(&martian{})

}
