package somePackage

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

// 正規表現
func ioutilPac() {
	fmt.Println("*** ioutilPac() ***")

	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content))

	if err := ioutil.WriteFile("./docs/ioutil_temp.txt", content, 0666); err != nil {
		log.Fatalln(err)
	}

	r := bytes.NewBuffer([]byte("abcde"))
	content2, _ := ioutil.ReadAll(r)
	fmt.Println(string(content2))

}
