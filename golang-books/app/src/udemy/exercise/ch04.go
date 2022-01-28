package exercise

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) Plus() int {
	return v.X + v.Y
}

func (v Vertex) String() string {
	return fmt.Sprintf("X is %d! Y is %d! \n", v.X, v.X)
}

func ch04() {

	v := Vertex{X: 3, Y: 4}
	fmt.Printf("%-15v %d \n", "ch04 Q1", v.Plus())
	fmt.Printf("%-15v %d \n", "ch04 Q2", v)
	fmt.Println(v)

}
