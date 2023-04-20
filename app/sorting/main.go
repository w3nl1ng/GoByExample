package main

import (
	"fmt"
	"sort"
)

func main() {
	str := []string{"c", "a", "b"}
	sort.Strings(str)
	fmt.Println(str)

	num := []int{2, 5, 3, 6}
	sort.Ints(num)
	fmt.Println(num)

	b := sort.StringsAreSorted(str)
	fmt.Println(b)
}