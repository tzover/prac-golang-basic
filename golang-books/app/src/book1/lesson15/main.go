package book1Lesson15

import (
	"fmt"
)

type (
	celsius    float64                        // °C
	fahrenheit float64                        // °F
	getRowFunc func(row int) (string, string) // 2つ文字列を返す
)

const (
	celsiusText    string = "°C"
	fahrenheitText string = "°F"
	line           string = "================================="
	rowFormat      string = "|     %-10s|     %-10s| \n"
	numFormat      string = "%.1f"
)

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

func drowTable(hdr1, hdr2 string, rows int, getRow getRowFunc) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2)
	}
	fmt.Println(line)
}

func ctof(row int) (string, string) {
	c := celsius(row*5 - 40)
	fmt.Println("c:", c)
	f := c.fahrenheit() //
	// fmt.Println("f:", f)
	cell1 := fmt.Sprintf(numFormat, c)
	cell2 := fmt.Sprintf(numFormat, f)

	return cell1, cell2
}

func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()

	cell1 := fmt.Sprintf(numFormat, f)
	cell2 := fmt.Sprintf(numFormat, c)

	return cell1, cell2
}

func Main() {
	fmt.Println("******* Package book1Lesson15 *******")

	drowTable(celsiusText, fahrenheitText, 29, ctof)
	fmt.Println()
	drowTable(fahrenheitText, celsiusText, 29, ftoc)

}
