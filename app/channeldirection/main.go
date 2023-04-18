package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings chan<- string, pongs <-chan string) {
	pings <- <-pongs
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pongs, pings)
	fmt.Println(<-pongs)
}