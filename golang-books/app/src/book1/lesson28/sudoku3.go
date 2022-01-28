package book1Lesson28

import (
	"fmt"
	"os"
	"strings"
)

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func sudoku3() {
	fmt.Println("----------")
	fmt.Println("sudoku3.go")
	fmt.Println("----------")

	var g Grid
	err := g.Set(10, 0, 5)

	if err != nil {

		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d errors(s) occurred : \n", len(errs))

			for _, e := range errs {
				fmt.Printf("- %v \n", e)
			}
		}

		os.Exit(1)
	}

}
