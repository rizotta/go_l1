package solutions

import (
	"fmt"
	"math/rand"
	"sync"
)

// Task07 - 7.	Реализовать конкурентную запись данных в map.
func Task07() {
	const N = 5
	arr := map[int]int{}

	// use Mutex to lock some part of memory
	mx := sync.Mutex{}

	// for waiting all goroutines
	wg := sync.WaitGroup{}
	wg.Add(N)

	// make N goroutines
	for i := 0; i < N; i++ {
		go func(i int) {
			mx.Lock()               // lock goroutine
			defer mx.Unlock()       // unlock when function is end
			arr[i] = rand.Intn(100) // add value in map
			wg.Done()               // reduce WaitGroup counter
		}(i)
	}

	wg.Wait() // wait all goroutines

	// read map
	for k, val := range arr {
		fmt.Printf("arr[%d] = %d\n", k, val)
	}
}
