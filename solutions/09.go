package solutions

import "fmt"

// Task09 - 9.	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
func Task09() {
	arr := []int{0, 1, 2, 3, 4, 5}

	// open write and read channels
	chWrite := make(chan int)
	chRead := make(chan int)

	// make goroutines
	go write(arr, chWrite)
	go calcAndWrite(chWrite, chRead)

	// show array values from chRead channel
	for val := range chRead {
		fmt.Println(val)
	}
}

// write values from arr to chWrite channel
func write(arr []int, chWrite chan int) {
	defer close(chWrite)
	for _, val := range arr {
		chWrite <- val
	}
}

// calcAndWrite - calculate sqr value from channel chWrite + write result to channel chRead
func calcAndWrite(chWrite chan int, chRead chan int) {
	defer close(chRead)
	for val := range chWrite {
		chRead <- val * 2
	}
}
