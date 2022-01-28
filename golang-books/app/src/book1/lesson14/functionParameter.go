package book1Lesson14

import (
	"fmt"
	"time"
)

func measureTemperature(samples int, sensor func() kelvin) {
	fmt.Println("----------")
	fmt.Println("functionParameter.go -> measureTemperature")
	fmt.Println("----------")

	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vo K\n", k)
		time.Sleep(time.Second)
	}
}
