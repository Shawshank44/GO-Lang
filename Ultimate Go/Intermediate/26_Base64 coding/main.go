package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("He~lo, Base64 encoding")

	// Encode Base64
	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	// Decode from Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(decoded))

	// Url safe encoding :
	UrlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL safe encoded : ", UrlSafeEncoded)

}
