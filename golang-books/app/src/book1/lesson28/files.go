package book1Lesson28

import (
	"fmt"
	"io/ioutil"
	"os"
)

func files() {
	fmt.Println("----------")
	fmt.Println("files.go")
	fmt.Println("----------")

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
