package array

import "fmt"

func pracByte() {
	fmt.Println("****** byte.go ******")

	b := []byte{72, 63}
	fmt.Println(b)
	fmt.Printf("b : type = %T \n", b)
	fmt.Println(string(b))
	fmt.Printf("string(b) : type = %T \n", string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))

}
