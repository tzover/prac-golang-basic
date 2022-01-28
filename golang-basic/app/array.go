package main

import "fmt"

func array() {
	var (
		a [5]int // a[0] - a[4]
	)

	b := []int{1, 3, 4, 7, 8}

	fmt.Println("-----array.go-----")
	fmt.Println(a)

	a[2] = 55
	fmt.Println(a)
	fmt.Println("a[2] : ", a[2])
	fmt.Println("a length : ", len(a))

	// range
	for id, numB := range b {
		fmt.Println("arrayB")
		fmt.Printf("[%d]", id)
		fmt.Println(numB)
	}
	// for _, numB := range b {
	// 	fmt.Println("arrayB")
	// 	fmt.Println(numB)
	// }

	fmt.Println("slice!!!!!")
	c := []int{1, 3, 4, 123, 312, 4, 534, 341}

	// slice
	cSlice := c[3:6] // 抜き出し
	fmt.Println("cSlice", cSlice)
	cSlice = append(cSlice, 2, 3, 67) // 追加

	// copy
	t := make([]int, len(c))
	fmt.Println("t", t) // [0,0,0,0,0,0,0,0]
	n := copy(t, cSlice)

	cSlice[2] = 999 // 全ての配列の値が変わってしまうことに注意

	fmt.Println("c : ", c)
	fmt.Println("c slice : ", cSlice)
	fmt.Println("c slice len : ", len(cSlice))
	fmt.Println("c slice cap : ", cap(cSlice))
	fmt.Println("n copy : ", n)

	// map
	mapEx1 := map[string]int{"tojima": 1, "yuya": 2}
	mapEx2 := make(map[string]int)
	mapEx2["tojima"] = 100
	mapEx2["yuya"] = 300

	fmt.Println("mapEx1", mapEx1)
	fmt.Println("mapEx2", mapEx2)
	fmt.Println(len(mapEx1))

	delete(mapEx1, "tojima")
	fmt.Println("mapEx1", mapEx1)

	v, ok := mapEx2["tojima"]
	fmt.Println("v : ", v)
	fmt.Println("ok : ", ok)

	mapEx3 := map[string]int{"tojima": 123, "yuya": 456}
	for k, v := range mapEx3 {
		fmt.Println("mapEx3", k, v)
	}
}
