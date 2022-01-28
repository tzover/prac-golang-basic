package hello

import (
	"fmt"
	"os/user"
	"strings"
	"time"
)

// 必ず最初に呼ばれる
func init() {
	fmt.Println("********** Packege hello -> hello.go **********")
	fmt.Println("init() は最初に呼ばれる")
	fmt.Println(user.Current())
}

func bazz() {
	fmt.Println("Bazz")
}

func hello() {
	fmt.Println("Hello World")
}

func pracStirngs() {
	var msg string = "Hello World"
	msg = strings.Replace(msg, "H", "X", 1)
	fmt.Println(msg)
	fmt.Println(strings.Contains(msg, "World")) // true

	fmt.Println("test" + "test")
	fmt.Println(`Test        Test
          Test
	Test`)

	fmt.Println(`"/"`)
}

func Main() {
	fmt.Println("Hello world!", "TEST", time.Now())
	bazz()
	hello()
	pracStirngs()
}
