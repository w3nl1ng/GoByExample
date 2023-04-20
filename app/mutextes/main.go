package main

import (
	"sync"
	"fmt"
)
type Container struct {
	mu sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counters[name]++
}

func main() {
	container := &Container{
		counters: map[string]int{"a": 0, "b": 0,},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, times int) {
		for i := 0; i < times; i++ {
			container.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	doIncrement("a", 1000)
	doIncrement("b", 1000)
	doIncrement("a", 1000)

	wg.Wait()
	fmt.Println(container.counters)
}