package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d start job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finish job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	const workNum = 5
	jobs := make(chan int, workNum)
	results := make(chan int, workNum)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 0; j < 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 0; a < workNum; a++ {
		<-results
	}
	close(results)
}