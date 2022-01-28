package book1Lesson28

import (
	"fmt"
	"os"
)

// func funcProverbs(name string) error {
// 	f, err := os.Create(name)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = fmt.Fprintln(f, "Errors are values")
// 	if err != nil {
// 		f.Close()
// 		return err
// 	}

// 	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully")
// 	f.Close()
// 	return err
// }

func proverbs() {
	fmt.Println("----------")
	fmt.Println("proverbs.go")
	fmt.Println("----------")

	err := funcProverbs("./src/book1/lesson28/proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
