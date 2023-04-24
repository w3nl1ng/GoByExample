package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")

	fmt.Println("FOO", os.Getenv("FOO"))
	fmt.Println("BAR", os.Getenv("BAR"))

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		fmt.Println(pair[0])
	}
}
