package statement

import (
	"fmt"
	"os"
)

// defer は関数の処理が終わった後に実行される

func foo() {
	defer fmt.Println("World Foo!")
	fmt.Println("Hello Foo.")
}
func pracDefer() {

	foo()

	defer fmt.Println("Hello")
	fmt.Println("World")

	readFile()
}

func readFile() {
	file, _ := os.Open("./main.go")
	defer file.Close() // 処理が終わったら閉じる

	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))

}
