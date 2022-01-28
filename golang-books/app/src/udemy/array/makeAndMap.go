package array

import "fmt"

func makeAndMap() {
	fmt.Println("****** makeAndMap.go ******")

	n := make([]int, 3, 5)
	fmt.Printf("len : %d, cap : %d, value : %v \n", len(n), cap(n), n)

	n = append(n, 0, 0)
	fmt.Printf("len : %d, cap : %d, value : %v \n", len(n), cap(n), n)

	n = append(n, 10, 20, 20, 21)
	fmt.Printf("len : %d, cap : %d, value : %v \n", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Printf("len : %d, cap : %d, value : %v \n", len(a), cap(a), a)

	b := make([]int, 0)
	fmt.Printf("len : %d, cap : %d, value : %v \n", len(b), cap(b), b)

	// c := make([]int, 5)
	c := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	// map

	m := map[string]int{"apple": 100, "banana": 300}
	fmt.Println("m : ", m)
	fmt.Println(`m["apple"]`, m["apple"])

	m["banana"] = 800
	fmt.Println("m : ", m)

	// nothing data
	fmt.Println(`m["nothing"]`, m["nothing"])

	// append
	m["abc"] = 900
	fmt.Println("m : ", m)

	v, ok := m["apple"]
	fmt.Println("v, ok : ", v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println("v2, ok2 : ", v2, ok2)

	// capの確認はできないがmakeで作成時に領域を確保するほうが速度が出やすい
	m2 := make(map[string]int, 5000)
	m2["a"], m2["b"] = 123, 345
	fmt.Println("m2 : ", m2)
	fmt.Printf("len(m2) : %d, value(m2) : %v \n", len(m2), m2)

	var s []int

	if s == nil {
		fmt.Println("Nil")
	}

}
