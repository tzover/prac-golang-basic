package main

import "fmt"

func process() {

	// score := 56

	fmt.Println("-----process.go-----")

	if score := 90; score > 80 {
		fmt.Println("Great!!!")
	} else if score > 60 {
		fmt.Println("Good!")
	} else {
		fmt.Println("so so...")
	}

	signal := "red"

	switch signal {
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Caution")
	case "blue":
		fmt.Println("Go!")
	default:
		break
	}

	score := 99

	switch {
	case score > 80:
		fmt.Println("Great")
	case score > 60:
		fmt.Println("Good!")
	default:
		fmt.Println("so so...")
	}

	for i := 0; i < 10; i++ {
		// if i == 3 { break } // 抜ける
		if i == 4 {
			continue
		} // Skip
		fmt.Printf("%d 回目のループ処理 \n", i)
	}

	// 上の省略記法
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// 無限ループ
	i = 0
	for {
		fmt.Println(i)
		i++

		if i == 5 {
			break
		}
	}
}
