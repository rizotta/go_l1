package solutions

import (
	"fmt"
	"sync"
)

// Task03 - Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
func Task03() {
	arr := [5]int{2, 4, 6, 8, 10}
	sum := 0
	var wg sync.WaitGroup // for waiting all goroutines
	for _, val := range arr {
		wg.Add(1)
		val := val
		go func() { // add random sum
			defer wg.Done()
			sum += val * val
		}()
	}
	wg.Wait()
	fmt.Printf("sum: %d\n", sum)
}
