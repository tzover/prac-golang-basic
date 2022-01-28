package book1Lesson22

import "math"

type world struct {
	radius float64
}

var (
	mars = world{radius: 3389.5}
)

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)

}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}
