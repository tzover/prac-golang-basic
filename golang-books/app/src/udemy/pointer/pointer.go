package pointer

import "fmt"

func one(x *int) {
	*x = 1
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 10000
}

func Main() {
	fmt.Println("********** Packege pounter -> pounter.go **********")

	var n int = 100

	one(&n)
	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n

	fmt.Println(p)
	fmt.Println(*p)

	// new make
	var p2 *int = new(int)
	fmt.Println(*p2)
	*p2++
	fmt.Println(*p2)

	var p3 *int
	fmt.Println(p3) // nil
	// panic
	// fmt.Println(*p3)
	// *p3++
	// fmt.Println(*p3)

	s := make([]int, 0)
	fmt.Printf("s : %T \n", s)

	m := make(map[string]int)
	fmt.Printf("m : %T \n", m)

	ch := make(chan int)
	fmt.Printf("ch : %T \n", ch)

	fmt.Printf("p2 : %T \n", p2)

	var st = new(struct{})
	fmt.Printf("st : %T \n", st)

	// struct
	v := Vertex{X: 1, Y: 2}
	changeVertex(v)
	fmt.Println("v :", v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Printf("%-10v %T %v \n", "v2 :", v2, v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Printf("%-10v %T %v \n", "v3 :", v3, v3)

	v4 := Vertex{}
	fmt.Printf("%-10v %T %v \n", "v4 :", v4, v4)

	var v5 Vertex
	fmt.Printf("%-10v %T %v \n", "v5 :", v5, v5)

	v6 := new(Vertex)
	fmt.Printf("%-10v %T %v \n", "v6 :", v6, v6)

	v7 := &Vertex{}
	changeVertex2(v7)
	fmt.Printf("%-10v %T %v \n", "v7 :", v7, v7)

}

type Vertex struct {
	// 変数を小文字にすると、privateになり、Mainからアクセスできなくなる
	X, Y int
	S    string
}
