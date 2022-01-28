package book1Lesson23

import (
	"fmt"
)

func Main() {
	fmt.Println("******* Package book1Lesson23 *******")
	compose()

	t := temperature{high: -1.0, low: -78.0}
	fmt.Printf("平均温度は %v °C\n", t.average())

	report := report{sol: 15, temperature: t}
	fmt.Printf("平均温度は %v °C\n", report.temperature.average())

	embed()

	report2 := report

	fmt.Println(report2.sol.days(1446))
	fmt.Println(report2.days(1446))

	d := report.days(1446)
	fmt.Println(d)

}
