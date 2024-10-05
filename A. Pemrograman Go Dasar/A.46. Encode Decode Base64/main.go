package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "john wick"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded:", decodedString)

	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	encodedString = string(encoded)
	fmt.Println(encodedString)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	decodedString = string(decoded)
	fmt.Println(decodedString)

	data = "https://kalipare.com/"

	encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	decodedString = string(decodedByte)
	fmt.Println(decodedString)
}
