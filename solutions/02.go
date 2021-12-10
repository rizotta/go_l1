package solutions

import (
	"fmt"
	"math"
	"sync"
)

// Task02 - Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.
func Task02() {
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup // for waiting all goroutines
	for _, val := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			sqr := math.Pow(float64(val), 2)
			fmt.Println(sqr) // show random sqr, like "fmt.Fprintln(os.Stdout, sqr)"
		}(val) // copy "k" variable because of goroutines
	}
	wg.Wait()
}
