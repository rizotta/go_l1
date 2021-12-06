package solutions

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

// Task06 - 6.	Реализовать все возможные способы остановки выполнения горутины.
func Task06() {
	const timeSec int64 = 3

	// 1 case: time limit
	timeLimit := time.NewTimer(time.Second * time.Duration(timeSec))
	go func() {
		fmt.Println("1 case start")
		for {
			select {
			case <-timeLimit.C:
				fmt.Println("1 case stop")
				return
			}
		}
	}()

	// 2 case: context WithTimeout done
	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second*time.Duration(timeSec))
	defer cancel1()
	go func() {
		fmt.Println("2 case start")
		for {
			select {
			case <-ctx1.Done():
				fmt.Println("2 case stop")
				return
			}
		}
	}()

	// 3 case: context WithTimeout + cancel with addition time
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second*time.Duration(timeSec))
	tick := time.NewTicker(time.Millisecond * 50)
	go func() {
		fmt.Println("3 case start")
		for {
			select {
			case _ = <-tick.C:
				cancel2()
			case <-ctx2.Done():
				fmt.Println("3 case stop")
				return
			}
		}
	}()

	// 4 case: context WithCancel
	ctx3, cancel3 := context.WithCancel(context.Background())

	go func() {
		fmt.Println("4 case start")
		cancel3()
		for {
			select {
			case <-ctx3.Done(): // if run cancel3
				fmt.Println("4 case stop")
				return
			}
		}
	}()

	// 5 case: use Goscheduler
	go func() {
		fmt.Println("5 case start")
		for i := 0; ; i++ {
			if i == 9999999 {
				fmt.Println("5 case stop")
				runtime.Gosched()
				return
			}
		}
	}()

	// 6 case: use time.After
	ch1 := make(chan bool)
	defer close(ch1)
	go func() {
		fmt.Println("6 case start")
		select {
		case <-time.After(time.Second * time.Duration(timeSec)):
			fmt.Println("6 case stop")
			return
		}
	}()

	// 7 case: use kill
	kill := make(chan interface{})
	defer close(kill)
	go func() {
		fmt.Println("7 case start")
		select {
		case <-kill:
			fmt.Println("7 case stop")
			return
		}
	}()

	// wait all cases
	time.Sleep(5 * time.Second)
}
