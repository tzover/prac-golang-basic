package book1Lesson10

import (
	"fmt"
	"strconv"
)

func pracItoa() {
	fmt.Println("----------")
	fmt.Println("itoa.go")
	fmt.Println("----------")

	countdown := 10

	str := "発射まで" + strconv.Itoa(countdown) + " sec "

	fmt.Println(str)

	countdown = 9

	str = fmt.Sprintf("発射まで %v sec", countdown)

	fmt.Println(str)

	countdown, err := strconv.Atoi("10a")

	if err != nil {
		fmt.Println("No Good")
	}

	fmt.Println(countdown)

}
