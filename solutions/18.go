package solutions

import (
	"fmt"
	"sync"
)

// Task18 - 18.	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.
func Task18() {
	c := newCounter()
	wg := &sync.WaitGroup{} // for waiting all goroutines

	// start 10 goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(worker int) {
			for j := 0; j < 5; j++ {
				c.Add()
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("counter:  %d\n", c.Get())
}

type Counter struct {
	sum int
	mu  sync.Mutex // use mutex to lock value
}

func (c *Counter) Add() {
	c.mu.Lock()
	c.sum++
	c.mu.Unlock()
}

func (c *Counter) Get() int {
	return c.sum
}

func newCounter() *Counter {
	return &Counter{
		0,
		sync.Mutex{},
	}
}
