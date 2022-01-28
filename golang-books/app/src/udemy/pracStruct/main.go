package pracStruct

import "fmt"

func Main() {
	fmt.Println("********** Packege pracStruct -> main.go **********")

	// v := Vertex{2, 3}
	v := New(5, 6, 7)
	v.scale3D(10)
	fmt.Println(v.area())
	fmt.Println(v.area3D())

	noStruct()
	pracInterface()
	pracInterface2()
	stringPrint()

}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func (v Vertex) area() int {
	return v.X * v.Y
}

// func (v *Vertex) scale(i int) {
// 	v.X = v.X * i
// 	v.Y = v.Y * i
// }

type Vertex struct {
	X, Y int
}
