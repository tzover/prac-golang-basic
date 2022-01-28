package array

import (
	"fmt"
)

func slice() {
	fmt.Println("****** slice.go ******")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("slice :", slice)
	fmt.Println("slice[2] :", slice[2])
	fmt.Println("slice[2:4] :", slice[2:4])
	fmt.Println("slice[4:] :", slice[4:])
	fmt.Println("slice[:2] :", slice[:2])
	fmt.Println("slice[:] :", slice[:])

	slice[2] = 55555
	fmt.Println("slice :", slice)

	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	fmt.Println("board :", board)

	board = append(board, []int{9, 10, 2})

	fmt.Println("board :", board)

}
