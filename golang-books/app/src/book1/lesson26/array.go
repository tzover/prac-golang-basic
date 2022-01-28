package book1Lesson26

import "fmt"

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
	// ...
}

func pracArray() {
	fmt.Println("----------")
	fmt.Println("array.go")
	fmt.Println("----------")

	var board [8][8]rune
	reset(&board)
	fmt.Printf("%c \n", board[0][0])

}
