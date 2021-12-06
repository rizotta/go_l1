package solutions

import (
	"fmt"
	"math/rand"
	"time"
)

// Task04 - 4.	Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров,
// которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте
func Task04() {
	// request worker limit
	var N int
	fmt.Print("enter number of workers (ex. '3') => ")
	fmt.Scanf("%d\n", &N)

	// open and close channels
	mainChan := make(chan int)
	defer close(mainChan)

	// read available value from channel
	for i := 0; i < N; i++ {
		go worker(i, mainChan)
	}

	// push random value in channel
	for {
		mainChan <- rand.Intn(100)
		time.Sleep(time.Second)
	}
}

func worker(i int, mainChan chan int) {
	for {
		val := <-mainChan // get value from channel
		fmt.Printf("worker: %d, value: %d\n", i, val)
	}
}
