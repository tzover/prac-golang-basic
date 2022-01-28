package book1Lesson21

import (
	"encoding/json"
	"fmt"
	"os"
)

func pracJson() {
	fmt.Println("----------")
	fmt.Println("json.go")
	fmt.Println("----------")

	type location struct {
		// lat, long float64 // {}空になる
		Lat, Long float64
	}

	// curiosity := location{lat: -4.1212, long: 136.343}
	curiosity := location{Lat: -4.1212, Long: 136.343}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
