package errorHandle

import (
	"fmt"
	"log"
	"os"
)

func Main() {
	fmt.Println("********** Packege errorHandle -> errorHandle.go **********")

	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatalln("Error", err)
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error", err)
	}

	fmt.Println(count, string(data))

	if err = os.Chdir("src"); err != nil {
		log.Fatalln("Error", err)
	}

	// custum Error
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

	// * & でfunctionを構成する理由　エラー内容の比較が難しい
	e1 := UserNotFound{"Mike"}
	e2 := UserNotFound{"Mike"}
	fmt.Println(e1 == e2) // true

	e3 := &UserNotFound{"Mike"}
	e4 := &UserNotFound{"Mike"}
	fmt.Println(e3 == e4) // false
}
