package book1Lesson16

import "fmt"

func chess() {
	fmt.Println("----------")
	fmt.Println("chess.go")
	fmt.Println("----------")

	var board [8][8]string

	board[0][0] = "r"
	board[0][7] = "r"

	for column := range board[1] {
		board[1][column] = "p"
	}
	fmt.Println(board)

}
