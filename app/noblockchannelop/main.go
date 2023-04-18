package main

import "fmt"

func main() {
	message := make(chan string)
	signal := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println(msg)
	default:
		fmt.Println("no message to receive")
	}

	msg := "hi"
	select {
	case message <-msg:
		fmt.Println(<-message)
	default:
		fmt.Println("cannot send to message")
	}

	select {
	case msg := <-message:
		fmt.Println(msg)
	case sig := <-signal:
		fmt.Println(sig)
	default:
		fmt.Println("no activity")
	}
}