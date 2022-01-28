package book1Lesson28

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invaild digit")
)

func sudoku2() {
	fmt.Println("----------")
	fmt.Println("sudoku2.go")
	fmt.Println("----------")

	var g Grid
	err := g.Set(10, 0, 5)

	if err != nil {

		switch err {
		case ErrBounds, ErrDigit:
			fmt.Println("Les erreurs de parametres hors limites.")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}

}
