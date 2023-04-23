package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		up := strings.ToUpper(sc.Text())
		fmt.Println(up)
	}

	if err := sc.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}