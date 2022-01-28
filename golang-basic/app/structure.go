package main

import "fmt"

type user struct {
	name string
	score int
}

func (u user) show() {
	fmt.Printf("name: %s, score: %d \n", u.name, u.score)
}
func (u *user) hit() { // 参照(pointer)渡し
	u.score++
}

func structure() {
	yuya := new(user)

	fmt.Println("-----structure.go-----")
	fmt.Println(yuya)
	
	yuya.name = "yuya"
	yuya.score = 100
	fmt.Println(yuya)
	
	tojima := user{ name: "tojima", score: 1000}
	fmt.Println(tojima)

	u := user{ name: "tojima", score: 99 }
	u.hit()
	u.show()


}