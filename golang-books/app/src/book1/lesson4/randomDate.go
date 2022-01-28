package book1Lesson4

import (
	"fmt"
	"math/rand"
	"time"
)

func randomDate() {
	fmt.Println("----------")
	fmt.Println("randomDate.go")
	fmt.Println("----------")

	rand.Seed(time.Now().UnixNano())

	var (
		year        int    = 2021
		era         string = "AD"
		month       int    = rand.Intn(12) + 1
		daysInMonth int    = 31
	)

	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1

	fmt.Println(era, year, month, day)

}
