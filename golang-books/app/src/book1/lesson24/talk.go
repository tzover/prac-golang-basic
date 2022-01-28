package book1Lesson24

import (
	"fmt"
	"strings"
)

var t interface {
	talk() string
}

type martian struct{}

type laser int

func (m martian) talk() string {
	fmt.Println("----------")
	fmt.Println("talk.go")
	fmt.Println("----------")
	return "nack nack"
}

func (l laser) talk() string {
	fmt.Println("----------")
	fmt.Println("talk.go")
	fmt.Println("----------")
	return strings.Repeat("pew ", int(l))
}
