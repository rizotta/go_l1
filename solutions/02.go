package solutions

import (
	"fmt"
	"l1/pkg"
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
			sqr := pkg.Sqr(val)
			fmt.Println(sqr) // show random sqr. Like "fmt.Fprintln(os.Stdout, sqr)"
		}(val) // copy "k" variable because of goroutines
	}
	wg.Wait()
}
