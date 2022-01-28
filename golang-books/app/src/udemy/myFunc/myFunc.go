package myFunc

import "fmt"

func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int) {
	result = price * item
	return result
}

// クロージャー
func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func Main() {
	fmt.Println("********** Packege myFunc -> myFunc.go **********")

	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	f := func(num int) {
		fmt.Println("inner func", num)
	}
	f(100)

	func(num int) {
		fmt.Println("inner func", num)
	}(200)

	// main関数の中で複数のカウンターを使いたい時にこの使い方をすると
	// 使用するカウンター数分、変数を用意しなくて良い

	counter := incrementGenerator()

	fmt.Println("1回目のcounter()", counter())
	fmt.Println("2回目のcounter()", counter())
	fmt.Println("3回目のcounter()", counter())
	fmt.Println("4回目のcounter()", counter())

	// pi
	c1 := circleArea(3.14)
	// radius
	fmt.Println(c1(2))
	fmt.Println(c1(3))

	c2 := circleArea(3)
	fmt.Println(c2(2))
	fmt.Println(c2(3))

	multipleArgs(10, 20, 30)
	multipleArgs(10, 20, 30, 19, 27)
}
