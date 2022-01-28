package pracStruct

import "fmt"

// all acept
func do(i interface{}) {

	// タイプアサーションをしないと演算できない
	// ii := i * 2 // Error
	// ii := i.(int)
	// ii *= 2
	// fmt.Println(ii)

	// case of string
	// ss := i.(string)
	// fmt.Println(ss + "!")

	// all acept
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know %T \n", v)
	}
}

func pracInterface2() {
	do(10)
	do("Mike")
	do(true)
}
