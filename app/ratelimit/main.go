package main

import (
	"fmt"
	"time"
)

func main() {
	request := make(chan int, 5)
	for i := 0; i < 5; i++ {
		request <- i
	}
	close(request)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range request {
		<- limiter
		fmt.Println("request:", req, time.Now())
	}

	fmt.Println()

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	request2 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		request2 <- i
	}
	close(request2)

	go func(){
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	for req := range request2 {
		<- burstyLimiter
		fmt.Println("request:", req, time.Now())
	}
}