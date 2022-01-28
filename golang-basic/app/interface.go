package main

import "fmt"

type greeter interface {
	greet()
}

func show(t interface{}) {
	// 型アサーション
	_, ok := t.(japanese)
	if ok {
		fmt.Println("I am Japanese")
	} else {
		fmt.Println("I am not Japanese")
	}
	// 型Switch
	switch t.(type) {
		case japanese:
			fmt.Println("I love Japanese")
		case american:
			fmt.Println("I love American")
		default:
			break
	}
}

type japanese struct {}
type american struct {}

func (j japanese) greet() {
	fmt.Println("こんにちは")
}
func (a american) greet() {
	fmt.Println("Hello")
}

func interfaceFunc() {
	greeters := []greeter{japanese{}, american{}}
	fmt.Println("-----interface.go-----")
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}