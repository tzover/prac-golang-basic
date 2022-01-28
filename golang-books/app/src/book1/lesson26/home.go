package book1Lesson26

import "fmt"

func home() {
	fmt.Println("----------")
	fmt.Println("home.go")
	fmt.Println("----------")

	canada := "Canada"

	var home *string
	fmt.Printf("home は %T . \n", home)

	home = &canada

	fmt.Println("&home(自分のアドレス) : ", &home)
	fmt.Println("home(canadaのaddress) : ", home)
	fmt.Println("home(canadaのvalue) : ", *home)

}
