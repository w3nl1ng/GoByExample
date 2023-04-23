package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"

	dEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(dEnc)

	dDec, err := base64.StdEncoding.DecodeString(dEnc)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dDec))

	urlEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(urlEnc)

	urlDec, _ := base64.URLEncoding.DecodeString(urlEnc)
	fmt.Println(string(urlDec))
}