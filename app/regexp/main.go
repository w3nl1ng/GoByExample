package main

import (
	"fmt"
	"regexp"
	"bytes"
)

func main() {
	isMatch, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(isMatch)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))

	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch", -1))

	fmt.Println(r.FindAllString("peach punch panch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("r:", r)

	r.ReplaceAllString("a peach", "<fruit>")

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}