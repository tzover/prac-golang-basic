package book1Lesson11

import (
	"fmt"
	"strings"
)

func Main() {
	fmt.Println("******* Package book1Lesson11 *******")
	decipher()
	cipher()
}

func decipher() {
	cipherText := "NSWLB"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0

	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		c = (c-k+26)%26 + 'A'

		message += string(c)

		keyIndex++

		keyIndex %= len(keyword)
	}
	fmt.Println(message)
}

func cipher() {
	message := "Hello"
	keyword := "golang"
	keyIndex := 0
	cipherText := ""

	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

	for i := 0; i < len(message); i++ {
		c := message[i]

		if c >= 'A' && c <= 'Z' {
			c -= 'A'
			k := keyword[keyIndex] - 'A'

			c = (c+k)%26 + 'A'

			keyIndex++
			keyIndex %= len(keyword)
		}
		cipherText += string(c)
	}
	fmt.Println(cipherText)
}
