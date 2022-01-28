package book1Lesson3

import "fmt"

func leap() {
	fmt.Println("----------")
	fmt.Println("leap.go")
	fmt.Println("----------")

	fmt.Println("2100年はうるう年ですか？")

	var (
		targetYear = 2100
		leap       = targetYear%400 == 0 || (targetYear%4 == 0 && targetYear%100 != 0)
	)

	if leap {
		fmt.Printf("Yes. %v is a leap year \n", targetYear)
	} else {
		fmt.Println("No. A common year")
	}

}
