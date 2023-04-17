package main

import "fmt"

func zeroVal(n int) {
	n = 0
}

func zeroPtr(p *int) {
	*p = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)

	fmt.Println("pointer:", &i)
}