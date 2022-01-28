package book1Lesson21

import (
	"fmt"
)

func sliceStruct() {
	fmt.Println("----------")
	fmt.Println("sliceStruct.go")
	fmt.Println("----------")

	// この使い方はやめたほうがいい
	// lats := []float64{-4.5895, -14.5684, -1.9462}
	// longs := []float64{137.4417, 175.472636, 354.4734}

	type location struct {
		name      string
		lat, long float64
	}

	locations := []location{
		{name: "testA", lat: -4.5895, long: 137.4417},
		{name: "testB", lat: -14.5682, long: 175.3183},
		{name: "testC", lat: -1.9662, long: 354.3113},
	}

	fmt.Printf("%+v \n", locations)

}
