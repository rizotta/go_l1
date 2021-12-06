package solutions

import (
	"fmt"
	"time"
)

// Task25 - 25.	Реализовать собственную функцию sleep.
func Task25() {
	fmt.Println("start")
	sleep(2)
	fmt.Println("end")
}

func sleep(s int) {
	<-time.After(time.Duration(s) * time.Second)
}
