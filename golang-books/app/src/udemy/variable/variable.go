package variable

import (
	"fmt"
)

// global variable
const Pi = 3.14

const (
	Username = "test_user"
	Password = "test_pass"
)

func foo() {
	xi := 1
	xi = 3
	xf64 := 3.2
	xs := "test"
	xt, xf := true, false

	fmt.Println(xi, xf64, xs, xt, xf)
	// 型チェック
	fmt.Printf("%T\n", xi)
	fmt.Printf("%T\n", xf64)
}

func variable() {
	fmt.Println("-----------")

	fmt.Println("1 + 1", 1+1)
	fmt.Println("10 - 1", 10-1)
	fmt.Println("10 / 2", 10/2)
	fmt.Println("10 / 3", 10/3)
	fmt.Println("10.0 / 3", 10.0/3)
	fmt.Println("10 / 3.0", 10/3.0)
	fmt.Println("10 % 2", 10%2)
	fmt.Println("10 % 3", 10%3)

	fmt.Println("-----------")
}

func Main() {
	fmt.Println("********** Packege variable -> variable.go **********")

	fmt.Println(Username, Password, Pi)

	var (
		i    int     = 1
		f64  float64 = 3.2
		s    string  = "test"
		t, f bool    = true, false
	)

	fmt.Println(i, f64, s, t, f)

	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)

	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type = %T value = %v \n", u8, u8)

	foo()
	variable()

}
