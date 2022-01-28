package book1Lesson6

import (
	"fmt"
)

func third() {
	fmt.Println("----------")
	fmt.Println("third.go")
	fmt.Println("----------")

	third := 1.0 / 3
	fmt.Println("third(println) : ", third)

	fmt.Printf("third(printf -> default) : %f test\n", third)
	fmt.Printf("third(printf -> .3 ) : %.3f test\n", third)
	fmt.Printf("third(printf -> 40.2 ) : %40.2f test\n", third)
	fmt.Printf("third(printf -> 050.2 ) : %050.2f test\n", third)

}
