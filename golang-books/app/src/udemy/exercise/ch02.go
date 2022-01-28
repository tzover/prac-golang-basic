package exercise

import "fmt"

func searchMinNum(num []int) (minNum int) {

	minNum = num[0]

	for _, val := range num {
		if minNum > val {
			minNum = val
		}
	}
	return
}

func calcSum(num map[string]int) (sumNum int) {
	for _, val := range num {
		sumNum += val
	}
	return
}

func ch02() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	minNum := searchMinNum(l)
	fmt.Printf("%-15v l のスライスの中で一番小さい数は %d \n", "ch02 Q1", minNum)

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grape":  150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	sumNum := calcSum(m)
	fmt.Printf("%-15v m のmapの合計値は %d \n", "ch02 Q2", sumNum)
}
