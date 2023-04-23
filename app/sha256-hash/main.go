package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := sha256.New()

	data := "this is sha256"
	s.Write([]byte(data))

	bs := s.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}