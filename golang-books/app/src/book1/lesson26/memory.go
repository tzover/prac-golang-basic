package book1Lesson26

import "fmt"

func memory() {
	fmt.Println("----------")
	fmt.Println("memory.go")
	fmt.Println("----------")

	answer := 40
	fmt.Println("answerのValueは", answer)
	fmt.Println("answerのメモリアドレス（RAMに保存されている）は", &answer)

	fmt.Println("変数addressにanswerのアドレスを代入")
	address := &answer
	fmt.Println("fmt.Println(address) : ", address)
	fmt.Println("fmt.Println(&address) : ", &address)
	fmt.Println("fmt.Println(*address) : ", *address)

}
