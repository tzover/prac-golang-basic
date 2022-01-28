package book1Lesson28

import (
	"fmt"
	"os"
)

const (
	rows    = 9
	columns = 9
)

type Grid [rows][columns]int8
type error interface {
	Error() string
}

func validDigit(digit int8) bool {
	return digit > 1 && digit <= 9
}

func (g *Grid) Set(row, column int, digit int8) error {

	var errs SudokuError

	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}

	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}

	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func sudoku() {
	fmt.Println("----------")
	fmt.Println("sudoku.go")
	fmt.Println("----------")

	var g Grid
	err := g.Set(10, 0, 5)

	if err != nil {
		fmt.Printf("An error occurred : %v . \n", err)
		os.Exit(1)
	}

}
