package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byLength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

func main() {
	s := byLength{"aaa", "ab", "abcd", "e"}
	sort.Sort(s)
	fmt.Println(s)
}