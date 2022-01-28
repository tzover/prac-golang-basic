package somePackage

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
}

// Hash
func hmacPac() {
	fmt.Println("*** hmacPac() ***")

	const (
		apiKey    = "User1Key"
		apiSecret = "User1Secret"
	)

	data := []byte("data")

	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)

	sign := hex.EncodeToString(h.Sum(nil))

	fmt.Println("sign", sign)

	Server(apiKey, sign, data)
}
