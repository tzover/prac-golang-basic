package book1Lesson14

import (
	"fmt"
	"math/rand"
)

func fakeSensor() kelvin {
	fmt.Println("----------")
	fmt.Println("sensor.go -> fakeSensor")
	fmt.Println("----------")

	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	fmt.Println("----------")
	fmt.Println("sensor.go -> realSensor")
	fmt.Println("----------")
	return 0
}
