package exercise

import "fmt"

func ch03() {
	var i int = 10
	var p *int = &i
	var j int = *p
	fmt.Printf("%-15v %d \n", "ch03 Q1", j)

	var i2 int = 100
	var j2 int = 200
	var p1 *int
	var p2 *int
	p1 = &i2
	p2 = &j2
	fmt.Println(*p1, *p2) // 100, 200
	i2 = *p1 + *p2
	fmt.Println(i2)       // 300
	fmt.Println(*p1, *p2) // 300, 200
	p2 = p1
	fmt.Println(*p1, *p2) // 300, 300
	j2 = *p2 + i2
	fmt.Printf("%-15v %d \n", "ch03 Q2", j2)
}
