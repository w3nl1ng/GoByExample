package main

import (
	"fmt"
	"regexp"
)

func main() {
	isMatch, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(isMatch)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
}