package solutions

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Task05 - 5.	Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
// По истечению N секунд программа должна завершаться
func Task05() {
	// request time limit
	var N int
	fmt.Print("enter limit of time (ex. '5') => ")
	fmt.Scanf("%d\n", &N)
	timeLimit := time.NewTimer(time.Second * time.Duration(N))

	// open channel
	mainChan := make(chan int)

	// use WaitGroup for waiting all goroutines
	wg := sync.WaitGroup{}
	wg.Add(1)

	// read value in channels
	go func() {
		for {
			val, notEmpty := <-mainChan
			if notEmpty {
				fmt.Printf("value: %d\n", val)
			} else {
				// all goroutines id done
				wg.Done()
				return
			}
		}
	}()

L:
	for {
		select {
		case <-timeLimit.C:
			// after the end of time limit
			fmt.Println("end of time limit")
			close(mainChan)
			break L
		default:
			// before time limit
			mainChan <- rand.Intn(100) // push value in channels
			time.Sleep(time.Second)
		}
	}

	wg.Wait()
}
