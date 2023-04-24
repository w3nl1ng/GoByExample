package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numPtr := flag.Int("num", 42, "a number")
	forkPtr := flag.Bool("fork", false, "whether to fork")

	var stringVar string
	flag.StringVar(&stringVar, "fruit", "", "a fruit")

	flag.Parse()

	fmt.Println(*wordPtr)
	fmt.Println(*numPtr)
	fmt.Println(*forkPtr)
	fmt.Println(stringVar)

	fmt.Println(flag.Args())
}
