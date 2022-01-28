package statement

import "fmt"

func forLoop() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i == 4 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)

	}
	// 無限ループ
	forNum := 0
	for {
		fmt.Println("forNum")
		forNum++

		if forNum == 5 {
			fmt.Println("break")
			break
		}
	}
}

func forRange() {
	l := []string{"Java", "Golang", "python"}

	for i, v := range l {
		fmt.Println(i, v)
	}
	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
	// key だけ取り出し
	for k := range m {
		fmt.Println(k)
	}
	// value だけ取り出し
	for _, v := range m {
		fmt.Println(v)
	}
}
