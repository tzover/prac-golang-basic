package book1Lesson29

import (
	"errors"
	"fmt"
	"os"
)

const (
	rows, columns = 9, 9
	empty         = 0
)

type Cell struct {
	digit int8
	fixed bool
}

type Grid [rows][columns]Cell

var (
	ErrBounds     = errors.New("out of bounds")
	ErrDigit      = errors.New("invalid digit")
	ErrInRow      = errors.New("digit already present in this row")
	ErrInColumn   = errors.New("digit already present in this column")
	ErrInRegion   = errors.New("digit already present in this region")
	ErrFixedDigit = errors.New("initial digits cannot be overwritten")
)

func NewSudoku(digits [rows][columns]int8) *Grid {
	var grid Grid

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			d := digits[r][c]

			if d != empty {
				grid[r][c].digit = d
				grid[r][c].fixed = true
			}
		}
	}
	return &grid
}

func (g *Grid) Set(row, column int, digit int8) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case !validDigit(digit):
		return ErrDigit
	case g.isFixed(row, column):
		return ErrFixedDigit
	case g.inRow(row, digit):
		return ErrInRow
	case g.inColumn(column, digit):
		return ErrInColumn
	case g.inRegion(row, column, digit):
		return ErrInRegion
	}

	g[row][column].digit = digit
	return nil
}

func (g *Grid) Clear(row, column int) error {
	switch {
	case !inBounds(row, column):
		return ErrBounds
	case g.isFixed(row, column):
		return ErrFixedDigit
	}

	g[row][column].digit = empty
	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows || column < 0 || column >= columns {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < columns; c++ {
		if g[row][c].digit == digit {
			return false
		}
	}
	return true
}

func (g *Grid) inColumn(column int, digit int8) bool {
	for r := 0; r < columns; r++ {
		if g[r][column].digit == digit {
			return false
		}
	}
	return true
}

func (g *Grid) inRegion(row, column int, digit int8) bool {
	startRow, startColumn := row/3*3, column/3*3

	for r := startRow; r < startRow+3; r++ {
		for c := startColumn; c < startColumn+3; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}

func (g *Grid) isFixed(row, column int) bool {
	return g[row][column].fixed
}

func Main() {
	fmt.Println("******* Package book1Lesson29 *******")

	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.Set(0, 0, 30)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range s {
		fmt.Println(row)
	}

}
