package main

import (
	"time"
	"fmt"
)

func main() {

	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)

	go func() {
		time.Sleep(1*time.Second)
		chan1 <- "chan1"
	}()

	go func() {
		time.Sleep(2*time.Second)
		chan2 <- "chan2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}
}