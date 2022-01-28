package book1Lesson21

import (
	"encoding/json"
	"fmt"
)

func jsonTags() {
	fmt.Println("----------")
	fmt.Println("jsonTags.go")
	fmt.Println("----------")

	type location struct {
		// lat, long float64 // {}空になる
		// Lat, Long float64
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	// curiosity := location{lat: -4.1212, long: 136.343}
	curiosity := location{Lat: -4.1212, Long: 136.343}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))
}
