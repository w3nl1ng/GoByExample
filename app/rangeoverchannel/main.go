package main

import "fmt"

func main() {
	chan1 := make(chan string, 2)

	chan1 <- "one"
	chan1 <- "two"
	close(chan1)

	for element := range chan1 {
		fmt.Println(element)
	}
}