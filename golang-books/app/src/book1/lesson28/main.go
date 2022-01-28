package book1Lesson28

import "fmt"

func Main() {
	fmt.Println("******* Package book1Lesson28 *******")

	files()
	proverbs()

	sudoku()
	sudoku2()
	sudoku3()

	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}

	}()

	panic("I forgot my towl")
}
