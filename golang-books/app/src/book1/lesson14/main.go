package book1Lesson14

import (
	"fmt"
)

type kelvin float64
type sensorFunc func() kelvin

func Main() {
	fmt.Println("******* Package book1Lesson14 *******")

	sensor := fakeSensor()
	fmt.Println(sensor)

	sensor = realSensor()
	fmt.Println(sensor)

	measureTemperature(3, fakeSensor)

	f()

	funcVar()

	anonymous()

	sensorFunc := calibrate(realSensor, 5)
	fmt.Println(sensorFunc())

}
