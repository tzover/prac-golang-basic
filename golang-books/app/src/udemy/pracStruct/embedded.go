package pracStruct

type Vertex3D struct {
	Vertex
	Z int
}

func (v Vertex3D) area3D() int {
	return v.X * v.Y * v.Z
}

func (v *Vertex3D) scale3D(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
	v.Z = v.Z * i
}
