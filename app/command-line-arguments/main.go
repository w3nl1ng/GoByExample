package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	fmt.Println(argsWithProg)
	for _, arg := range argsWithoutProg {
		fmt.Println(arg)
	}
}
