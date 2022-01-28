package somePackage

import (
	"fmt"
	"regexp"
)

// 正規表現
func regexPac() {
	fmt.Println("*** regexPac() ***")

	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match) // true

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	// s := "/view/test"
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)

	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

}
