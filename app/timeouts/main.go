package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string, 1)
	go func() {
		time.Sleep(2*time.Second)
		chan1 <- "chan1"
	}()

	select {
	case result1 := <-chan1:
		fmt.Println(result1)
	case <-time.After(1*time.Second):
		fmt.Println("timeout 1s")
	}

	chan2 := make(chan string, 1)
	go func() {
		time.Sleep(2*time.Second)
		chan2<- "chan2"
	}()

	select {
	case result2 := <-chan2:
		fmt.Println(result2)
	case <-time.After(3*time.Second):
		fmt.Println("timeout 3s")
	}
}