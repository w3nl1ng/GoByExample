package main

import (
	"fmt"
	"sync"
	"time"
)

func work(i int) {
	fmt.Printf("worker %d starting\n", i)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", i)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		i := i

		go func() {
			defer wg.Done()
			work(i)
		}()
	}

	wg.Wait()
}